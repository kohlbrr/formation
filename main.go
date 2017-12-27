package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jmcgill/formation/aws"
	"github.com/jmcgill/formation/core"
	"os"
	"strings"

	"github.com/hashicorp/terraform/config/configschema"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	aws2 "github.com/terraform-providers/terraform-provider-aws/aws"
)

// TODO(jimmy): Decide exactly where this belongs
//func RegisterImporters() map[string]core.Importer {
//	return map[string]core.Importer{
//		"aws_s3_bucket": &aws.S3BucketImporter{},
//		// "aws_iam_role": &aws.IAmRoleImporter{},
//	}
//}

type UIInput struct {
}

func (u *UIInput) Input(opts *terraform.InputOpts) (string, error) {
	fmt.Printf("Asking for input %s\n", opts.Query)
	return "us-west-2", nil
}

// Compare a converted resource to the Terraform Schema for that resource type, and
// mark any fields which we know are computed. These are fields which we want to exist in our
// .tfstate files, but not in our .tf files
func MarkComputedFields(r *core.InlineResource, s *configschema.Block) *core.InlineResource {
	for _, f := range r.Fields {
		// This is a scalar object
		if f.FieldType == core.SCALAR && f.Path != "id" {
			f.Computed = s.Attributes[f.Key].Computed
			continue
		}

		// The children of maps are not known ahead of time, so we can't walk further
		// in the schema.
		if f.FieldType == core.MAP {
			continue
		}

		if f.FieldType == core.LIST {
			MarkComputedFields(f.NestedValue, &s.BlockTypes[f.Key].Block)
		}

		if f.FieldType == core.NESTED {
			MarkComputedFields(f.NestedValue, s)
		}
	}

	return r
}

func ValueSet(r *core.InlineResource, key string) bool {
	fmt.Printf("Looking for value set %s\n", key)
	for _, f := range r.Fields {
		fmt.Printf("Checking whether value is set - comparing %s to %s\n", f.Key, key)
		if f.Key == key {
			return true
		}
	}
	return false
}

func AppendToInstancePath(prefix string, suffix string) string {
	if prefix == "" {
		return suffix
	}
	return prefix + "." + suffix
}

func DecorateWithDefaultFields(instanceState *terraform.InstanceState, r *core.InlineResource, terraformSchema map[string]*schema.Schema, path string) *core.InlineResource {
	// Are there any default fields at this level in the schema that are not set in our resource?
	for key, v := range terraformSchema {
		if v.Default == nil {
			continue
		}

		if ValueSet(r, key) {
			continue
		}

		fmt.Printf("****** TRYING TO SET DEFAULT FIELD %s\n", key)

		var defaultValue string

		// TODO(jimmy): Switch to a TYPE enum - can pull this straight out of the schema object
		isBool := false

		switch v.Default.(type) {
		case bool:
			defaultValue = strconv.FormatBool(v.Default.(bool))
			isBool = true
		case string:
			defaultValue = v.Default.(string)
		}

		instancePath := AppendToInstancePath(path, key)
		instanceState.Attributes[instancePath] = defaultValue

		field := &core.Field{
			FieldType: core.SCALAR,
			Path: instancePath,
			Key:     key,
			ScalarValue: &core.ScalarValue{
				StringValue: defaultValue,
				IsBool:      isBool,
			},
		}
		r.Fields = append(r.Fields, field)
	}

	for _, f := range r.Fields {
		instancePath := AppendToInstancePath(path, f.Key)

		// TODO(jimmy): Work out how default values should work for scalar lists
		// HACK: Assume all resources are nested
		if f.FieldType == core.LIST {
			// Scalar lists will never have a default value, so we don't need to expand the parent key when
			// recursing into a list. Instead we will rely on each nested object to do this.
			fmt.Printf("Recursing into list\n")
			DecorateWithDefaultFields(instanceState, f.NestedValue.Fields[0].NestedValue, terraformSchema[f.Key].Elem.(*schema.Resource).Schema, instancePath)
		}

		//if f.FieldType == core.NESTED {
		//	fmt.Printf("Recursing into nested resource - instance state is now %s\n", instanceState)
		//	// TODO(jimmy): parentKey needs to be the path to the nested objecdt e.g. foo.1234 but I don't
		//	// currently set this while parsing (bucause it would require lookahead). Fix this - perhaps by having
		//	// the first child discovered fill it in.
		//	//
		//	// This depends on whether there is ever a list with a required-and-default set of keys that is not filled
		//	// in by importing. I need to work through some more concrete examples first.
		//	// TODO(jimmy): This is probably wrong?
		//	DecorateWithDefaultFields(instanceState, f.NestedValue, terraformSchema, instancePath)
		//}
	}

	return r
}

