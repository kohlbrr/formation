package aws

import (
	"github.com/jmcgill/formation/core"
	//"github.com/aws/aws-sdk-go/aws"
)

type AwsOpsworksInstanceImporter struct {
}

// Lists all resources of this type
func (*AwsOpsworksInstanceImporter) Describe(meta interface{}) ([]*core.Instance, error) {
	return nil, nil
	//svc :=  meta.(*AWSClient).opsworksconn

	// Add code to list resources here
	//result, err := svc.ListBuckets(nil)
	//if err != nil {
	//  return nil, err
	//}

	//existingInstances := ... // e.g. result.Buckets
	//instances := make([]*core.Instance, len(existingInstances))
	//for i, existingInstance := range existingInstances {
	//	instances[i] = &core.Instance{
	//		Name: strings.Replace(aws.StringValue(existingInstance.Name), "-", "_", -1),
	//		ID:   aws.StringValue(existingInstance.Name),
	//	}
	//}

	// return instances, nil
}

// Describes which other resources this resource can reference
func (*AwsOpsworksInstanceImporter) Links() map[string]string {
	return map[string]string{}
}
