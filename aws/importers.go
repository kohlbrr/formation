package aws

import "github.com/jmcgill/formation/core"

func Importers() map[string]core.Importer {
    return map[string]core.Importer {
        // "aws_dms_replication_task": &AwsDmsReplicationTaskImporter{},
        // "aws_opsworks_memcached_layer": &AwsOpsworksMemcachedLayerImporter{},
        // "aws_inspector_assessment_target": &AwsInspectorAssessmentTargetImporter{},
        // "aws_route53_zone": &AwsRoute53ZoneImporter{},
        // "aws_vpc_endpoint": &AwsVpcEndpointImporter{},
        // "aws_iam_user_policy": &AwsIamUserPolicyImporter{},
        // "aws_ses_receipt_rule_set": &AwsSesReceiptRuleSetImporter{},
        // "aws_network_interface_sg_attachment": &AwsNetworkInterfaceSgAttachmentImporter{},
        // "aws_redshift_cluster": &AwsRedshiftClusterImporter{},
        // "aws_network_acl_rule": &AwsNetworkAclRuleImporter{},
        // "aws_elastic_beanstalk_environment": &AwsElasticBeanstalkEnvironmentImporter{},
        // "aws_dms_replication_instance": &AwsDmsReplicationInstanceImporter{},
        // "aws_internet_gateway": &AwsInternetGatewayImporter{},
        // "aws_ssm_association": &AwsSsmAssociationImporter{},
        // "aws_opsworks_rds_db_instance": &AwsOpsworksRdsDbInstanceImporter{},
        // "aws_route53_zone_association": &AwsRoute53ZoneAssociationImporter{},
        // "aws_sqs_queue_policy": &AwsSqsQueuePolicyImporter{},
        // "aws_wafregional_ipset": &AwsWafregionalIpsetImporter{},
        // "aws_alb_listener_rule": &AwsAlbListenerRuleImporter{},
        // "aws_dynamodb_table": &AwsDynamodbTableImporter{},
        // "aws_codepipeline": &AwsCodepipelineImporter{},
        // "aws_eip": &AwsEipImporter{},
        // "aws_simpledb_domain": &AwsSimpledbDomainImporter{},
        // "aws_inspector_assessment_template": &AwsInspectorAssessmentTemplateImporter{},
        // "aws_ecs_service": &AwsEcsServiceImporter{},
        // "aws_iam_user_ssh_key": &AwsIamUserSshKeyImporter{},
        // "aws_sns_topic_policy": &AwsSnsTopicPolicyImporter{},
        // "aws_ssm_patch_baseline": &AwsSsmPatchBaselineImporter{},
        // "aws_db_parameter_group": &AwsDbParameterGroupImporter{},
        // "aws_directory_service_directory": &AwsDirectoryServiceDirectoryImporter{},
        // "aws_ebs_volume": &AwsEbsVolumeImporter{},
        // "aws_lambda_permission": &AwsLambdaPermissionImporter{},
        // "aws_elasticache_parameter_group": &AwsElasticacheParameterGroupImporter{},
        // "aws_appautoscaling_policy": &AwsAppautoscalingPolicyImporter{},
        // "aws_db_security_group": &AwsDbSecurityGroupImporter{},
        // "aws_api_gateway_integration": &AwsApiGatewayIntegrationImporter{},
        // "aws_ses_active_receipt_rule_set": &AwsSesActiveReceiptRuleSetImporter{},
        // "aws_default_network_acl": &AwsDefaultNetworkAclImporter{},
        // "aws_cloudwatch_event_rule": &AwsCloudwatchEventRuleImporter{},
        // "aws_vpc_peering_connection": &AwsVpcPeeringConnectionImporter{},
        // "aws_autoscaling_lifecycle_hook": &AwsAutoscalingLifecycleHookImporter{},
        // "aws_opsworks_custom_layer": &AwsOpsworksCustomLayerImporter{},
        // "aws_network_interface": &AwsNetworkInterfaceImporter{},
        // "aws_instance": &AwsInstanceImporter{},
        // "aws_codecommit_repository": &AwsCodecommitRepositoryImporter{},
        // "aws_api_gateway_method_response": &AwsApiGatewayMethodResponseImporter{},
        // "aws_api_gateway_method": &AwsApiGatewayMethodImporter{},
        // "aws_sqs_queue": &AwsSqsQueueImporter{},
        // "aws_security_group_rule": &AwsSecurityGroupRuleImporter{},
        // "aws_flow_log": &AwsFlowLogImporter{},
        // "aws_wafregional_byte_match_set": &AwsWafregionalByteMatchSetImporter{},
        // "aws_iam_role_policy_attachment": &AwsIamRolePolicyAttachmentImporter{},
        // "aws_lambda_alias": &AwsLambdaAliasImporter{},
        // "aws_cloudwatch_event_target": &AwsCloudwatchEventTargetImporter{},
        // "aws_lb_cookie_stickiness_policy": &AwsLbCookieStickinessPolicyImporter{},
        // "aws_vpn_gateway_attachment": &AwsVpnGatewayAttachmentImporter{},
        // "aws_api_gateway_model": &AwsApiGatewayModelImporter{},
        // "aws_sfn_activity": &AwsSfnActivityImporter{},
        // "aws_security_group_rule_migrate": &AwsSecurityGroupRuleMigrateImporter{},
        // "aws_ses_receipt_rule": &AwsSesReceiptRuleImporter{},
        // "aws_s3_bucket_object": &AwsS3BucketObjectImporter{},
        "aws_iam_role": &AwsIamRoleImporter{},
        // "aws_alb_target_group_attachment": &AwsAlbTargetGroupAttachmentImporter{},
        // "aws_elastic_beanstalk_application": &AwsElasticBeanstalkApplicationImporter{},
        // "aws_kms_key": &AwsKmsKeyImporter{},
        // "aws_iam_policy": &AwsIamPolicyImporter{},
        // "aws_ecr_repository_policy": &AwsEcrRepositoryPolicyImporter{},
        // "aws_db_event_subscription": &AwsDbEventSubscriptionImporter{},
        // "aws_default_security_group": &AwsDefaultSecurityGroupImporter{},
        // "aws_rds_cluster_parameter_group": &AwsRdsClusterParameterGroupImporter{},
        // "aws_api_gateway_usage_plan": &AwsApiGatewayUsagePlanImporter{},
        // "aws_autoscaling_policy": &AwsAutoscalingPolicyImporter{},
        // "aws_key_pair_migrate": &AwsKeyPairMigrateImporter{},
        // "aws_iam_account_password_policy": &AwsIamAccountPasswordPolicyImporter{},
        // "aws_ssm_maintenance_window_task": &AwsSsmMaintenanceWindowTaskImporter{},
        // "aws_spot_datafeed_subscription": &AwsSpotDatafeedSubscriptionImporter{},
        // "aws_spot_instance_request": &AwsSpotInstanceRequestImporter{},
        // "aws_eip_association": &AwsEipAssociationImporter{},
        // "aws_default_vpc": &AwsDefaultVpcImporter{},
        // "aws_app_cookie_stickiness_policy": &AwsAppCookieStickinessPolicyImporter{},
        // "aws_efs_file_system": &AwsEfsFileSystemImporter{},
        // "aws_devicefarm_project": &AwsDevicefarmProjectImporter{},
        // "aws_s3_bucket_notification": &AwsS3BucketNotificationImporter{},
        // "aws_route": &AwsRouteImporter{},
        // "aws_cloudfront_distribution": &AwsCloudfrontDistributionImporter{},
        // "aws_codebuild_project_migrate": &AwsCodebuildProjectMigrateImporter{},
        // "aws_opsworks_java_app_layer": &AwsOpsworksJavaAppLayerImporter{},
        // "aws_iam_user": &AwsIamUserImporter{},
        // "aws_iam_user_login_profile": &AwsIamUserLoginProfileImporter{},
        // "aws_elasticsearch_domain": &AwsElasticsearchDomainImporter{},
        // "aws_inspector_group": &AwsInspectorGroupImporter{},
        // "aws_api_gateway_integration_response": &AwsApiGatewayIntegrationResponseImporter{},
        // "aws_elastic_transcoder_preset": &AwsElasticTranscoderPresetImporter{},
        // "aws_waf_byte_match_set": &AwsWafByteMatchSetImporter{},
        // "aws_codedeploy_deployment_group": &AwsCodedeployDeploymentGroupImporter{},
        // "aws_iam_policy_attachment": &AwsIamPolicyAttachmentImporter{},
        // "aws_route_table_association": &AwsRouteTableAssociationImporter{},
        // "aws_waf_sql_injection_match_set": &AwsWafSqlInjectionMatchSetImporter{},
        // "aws_api_gateway_usage_plan_key": &AwsApiGatewayUsagePlanKeyImporter{},
        // "aws_opsworks_rails_app_layer": &AwsOpsworksRailsAppLayerImporter{},
        // "aws_elasticache_subnet_group": &AwsElasticacheSubnetGroupImporter{},
        // "aws_subnet_migrate": &AwsSubnetMigrateImporter{},
        // "aws_cloudwatch_metric_alarm_migrate": &AwsCloudwatchMetricAlarmMigrateImporter{},
        // "aws_vpc_endpoint_route_table_association": &AwsVpcEndpointRouteTableAssociationImporter{},
        // "aws_vpc_peering_connection_accepter": &AwsVpcPeeringConnectionAccepterImporter{},
        // "aws_ssm_parameter": &AwsSsmParameterImporter{},
        // "aws_api_gateway_client_certificate": &AwsApiGatewayClientCertificateImporter{},
        // "aws_waf_ipset": &AwsWafIpsetImporter{},
        // "aws_elb": &AwsElbImporter{},
        // "aws_codebuild_project": &AwsCodebuildProjectImporter{},
        // "aws_elastic_beanstalk_environment_migrate": &AwsElasticBeanstalkEnvironmentMigrateImporter{},
        // "aws_api_gateway_rest_api": &AwsApiGatewayRestApiImporter{},
        // "aws_lightsail_instance": &AwsLightsailInstanceImporter{},
        // "aws_network_interface_attachment": &AwsNetworkInterfaceAttachmentImporter{},
        // "aws_ami_copy": &AwsAmiCopyImporter{},
        // "aws_elb_attachment": &AwsElbAttachmentImporter{},
        // "aws_emr_instance_group": &AwsEmrInstanceGroupImporter{},
        // "aws_waf_web_acl": &AwsWafWebAclImporter{},
        // "aws_kms_alias": &AwsKmsAliasImporter{},
        // "aws_iam_group_membership": &AwsIamGroupMembershipImporter{},
        // "aws_route_table": &AwsRouteTableImporter{},
        // "aws_ami_from_instance": &AwsAmiFromInstanceImporter{},
        // "aws_opsworks_nodejs_app_layer": &AwsOpsworksNodejsAppLayerImporter{},
        // "aws_dms_certificate": &AwsDmsCertificateImporter{},
        // "aws_config_configuration_recorder_status": &AwsConfigConfigurationRecorderStatusImporter{},
        // "aws_proxy_protocol_policy": &AwsProxyProtocolPolicyImporter{},
        // "aws_s3_bucket_policy": &AwsS3BucketPolicyImporter{},
        // "aws_dynamodb_table_migrate": &AwsDynamodbTableMigrateImporter{},
        // "aws_cloudwatch_log_destination": &AwsCloudwatchLogDestinationImporter{},
        "aws_s3_bucket": &AwsS3BucketImporter{},
        // "aws_db_instance": &AwsDbInstanceImporter{},
        // "aws_cloudwatch_log_destination_policy": &AwsCloudwatchLogDestinationPolicyImporter{},
        // "aws_cloudfront_origin_access_identity": &AwsCloudfrontOriginAccessIdentityImporter{},
        // "aws_load_balancer_policy": &AwsLoadBalancerPolicyImporter{},
        // "aws_vpc_migrate": &AwsVpcMigrateImporter{},
        // "aws_lambda_function": &AwsLambdaFunctionImporter{},
        // "aws_iam_role_policy": &AwsIamRolePolicyImporter{},
        // "aws_autoscaling_attachment": &AwsAutoscalingAttachmentImporter{},
        // "aws_cloudwatch_log_metric_filter": &AwsCloudwatchLogMetricFilterImporter{},
        // "aws_opsworks_instance": &AwsOpsworksInstanceImporter{},
        // "aws_default_vpc_dhcp_options": &AwsDefaultVpcDhcpOptionsImporter{},
        // "aws_api_gateway_domain_name": &AwsApiGatewayDomainNameImporter{},
        // "aws_vpn_connection": &AwsVpnConnectionImporter{},
        // "aws_api_gateway_stage": &AwsApiGatewayStageImporter{},
        // "aws_lightsail_key_pair": &AwsLightsailKeyPairImporter{},
        // "aws_dms_replication_subnet_group": &AwsDmsReplicationSubnetGroupImporter{},
        // "aws_launch_configuration": &AwsLaunchConfigurationImporter{},
        // "aws_api_gateway_account": &AwsApiGatewayAccountImporter{},
        // "aws_cloudwatch_log_group": &AwsCloudwatchLogGroupImporter{},
        // "aws_ami": &AwsAmiImporter{},
        // "aws_dms_endpoint": &AwsDmsEndpointImporter{},
        // "aws_opsworks_ganglia_layer": &AwsOpsworksGangliaLayerImporter{},
        // "aws_rds_cluster_instance": &AwsRdsClusterInstanceImporter{},
        // "aws_api_gateway_request_validator": &AwsApiGatewayRequestValidatorImporter{},
        // "aws_elasticsearch_domain_policy": &AwsElasticsearchDomainPolicyImporter{},
        // "aws_api_gateway_deployment": &AwsApiGatewayDeploymentImporter{},
        // "aws_ebs_snapshot": &AwsEbsSnapshotImporter{},
        // "aws_cloudtrail": &AwsCloudtrailImporter{},
        // "aws_config_delivery_channel": &AwsConfigDeliveryChannelImporter{},
        // "aws_security_group": &AwsSecurityGroupImporter{},
        // "aws_instance_migrate": &AwsInstanceMigrateImporter{},
        // "aws_codedeploy_deployment_config": &AwsCodedeployDeploymentConfigImporter{},
        // "aws_ssm_maintenance_window": &AwsSsmMaintenanceWindowImporter{},
        // "aws_network_acl": &AwsNetworkAclImporter{},
        // "aws_db_subnet_group": &AwsDbSubnetGroupImporter{},
        // "aws_ecs_cluster": &AwsEcsClusterImporter{},
        // "aws_kinesis_stream": &AwsKinesisStreamImporter{},
        // "aws_opsworks_stack": &AwsOpsworksStackImporter{},
        // "aws_db_snapshot": &AwsDbSnapshotImporter{},
        // "aws_cloudformation_stack": &AwsCloudformationStackImporter{},
        // "aws_vpn_gateway_route_propagation": &AwsVpnGatewayRoutePropagationImporter{},
        // "aws_opsworks_application": &AwsOpsworksApplicationImporter{},
        // "aws_cognito_identity_pool": &AwsCognitoIdentityPoolImporter{},
        // "aws_api_gateway_authorizer": &AwsApiGatewayAuthorizerImporter{},
        // "aws_codedeploy_app": &AwsCodedeployAppImporter{},
        // "aws_spot_fleet_request_migrate": &AwsSpotFleetRequestMigrateImporter{},
        // "aws_route53_delegation_set": &AwsRoute53DelegationSetImporter{},
        // "aws_cloudwatch_log_stream": &AwsCloudwatchLogStreamImporter{},
        // "aws_iam_access_key": &AwsIamAccessKeyImporter{},
        // "aws_route53_health_check": &AwsRoute53HealthCheckImporter{},
        // "aws_waf_rule": &AwsWafRuleImporter{},
        // "aws_autoscaling_group_waiting": &AwsAutoscalingGroupWaitingImporter{},
        // "aws_volume_attachment": &AwsVolumeAttachmentImporter{},
        // "aws_snapshot_create_volume_permission": &AwsSnapshotCreateVolumePermissionImporter{},
        // "aws_lightsail_static_ip_attachment": &AwsLightsailStaticIpAttachmentImporter{},
        // "aws_vpn_gateway": &AwsVpnGatewayImporter{},
        // "aws_egress_only_internet_gateway": &AwsEgressOnlyInternetGatewayImporter{},
        // "aws_opsworks_static_web_layer": &AwsOpsworksStaticWebLayerImporter{},
        // "aws_config_configuration_recorder": &AwsConfigConfigurationRecorderImporter{},
        // "aws_iam_saml_provider": &AwsIamSamlProviderImporter{},
        // "aws_opsworks_php_app_layer": &AwsOpsworksPhpAppLayerImporter{},
        // "aws_efs_mount_target": &AwsEfsMountTargetImporter{},
        // "aws_opsworks_user_profile": &AwsOpsworksUserProfileImporter{},
        // "aws_redshift_subnet_group": &AwsRedshiftSubnetGroupImporter{},
        // "aws_load_balancer_backend_server_policy": &AwsLoadBalancerBackendServerPolicyImporter{},
        // "aws_lightsail_domain": &AwsLightsailDomainImporter{},
        // "aws_ses_domain_identity": &AwsSesDomainIdentityImporter{},
        // "aws_main_route_table_association": &AwsMainRouteTableAssociationImporter{},
        // "aws_subnet": &AwsSubnetImporter{},
        // "aws_customer_gateway": &AwsCustomerGatewayImporter{},
        // "aws_glacier_vault": &AwsGlacierVaultImporter{},
        // "aws_autoscaling_schedule": &AwsAutoscalingScheduleImporter{},
        // "aws_ssm_patch_group": &AwsSsmPatchGroupImporter{},
        // "aws_codecommit_trigger": &AwsCodecommitTriggerImporter{},
        // "aws_iam_instance_profile": &AwsIamInstanceProfileImporter{},
        // "aws_elastic_beanstalk_application_version": &AwsElasticBeanstalkApplicationVersionImporter{},
        // "aws_alb_target_group": &AwsAlbTargetGroupImporter{},
        // "aws_elastic_beanstalk_configuration_template": &AwsElasticBeanstalkConfigurationTemplateImporter{},
        // "aws_ses_receipt_filter": &AwsSesReceiptFilterImporter{},
        // "aws_route53_record_migrate": &AwsRoute53RecordMigrateImporter{},
        // "aws_default_route_table": &AwsDefaultRouteTableImporter{},
        // "aws_autoscaling_notification": &AwsAutoscalingNotificationImporter{},
        // "aws_key_pair": &AwsKeyPairImporter{},
        // "aws_lambda_event_source_mapping": &AwsLambdaEventSourceMappingImporter{},
        // "aws_sfn_state_machine": &AwsSfnStateMachineImporter{},
        // "aws_db_option_group": &AwsDbOptionGroupImporter{},
        // "vpn_connection_route": &VpnConnectionRouteImporter{},
        // "aws_load_balancer_listener_policy": &AwsLoadBalancerListenerPolicyImporter{},
        // "aws_config_config_rule": &AwsConfigConfigRuleImporter{},
        // "aws_default_subnet": &AwsDefaultSubnetImporter{},
        // "aws_elastic_transcoder_pipeline": &AwsElasticTranscoderPipelineImporter{},
        // "aws_ecs_task_definition": &AwsEcsTaskDefinitionImporter{},
        // "aws_ecr_repository": &AwsEcrRepositoryImporter{},
        // "aws_redshift_parameter_group": &AwsRedshiftParameterGroupImporter{},
        // "aws_spot_fleet_request": &AwsSpotFleetRequestImporter{},
        // "aws_api_gateway_base_path_mapping": &AwsApiGatewayBasePathMappingImporter{},
        // "aws_iam_user_policy_attachment": &AwsIamUserPolicyAttachmentImporter{},
        // "aws_vpc": &AwsVpcImporter{},
        // "aws_ses_event_destination": &AwsSesEventDestinationImporter{},
        // "aws_cloudwatch_log_subscription_filter": &AwsCloudwatchLogSubscriptionFilterImporter{},
        // "aws_alb_listener": &AwsAlbListenerImporter{},
        // "aws_iam_group": &AwsIamGroupImporter{},
        // "aws_emr_cluster": &AwsEmrClusterImporter{},
        // "aws_iam_group_policy_attachment": &AwsIamGroupPolicyAttachmentImporter{},
        // "aws_cloudwatch_metric_alarm": &AwsCloudwatchMetricAlarmImporter{},
        // "aws_iam_account_alias": &AwsIamAccountAliasImporter{},
        // "aws_emr_security_configuration": &AwsEmrSecurityConfigurationImporter{},
        // "aws_iam_group_policy": &AwsIamGroupPolicyImporter{},
        // "aws_kinesis_firehose_delivery_stream": &AwsKinesisFirehoseDeliveryStreamImporter{},
        // "aws_autoscaling_group": &AwsAutoscalingGroupImporter{},
        // "aws_opsworks_haproxy_layer": &AwsOpsworksHaproxyLayerImporter{},
        // "aws_waf_xss_match_set": &AwsWafXssMatchSetImporter{},
        // "aws_ssm_maintenance_window_target": &AwsSsmMaintenanceWindowTargetImporter{},
        // "aws_elasticache_replication_group": &AwsElasticacheReplicationGroupImporter{},
        // "aws_vpc_dhcp_options_association": &AwsVpcDhcpOptionsAssociationImporter{},
        // "aws_opsworks_permission": &AwsOpsworksPermissionImporter{},
        // "aws_kinesis_firehose_delivery_stream_migrate": &AwsKinesisFirehoseDeliveryStreamMigrateImporter{},
        // "aws_api_gateway_resource": &AwsApiGatewayResourceImporter{},
        // "aws_opsworks_mysql_layer": &AwsOpsworksMysqlLayerImporter{},
        // "aws_iam_server_certificate": &AwsIamServerCertificateImporter{},
        // "aws_lb_ssl_negotiation_policy": &AwsLbSslNegotiationPolicyImporter{},
        // "aws_ami_launch_permission": &AwsAmiLaunchPermissionImporter{},
        // "aws_route53_record": &AwsRoute53RecordImporter{},
        // "aws_elasticache_cluster": &AwsElasticacheClusterImporter{},
        // "aws_redshift_security_group": &AwsRedshiftSecurityGroupImporter{},
        // "aws_rds_cluster": &AwsRdsClusterImporter{},
        // "aws_ssm_document": &AwsSsmDocumentImporter{},
        // "aws_ssm_activation": &AwsSsmActivationImporter{},
        // "aws_lightsail_static_ip": &AwsLightsailStaticIpImporter{},
        // "aws_sns_topic": &AwsSnsTopicImporter{},
        // "aws_api_gateway_method_settings": &AwsApiGatewayMethodSettingsImporter{},
        // "aws_elasticache_security_group": &AwsElasticacheSecurityGroupImporter{},
        // "aws_placement_group": &AwsPlacementGroupImporter{},
        // "aws_waf_size_constraint_set": &AwsWafSizeConstraintSetImporter{},
        // "aws_ses_configuration_set": &AwsSesConfigurationSetImporter{},
        // "aws_vpc_dhcp_options": &AwsVpcDhcpOptionsImporter{},
        // "aws_iam_openid_connect_provider": &AwsIamOpenidConnectProviderImporter{},
        // "aws_sns_topic_subscription": &AwsSnsTopicSubscriptionImporter{},
        // "aws_alb": &AwsAlbImporter{},
        // "aws_appautoscaling_target": &AwsAppautoscalingTargetImporter{},
        // "aws_api_gateway_api_key": &AwsApiGatewayApiKeyImporter{},
        // "aws_nat_gateway": &AwsNatGatewayImporter{},
    }
}