type FieldIndexEntry struct {
	path     string
	resource *core.Resource
}
type FieldIndex map[string][]*FieldIndexEntry

func IndexResource(resource *core.Resource, index FieldIndex) {
	IndexFields(resource, resource.Fields, index)
}

// Index all fields
// TODO(jimmy): Implement a visitor pattern that allows this to be done during the parsing pass.
func IndexFields(resource *core.Resource, r *core.InlineResource, index FieldIndex) {
	for _, f := range r.Fields {
		// This is a scalar object
		if f.FieldType == core.SCALAR {
			// Don't index null values
			if f.ScalarValue.StringValue == "" {
				continue
			}

			fmt.Printf("**** %s : %s\n", f.Path, f.ScalarValue.StringValue)
			e := &FieldIndexEntry{
				path:     resource.Type + "." + f.Path,
				resource: resource,
			}
			index[f.ScalarValue.StringValue] = append(index[f.ScalarValue.StringValue], e)
			continue
		}

		// The children of maps are not known ahead of time, and are never referenced from other resources.
		if f.FieldType == core.MAP {
			continue
		}

		if f.FieldType == core.LIST || f.FieldType == core.NESTED {

			IndexFields(resource, f.NestedValue, index)
		}
	}
}

func FindLink(index FieldIndex, value string, allowedPath string) (*core.Resource, bool) {
	fmt.Printf("Looking for link for value: %s\n", value)
	if fields, ok := index[value]; ok {
		fmt.Printf("Yup - that is in our reverse index!\n")
		for _, field := range fields {
			fmt.Printf("Performing allowed path comparison %s - %s\n", field.path, allowedPath)
			if field.path == allowedPath {
				fmt.Printf("Paths matched!\n")
				return field.resource, true
			}
		}
	}

	return nil, false
}

func LinkFields(r *core.InlineResource, links map[string]string, index FieldIndex) *core.InlineResource {
	RecursivelyLinkFields(r, links, index, "")
	return r
}

func RecursivelyLinkFields(r *core.InlineResource, links map[string]string, index FieldIndex, path string) {
	for _, f := range r.Fields {
		// This is a scalar object
		if f.FieldType == core.SCALAR {
			var z string
			if path != "" {
				z = path + "." + f.Key
			} else {
				z = f.Key
			}

			fmt.Printf("Looking for a rule for %s\n", z)

			// This field _can_ link to another resource
			if allowedPath, ok := links[z]; ok {
				fmt.Printf("Found rule: %s\n", allowedPath)
				if resource, ok := FindLink(index, f.ScalarValue.StringValue, allowedPath); ok {
					fmt.Printf("Found a valid link\n")
					// Substitute in the resource name
					resolvedPath := strings.Replace(allowedPath, resource.Type, resource.Type+"."+resource.Name, 1)
					f.Link = resolvedPath
				}
			}
			continue
		}

		// Can a Map ever contain a field which links to another resource? TBD
		if f.FieldType == core.MAP {
			continue
		}

		if f.FieldType == core.LIST {
			var z string
			if path != "" {
				z = path + "." + f.Key
			} else {
				z = f.Key
			}

			RecursivelyLinkFields(f.NestedValue, links, index, z)
		}

		if f.FieldType == core.NESTED {
			RecursivelyLinkFields(f.NestedValue, links, index, path)
		}
	}
}

type ImportedResource struct {
	resource *core.Resource
	state *terraform.InstanceState
}

func main() {
	fmt.Println("Welcome to formation")

	// TODO(jimmy): Bundle these into an object
	allResources := make(map[string][]*ImportedResource)

	// TODO(JIMMY): hide this away in a struct
	index := make(FieldIndex)

	importers := aws.Importers()

	// Configure Terraform Plugin
	provider := aws2.Provider()

	// A duplicate Provider which is guaranteed to be configured in the same way as
	// the AWS probvider above. This is needed so that we can access otherwisr private
	// fields that are configured during initialization.
	localProvider := aws.Provider()
	c := terraform.NewResourceConfig(nil)
	localProvider.Input(&UIInput{}, c)

	err := localProvider.Configure(c)
	if err != nil {
		fmt.Printf("Error configuring internal provider")
	}
	localSchemaProvider := localProvider.(*schema.Provider)

	// spew.Dump(x)

	//v := reflect.ValueOf(provider)
	//x := (v.FieldByName("Schema").Interface()).(map[string]*schema.Schema)
	// fmt.Println(y.Interface())

	//x := (*schema.Provider)(provider)
	//x := (*schema.Provider)(unsafe.Pointer(&provider))
	//for key, _ := range x {
	//	fmt.Printf("%s\n", key)
	//}
	//spew.Dump((*x).Schema)

	// Why doesn't this return the Schema I want?

	c = terraform.NewResourceConfig(nil)
	provider.Input(&UIInput{}, c)
	err = provider.Configure(c)
	if err != nil {
		fmt.Printf("Error in configuration: %s", err)
	}

	// For each importer
	for resourceType, importer := range importers {
		// TODO(jimmy): Initialize importer with our provider configuration so that
		// we're guaranteed to be using the same AWS credentials.
		instances, err := importer.Describe(localSchemaProvider.Meta())
		if err != nil {
			panic(err)
		}

		for _, instance := range instances {
			instanceInfo := &terraform.InstanceInfo{
				// Id is a unique name to represent this instance. This is not related
				// to InstanceState.ID in any way.
				Id: instance.ID,

				// Type is the resource type of this instance
				Type: resourceType,
			}

			instancesToImport, err := provider.ImportState(instanceInfo, instance.ID)
			if err != nil {
				fmt.Printf("Error importing instances: %s", err)
			}

			// TODO(jimmy): It's not always safe to assume that import returns a single reso	urce.
			// If it returns multiple, do we need to refresh each one individually? I need a good
			// case study here before guessing at the behaviour.
			instanceState, err := provider.Refresh(instanceInfo, instancesToImport[0])
			if err != nil {
				fmt.Printf("Error refreshing Instance State: %s", err)
				panic("Error refreshing Instance State")
			}

			// Convert this resource from Terraform's internal format to a Formation Resource
			parser := core.InstanceStateParser{}
			resource := parser.Parse(instanceState)

			/// Fill in name and type
			resource.Name = instance.Name
			resource.Type = resourceType

			// Get the schema for this resource
			request := &terraform.ProviderSchemaRequest{
				ResourceTypes: []string{resourceType},
			}
			s, _ := provider.GetSchema(request)

			// Mark computed fields - we don't want to output these
			MarkComputedFields(resource.Fields, s.ResourceTypes[resourceType])

			// To get the resource schema we need to poke into the internal implementation of the AWS provider
			schemaProvider := provider.(*schema.Provider)
			DecorateWithDefaultFields(instanceState, resource.Fields, schemaProvider.ResourcesMap[resourceType].Schema, "")

			// Store this resource for later
			allResources[resourceType] = append(allResources[resourceType], &ImportedResource{
				resource: resource,
				state: instanceState,
			})

			// Index this resource
			IndexFields(resource, resource.Fields, index)
		}
	}

	fmt.Printf("Finished importing and indexing fields\n")

	// At this point, all resources have been index
	for resourceType, resources := range allResources {
		fmt.Printf("Resource type is: %s, %i\n", resourceType, len(resources))
		f, err := os.Create(resourceType + ".tf")
		defer f.Close()

		if err != nil {
			fmt.Printf("Error creating file for resource: %s\n", resourceType)
			return
		}

		fmt.Printf("Getting here too\n")

		for i, importedResource := range resources {
			resource := importedResource.resource
			LinkFields(resource.Fields, importers[resource.Type].Links(), index)

			// TODO(jimmy): Print to a file
			printer := core.Printer{}
			printer.PrintToFile(f, resource)

			// Space out resources for readability
			if i != len(resources)-1 {
				fmt.Fprint(f, "\n\n")
			}
		}
	}

	// Write TFState
	state := terraform.State{
		Version: 3,

		// TODO(jimmy): Pull this out of the terraform Context object
		TFVersion: "0.11.1",

		Serial: 1,

		// TODO(jimmy): Work out how this is generated
		Lineage: "Formation",

		Modules: []*terraform.ModuleState{
			{
				Path: []string{
					"root",
				},
			},
		},
	}

	state.Modules[0].Resources = make(map[string]*terraform.ResourceState)
	for _, resources := range allResources {
		for _, importedResource := range resources {
			resource := importedResource.resource
			r := &terraform.ResourceState{
				Type:    resource.Type,
				Primary: importedResource.state,
				Provider: "provider.aws",
			}
			state.Modules[0].Resources[resource.Type+"."+resource.Name] = r
		}
	}

	f, err := os.Create("terraform.tfstate")
	if err != nil {
		panic("Failure to create TFState file")
	}
	defer f.Close()

	j, _ := json.MarshalIndent(state, "", "    ")
	f.Write(j)
	fmt.Printf("Def getting here...\n")
}
