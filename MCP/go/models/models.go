package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeAssociationExecutionsRequest represents the DescribeAssociationExecutionsRequest schema from the OpenAPI specification
type DescribeAssociationExecutionsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Associationid interface{} `json:"AssociationId"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeregisterPatchBaselineForPatchGroupResult represents the DeregisterPatchBaselineForPatchGroupResult schema from the OpenAPI specification
type DeregisterPatchBaselineForPatchGroupResult struct {
	Patchgroup interface{} `json:"PatchGroup,omitempty"`
	Baselineid interface{} `json:"BaselineId,omitempty"`
}

// GetMaintenanceWindowResult represents the GetMaintenanceWindowResult schema from the OpenAPI specification
type GetMaintenanceWindowResult struct {
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Scheduletimezone interface{} `json:"ScheduleTimezone,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Nextexecutiontime interface{} `json:"NextExecutionTime,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Modifieddate interface{} `json:"ModifiedDate,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Allowunassociatedtargets interface{} `json:"AllowUnassociatedTargets,omitempty"`
	Cutoff interface{} `json:"Cutoff,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
}

// DescribeMaintenanceWindowTargetsRequest represents the DescribeMaintenanceWindowTargetsRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowTargetsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowid interface{} `json:"WindowId"`
}

// DescribeInstancePatchStatesForPatchGroupRequest represents the DescribeInstancePatchStatesForPatchGroupRequest schema from the OpenAPI specification
type DescribeInstancePatchStatesForPatchGroupRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Patchgroup interface{} `json:"PatchGroup"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// NotificationConfig represents the NotificationConfig schema from the OpenAPI specification
type NotificationConfig struct {
	Notificationarn interface{} `json:"NotificationArn,omitempty"`
	Notificationevents interface{} `json:"NotificationEvents,omitempty"`
	Notificationtype interface{} `json:"NotificationType,omitempty"`
}

// ListOpsItemRelatedItemsRequest represents the ListOpsItemRelatedItemsRequest schema from the OpenAPI specification
type ListOpsItemRelatedItemsRequest struct {
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SendCommandRequest represents the SendCommandRequest schema from the OpenAPI specification
type SendCommandRequest struct {
	Documenthashtype interface{} `json:"DocumentHashType,omitempty"`
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Documenthash interface{} `json:"DocumentHash,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Outputs3bucketname interface{} `json:"OutputS3BucketName,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Documentname interface{} `json:"DocumentName"`
	Notificationconfig interface{} `json:"NotificationConfig,omitempty"`
	Outputs3keyprefix interface{} `json:"OutputS3KeyPrefix,omitempty"`
	Cloudwatchoutputconfig interface{} `json:"CloudWatchOutputConfig,omitempty"`
	Outputs3region interface{} `json:"OutputS3Region,omitempty"`
	Timeoutseconds interface{} `json:"TimeoutSeconds,omitempty"`
}

// S3OutputUrl represents the S3OutputUrl schema from the OpenAPI specification
type S3OutputUrl struct {
	Outputurl interface{} `json:"OutputUrl,omitempty"`
}

// DescribeInstanceAssociationsStatusRequest represents the DescribeInstanceAssociationsStatusRequest schema from the OpenAPI specification
type DescribeInstanceAssociationsStatusRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ResultAttribute represents the ResultAttribute schema from the OpenAPI specification
type ResultAttribute struct {
	Typename interface{} `json:"TypeName"`
}

// ResourceComplianceSummaryItem represents the ResourceComplianceSummaryItem schema from the OpenAPI specification
type ResourceComplianceSummaryItem struct {
	Compliancetype interface{} `json:"ComplianceType,omitempty"`
	Compliantsummary interface{} `json:"CompliantSummary,omitempty"`
	Executionsummary interface{} `json:"ExecutionSummary,omitempty"`
	Noncompliantsummary interface{} `json:"NonCompliantSummary,omitempty"`
	Overallseverity interface{} `json:"OverallSeverity,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// PutComplianceItemsRequest represents the PutComplianceItemsRequest schema from the OpenAPI specification
type PutComplianceItemsRequest struct {
	Executionsummary interface{} `json:"ExecutionSummary"`
	Itemcontenthash interface{} `json:"ItemContentHash,omitempty"`
	Items interface{} `json:"Items"`
	Resourceid interface{} `json:"ResourceId"`
	Resourcetype interface{} `json:"ResourceType"`
	Uploadtype interface{} `json:"UploadType,omitempty"`
	Compliancetype interface{} `json:"ComplianceType"`
}

// PatchGroupPatchBaselineMapping represents the PatchGroupPatchBaselineMapping schema from the OpenAPI specification
type PatchGroupPatchBaselineMapping struct {
	Baselineidentity interface{} `json:"BaselineIdentity,omitempty"`
	Patchgroup interface{} `json:"PatchGroup,omitempty"`
}

// DescribeMaintenanceWindowExecutionTasksResult represents the DescribeMaintenanceWindowExecutionTasksResult schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionTasksResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowexecutiontaskidentities interface{} `json:"WindowExecutionTaskIdentities,omitempty"`
}

// ListCommandsResult represents the ListCommandsResult schema from the OpenAPI specification
type ListCommandsResult struct {
	Commands interface{} `json:"Commands,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListResourceDataSyncResult represents the ListResourceDataSyncResult schema from the OpenAPI specification
type ListResourceDataSyncResult struct {
	Resourcedatasyncitems interface{} `json:"ResourceDataSyncItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetDeployablePatchSnapshotForInstanceResult represents the GetDeployablePatchSnapshotForInstanceResult schema from the OpenAPI specification
type GetDeployablePatchSnapshotForInstanceResult struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Product interface{} `json:"Product,omitempty"`
	Snapshotdownloadurl interface{} `json:"SnapshotDownloadUrl,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
}

// MaintenanceWindowTaskInvocationParameters represents the MaintenanceWindowTaskInvocationParameters schema from the OpenAPI specification
type MaintenanceWindowTaskInvocationParameters struct {
	Lambda interface{} `json:"Lambda,omitempty"`
	Runcommand interface{} `json:"RunCommand,omitempty"`
	Stepfunctions interface{} `json:"StepFunctions,omitempty"`
	Automation interface{} `json:"Automation,omitempty"`
}

// DescribeDocumentPermissionResponse represents the DescribeDocumentPermissionResponse schema from the OpenAPI specification
type DescribeDocumentPermissionResponse struct {
	Accountids interface{} `json:"AccountIds,omitempty"`
	Accountsharinginfolist interface{} `json:"AccountSharingInfoList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InstancePatchStateFilter represents the InstancePatchStateFilter schema from the OpenAPI specification
type InstancePatchStateFilter struct {
	TypeField interface{} `json:"Type"`
	Values interface{} `json:"Values"`
	Key interface{} `json:"Key"`
}

// GetOpsMetadataResult represents the GetOpsMetadataResult schema from the OpenAPI specification
type GetOpsMetadataResult struct {
	Metadata interface{} `json:"Metadata,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
}

// UpdateDocumentMetadataRequest represents the UpdateDocumentMetadataRequest schema from the OpenAPI specification
type UpdateDocumentMetadataRequest struct {
	Documentreviews interface{} `json:"DocumentReviews"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name"`
}

// ComplianceExecutionSummary represents the ComplianceExecutionSummary schema from the OpenAPI specification
type ComplianceExecutionSummary struct {
	Executiontype interface{} `json:"ExecutionType,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Executiontime interface{} `json:"ExecutionTime"`
}

// ResourceDataSyncSource represents the ResourceDataSyncSource schema from the OpenAPI specification
type ResourceDataSyncSource struct {
	Includefutureregions interface{} `json:"IncludeFutureRegions,omitempty"`
	Sourceregions interface{} `json:"SourceRegions"`
	Sourcetype interface{} `json:"SourceType"`
	Awsorganizationssource interface{} `json:"AwsOrganizationsSource,omitempty"`
	Enableallopsdatasources interface{} `json:"EnableAllOpsDataSources,omitempty"`
}

// DescribePatchPropertiesRequest represents the DescribePatchPropertiesRequest schema from the OpenAPI specification
type DescribePatchPropertiesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem"`
	Patchset interface{} `json:"PatchSet,omitempty"`
	Property interface{} `json:"Property"`
}

// UpdateManagedInstanceRoleResult represents the UpdateManagedInstanceRoleResult schema from the OpenAPI specification
type UpdateManagedInstanceRoleResult struct {
}

// InstancePatchState represents the InstancePatchState schema from the OpenAPI specification
type InstancePatchState struct {
	Patchgroup interface{} `json:"PatchGroup"`
	Installedcount interface{} `json:"InstalledCount,omitempty"`
	Operationstarttime interface{} `json:"OperationStartTime"`
	Securitynoncompliantcount interface{} `json:"SecurityNonCompliantCount,omitempty"`
	Installoverridelist interface{} `json:"InstallOverrideList,omitempty"`
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Rebootoption interface{} `json:"RebootOption,omitempty"`
	Snapshotid interface{} `json:"SnapshotId,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
	Unreportednotapplicablecount interface{} `json:"UnreportedNotApplicableCount,omitempty"`
	Operationendtime interface{} `json:"OperationEndTime"`
	Missingcount interface{} `json:"MissingCount,omitempty"`
	Operation interface{} `json:"Operation"`
	Failedcount interface{} `json:"FailedCount,omitempty"`
	Notapplicablecount interface{} `json:"NotApplicableCount,omitempty"`
	Installedpendingrebootcount interface{} `json:"InstalledPendingRebootCount,omitempty"`
	Baselineid interface{} `json:"BaselineId"`
	Criticalnoncompliantcount interface{} `json:"CriticalNonCompliantCount,omitempty"`
	Installedrejectedcount interface{} `json:"InstalledRejectedCount,omitempty"`
	Lastnorebootinstalloperationtime interface{} `json:"LastNoRebootInstallOperationTime,omitempty"`
	Installedothercount interface{} `json:"InstalledOtherCount,omitempty"`
	Othernoncompliantcount interface{} `json:"OtherNonCompliantCount,omitempty"`
}

// InventoryItem represents the InventoryItem schema from the OpenAPI specification
type InventoryItem struct {
	Contenthash interface{} `json:"ContentHash,omitempty"`
	Context interface{} `json:"Context,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion"`
	Typename interface{} `json:"TypeName"`
	Capturetime interface{} `json:"CaptureTime"`
	Content interface{} `json:"Content,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// AssociationVersionInfo represents the AssociationVersionInfo schema from the OpenAPI specification
type AssociationVersionInfo struct {
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Complianceseverity interface{} `json:"ComplianceSeverity,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Synccompliance interface{} `json:"SyncCompliance,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Applyonlyatcroninterval interface{} `json:"ApplyOnlyAtCronInterval,omitempty"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Calendarnames interface{} `json:"CalendarNames,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
}

// DeleteResourceDataSyncRequest represents the DeleteResourceDataSyncRequest schema from the OpenAPI specification
type DeleteResourceDataSyncRequest struct {
	Syncname interface{} `json:"SyncName"`
	Synctype interface{} `json:"SyncType,omitempty"`
}

// Patch represents the Patch schema from the OpenAPI specification
type Patch struct {
	Name interface{} `json:"Name,omitempty"`
	Cveids interface{} `json:"CVEIds,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Epoch interface{} `json:"Epoch,omitempty"`
	Product interface{} `json:"Product,omitempty"`
	Language interface{} `json:"Language,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Msrcseverity interface{} `json:"MsrcSeverity,omitempty"`
	Vendor interface{} `json:"Vendor,omitempty"`
	Arch interface{} `json:"Arch,omitempty"`
	Contenturl interface{} `json:"ContentUrl,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Releasedate interface{} `json:"ReleaseDate,omitempty"`
	Kbnumber interface{} `json:"KbNumber,omitempty"`
	Repository interface{} `json:"Repository,omitempty"`
	Advisoryids interface{} `json:"AdvisoryIds,omitempty"`
	Bugzillaids interface{} `json:"BugzillaIds,omitempty"`
	Productfamily interface{} `json:"ProductFamily,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Classification interface{} `json:"Classification,omitempty"`
	Msrcnumber interface{} `json:"MsrcNumber,omitempty"`
	Release interface{} `json:"Release,omitempty"`
}

// OpsEntityItemMap represents the OpsEntityItemMap schema from the OpenAPI specification
type OpsEntityItemMap struct {
}

// Runbook represents the Runbook schema from the OpenAPI specification
type Runbook struct {
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Targetparametername interface{} `json:"TargetParameterName,omitempty"`
	Documentname interface{} `json:"DocumentName"`
}

// AssociateOpsItemRelatedItemResponse represents the AssociateOpsItemRelatedItemResponse schema from the OpenAPI specification
type AssociateOpsItemRelatedItemResponse struct {
	Associationid interface{} `json:"AssociationId,omitempty"`
}

// UpdateDocumentResult represents the UpdateDocumentResult schema from the OpenAPI specification
type UpdateDocumentResult struct {
	Documentdescription interface{} `json:"DocumentDescription,omitempty"`
}

// PatchSource represents the PatchSource schema from the OpenAPI specification
type PatchSource struct {
	Configuration interface{} `json:"Configuration"`
	Name interface{} `json:"Name"`
	Products interface{} `json:"Products"`
}

// DocumentVersionInfo represents the DocumentVersionInfo schema from the OpenAPI specification
type DocumentVersionInfo struct {
	Displayname interface{} `json:"DisplayName,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Reviewstatus interface{} `json:"ReviewStatus,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Isdefaultversion interface{} `json:"IsDefaultVersion,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Statusinformation interface{} `json:"StatusInformation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DescribeMaintenanceWindowScheduleRequest represents the DescribeMaintenanceWindowScheduleRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowScheduleRequest struct {
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetInventoryResult represents the GetInventoryResult schema from the OpenAPI specification
type GetInventoryResult struct {
	Entities interface{} `json:"Entities,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteResourceDataSyncResult represents the DeleteResourceDataSyncResult schema from the OpenAPI specification
type DeleteResourceDataSyncResult struct {
}

// UpdateManagedInstanceRoleRequest represents the UpdateManagedInstanceRoleRequest schema from the OpenAPI specification
type UpdateManagedInstanceRoleRequest struct {
	Iamrole interface{} `json:"IamRole"`
	Instanceid interface{} `json:"InstanceId"`
}

// Parameters represents the Parameters schema from the OpenAPI specification
type Parameters struct {
}

// PatchPropertyEntry represents the PatchPropertyEntry schema from the OpenAPI specification
type PatchPropertyEntry struct {
}

// UpdateOpsMetadataResult represents the UpdateOpsMetadataResult schema from the OpenAPI specification
type UpdateOpsMetadataResult struct {
	Opsmetadataarn interface{} `json:"OpsMetadataArn,omitempty"`
}

// InventoryItemSchema represents the InventoryItemSchema schema from the OpenAPI specification
type InventoryItemSchema struct {
	Version interface{} `json:"Version,omitempty"`
	Attributes interface{} `json:"Attributes"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Typename interface{} `json:"TypeName"`
}

// DescribeAssociationExecutionTargetsResult represents the DescribeAssociationExecutionTargetsResult schema from the OpenAPI specification
type DescribeAssociationExecutionTargetsResult struct {
	Associationexecutiontargets interface{} `json:"AssociationExecutionTargets,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SessionManagerOutputUrl represents the SessionManagerOutputUrl schema from the OpenAPI specification
type SessionManagerOutputUrl struct {
	S3outputurl interface{} `json:"S3OutputUrl,omitempty"`
	Cloudwatchoutputurl interface{} `json:"CloudWatchOutputUrl,omitempty"`
}

// GetResourcePoliciesRequest represents the GetResourcePoliciesRequest schema from the OpenAPI specification
type GetResourcePoliciesRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// ResumeSessionResponse represents the ResumeSessionResponse schema from the OpenAPI specification
type ResumeSessionResponse struct {
	Tokenvalue interface{} `json:"TokenValue,omitempty"`
	Sessionid interface{} `json:"SessionId,omitempty"`
	Streamurl interface{} `json:"StreamUrl,omitempty"`
}

// ParameterHistory represents the ParameterHistory schema from the OpenAPI specification
type ParameterHistory struct {
	Datatype interface{} `json:"DataType,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Labels interface{} `json:"Labels,omitempty"`
	Tier interface{} `json:"Tier,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Allowedpattern interface{} `json:"AllowedPattern,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Lastmodifieddate interface{} `json:"LastModifiedDate,omitempty"`
	Lastmodifieduser interface{} `json:"LastModifiedUser,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// InventoryResultItem represents the InventoryResultItem schema from the OpenAPI specification
type InventoryResultItem struct {
	Capturetime interface{} `json:"CaptureTime,omitempty"`
	Content interface{} `json:"Content"`
	Contenthash interface{} `json:"ContentHash,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion"`
	Typename interface{} `json:"TypeName"`
}

// ListOpsItemRelatedItemsResponse represents the ListOpsItemRelatedItemsResponse schema from the OpenAPI specification
type ListOpsItemRelatedItemsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Summaries interface{} `json:"Summaries,omitempty"`
}

// InstanceAssociationStatusInfo represents the InstanceAssociationStatusInfo schema from the OpenAPI specification
type InstanceAssociationStatusInfo struct {
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Executiondate interface{} `json:"ExecutionDate,omitempty"`
	Executionsummary interface{} `json:"ExecutionSummary,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Outputurl interface{} `json:"OutputUrl,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Detailedstatus interface{} `json:"DetailedStatus,omitempty"`
}

// CreateActivationResult represents the CreateActivationResult schema from the OpenAPI specification
type CreateActivationResult struct {
	Activationcode interface{} `json:"ActivationCode,omitempty"`
	Activationid interface{} `json:"ActivationId,omitempty"`
}

// DescribeAssociationResult represents the DescribeAssociationResult schema from the OpenAPI specification
type DescribeAssociationResult struct {
	Associationdescription interface{} `json:"AssociationDescription,omitempty"`
}

// DeleteMaintenanceWindowResult represents the DeleteMaintenanceWindowResult schema from the OpenAPI specification
type DeleteMaintenanceWindowResult struct {
	Windowid interface{} `json:"WindowId,omitempty"`
}

// RemoveTagsFromResourceResult represents the RemoveTagsFromResourceResult schema from the OpenAPI specification
type RemoveTagsFromResourceResult struct {
}

// CreateMaintenanceWindowResult represents the CreateMaintenanceWindowResult schema from the OpenAPI specification
type CreateMaintenanceWindowResult struct {
	Windowid interface{} `json:"WindowId,omitempty"`
}

// MaintenanceWindowFilter represents the MaintenanceWindowFilter schema from the OpenAPI specification
type MaintenanceWindowFilter struct {
	Key interface{} `json:"Key,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// ListAssociationsRequest represents the ListAssociationsRequest schema from the OpenAPI specification
type ListAssociationsRequest struct {
	Associationfilterlist interface{} `json:"AssociationFilterList,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InstanceAssociation represents the InstanceAssociation schema from the OpenAPI specification
type InstanceAssociation struct {
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Content interface{} `json:"Content,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
}

// ListCommandInvocationsRequest represents the ListCommandInvocationsRequest schema from the OpenAPI specification
type ListCommandInvocationsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Details interface{} `json:"Details,omitempty"`
}

// RegisterDefaultPatchBaselineRequest represents the RegisterDefaultPatchBaselineRequest schema from the OpenAPI specification
type RegisterDefaultPatchBaselineRequest struct {
	Baselineid interface{} `json:"BaselineId"`
}

// RegistrationMetadataItem represents the RegistrationMetadataItem schema from the OpenAPI specification
type RegistrationMetadataItem struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// PatchStatus represents the PatchStatus schema from the OpenAPI specification
type PatchStatus struct {
	Approvaldate interface{} `json:"ApprovalDate,omitempty"`
	Compliancelevel interface{} `json:"ComplianceLevel,omitempty"`
	Deploymentstatus interface{} `json:"DeploymentStatus,omitempty"`
}

// DeleteAssociationRequest represents the DeleteAssociationRequest schema from the OpenAPI specification
type DeleteAssociationRequest struct {
	Name interface{} `json:"Name,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
}

// AttachmentsSource represents the AttachmentsSource schema from the OpenAPI specification
type AttachmentsSource struct {
	Key interface{} `json:"Key,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// UpdateMaintenanceWindowResult represents the UpdateMaintenanceWindowResult schema from the OpenAPI specification
type UpdateMaintenanceWindowResult struct {
	Name interface{} `json:"Name,omitempty"`
	Scheduletimezone interface{} `json:"ScheduleTimezone,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Cutoff interface{} `json:"Cutoff,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Allowunassociatedtargets interface{} `json:"AllowUnassociatedTargets,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
}

// DescribeMaintenanceWindowExecutionsResult represents the DescribeMaintenanceWindowExecutionsResult schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionsResult struct {
	Windowexecutions interface{} `json:"WindowExecutions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MaintenanceWindowStepFunctionsParameters represents the MaintenanceWindowStepFunctionsParameters schema from the OpenAPI specification
type MaintenanceWindowStepFunctionsParameters struct {
	Input interface{} `json:"Input,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// Association represents the Association schema from the OpenAPI specification
type Association struct {
	Associationid interface{} `json:"AssociationId,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Lastexecutiondate interface{} `json:"LastExecutionDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Overview interface{} `json:"Overview,omitempty"`
}

// DeleteParameterRequest represents the DeleteParameterRequest schema from the OpenAPI specification
type DeleteParameterRequest struct {
	Name interface{} `json:"Name"`
}

// DescribePatchGroupStateResult represents the DescribePatchGroupStateResult schema from the OpenAPI specification
type DescribePatchGroupStateResult struct {
	Instanceswithmissingpatches interface{} `json:"InstancesWithMissingPatches,omitempty"`
	Instanceswithnotapplicablepatches interface{} `json:"InstancesWithNotApplicablePatches,omitempty"`
	Instanceswithunreportednotapplicablepatches interface{} `json:"InstancesWithUnreportedNotApplicablePatches,omitempty"`
	Instances interface{} `json:"Instances,omitempty"`
	Instanceswithcriticalnoncompliantpatches interface{} `json:"InstancesWithCriticalNonCompliantPatches,omitempty"`
	Instanceswithinstalledpendingrebootpatches interface{} `json:"InstancesWithInstalledPendingRebootPatches,omitempty"`
	Instanceswithinstalledrejectedpatches interface{} `json:"InstancesWithInstalledRejectedPatches,omitempty"`
	Instanceswithinstalledpatches interface{} `json:"InstancesWithInstalledPatches,omitempty"`
	Instanceswithothernoncompliantpatches interface{} `json:"InstancesWithOtherNonCompliantPatches,omitempty"`
	Instanceswithsecuritynoncompliantpatches interface{} `json:"InstancesWithSecurityNonCompliantPatches,omitempty"`
	Instanceswithfailedpatches interface{} `json:"InstancesWithFailedPatches,omitempty"`
	Instanceswithinstalledotherpatches interface{} `json:"InstancesWithInstalledOtherPatches,omitempty"`
}

// LabelParameterVersionRequest represents the LabelParameterVersionRequest schema from the OpenAPI specification
type LabelParameterVersionRequest struct {
	Parameterversion interface{} `json:"ParameterVersion,omitempty"`
	Labels interface{} `json:"Labels"`
	Name interface{} `json:"Name"`
}

// DeregisterManagedInstanceResult represents the DeregisterManagedInstanceResult schema from the OpenAPI specification
type DeregisterManagedInstanceResult struct {
}

// GetCalendarStateResponse represents the GetCalendarStateResponse schema from the OpenAPI specification
type GetCalendarStateResponse struct {
	Attime interface{} `json:"AtTime,omitempty"`
	Nexttransitiontime interface{} `json:"NextTransitionTime,omitempty"`
	State interface{} `json:"State,omitempty"`
}

// DeletePatchBaselineResult represents the DeletePatchBaselineResult schema from the OpenAPI specification
type DeletePatchBaselineResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
}

// OpsItemEventSummary represents the OpsItemEventSummary schema from the OpenAPI specification
type OpsItemEventSummary struct {
	Source interface{} `json:"Source,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Detail interface{} `json:"Detail,omitempty"`
	Detailtype interface{} `json:"DetailType,omitempty"`
	Eventid interface{} `json:"EventId,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
}

// ListDocumentsResult represents the ListDocumentsResult schema from the OpenAPI specification
type ListDocumentsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Documentidentifiers interface{} `json:"DocumentIdentifiers,omitempty"`
}

// DescribeActivationsFilter represents the DescribeActivationsFilter schema from the OpenAPI specification
type DescribeActivationsFilter struct {
	Filterkey interface{} `json:"FilterKey,omitempty"`
	Filtervalues interface{} `json:"FilterValues,omitempty"`
}

// ParameterInlinePolicy represents the ParameterInlinePolicy schema from the OpenAPI specification
type ParameterInlinePolicy struct {
	Policystatus interface{} `json:"PolicyStatus,omitempty"`
	Policytext interface{} `json:"PolicyText,omitempty"`
	Policytype interface{} `json:"PolicyType,omitempty"`
}

// AttachmentContent represents the AttachmentContent schema from the OpenAPI specification
type AttachmentContent struct {
	Hash interface{} `json:"Hash,omitempty"`
	Hashtype interface{} `json:"HashType,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Size interface{} `json:"Size,omitempty"`
	Url interface{} `json:"Url,omitempty"`
}

// ListOpsMetadataRequest represents the ListOpsMetadataRequest schema from the OpenAPI specification
type ListOpsMetadataRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateActivationRequest represents the CreateActivationRequest schema from the OpenAPI specification
type CreateActivationRequest struct {
	Defaultinstancename interface{} `json:"DefaultInstanceName,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Expirationdate interface{} `json:"ExpirationDate,omitempty"`
	Iamrole interface{} `json:"IamRole"`
	Registrationlimit interface{} `json:"RegistrationLimit,omitempty"`
	Registrationmetadata interface{} `json:"RegistrationMetadata,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// OpsItem represents the OpsItem schema from the OpenAPI specification
type OpsItem struct {
	Notifications interface{} `json:"Notifications,omitempty"`
	Plannedendtime interface{} `json:"PlannedEndTime,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Plannedstarttime interface{} `json:"PlannedStartTime,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Operationaldata interface{} `json:"OperationalData,omitempty"`
	Opsitemtype interface{} `json:"OpsItemType,omitempty"`
	Actualendtime interface{} `json:"ActualEndTime,omitempty"`
	Opsitemarn interface{} `json:"OpsItemArn,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Actualstarttime interface{} `json:"ActualStartTime,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Relatedopsitems interface{} `json:"RelatedOpsItems,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
}

// AutomationParameterMap represents the AutomationParameterMap schema from the OpenAPI specification
type AutomationParameterMap struct {
}

// DeregisterPatchBaselineForPatchGroupRequest represents the DeregisterPatchBaselineForPatchGroupRequest schema from the OpenAPI specification
type DeregisterPatchBaselineForPatchGroupRequest struct {
	Baselineid interface{} `json:"BaselineId"`
	Patchgroup interface{} `json:"PatchGroup"`
}

// StartAutomationExecutionRequest represents the StartAutomationExecutionRequest schema from the OpenAPI specification
type StartAutomationExecutionRequest struct {
	Mode interface{} `json:"Mode,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Documentname interface{} `json:"DocumentName"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Targetparametername interface{} `json:"TargetParameterName,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
}

// CreatePatchBaselineRequest represents the CreatePatchBaselineRequest schema from the OpenAPI specification
type CreatePatchBaselineRequest struct {
	Approvalrules interface{} `json:"ApprovalRules,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Globalfilters interface{} `json:"GlobalFilters,omitempty"`
	Name interface{} `json:"Name"`
	Rejectedpatches interface{} `json:"RejectedPatches,omitempty"`
	Approvedpatches interface{} `json:"ApprovedPatches,omitempty"`
	Approvedpatchescompliancelevel interface{} `json:"ApprovedPatchesComplianceLevel,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Rejectedpatchesaction interface{} `json:"RejectedPatchesAction,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Approvedpatchesenablenonsecurity interface{} `json:"ApprovedPatchesEnableNonSecurity,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Sources interface{} `json:"Sources,omitempty"`
}

// GetResourcePoliciesResponseEntry represents the GetResourcePoliciesResponseEntry schema from the OpenAPI specification
type GetResourcePoliciesResponseEntry struct {
	Policy interface{} `json:"Policy,omitempty"`
	Policyhash interface{} `json:"PolicyHash,omitempty"`
	Policyid interface{} `json:"PolicyId,omitempty"`
}

// FailedCreateAssociation represents the FailedCreateAssociation schema from the OpenAPI specification
type FailedCreateAssociation struct {
	Fault interface{} `json:"Fault,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Entry interface{} `json:"Entry,omitempty"`
}

// RegisterTargetWithMaintenanceWindowRequest represents the RegisterTargetWithMaintenanceWindowRequest schema from the OpenAPI specification
type RegisterTargetWithMaintenanceWindowRequest struct {
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Resourcetype interface{} `json:"ResourceType"`
	Targets interface{} `json:"Targets"`
	Windowid interface{} `json:"WindowId"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// DeleteResourcePolicyRequest represents the DeleteResourcePolicyRequest schema from the OpenAPI specification
type DeleteResourcePolicyRequest struct {
	Policyhash interface{} `json:"PolicyHash"`
	Policyid interface{} `json:"PolicyId"`
	Resourcearn interface{} `json:"ResourceArn"`
}

// InstanceInformationFilter represents the InstanceInformationFilter schema from the OpenAPI specification
type InstanceInformationFilter struct {
	Key interface{} `json:"key"`
	Valueset interface{} `json:"valueSet"`
}

// DescribeInstanceAssociationsStatusResult represents the DescribeInstanceAssociationsStatusResult schema from the OpenAPI specification
type DescribeInstanceAssociationsStatusResult struct {
	Instanceassociationstatusinfos interface{} `json:"InstanceAssociationStatusInfos,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// OpsItemRelatedItemSummary represents the OpsItemRelatedItemSummary schema from the OpenAPI specification
type OpsItemRelatedItemSummary struct {
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Lastmodifiedby OpsItemIdentity `json:"LastModifiedBy,omitempty"` // Information about the user or resource that created an OpsItem event.
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Resourceuri interface{} `json:"ResourceUri,omitempty"`
	Createdby OpsItemIdentity `json:"CreatedBy,omitempty"` // Information about the user or resource that created an OpsItem event.
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Associationtype interface{} `json:"AssociationType,omitempty"`
}

// DeleteMaintenanceWindowRequest represents the DeleteMaintenanceWindowRequest schema from the OpenAPI specification
type DeleteMaintenanceWindowRequest struct {
	Windowid interface{} `json:"WindowId"`
}

// MaintenanceWindowExecution represents the MaintenanceWindowExecution schema from the OpenAPI specification
type MaintenanceWindowExecution struct {
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
}

// OpsMetadata represents the OpsMetadata schema from the OpenAPI specification
type OpsMetadata struct {
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Creationdate interface{} `json:"CreationDate,omitempty"`
	Lastmodifieddate interface{} `json:"LastModifiedDate,omitempty"`
	Lastmodifieduser interface{} `json:"LastModifiedUser,omitempty"`
	Opsmetadataarn interface{} `json:"OpsMetadataArn,omitempty"`
}

// ComplianceStringFilter represents the ComplianceStringFilter schema from the OpenAPI specification
type ComplianceStringFilter struct {
	Key interface{} `json:"Key,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// DeleteDocumentRequest represents the DeleteDocumentRequest schema from the OpenAPI specification
type DeleteDocumentRequest struct {
	Versionname interface{} `json:"VersionName,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Force interface{} `json:"Force,omitempty"`
	Name interface{} `json:"Name"`
}

// DescribePatchGroupsResult represents the DescribePatchGroupsResult schema from the OpenAPI specification
type DescribePatchGroupsResult struct {
	Mappings interface{} `json:"Mappings,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// OutputSource represents the OutputSource schema from the OpenAPI specification
type OutputSource struct {
	Outputsourceid interface{} `json:"OutputSourceId,omitempty"`
	Outputsourcetype interface{} `json:"OutputSourceType,omitempty"`
}

// CommandFilter represents the CommandFilter schema from the OpenAPI specification
type CommandFilter struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// ListComplianceSummariesRequest represents the ListComplianceSummariesRequest schema from the OpenAPI specification
type ListComplianceSummariesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeleteOpsMetadataResult represents the DeleteOpsMetadataResult schema from the OpenAPI specification
type DeleteOpsMetadataResult struct {
}

// DeregisterManagedInstanceRequest represents the DeregisterManagedInstanceRequest schema from the OpenAPI specification
type DeregisterManagedInstanceRequest struct {
	Instanceid interface{} `json:"InstanceId"`
}

// GetOpsSummaryResult represents the GetOpsSummaryResult schema from the OpenAPI specification
type GetOpsSummaryResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Entities interface{} `json:"Entities,omitempty"`
}

// UpdateAssociationRequest represents the UpdateAssociationRequest schema from the OpenAPI specification
type UpdateAssociationRequest struct {
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Associationid interface{} `json:"AssociationId"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Applyonlyatcroninterval interface{} `json:"ApplyOnlyAtCronInterval,omitempty"`
	Complianceseverity interface{} `json:"ComplianceSeverity,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Calendarnames interface{} `json:"CalendarNames,omitempty"`
	Synccompliance interface{} `json:"SyncCompliance,omitempty"`
	Alarmconfiguration AlarmConfiguration `json:"AlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Automationtargetparametername interface{} `json:"AutomationTargetParameterName,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
}

// GetAutomationExecutionResult represents the GetAutomationExecutionResult schema from the OpenAPI specification
type GetAutomationExecutionResult struct {
	Automationexecution interface{} `json:"AutomationExecution,omitempty"`
}

// DescribeSessionsRequest represents the DescribeSessionsRequest schema from the OpenAPI specification
type DescribeSessionsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	State interface{} `json:"State"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeregisterTargetFromMaintenanceWindowResult represents the DeregisterTargetFromMaintenanceWindowResult schema from the OpenAPI specification
type DeregisterTargetFromMaintenanceWindowResult struct {
	Windowid interface{} `json:"WindowId,omitempty"`
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
}

// ProgressCounters represents the ProgressCounters schema from the OpenAPI specification
type ProgressCounters struct {
	Cancelledsteps interface{} `json:"CancelledSteps,omitempty"`
	Failedsteps interface{} `json:"FailedSteps,omitempty"`
	Successsteps interface{} `json:"SuccessSteps,omitempty"`
	Timedoutsteps interface{} `json:"TimedOutSteps,omitempty"`
	Totalsteps interface{} `json:"TotalSteps,omitempty"`
}

// CreatePatchBaselineResult represents the CreatePatchBaselineResult schema from the OpenAPI specification
type CreatePatchBaselineResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
}

// DescribeInstancePatchStatesResult represents the DescribeInstancePatchStatesResult schema from the OpenAPI specification
type DescribeInstancePatchStatesResult struct {
	Instancepatchstates interface{} `json:"InstancePatchStates,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InventoryDeletionSummary represents the InventoryDeletionSummary schema from the OpenAPI specification
type InventoryDeletionSummary struct {
	Remainingcount interface{} `json:"RemainingCount,omitempty"`
	Summaryitems interface{} `json:"SummaryItems,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// DescribeEffectivePatchesForPatchBaselineRequest represents the DescribeEffectivePatchesForPatchBaselineRequest schema from the OpenAPI specification
type DescribeEffectivePatchesForPatchBaselineRequest struct {
	Baselineid interface{} `json:"BaselineId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetMaintenanceWindowExecutionTaskResult represents the GetMaintenanceWindowExecutionTaskResult schema from the OpenAPI specification
type GetMaintenanceWindowExecutionTaskResult struct {
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
	Taskexecutionid interface{} `json:"TaskExecutionId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
}

// ComplianceItem represents the ComplianceItem schema from the OpenAPI specification
type ComplianceItem struct {
	Details interface{} `json:"Details,omitempty"`
	Compliancetype interface{} `json:"ComplianceType,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Executionsummary interface{} `json:"ExecutionSummary,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// OpsItemSummary represents the OpsItemSummary schema from the OpenAPI specification
type OpsItemSummary struct {
	Actualendtime interface{} `json:"ActualEndTime,omitempty"`
	Operationaldata interface{} `json:"OperationalData,omitempty"`
	Opsitemtype interface{} `json:"OpsItemType,omitempty"`
	Source interface{} `json:"Source,omitempty"`
	Lastmodifiedtime interface{} `json:"LastModifiedTime,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Lastmodifiedby interface{} `json:"LastModifiedBy,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Plannedstarttime interface{} `json:"PlannedStartTime,omitempty"`
	Createdby interface{} `json:"CreatedBy,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Actualstarttime interface{} `json:"ActualStartTime,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Plannedendtime interface{} `json:"PlannedEndTime,omitempty"`
}

// InventoryResultItemMap represents the InventoryResultItemMap schema from the OpenAPI specification
type InventoryResultItemMap struct {
}

// ListOpsMetadataResult represents the ListOpsMetadataResult schema from the OpenAPI specification
type ListOpsMetadataResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Opsmetadatalist interface{} `json:"OpsMetadataList,omitempty"`
}

// UpdateAssociationResult represents the UpdateAssociationResult schema from the OpenAPI specification
type UpdateAssociationResult struct {
	Associationdescription interface{} `json:"AssociationDescription,omitempty"`
}

// ResourceDataSyncDestinationDataSharing represents the ResourceDataSyncDestinationDataSharing schema from the OpenAPI specification
type ResourceDataSyncDestinationDataSharing struct {
	Destinationdatasharingtype interface{} `json:"DestinationDataSharingType,omitempty"`
}

// DescribeInstancePatchesResult represents the DescribeInstancePatchesResult schema from the OpenAPI specification
type DescribeInstancePatchesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Patches interface{} `json:"Patches,omitempty"`
}

// ComplianceSummaryItem represents the ComplianceSummaryItem schema from the OpenAPI specification
type ComplianceSummaryItem struct {
	Compliancetype interface{} `json:"ComplianceType,omitempty"`
	Compliantsummary interface{} `json:"CompliantSummary,omitempty"`
	Noncompliantsummary interface{} `json:"NonCompliantSummary,omitempty"`
}

// OpsItemIdentity represents the OpsItemIdentity schema from the OpenAPI specification
type OpsItemIdentity struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// SessionManagerParameters represents the SessionManagerParameters schema from the OpenAPI specification
type SessionManagerParameters struct {
}

// UpdatePatchBaselineResult represents the UpdatePatchBaselineResult schema from the OpenAPI specification
type UpdatePatchBaselineResult struct {
	Rejectedpatches interface{} `json:"RejectedPatches,omitempty"`
	Rejectedpatchesaction interface{} `json:"RejectedPatchesAction,omitempty"`
	Sources interface{} `json:"Sources,omitempty"`
	Approvalrules interface{} `json:"ApprovalRules,omitempty"`
	Globalfilters interface{} `json:"GlobalFilters,omitempty"`
	Approvedpatches interface{} `json:"ApprovedPatches,omitempty"`
	Approvedpatchesenablenonsecurity interface{} `json:"ApprovedPatchesEnableNonSecurity,omitempty"`
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Approvedpatchescompliancelevel interface{} `json:"ApprovedPatchesComplianceLevel,omitempty"`
	Modifieddate interface{} `json:"ModifiedDate,omitempty"`
}

// CreateOpsItemResponse represents the CreateOpsItemResponse schema from the OpenAPI specification
type CreateOpsItemResponse struct {
	Opsitemarn interface{} `json:"OpsItemArn,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
}

// GetParameterHistoryRequest represents the GetParameterHistoryRequest schema from the OpenAPI specification
type GetParameterHistoryRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Name interface{} `json:"Name"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Withdecryption interface{} `json:"WithDecryption,omitempty"`
}

// ResetServiceSettingRequest represents the ResetServiceSettingRequest schema from the OpenAPI specification
type ResetServiceSettingRequest struct {
	Settingid interface{} `json:"SettingId"`
}

// DeleteAssociationResult represents the DeleteAssociationResult schema from the OpenAPI specification
type DeleteAssociationResult struct {
}

// DeleteParametersRequest represents the DeleteParametersRequest schema from the OpenAPI specification
type DeleteParametersRequest struct {
	Names interface{} `json:"Names"`
}

// DescribeMaintenanceWindowExecutionsRequest represents the DescribeMaintenanceWindowExecutionsRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowid interface{} `json:"WindowId"`
}

// PatchFilter represents the PatchFilter schema from the OpenAPI specification
type PatchFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// FailureDetails represents the FailureDetails schema from the OpenAPI specification
type FailureDetails struct {
	Details interface{} `json:"Details,omitempty"`
	Failurestage interface{} `json:"FailureStage,omitempty"`
	Failuretype interface{} `json:"FailureType,omitempty"`
}

// CreateDocumentResult represents the CreateDocumentResult schema from the OpenAPI specification
type CreateDocumentResult struct {
	Documentdescription interface{} `json:"DocumentDescription,omitempty"`
}

// OpsMetadataFilter represents the OpsMetadataFilter schema from the OpenAPI specification
type OpsMetadataFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// OpsItemDataValue represents the OpsItemDataValue schema from the OpenAPI specification
type OpsItemDataValue struct {
	TypeField interface{} `json:"Type,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// SendCommandResult represents the SendCommandResult schema from the OpenAPI specification
type SendCommandResult struct {
	Command interface{} `json:"Command,omitempty"`
}

// MaintenanceWindowExecutionTaskIdentity represents the MaintenanceWindowExecutionTaskIdentity schema from the OpenAPI specification
type MaintenanceWindowExecutionTaskIdentity struct {
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Taskexecutionid interface{} `json:"TaskExecutionId,omitempty"`
	Tasktype interface{} `json:"TaskType,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// ListComplianceSummariesResult represents the ListComplianceSummariesResult schema from the OpenAPI specification
type ListComplianceSummariesResult struct {
	Compliancesummaryitems interface{} `json:"ComplianceSummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetConnectionStatusResponse represents the GetConnectionStatusResponse schema from the OpenAPI specification
type GetConnectionStatusResponse struct {
	Status interface{} `json:"Status,omitempty"`
	Target interface{} `json:"Target,omitempty"`
}

// CommandInvocation represents the CommandInvocation schema from the OpenAPI specification
type CommandInvocation struct {
	Commandplugins interface{} `json:"CommandPlugins,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Traceoutput interface{} `json:"TraceOutput,omitempty"`
	Cloudwatchoutputconfig interface{} `json:"CloudWatchOutputConfig,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Notificationconfig interface{} `json:"NotificationConfig,omitempty"`
	Standarderrorurl interface{} `json:"StandardErrorUrl,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Requesteddatetime interface{} `json:"RequestedDateTime,omitempty"`
	Standardoutputurl interface{} `json:"StandardOutputUrl,omitempty"`
	Instancename interface{} `json:"InstanceName,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// AssociationOverview represents the AssociationOverview schema from the OpenAPI specification
type AssociationOverview struct {
	Associationstatusaggregatedcount interface{} `json:"AssociationStatusAggregatedCount,omitempty"`
	Detailedstatus interface{} `json:"DetailedStatus,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// GetParametersRequest represents the GetParametersRequest schema from the OpenAPI specification
type GetParametersRequest struct {
	Names interface{} `json:"Names"`
	Withdecryption interface{} `json:"WithDecryption,omitempty"`
}

// StartSessionRequest represents the StartSessionRequest schema from the OpenAPI specification
type StartSessionRequest struct {
	Target interface{} `json:"Target"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Reason interface{} `json:"Reason,omitempty"`
}

// DescribeMaintenanceWindowsRequest represents the DescribeMaintenanceWindowsRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StartChangeRequestExecutionResult represents the StartChangeRequestExecutionResult schema from the OpenAPI specification
type StartChangeRequestExecutionResult struct {
	Automationexecutionid interface{} `json:"AutomationExecutionId,omitempty"`
}

// DescribeAutomationStepExecutionsRequest represents the DescribeAutomationStepExecutionsRequest schema from the OpenAPI specification
type DescribeAutomationStepExecutionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Reverseorder interface{} `json:"ReverseOrder,omitempty"`
	Automationexecutionid interface{} `json:"AutomationExecutionId"`
	Filters interface{} `json:"Filters,omitempty"`
}

// OpsResultAttribute represents the OpsResultAttribute schema from the OpenAPI specification
type OpsResultAttribute struct {
	Typename interface{} `json:"TypeName"`
}

// TerminateSessionResponse represents the TerminateSessionResponse schema from the OpenAPI specification
type TerminateSessionResponse struct {
	Sessionid interface{} `json:"SessionId,omitempty"`
}

// DescribeMaintenanceWindowTasksRequest represents the DescribeMaintenanceWindowTasksRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowTasksRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowid interface{} `json:"WindowId"`
}

// UpdateResourceDataSyncRequest represents the UpdateResourceDataSyncRequest schema from the OpenAPI specification
type UpdateResourceDataSyncRequest struct {
	Syncname interface{} `json:"SyncName"`
	Syncsource interface{} `json:"SyncSource"`
	Synctype interface{} `json:"SyncType"`
}

// GetOpsSummaryRequest represents the GetOpsSummaryRequest schema from the OpenAPI specification
type GetOpsSummaryRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resultattributes interface{} `json:"ResultAttributes,omitempty"`
	Syncname interface{} `json:"SyncName,omitempty"`
	Aggregators interface{} `json:"Aggregators,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// MetadataValue represents the MetadataValue schema from the OpenAPI specification
type MetadataValue struct {
	Value interface{} `json:"Value,omitempty"`
}

// UpdateDocumentRequest represents the UpdateDocumentRequest schema from the OpenAPI specification
type UpdateDocumentRequest struct {
	Targettype interface{} `json:"TargetType,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Attachments interface{} `json:"Attachments,omitempty"`
	Content interface{} `json:"Content"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name"`
}

// DeregisterTaskFromMaintenanceWindowResult represents the DeregisterTaskFromMaintenanceWindowResult schema from the OpenAPI specification
type DeregisterTaskFromMaintenanceWindowResult struct {
	Windowid interface{} `json:"WindowId,omitempty"`
	Windowtaskid interface{} `json:"WindowTaskId,omitempty"`
}

// ReviewInformation represents the ReviewInformation schema from the OpenAPI specification
type ReviewInformation struct {
	Reviewer interface{} `json:"Reviewer,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Reviewedtime interface{} `json:"ReviewedTime,omitempty"`
}

// ListDocumentsRequest represents the ListDocumentsRequest schema from the OpenAPI specification
type ListDocumentsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Documentfilterlist interface{} `json:"DocumentFilterList,omitempty"`
}

// MetadataMap represents the MetadataMap schema from the OpenAPI specification
type MetadataMap struct {
}

// ParameterMetadata represents the ParameterMetadata schema from the OpenAPI specification
type ParameterMetadata struct {
	Allowedpattern interface{} `json:"AllowedPattern,omitempty"`
	Lastmodifieduser interface{} `json:"LastModifiedUser,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
	Tier interface{} `json:"Tier,omitempty"`
	Datatype interface{} `json:"DataType,omitempty"`
	Lastmodifieddate interface{} `json:"LastModifiedDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// RegisterTaskWithMaintenanceWindowResult represents the RegisterTaskWithMaintenanceWindowResult schema from the OpenAPI specification
type RegisterTaskWithMaintenanceWindowResult struct {
	Windowtaskid interface{} `json:"WindowTaskId,omitempty"`
}

// RegisterDefaultPatchBaselineResult represents the RegisterDefaultPatchBaselineResult schema from the OpenAPI specification
type RegisterDefaultPatchBaselineResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
}

// CreateResourceDataSyncRequest represents the CreateResourceDataSyncRequest schema from the OpenAPI specification
type CreateResourceDataSyncRequest struct {
	Syncname interface{} `json:"SyncName"`
	Syncsource interface{} `json:"SyncSource,omitempty"`
	Synctype interface{} `json:"SyncType,omitempty"`
	S3destination interface{} `json:"S3Destination,omitempty"`
}

// StopAutomationExecutionRequest represents the StopAutomationExecutionRequest schema from the OpenAPI specification
type StopAutomationExecutionRequest struct {
	TypeField interface{} `json:"Type,omitempty"`
	Automationexecutionid interface{} `json:"AutomationExecutionId"`
}

// UpdateOpsItemRequest represents the UpdateOpsItemRequest schema from the OpenAPI specification
type UpdateOpsItemRequest struct {
	Actualendtime interface{} `json:"ActualEndTime,omitempty"`
	Plannedstarttime interface{} `json:"PlannedStartTime,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Operationaldata interface{} `json:"OperationalData,omitempty"`
	Notifications interface{} `json:"Notifications,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Opsitemarn interface{} `json:"OpsItemArn,omitempty"`
	Relatedopsitems interface{} `json:"RelatedOpsItems,omitempty"`
	Actualstarttime interface{} `json:"ActualStartTime,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Opsitemid interface{} `json:"OpsItemId"`
	Description interface{} `json:"Description,omitempty"`
	Title interface{} `json:"Title,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Operationaldatatodelete interface{} `json:"OperationalDataToDelete,omitempty"`
	Plannedendtime interface{} `json:"PlannedEndTime,omitempty"`
}

// PatchRuleGroup represents the PatchRuleGroup schema from the OpenAPI specification
type PatchRuleGroup struct {
	Patchrules interface{} `json:"PatchRules"`
}

// CompliantSummary represents the CompliantSummary schema from the OpenAPI specification
type CompliantSummary struct {
	Severitysummary interface{} `json:"SeveritySummary,omitempty"`
	Compliantcount interface{} `json:"CompliantCount,omitempty"`
}

// DescribeAutomationStepExecutionsResult represents the DescribeAutomationStepExecutionsResult schema from the OpenAPI specification
type DescribeAutomationStepExecutionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Stepexecutions interface{} `json:"StepExecutions,omitempty"`
}

// GetDocumentResult represents the GetDocumentResult schema from the OpenAPI specification
type GetDocumentResult struct {
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusinformation interface{} `json:"StatusInformation,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Attachmentscontent interface{} `json:"AttachmentsContent,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Requires interface{} `json:"Requires,omitempty"`
	Content interface{} `json:"Content,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Documenttype interface{} `json:"DocumentType,omitempty"`
	Reviewstatus interface{} `json:"ReviewStatus,omitempty"`
}

// ResumeSessionRequest represents the ResumeSessionRequest schema from the OpenAPI specification
type ResumeSessionRequest struct {
	Sessionid interface{} `json:"SessionId"`
}

// ListResourceDataSyncRequest represents the ListResourceDataSyncRequest schema from the OpenAPI specification
type ListResourceDataSyncRequest struct {
	Synctype interface{} `json:"SyncType,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// AlarmStateInformation represents the AlarmStateInformation schema from the OpenAPI specification
type AlarmStateInformation struct {
	Name interface{} `json:"Name"`
	State interface{} `json:"State"`
}

// AssociationExecutionTarget represents the AssociationExecutionTarget schema from the OpenAPI specification
type AssociationExecutionTarget struct {
	Detailedstatus interface{} `json:"DetailedStatus,omitempty"`
	Lastexecutiondate interface{} `json:"LastExecutionDate,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Outputsource interface{} `json:"OutputSource,omitempty"`
	Resourceid interface{} `json:"ResourceId,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
}

// AssociateOpsItemRelatedItemRequest represents the AssociateOpsItemRelatedItemRequest schema from the OpenAPI specification
type AssociateOpsItemRelatedItemRequest struct {
	Resourceuri interface{} `json:"ResourceUri"`
	Associationtype interface{} `json:"AssociationType"`
	Opsitemid interface{} `json:"OpsItemId"`
	Resourcetype interface{} `json:"ResourceType"`
}

// DescribePatchGroupStateRequest represents the DescribePatchGroupStateRequest schema from the OpenAPI specification
type DescribePatchGroupStateRequest struct {
	Patchgroup interface{} `json:"PatchGroup"`
}

// ListResourceComplianceSummariesResult represents the ListResourceComplianceSummariesResult schema from the OpenAPI specification
type ListResourceComplianceSummariesResult struct {
	Resourcecompliancesummaryitems interface{} `json:"ResourceComplianceSummaryItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MaintenanceWindowTaskParameterValueExpression represents the MaintenanceWindowTaskParameterValueExpression schema from the OpenAPI specification
type MaintenanceWindowTaskParameterValueExpression struct {
	Values interface{} `json:"Values,omitempty"`
}

// ResourceDataSyncAwsOrganizationsSource represents the ResourceDataSyncAwsOrganizationsSource schema from the OpenAPI specification
type ResourceDataSyncAwsOrganizationsSource struct {
	Organizationalunits interface{} `json:"OrganizationalUnits,omitempty"`
	Organizationsourcetype interface{} `json:"OrganizationSourceType"`
}

// DocumentReviewCommentSource represents the DocumentReviewCommentSource schema from the OpenAPI specification
type DocumentReviewCommentSource struct {
	Content interface{} `json:"Content,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// DeleteInventoryRequest represents the DeleteInventoryRequest schema from the OpenAPI specification
type DeleteInventoryRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Dryrun interface{} `json:"DryRun,omitempty"`
	Schemadeleteoption interface{} `json:"SchemaDeleteOption,omitempty"`
	Typename interface{} `json:"TypeName"`
}

// StepExecutionFilter represents the StepExecutionFilter schema from the OpenAPI specification
type StepExecutionFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// DescribeEffectiveInstanceAssociationsResult represents the DescribeEffectiveInstanceAssociationsResult schema from the OpenAPI specification
type DescribeEffectiveInstanceAssociationsResult struct {
	Associations interface{} `json:"Associations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RegisterTargetWithMaintenanceWindowResult represents the RegisterTargetWithMaintenanceWindowResult schema from the OpenAPI specification
type RegisterTargetWithMaintenanceWindowResult struct {
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
}

// DescribeMaintenanceWindowExecutionTaskInvocationsResult represents the DescribeMaintenanceWindowExecutionTaskInvocationsResult schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionTaskInvocationsResult struct {
	Windowexecutiontaskinvocationidentities interface{} `json:"WindowExecutionTaskInvocationIdentities,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetParameterResult represents the GetParameterResult schema from the OpenAPI specification
type GetParameterResult struct {
	Parameter interface{} `json:"Parameter,omitempty"`
}

// DescribePatchBaselinesResult represents the DescribePatchBaselinesResult schema from the OpenAPI specification
type DescribePatchBaselinesResult struct {
	Baselineidentities interface{} `json:"BaselineIdentities,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeInventoryDeletionsRequest represents the DescribeInventoryDeletionsRequest schema from the OpenAPI specification
type DescribeInventoryDeletionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Deletionid interface{} `json:"DeletionId,omitempty"`
}

// DocumentIdentifier represents the DocumentIdentifier schema from the OpenAPI specification
type DocumentIdentifier struct {
	Requires interface{} `json:"Requires,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Author interface{} `json:"Author,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion,omitempty"`
	Targettype interface{} `json:"TargetType,omitempty"`
	Platformtypes interface{} `json:"PlatformTypes,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Reviewstatus interface{} `json:"ReviewStatus,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Documenttype interface{} `json:"DocumentType,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
}

// MaintenanceWindowTaskParameters represents the MaintenanceWindowTaskParameters schema from the OpenAPI specification
type MaintenanceWindowTaskParameters struct {
}

// ListDocumentMetadataHistoryResponse represents the ListDocumentMetadataHistoryResponse schema from the OpenAPI specification
type ListDocumentMetadataHistoryResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Author interface{} `json:"Author,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Metadata interface{} `json:"Metadata,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// CreateAssociationBatchRequest represents the CreateAssociationBatchRequest schema from the OpenAPI specification
type CreateAssociationBatchRequest struct {
	Entries interface{} `json:"Entries"`
}

// MaintenanceWindowIdentityForTarget represents the MaintenanceWindowIdentityForTarget schema from the OpenAPI specification
type MaintenanceWindowIdentityForTarget struct {
	Name interface{} `json:"Name,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
}

// GetAutomationExecutionRequest represents the GetAutomationExecutionRequest schema from the OpenAPI specification
type GetAutomationExecutionRequest struct {
	Automationexecutionid interface{} `json:"AutomationExecutionId"`
}

// ListAssociationsResult represents the ListAssociationsResult schema from the OpenAPI specification
type ListAssociationsResult struct {
	Associations interface{} `json:"Associations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// StepExecution represents the StepExecution schema from the OpenAPI specification
type StepExecution struct {
	Validnextsteps interface{} `json:"ValidNextSteps,omitempty"`
	Responsecode interface{} `json:"ResponseCode,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Response interface{} `json:"Response,omitempty"`
	Stepstatus interface{} `json:"StepStatus,omitempty"`
	Executionendtime interface{} `json:"ExecutionEndTime,omitempty"`
	Failuremessage interface{} `json:"FailureMessage,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Stepexecutionid interface{} `json:"StepExecutionId,omitempty"`
	Timeoutseconds interface{} `json:"TimeoutSeconds,omitempty"`
	Iscritical interface{} `json:"IsCritical,omitempty"`
	Failuredetails interface{} `json:"FailureDetails,omitempty"`
	Stepname interface{} `json:"StepName,omitempty"`
	Targetlocation interface{} `json:"TargetLocation,omitempty"`
	Outputs interface{} `json:"Outputs,omitempty"`
	Action interface{} `json:"Action,omitempty"`
	Overriddenparameters interface{} `json:"OverriddenParameters,omitempty"`
	Executionstarttime interface{} `json:"ExecutionStartTime,omitempty"`
	Nextstep interface{} `json:"NextStep,omitempty"`
	Maxattempts interface{} `json:"MaxAttempts,omitempty"`
	Inputs interface{} `json:"Inputs,omitempty"`
	Isend interface{} `json:"IsEnd,omitempty"`
	Onfailure interface{} `json:"OnFailure,omitempty"`
}

// DescribeMaintenanceWindowTargetsResult represents the DescribeMaintenanceWindowTargetsResult schema from the OpenAPI specification
type DescribeMaintenanceWindowTargetsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
}

// TerminateSessionRequest represents the TerminateSessionRequest schema from the OpenAPI specification
type TerminateSessionRequest struct {
	Sessionid interface{} `json:"SessionId"`
}

// PatchFilterGroup represents the PatchFilterGroup schema from the OpenAPI specification
type PatchFilterGroup struct {
	Patchfilters interface{} `json:"PatchFilters"`
}

// GetDefaultPatchBaselineRequest represents the GetDefaultPatchBaselineRequest schema from the OpenAPI specification
type GetDefaultPatchBaselineRequest struct {
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
}

// GetParameterRequest represents the GetParameterRequest schema from the OpenAPI specification
type GetParameterRequest struct {
	Name interface{} `json:"Name"`
	Withdecryption interface{} `json:"WithDecryption,omitempty"`
}

// DescribeAvailablePatchesResult represents the DescribeAvailablePatchesResult schema from the OpenAPI specification
type DescribeAvailablePatchesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Patches interface{} `json:"Patches,omitempty"`
}

// DescribeOpsItemsRequest represents the DescribeOpsItemsRequest schema from the OpenAPI specification
type DescribeOpsItemsRequest struct {
	Opsitemfilters interface{} `json:"OpsItemFilters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetConnectionStatusRequest represents the GetConnectionStatusRequest schema from the OpenAPI specification
type GetConnectionStatusRequest struct {
	Target interface{} `json:"Target"`
}

// DescribeMaintenanceWindowExecutionTasksRequest represents the DescribeMaintenanceWindowExecutionTasksRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionTasksRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId"`
	Filters interface{} `json:"Filters,omitempty"`
}

// ResolvedTargets represents the ResolvedTargets schema from the OpenAPI specification
type ResolvedTargets struct {
	Parametervalues interface{} `json:"ParameterValues,omitempty"`
	Truncated interface{} `json:"Truncated,omitempty"`
}

// CreateDocumentRequest represents the CreateDocumentRequest schema from the OpenAPI specification
type CreateDocumentRequest struct {
	Versionname interface{} `json:"VersionName,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Requires interface{} `json:"Requires,omitempty"`
	Targettype interface{} `json:"TargetType,omitempty"`
	Attachments interface{} `json:"Attachments,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Documenttype interface{} `json:"DocumentType,omitempty"`
	Name interface{} `json:"Name"`
	Content interface{} `json:"Content"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListAssociationVersionsRequest represents the ListAssociationVersionsRequest schema from the OpenAPI specification
type ListAssociationVersionsRequest struct {
	Associationid interface{} `json:"AssociationId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeAutomationExecutionsRequest represents the DescribeAutomationExecutionsRequest schema from the OpenAPI specification
type DescribeAutomationExecutionsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CancelCommandResult represents the CancelCommandResult schema from the OpenAPI specification
type CancelCommandResult struct {
}

// StopAutomationExecutionResult represents the StopAutomationExecutionResult schema from the OpenAPI specification
type StopAutomationExecutionResult struct {
}

// PutParameterRequest represents the PutParameterRequest schema from the OpenAPI specification
type PutParameterRequest struct {
	Description interface{} `json:"Description,omitempty"`
	Overwrite interface{} `json:"Overwrite,omitempty"`
	Policies interface{} `json:"Policies,omitempty"`
	Allowedpattern interface{} `json:"AllowedPattern,omitempty"`
	Keyid interface{} `json:"KeyId,omitempty"`
	Value interface{} `json:"Value"`
	Datatype interface{} `json:"DataType,omitempty"`
	Name interface{} `json:"Name"`
	Tags interface{} `json:"Tags,omitempty"`
	Tier interface{} `json:"Tier,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// LabelParameterVersionResult represents the LabelParameterVersionResult schema from the OpenAPI specification
type LabelParameterVersionResult struct {
	Invalidlabels interface{} `json:"InvalidLabels,omitempty"`
	Parameterversion interface{} `json:"ParameterVersion,omitempty"`
}

// RegisterTaskWithMaintenanceWindowRequest represents the RegisterTaskWithMaintenanceWindowRequest schema from the OpenAPI specification
type RegisterTaskWithMaintenanceWindowRequest struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Taskarn interface{} `json:"TaskArn"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Tasktype interface{} `json:"TaskType"`
	Priority interface{} `json:"Priority,omitempty"`
	Windowid interface{} `json:"WindowId"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Logginginfo interface{} `json:"LoggingInfo,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Cutoffbehavior interface{} `json:"CutoffBehavior,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Taskinvocationparameters interface{} `json:"TaskInvocationParameters,omitempty"`
}

// DescribeAssociationExecutionsResult represents the DescribeAssociationExecutionsResult schema from the OpenAPI specification
type DescribeAssociationExecutionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Associationexecutions interface{} `json:"AssociationExecutions,omitempty"`
}

// AlarmConfiguration represents the AlarmConfiguration schema from the OpenAPI specification
type AlarmConfiguration struct {
	Alarms interface{} `json:"Alarms"`
	Ignorepollalarmfailure interface{} `json:"IgnorePollAlarmFailure,omitempty"`
}

// CreateOpsMetadataResult represents the CreateOpsMetadataResult schema from the OpenAPI specification
type CreateOpsMetadataResult struct {
	Opsmetadataarn interface{} `json:"OpsMetadataArn,omitempty"`
}

// UpdateDocumentDefaultVersionResult represents the UpdateDocumentDefaultVersionResult schema from the OpenAPI specification
type UpdateDocumentDefaultVersionResult struct {
	Description interface{} `json:"Description,omitempty"`
}

// StartChangeRequestExecutionRequest represents the StartChangeRequestExecutionRequest schema from the OpenAPI specification
type StartChangeRequestExecutionRequest struct {
	Documentname interface{} `json:"DocumentName"`
	Changedetails interface{} `json:"ChangeDetails,omitempty"`
	Scheduledtime interface{} `json:"ScheduledTime,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Runbooks interface{} `json:"Runbooks"`
	Scheduledendtime interface{} `json:"ScheduledEndTime,omitempty"`
	Autoapprove interface{} `json:"AutoApprove,omitempty"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Changerequestname interface{} `json:"ChangeRequestName,omitempty"`
}

// UpdateDocumentMetadataResponse represents the UpdateDocumentMetadataResponse schema from the OpenAPI specification
type UpdateDocumentMetadataResponse struct {
}

// ResetServiceSettingResult represents the ResetServiceSettingResult schema from the OpenAPI specification
type ResetServiceSettingResult struct {
	Servicesetting interface{} `json:"ServiceSetting,omitempty"`
}

// GetMaintenanceWindowExecutionResult represents the GetMaintenanceWindowExecutionResult schema from the OpenAPI specification
type GetMaintenanceWindowExecutionResult struct {
	Starttime interface{} `json:"StartTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Taskids interface{} `json:"TaskIds,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
}

// DescribeSessionsResponse represents the DescribeSessionsResponse schema from the OpenAPI specification
type DescribeSessionsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sessions interface{} `json:"Sessions,omitempty"`
}

// DescribePatchPropertiesResult represents the DescribePatchPropertiesResult schema from the OpenAPI specification
type DescribePatchPropertiesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Properties interface{} `json:"Properties,omitempty"`
}

// ListTagsForResourceResult represents the ListTagsForResourceResult schema from the OpenAPI specification
type ListTagsForResourceResult struct {
	Taglist interface{} `json:"TagList,omitempty"`
}

// CancelMaintenanceWindowExecutionRequest represents the CancelMaintenanceWindowExecutionRequest schema from the OpenAPI specification
type CancelMaintenanceWindowExecutionRequest struct {
	Windowexecutionid interface{} `json:"WindowExecutionId"`
}

// OpsItemFilter represents the OpsItemFilter schema from the OpenAPI specification
type OpsItemFilter struct {
	Key interface{} `json:"Key"`
	Operator interface{} `json:"Operator"`
	Values interface{} `json:"Values"`
}

// UpdateMaintenanceWindowTargetRequest represents the UpdateMaintenanceWindowTargetRequest schema from the OpenAPI specification
type UpdateMaintenanceWindowTargetRequest struct {
	Targets interface{} `json:"Targets,omitempty"`
	Windowid interface{} `json:"WindowId"`
	Windowtargetid interface{} `json:"WindowTargetId"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Replace interface{} `json:"Replace,omitempty"`
}

// InventoryItemEntry represents the InventoryItemEntry schema from the OpenAPI specification
type InventoryItemEntry struct {
}

// NormalStringMap represents the NormalStringMap schema from the OpenAPI specification
type NormalStringMap struct {
}

// InventoryAggregator represents the InventoryAggregator schema from the OpenAPI specification
type InventoryAggregator struct {
	Aggregators interface{} `json:"Aggregators,omitempty"`
	Expression interface{} `json:"Expression,omitempty"`
	Groups interface{} `json:"Groups,omitempty"`
}

// GetCommandInvocationRequest represents the GetCommandInvocationRequest schema from the OpenAPI specification
type GetCommandInvocationRequest struct {
	Commandid interface{} `json:"CommandId"`
	Instanceid interface{} `json:"InstanceId"`
	Pluginname interface{} `json:"PluginName,omitempty"`
}

// ModifyDocumentPermissionResponse represents the ModifyDocumentPermissionResponse schema from the OpenAPI specification
type ModifyDocumentPermissionResponse struct {
}

// GetMaintenanceWindowExecutionRequest represents the GetMaintenanceWindowExecutionRequest schema from the OpenAPI specification
type GetMaintenanceWindowExecutionRequest struct {
	Windowexecutionid interface{} `json:"WindowExecutionId"`
}

// AddTagsToResourceRequest represents the AddTagsToResourceRequest schema from the OpenAPI specification
type AddTagsToResourceRequest struct {
	Tags interface{} `json:"Tags"`
	Resourceid interface{} `json:"ResourceId"`
	Resourcetype interface{} `json:"ResourceType"`
}

// PutResourcePolicyResponse represents the PutResourcePolicyResponse schema from the OpenAPI specification
type PutResourcePolicyResponse struct {
	Policyhash interface{} `json:"PolicyHash,omitempty"`
	Policyid interface{} `json:"PolicyId,omitempty"`
}

// InventoryResultEntity represents the InventoryResultEntity schema from the OpenAPI specification
type InventoryResultEntity struct {
	Data interface{} `json:"Data,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// DescribeMaintenanceWindowsForTargetRequest represents the DescribeMaintenanceWindowsForTargetRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowsForTargetRequest struct {
	Resourcetype interface{} `json:"ResourceType"`
	Targets interface{} `json:"Targets"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MaintenanceWindowTask represents the MaintenanceWindowTask schema from the OpenAPI specification
type MaintenanceWindowTask struct {
	Windowid interface{} `json:"WindowId,omitempty"`
	Cutoffbehavior interface{} `json:"CutoffBehavior,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Logginginfo interface{} `json:"LoggingInfo,omitempty"`
	Windowtaskid interface{} `json:"WindowTaskId,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
}

// DocumentRequires represents the DocumentRequires schema from the OpenAPI specification
type DocumentRequires struct {
	Version interface{} `json:"Version,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Name interface{} `json:"Name"`
	Requiretype interface{} `json:"RequireType,omitempty"`
}

// OpsEntityItemEntry represents the OpsEntityItemEntry schema from the OpenAPI specification
type OpsEntityItemEntry struct {
}

// PatchOrchestratorFilter represents the PatchOrchestratorFilter schema from the OpenAPI specification
type PatchOrchestratorFilter struct {
	Key interface{} `json:"Key,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// DescribeInstancePatchesRequest represents the DescribeInstancePatchesRequest schema from the OpenAPI specification
type DescribeInstancePatchesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// MaintenanceWindowLambdaParameters represents the MaintenanceWindowLambdaParameters schema from the OpenAPI specification
type MaintenanceWindowLambdaParameters struct {
	Payload interface{} `json:"Payload,omitempty"`
	Qualifier interface{} `json:"Qualifier,omitempty"`
	Clientcontext interface{} `json:"ClientContext,omitempty"`
}

// Command represents the Command schema from the OpenAPI specification
type Command struct {
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Targetcount interface{} `json:"TargetCount,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Servicerole interface{} `json:"ServiceRole,omitempty"`
	Notificationconfig interface{} `json:"NotificationConfig,omitempty"`
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Outputs3bucketname interface{} `json:"OutputS3BucketName,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Timeoutseconds interface{} `json:"TimeoutSeconds,omitempty"`
	Expiresafter interface{} `json:"ExpiresAfter,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Deliverytimedoutcount interface{} `json:"DeliveryTimedOutCount,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Completedcount interface{} `json:"CompletedCount,omitempty"`
	Outputs3keyprefix interface{} `json:"OutputS3KeyPrefix,omitempty"`
	Errorcount interface{} `json:"ErrorCount,omitempty"`
	Outputs3region interface{} `json:"OutputS3Region,omitempty"`
	Cloudwatchoutputconfig interface{} `json:"CloudWatchOutputConfig,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Requesteddatetime interface{} `json:"RequestedDateTime,omitempty"`
}

// AssociationDescription represents the AssociationDescription schema from the OpenAPI specification
type AssociationDescription struct {
	Status interface{} `json:"Status,omitempty"`
	Date interface{} `json:"Date,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Synccompliance interface{} `json:"SyncCompliance,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Applyonlyatcroninterval interface{} `json:"ApplyOnlyAtCronInterval,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Lastupdateassociationdate interface{} `json:"LastUpdateAssociationDate,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Complianceseverity interface{} `json:"ComplianceSeverity,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Alarmconfiguration AlarmConfiguration `json:"AlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Calendarnames interface{} `json:"CalendarNames,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Overview interface{} `json:"Overview,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Lastexecutiondate interface{} `json:"LastExecutionDate,omitempty"`
	Lastsuccessfulexecutiondate interface{} `json:"LastSuccessfulExecutionDate,omitempty"`
	Automationtargetparametername interface{} `json:"AutomationTargetParameterName,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
}

// ListComplianceItemsRequest represents the ListComplianceItemsRequest schema from the OpenAPI specification
type ListComplianceItemsRequest struct {
	Resourceids interface{} `json:"ResourceIds,omitempty"`
	Resourcetypes interface{} `json:"ResourceTypes,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CreateAssociationBatchRequestEntry represents the CreateAssociationBatchRequestEntry schema from the OpenAPI specification
type CreateAssociationBatchRequestEntry struct {
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Alarmconfiguration AlarmConfiguration `json:"AlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Parameters interface{} `json:"Parameters,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Complianceseverity interface{} `json:"ComplianceSeverity,omitempty"`
	Applyonlyatcroninterval interface{} `json:"ApplyOnlyAtCronInterval,omitempty"`
	Automationtargetparametername interface{} `json:"AutomationTargetParameterName,omitempty"`
	Calendarnames interface{} `json:"CalendarNames,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Name interface{} `json:"Name"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Synccompliance interface{} `json:"SyncCompliance,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
}

// AssociationExecution represents the AssociationExecution schema from the OpenAPI specification
type AssociationExecution struct {
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Createdtime interface{} `json:"CreatedTime,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Resourcecountbystatus interface{} `json:"ResourceCountByStatus,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Detailedstatus interface{} `json:"DetailedStatus,omitempty"`
	Alarmconfiguration AlarmConfiguration `json:"AlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Associationid interface{} `json:"AssociationId,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Lastexecutiondate interface{} `json:"LastExecutionDate,omitempty"`
}

// OpsEntity represents the OpsEntity schema from the OpenAPI specification
type OpsEntity struct {
	Id interface{} `json:"Id,omitempty"`
	Data interface{} `json:"Data,omitempty"`
}

// UpdateOpsItemResponse represents the UpdateOpsItemResponse schema from the OpenAPI specification
type UpdateOpsItemResponse struct {
}

// GetDocumentRequest represents the GetDocumentRequest schema from the OpenAPI specification
type GetDocumentRequest struct {
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name"`
	Versionname interface{} `json:"VersionName,omitempty"`
}

// DescribeParametersRequest represents the DescribeParametersRequest schema from the OpenAPI specification
type DescribeParametersRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Parameterfilters interface{} `json:"ParameterFilters,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// GetInventorySchemaRequest represents the GetInventorySchemaRequest schema from the OpenAPI specification
type GetInventorySchemaRequest struct {
	Typename interface{} `json:"TypeName,omitempty"`
	Aggregator interface{} `json:"Aggregator,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Subtype interface{} `json:"SubType,omitempty"`
}

// MaintenanceWindowIdentity represents the MaintenanceWindowIdentity schema from the OpenAPI specification
type MaintenanceWindowIdentity struct {
	Cutoff interface{} `json:"Cutoff,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Nextexecutiontime interface{} `json:"NextExecutionTime,omitempty"`
	Scheduletimezone interface{} `json:"ScheduleTimezone,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
}

// DescribeInventoryDeletionsResult represents the DescribeInventoryDeletionsResult schema from the OpenAPI specification
type DescribeInventoryDeletionsResult struct {
	Inventorydeletions interface{} `json:"InventoryDeletions,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PatchBaselineIdentity represents the PatchBaselineIdentity schema from the OpenAPI specification
type PatchBaselineIdentity struct {
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Baselinedescription interface{} `json:"BaselineDescription,omitempty"`
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Baselinename interface{} `json:"BaselineName,omitempty"`
	Defaultbaseline interface{} `json:"DefaultBaseline,omitempty"`
}

// BaselineOverride represents the BaselineOverride schema from the OpenAPI specification
type BaselineOverride struct {
	Sources interface{} `json:"Sources,omitempty"`
	Approvedpatchesenablenonsecurity interface{} `json:"ApprovedPatchesEnableNonSecurity,omitempty"`
	Rejectedpatches interface{} `json:"RejectedPatches,omitempty"`
	Rejectedpatchesaction interface{} `json:"RejectedPatchesAction,omitempty"`
	Approvalrules PatchRuleGroup `json:"ApprovalRules,omitempty"` // A set of rules defining the approval rules for a patch baseline.
	Approvedpatchescompliancelevel interface{} `json:"ApprovedPatchesComplianceLevel,omitempty"`
	Globalfilters PatchFilterGroup `json:"GlobalFilters,omitempty"` // A set of patch filters, typically used for approval rules.
	Approvedpatches interface{} `json:"ApprovedPatches,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
}

// GetServiceSettingRequest represents the GetServiceSettingRequest schema from the OpenAPI specification
type GetServiceSettingRequest struct {
	Settingid interface{} `json:"SettingId"`
}

// InstanceAssociationOutputUrl represents the InstanceAssociationOutputUrl schema from the OpenAPI specification
type InstanceAssociationOutputUrl struct {
	S3outputurl interface{} `json:"S3OutputUrl,omitempty"`
}

// PutInventoryResult represents the PutInventoryResult schema from the OpenAPI specification
type PutInventoryResult struct {
	Message interface{} `json:"Message,omitempty"`
}

// TargetMap represents the TargetMap schema from the OpenAPI specification
type TargetMap struct {
}

// AssociationStatus represents the AssociationStatus schema from the OpenAPI specification
type AssociationStatus struct {
	Date interface{} `json:"Date"`
	Message interface{} `json:"Message"`
	Name interface{} `json:"Name"`
	Additionalinfo interface{} `json:"AdditionalInfo,omitempty"`
}

// DescribeEffectiveInstanceAssociationsRequest represents the DescribeEffectiveInstanceAssociationsRequest schema from the OpenAPI specification
type DescribeEffectiveInstanceAssociationsRequest struct {
	Instanceid interface{} `json:"InstanceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateAssociationStatusResult represents the UpdateAssociationStatusResult schema from the OpenAPI specification
type UpdateAssociationStatusResult struct {
	Associationdescription interface{} `json:"AssociationDescription,omitempty"`
}

// DescribeInstancePatchStatesForPatchGroupResult represents the DescribeInstancePatchStatesForPatchGroupResult schema from the OpenAPI specification
type DescribeInstancePatchStatesForPatchGroupResult struct {
	Instancepatchstates interface{} `json:"InstancePatchStates,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Target interface{} `json:"Target,omitempty"`
	Details interface{} `json:"Details,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Outputurl interface{} `json:"OutputUrl,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Maxsessionduration interface{} `json:"MaxSessionDuration,omitempty"`
	Reason interface{} `json:"Reason,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Sessionid interface{} `json:"SessionId,omitempty"`
}

// RegisterPatchBaselineForPatchGroupRequest represents the RegisterPatchBaselineForPatchGroupRequest schema from the OpenAPI specification
type RegisterPatchBaselineForPatchGroupRequest struct {
	Patchgroup interface{} `json:"PatchGroup"`
	Baselineid interface{} `json:"BaselineId"`
}

// DeleteOpsMetadataRequest represents the DeleteOpsMetadataRequest schema from the OpenAPI specification
type DeleteOpsMetadataRequest struct {
	Opsmetadataarn interface{} `json:"OpsMetadataArn"`
}

// SeveritySummary represents the SeveritySummary schema from the OpenAPI specification
type SeveritySummary struct {
	Criticalcount interface{} `json:"CriticalCount,omitempty"`
	Highcount interface{} `json:"HighCount,omitempty"`
	Informationalcount interface{} `json:"InformationalCount,omitempty"`
	Lowcount interface{} `json:"LowCount,omitempty"`
	Mediumcount interface{} `json:"MediumCount,omitempty"`
	Unspecifiedcount interface{} `json:"UnspecifiedCount,omitempty"`
}

// ModifyDocumentPermissionRequest represents the ModifyDocumentPermissionRequest schema from the OpenAPI specification
type ModifyDocumentPermissionRequest struct {
	Permissiontype interface{} `json:"PermissionType"`
	Shareddocumentversion interface{} `json:"SharedDocumentVersion,omitempty"`
	Accountidstoadd interface{} `json:"AccountIdsToAdd,omitempty"`
	Accountidstoremove interface{} `json:"AccountIdsToRemove,omitempty"`
	Name interface{} `json:"Name"`
}

// GetParametersResult represents the GetParametersResult schema from the OpenAPI specification
type GetParametersResult struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	Invalidparameters interface{} `json:"InvalidParameters,omitempty"`
}

// ListInventoryEntriesResult represents the ListInventoryEntriesResult schema from the OpenAPI specification
type ListInventoryEntriesResult struct {
	Typename interface{} `json:"TypeName,omitempty"`
	Capturetime interface{} `json:"CaptureTime,omitempty"`
	Entries interface{} `json:"Entries,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion,omitempty"`
}

// Parameter represents the Parameter schema from the OpenAPI specification
type Parameter struct {
	Lastmodifieddate interface{} `json:"LastModifiedDate,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Datatype interface{} `json:"DataType,omitempty"`
	Value interface{} `json:"Value,omitempty"`
	Selector interface{} `json:"Selector,omitempty"`
	Sourceresult interface{} `json:"SourceResult,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Arn interface{} `json:"ARN,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// InstanceAssociationStatusAggregatedCount represents the InstanceAssociationStatusAggregatedCount schema from the OpenAPI specification
type InstanceAssociationStatusAggregatedCount struct {
}

// InventoryItemContentContext represents the InventoryItemContentContext schema from the OpenAPI specification
type InventoryItemContentContext struct {
}

// UpdateMaintenanceWindowRequest represents the UpdateMaintenanceWindowRequest schema from the OpenAPI specification
type UpdateMaintenanceWindowRequest struct {
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Windowid interface{} `json:"WindowId"`
	Scheduletimezone interface{} `json:"ScheduleTimezone,omitempty"`
	Replace interface{} `json:"Replace,omitempty"`
	Schedule interface{} `json:"Schedule,omitempty"`
	Cutoff interface{} `json:"Cutoff,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Startdate interface{} `json:"StartDate,omitempty"`
	Allowunassociatedtargets interface{} `json:"AllowUnassociatedTargets,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Enabled interface{} `json:"Enabled,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
}

// UpdateOpsMetadataRequest represents the UpdateOpsMetadataRequest schema from the OpenAPI specification
type UpdateOpsMetadataRequest struct {
	Opsmetadataarn interface{} `json:"OpsMetadataArn"`
	Keystodelete interface{} `json:"KeysToDelete,omitempty"`
	Metadatatoupdate interface{} `json:"MetadataToUpdate,omitempty"`
}

// ResourceDataSyncItem represents the ResourceDataSyncItem schema from the OpenAPI specification
type ResourceDataSyncItem struct {
	Syncsource interface{} `json:"SyncSource,omitempty"`
	Synctype interface{} `json:"SyncType,omitempty"`
	Laststatus interface{} `json:"LastStatus,omitempty"`
	S3destination interface{} `json:"S3Destination,omitempty"`
	Syncname interface{} `json:"SyncName,omitempty"`
	Lastsuccessfulsynctime interface{} `json:"LastSuccessfulSyncTime,omitempty"`
	Lastsynctime interface{} `json:"LastSyncTime,omitempty"`
	Synccreatedtime interface{} `json:"SyncCreatedTime,omitempty"`
	Lastsyncstatusmessage interface{} `json:"LastSyncStatusMessage,omitempty"`
	Synclastmodifiedtime interface{} `json:"SyncLastModifiedTime,omitempty"`
}

// GetMaintenanceWindowExecutionTaskInvocationRequest represents the GetMaintenanceWindowExecutionTaskInvocationRequest schema from the OpenAPI specification
type GetMaintenanceWindowExecutionTaskInvocationRequest struct {
	Taskid interface{} `json:"TaskId"`
	Windowexecutionid interface{} `json:"WindowExecutionId"`
	Invocationid interface{} `json:"InvocationId"`
}

// GetOpsMetadataRequest represents the GetOpsMetadataRequest schema from the OpenAPI specification
type GetOpsMetadataRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Opsmetadataarn interface{} `json:"OpsMetadataArn"`
}

// CreateAssociationRequest represents the CreateAssociationRequest schema from the OpenAPI specification
type CreateAssociationRequest struct {
	Synccompliance interface{} `json:"SyncCompliance,omitempty"`
	Alarmconfiguration AlarmConfiguration `json:"AlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Name interface{} `json:"Name"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Outputlocation interface{} `json:"OutputLocation,omitempty"`
	Applyonlyatcroninterval interface{} `json:"ApplyOnlyAtCronInterval,omitempty"`
	Automationtargetparametername interface{} `json:"AutomationTargetParameterName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Scheduleexpression interface{} `json:"ScheduleExpression,omitempty"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Associationname interface{} `json:"AssociationName,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Complianceseverity interface{} `json:"ComplianceSeverity,omitempty"`
	Calendarnames interface{} `json:"CalendarNames,omitempty"`
}

// GetMaintenanceWindowRequest represents the GetMaintenanceWindowRequest schema from the OpenAPI specification
type GetMaintenanceWindowRequest struct {
	Windowid interface{} `json:"WindowId"`
}

// DescribeActivationsRequest represents the DescribeActivationsRequest schema from the OpenAPI specification
type DescribeActivationsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ListCommandsRequest represents the ListCommandsRequest schema from the OpenAPI specification
type ListCommandsRequest struct {
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
}

// DescribeInstancePatchStatesRequest represents the DescribeInstancePatchStatesRequest schema from the OpenAPI specification
type DescribeInstancePatchStatesRequest struct {
	Instanceids interface{} `json:"InstanceIds"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateMaintenanceWindowTargetResult represents the UpdateMaintenanceWindowTargetResult schema from the OpenAPI specification
type UpdateMaintenanceWindowTargetResult struct {
	Name interface{} `json:"Name,omitempty"`
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
}

// PutInventoryRequest represents the PutInventoryRequest schema from the OpenAPI specification
type PutInventoryRequest struct {
	Items interface{} `json:"Items"`
	Instanceid interface{} `json:"InstanceId"`
}

// InventoryGroup represents the InventoryGroup schema from the OpenAPI specification
type InventoryGroup struct {
	Filters interface{} `json:"Filters"`
	Name interface{} `json:"Name"`
}

// DeleteParameterResult represents the DeleteParameterResult schema from the OpenAPI specification
type DeleteParameterResult struct {
}

// InstanceInformationStringFilter represents the InstanceInformationStringFilter schema from the OpenAPI specification
type InstanceInformationStringFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// AssociationExecutionTargetsFilter represents the AssociationExecutionTargetsFilter schema from the OpenAPI specification
type AssociationExecutionTargetsFilter struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// OpsFilter represents the OpsFilter schema from the OpenAPI specification
type OpsFilter struct {
	TypeField interface{} `json:"Type,omitempty"`
	Values interface{} `json:"Values"`
	Key interface{} `json:"Key"`
}

// ListDocumentVersionsRequest represents the ListDocumentVersionsRequest schema from the OpenAPI specification
type ListDocumentVersionsRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Name interface{} `json:"Name"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ServiceSetting represents the ServiceSetting schema from the OpenAPI specification
type ServiceSetting struct {
	Status interface{} `json:"Status,omitempty"`
	Arn interface{} `json:"ARN,omitempty"`
	Lastmodifieddate interface{} `json:"LastModifiedDate,omitempty"`
	Lastmodifieduser interface{} `json:"LastModifiedUser,omitempty"`
	Settingid interface{} `json:"SettingId,omitempty"`
	Settingvalue interface{} `json:"SettingValue,omitempty"`
}

// ResourceDataSyncOrganizationalUnit represents the ResourceDataSyncOrganizationalUnit schema from the OpenAPI specification
type ResourceDataSyncOrganizationalUnit struct {
	Organizationalunitid interface{} `json:"OrganizationalUnitId,omitempty"`
}

// UpdateServiceSettingResult represents the UpdateServiceSettingResult schema from the OpenAPI specification
type UpdateServiceSettingResult struct {
}

// DocumentDefaultVersionDescription represents the DocumentDefaultVersionDescription schema from the OpenAPI specification
type DocumentDefaultVersionDescription struct {
	Name interface{} `json:"Name,omitempty"`
	Defaultversion interface{} `json:"DefaultVersion,omitempty"`
	Defaultversionname interface{} `json:"DefaultVersionName,omitempty"`
}

// AssociationFilter represents the AssociationFilter schema from the OpenAPI specification
type AssociationFilter struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// InstanceAssociationOutputLocation represents the InstanceAssociationOutputLocation schema from the OpenAPI specification
type InstanceAssociationOutputLocation struct {
	S3location interface{} `json:"S3Location,omitempty"`
}

// GetMaintenanceWindowExecutionTaskRequest represents the GetMaintenanceWindowExecutionTaskRequest schema from the OpenAPI specification
type GetMaintenanceWindowExecutionTaskRequest struct {
	Taskid interface{} `json:"TaskId"`
	Windowexecutionid interface{} `json:"WindowExecutionId"`
}

// GetCalendarStateRequest represents the GetCalendarStateRequest schema from the OpenAPI specification
type GetCalendarStateRequest struct {
	Attime interface{} `json:"AtTime,omitempty"`
	Calendarnames interface{} `json:"CalendarNames"`
}

// DescribeEffectivePatchesForPatchBaselineResult represents the DescribeEffectivePatchesForPatchBaselineResult schema from the OpenAPI specification
type DescribeEffectivePatchesForPatchBaselineResult struct {
	Effectivepatches interface{} `json:"EffectivePatches,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteResourcePolicyResponse represents the DeleteResourcePolicyResponse schema from the OpenAPI specification
type DeleteResourcePolicyResponse struct {
}

// ScheduledWindowExecution represents the ScheduledWindowExecution schema from the OpenAPI specification
type ScheduledWindowExecution struct {
	Name interface{} `json:"Name,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Executiontime interface{} `json:"ExecutionTime,omitempty"`
}

// ListCommandInvocationsResult represents the ListCommandInvocationsResult schema from the OpenAPI specification
type ListCommandInvocationsResult struct {
	Commandinvocations interface{} `json:"CommandInvocations,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeActivationsResult represents the DescribeActivationsResult schema from the OpenAPI specification
type DescribeActivationsResult struct {
	Activationlist interface{} `json:"ActivationList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeMaintenanceWindowScheduleResult represents the DescribeMaintenanceWindowScheduleResult schema from the OpenAPI specification
type DescribeMaintenanceWindowScheduleResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Scheduledwindowexecutions interface{} `json:"ScheduledWindowExecutions,omitempty"`
}

// DescribePatchBaselinesRequest represents the DescribePatchBaselinesRequest schema from the OpenAPI specification
type DescribePatchBaselinesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DeleteInventoryResult represents the DeleteInventoryResult schema from the OpenAPI specification
type DeleteInventoryResult struct {
	Deletionid interface{} `json:"DeletionId,omitempty"`
	Deletionsummary interface{} `json:"DeletionSummary,omitempty"`
	Typename interface{} `json:"TypeName,omitempty"`
}

// GetInventorySchemaResult represents the GetInventorySchemaResult schema from the OpenAPI specification
type GetInventorySchemaResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Schemas interface{} `json:"Schemas,omitempty"`
}

// DescribePatchGroupsRequest represents the DescribePatchGroupsRequest schema from the OpenAPI specification
type DescribePatchGroupsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetPatchBaselineForPatchGroupResult represents the GetPatchBaselineForPatchGroupResult schema from the OpenAPI specification
type GetPatchBaselineForPatchGroupResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Patchgroup interface{} `json:"PatchGroup,omitempty"`
}

// DocumentReviews represents the DocumentReviews schema from the OpenAPI specification
type DocumentReviews struct {
	Action interface{} `json:"Action"`
	Comment interface{} `json:"Comment,omitempty"`
}

// StartAutomationExecutionResult represents the StartAutomationExecutionResult schema from the OpenAPI specification
type StartAutomationExecutionResult struct {
	Automationexecutionid interface{} `json:"AutomationExecutionId,omitempty"`
}

// CloudWatchOutputConfig represents the CloudWatchOutputConfig schema from the OpenAPI specification
type CloudWatchOutputConfig struct {
	Cloudwatchloggroupname interface{} `json:"CloudWatchLogGroupName,omitempty"`
	Cloudwatchoutputenabled interface{} `json:"CloudWatchOutputEnabled,omitempty"`
}

// TargetLocation represents the TargetLocation schema from the OpenAPI specification
type TargetLocation struct {
	Regions interface{} `json:"Regions,omitempty"`
	Targetlocationalarmconfiguration AlarmConfiguration `json:"TargetLocationAlarmConfiguration,omitempty"` // The details for the CloudWatch alarm you want to apply to an automation or command.
	Targetlocationmaxconcurrency interface{} `json:"TargetLocationMaxConcurrency,omitempty"`
	Targetlocationmaxerrors interface{} `json:"TargetLocationMaxErrors,omitempty"`
	Accounts interface{} `json:"Accounts,omitempty"`
	Executionrolename interface{} `json:"ExecutionRoleName,omitempty"`
}

// CreateAssociationBatchResult represents the CreateAssociationBatchResult schema from the OpenAPI specification
type CreateAssociationBatchResult struct {
	Failed interface{} `json:"Failed,omitempty"`
	Successful interface{} `json:"Successful,omitempty"`
}

// OpsAggregatorValueMap represents the OpsAggregatorValueMap schema from the OpenAPI specification
type OpsAggregatorValueMap struct {
}

// DescribeDocumentPermissionRequest represents the DescribeDocumentPermissionRequest schema from the OpenAPI specification
type DescribeDocumentPermissionRequest struct {
	Name interface{} `json:"Name"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Permissiontype interface{} `json:"PermissionType"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// InstanceAggregatedAssociationOverview represents the InstanceAggregatedAssociationOverview schema from the OpenAPI specification
type InstanceAggregatedAssociationOverview struct {
	Detailedstatus interface{} `json:"DetailedStatus,omitempty"`
	Instanceassociationstatusaggregatedcount interface{} `json:"InstanceAssociationStatusAggregatedCount,omitempty"`
}

// DeleteActivationRequest represents the DeleteActivationRequest schema from the OpenAPI specification
type DeleteActivationRequest struct {
	Activationid interface{} `json:"ActivationId"`
}

// GetParametersByPathResult represents the GetParametersByPathResult schema from the OpenAPI specification
type GetParametersByPathResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
}

// AutomationExecutionMetadata represents the AutomationExecutionMetadata schema from the OpenAPI specification
type AutomationExecutionMetadata struct {
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Runbooks interface{} `json:"Runbooks,omitempty"`
	Executedby interface{} `json:"ExecutedBy,omitempty"`
	Currentaction interface{} `json:"CurrentAction,omitempty"`
	Currentstepname interface{} `json:"CurrentStepName,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Targetparametername interface{} `json:"TargetParameterName,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Target interface{} `json:"Target,omitempty"`
	Automationexecutionstatus interface{} `json:"AutomationExecutionStatus,omitempty"`
	Failuremessage interface{} `json:"FailureMessage,omitempty"`
	Automationsubtype interface{} `json:"AutomationSubtype,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Logfile interface{} `json:"LogFile,omitempty"`
	Scheduledtime interface{} `json:"ScheduledTime,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Outputs interface{} `json:"Outputs,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Automationexecutionid interface{} `json:"AutomationExecutionId,omitempty"`
	Executionstarttime interface{} `json:"ExecutionStartTime,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Executionendtime interface{} `json:"ExecutionEndTime,omitempty"`
	Automationtype interface{} `json:"AutomationType,omitempty"`
	Parentautomationexecutionid interface{} `json:"ParentAutomationExecutionId,omitempty"`
	Changerequestname interface{} `json:"ChangeRequestName,omitempty"`
	Resolvedtargets interface{} `json:"ResolvedTargets,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
}

// UpdateDocumentDefaultVersionRequest represents the UpdateDocumentDefaultVersionRequest schema from the OpenAPI specification
type UpdateDocumentDefaultVersionRequest struct {
	Documentversion interface{} `json:"DocumentVersion"`
	Name interface{} `json:"Name"`
}

// Alarm represents the Alarm schema from the OpenAPI specification
type Alarm struct {
	Name interface{} `json:"Name"`
}

// EffectivePatch represents the EffectivePatch schema from the OpenAPI specification
type EffectivePatch struct {
	Patch interface{} `json:"Patch,omitempty"`
	Patchstatus interface{} `json:"PatchStatus,omitempty"`
}

// LoggingInfo represents the LoggingInfo schema from the OpenAPI specification
type LoggingInfo struct {
	S3region interface{} `json:"S3Region"`
	S3bucketname interface{} `json:"S3BucketName"`
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"`
}

// ResourceDataSyncS3Destination represents the ResourceDataSyncS3Destination schema from the OpenAPI specification
type ResourceDataSyncS3Destination struct {
	Bucketname interface{} `json:"BucketName"`
	Destinationdatasharing interface{} `json:"DestinationDataSharing,omitempty"`
	Prefix interface{} `json:"Prefix,omitempty"`
	Region interface{} `json:"Region"`
	Syncformat interface{} `json:"SyncFormat"`
	Awskmskeyarn interface{} `json:"AWSKMSKeyARN,omitempty"`
}

// ParameterStringFilter represents the ParameterStringFilter schema from the OpenAPI specification
type ParameterStringFilter struct {
	Key interface{} `json:"Key"`
	Option interface{} `json:"Option,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// ListInventoryEntriesRequest represents the ListInventoryEntriesRequest schema from the OpenAPI specification
type ListInventoryEntriesRequest struct {
	Typename interface{} `json:"TypeName"`
	Filters interface{} `json:"Filters,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// CancelCommandRequest represents the CancelCommandRequest schema from the OpenAPI specification
type CancelCommandRequest struct {
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Commandid interface{} `json:"CommandId"`
}

// GetResourcePoliciesResponse represents the GetResourcePoliciesResponse schema from the OpenAPI specification
type GetResourcePoliciesResponse struct {
	Policies interface{} `json:"Policies,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteDocumentResult represents the DeleteDocumentResult schema from the OpenAPI specification
type DeleteDocumentResult struct {
}

// UpdateMaintenanceWindowTaskRequest represents the UpdateMaintenanceWindowTaskRequest schema from the OpenAPI specification
type UpdateMaintenanceWindowTaskRequest struct {
	Taskinvocationparameters interface{} `json:"TaskInvocationParameters,omitempty"`
	Windowtaskid interface{} `json:"WindowTaskId"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Replace interface{} `json:"Replace,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
	Windowid interface{} `json:"WindowId"`
	Description interface{} `json:"Description,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Cutoffbehavior interface{} `json:"CutoffBehavior,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Logginginfo interface{} `json:"LoggingInfo,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
}

// PatchRule represents the PatchRule schema from the OpenAPI specification
type PatchRule struct {
	Enablenonsecurity interface{} `json:"EnableNonSecurity,omitempty"`
	Patchfiltergroup interface{} `json:"PatchFilterGroup"`
	Approveafterdays interface{} `json:"ApproveAfterDays,omitempty"`
	Approveuntildate interface{} `json:"ApproveUntilDate,omitempty"`
	Compliancelevel interface{} `json:"ComplianceLevel,omitempty"`
}

// UnlabelParameterVersionResult represents the UnlabelParameterVersionResult schema from the OpenAPI specification
type UnlabelParameterVersionResult struct {
	Removedlabels interface{} `json:"RemovedLabels,omitempty"`
	Invalidlabels interface{} `json:"InvalidLabels,omitempty"`
}

// DocumentReviewerResponseSource represents the DocumentReviewerResponseSource schema from the OpenAPI specification
type DocumentReviewerResponseSource struct {
	Comment interface{} `json:"Comment,omitempty"`
	Createtime interface{} `json:"CreateTime,omitempty"`
	Reviewstatus interface{} `json:"ReviewStatus,omitempty"`
	Reviewer interface{} `json:"Reviewer,omitempty"`
	Updatedtime interface{} `json:"UpdatedTime,omitempty"`
}

// AttachmentInformation represents the AttachmentInformation schema from the OpenAPI specification
type AttachmentInformation struct {
	Name interface{} `json:"Name,omitempty"`
}

// DisassociateOpsItemRelatedItemRequest represents the DisassociateOpsItemRelatedItemRequest schema from the OpenAPI specification
type DisassociateOpsItemRelatedItemRequest struct {
	Associationid interface{} `json:"AssociationId"`
	Opsitemid interface{} `json:"OpsItemId"`
}

// PatchComplianceData represents the PatchComplianceData schema from the OpenAPI specification
type PatchComplianceData struct {
	State interface{} `json:"State"`
	Title interface{} `json:"Title"`
	Cveids interface{} `json:"CVEIds,omitempty"`
	Classification interface{} `json:"Classification"`
	Installedtime interface{} `json:"InstalledTime"`
	Kbid interface{} `json:"KBId"`
	Severity interface{} `json:"Severity"`
}

// ListResourceComplianceSummariesRequest represents the ListResourceComplianceSummariesRequest schema from the OpenAPI specification
type ListResourceComplianceSummariesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DescribeDocumentResult represents the DescribeDocumentResult schema from the OpenAPI specification
type DescribeDocumentResult struct {
	Document interface{} `json:"Document,omitempty"`
}

// DescribeAssociationExecutionTargetsRequest represents the DescribeAssociationExecutionTargetsRequest schema from the OpenAPI specification
type DescribeAssociationExecutionTargetsRequest struct {
	Associationid interface{} `json:"AssociationId"`
	Executionid interface{} `json:"ExecutionId"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MaintenanceWindowAutomationParameters represents the MaintenanceWindowAutomationParameters schema from the OpenAPI specification
type MaintenanceWindowAutomationParameters struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
}

// DescribeAssociationRequest represents the DescribeAssociationRequest schema from the OpenAPI specification
type DescribeAssociationRequest struct {
	Associationid interface{} `json:"AssociationId,omitempty"`
	Associationversion interface{} `json:"AssociationVersion,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ComplianceItemEntry represents the ComplianceItemEntry schema from the OpenAPI specification
type ComplianceItemEntry struct {
	Title interface{} `json:"Title,omitempty"`
	Details interface{} `json:"Details,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Severity interface{} `json:"Severity"`
	Status interface{} `json:"Status"`
}

// DocumentParameter represents the DocumentParameter schema from the OpenAPI specification
type DocumentParameter struct {
	Defaultvalue interface{} `json:"DefaultValue,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// StartAssociationsOnceRequest represents the StartAssociationsOnceRequest schema from the OpenAPI specification
type StartAssociationsOnceRequest struct {
	Associationids interface{} `json:"AssociationIds"`
}

// DescribeMaintenanceWindowsForTargetResult represents the DescribeMaintenanceWindowsForTargetResult schema from the OpenAPI specification
type DescribeMaintenanceWindowsForTargetResult struct {
	Windowidentities interface{} `json:"WindowIdentities,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DescribeDocumentRequest represents the DescribeDocumentRequest schema from the OpenAPI specification
type DescribeDocumentRequest struct {
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Name interface{} `json:"Name"`
	Versionname interface{} `json:"VersionName,omitempty"`
}

// CancelMaintenanceWindowExecutionResult represents the CancelMaintenanceWindowExecutionResult schema from the OpenAPI specification
type CancelMaintenanceWindowExecutionResult struct {
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
}

// AssociationStatusAggregatedCount represents the AssociationStatusAggregatedCount schema from the OpenAPI specification
type AssociationStatusAggregatedCount struct {
}

// CreateMaintenanceWindowRequest represents the CreateMaintenanceWindowRequest schema from the OpenAPI specification
type CreateMaintenanceWindowRequest struct {
	Startdate interface{} `json:"StartDate,omitempty"`
	Scheduletimezone interface{} `json:"ScheduleTimezone,omitempty"`
	Enddate interface{} `json:"EndDate,omitempty"`
	Name interface{} `json:"Name"`
	Schedule interface{} `json:"Schedule"`
	Tags interface{} `json:"Tags,omitempty"`
	Allowunassociatedtargets interface{} `json:"AllowUnassociatedTargets"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Duration interface{} `json:"Duration"`
	Cutoff interface{} `json:"Cutoff"`
	Scheduleoffset interface{} `json:"ScheduleOffset,omitempty"`
}

// ListDocumentMetadataHistoryRequest represents the ListDocumentMetadataHistoryRequest schema from the OpenAPI specification
type ListDocumentMetadataHistoryRequest struct {
	Name interface{} `json:"Name"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Metadata interface{} `json:"Metadata"`
}

// DescribeOpsItemsResponse represents the DescribeOpsItemsResponse schema from the OpenAPI specification
type DescribeOpsItemsResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Opsitemsummaries interface{} `json:"OpsItemSummaries,omitempty"`
}

// OpsItemNotification represents the OpsItemNotification schema from the OpenAPI specification
type OpsItemNotification struct {
	Arn interface{} `json:"Arn,omitempty"`
}

// InventoryDeletionSummaryItem represents the InventoryDeletionSummaryItem schema from the OpenAPI specification
type InventoryDeletionSummaryItem struct {
	Remainingcount interface{} `json:"RemainingCount,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Count interface{} `json:"Count,omitempty"`
}

// MaintenanceWindowExecutionTaskInvocationIdentity represents the MaintenanceWindowExecutionTaskInvocationIdentity schema from the OpenAPI specification
type MaintenanceWindowExecutionTaskInvocationIdentity struct {
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Tasktype interface{} `json:"TaskType,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
	Invocationid interface{} `json:"InvocationId,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Taskexecutionid interface{} `json:"TaskExecutionId,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
}

// S3OutputLocation represents the S3OutputLocation schema from the OpenAPI specification
type S3OutputLocation struct {
	Outputs3bucketname interface{} `json:"OutputS3BucketName,omitempty"`
	Outputs3keyprefix interface{} `json:"OutputS3KeyPrefix,omitempty"`
	Outputs3region interface{} `json:"OutputS3Region,omitempty"`
}

// CreateOpsItemRequest represents the CreateOpsItemRequest schema from the OpenAPI specification
type CreateOpsItemRequest struct {
	Priority interface{} `json:"Priority,omitempty"`
	Severity interface{} `json:"Severity,omitempty"`
	Actualendtime interface{} `json:"ActualEndTime,omitempty"`
	Title interface{} `json:"Title"`
	Description interface{} `json:"Description"`
	Source interface{} `json:"Source"`
	Actualstarttime interface{} `json:"ActualStartTime,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Plannedstarttime interface{} `json:"PlannedStartTime,omitempty"`
	Relatedopsitems interface{} `json:"RelatedOpsItems,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Notifications interface{} `json:"Notifications,omitempty"`
	Plannedendtime interface{} `json:"PlannedEndTime,omitempty"`
	Opsitemtype interface{} `json:"OpsItemType,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Operationaldata interface{} `json:"OperationalData,omitempty"`
}

// UpdateResourceDataSyncResult represents the UpdateResourceDataSyncResult schema from the OpenAPI specification
type UpdateResourceDataSyncResult struct {
}

// NonCompliantSummary represents the NonCompliantSummary schema from the OpenAPI specification
type NonCompliantSummary struct {
	Noncompliantcount interface{} `json:"NonCompliantCount,omitempty"`
	Severitysummary interface{} `json:"SeveritySummary,omitempty"`
}

// DocumentFilter represents the DocumentFilter schema from the OpenAPI specification
type DocumentFilter struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// UnlabelParameterVersionRequest represents the UnlabelParameterVersionRequest schema from the OpenAPI specification
type UnlabelParameterVersionRequest struct {
	Labels interface{} `json:"Labels"`
	Name interface{} `json:"Name"`
	Parameterversion interface{} `json:"ParameterVersion"`
}

// DeleteParametersResult represents the DeleteParametersResult schema from the OpenAPI specification
type DeleteParametersResult struct {
	Deletedparameters interface{} `json:"DeletedParameters,omitempty"`
	Invalidparameters interface{} `json:"InvalidParameters,omitempty"`
}

// StartAssociationsOnceResult represents the StartAssociationsOnceResult schema from the OpenAPI specification
type StartAssociationsOnceResult struct {
}

// GetServiceSettingResult represents the GetServiceSettingResult schema from the OpenAPI specification
type GetServiceSettingResult struct {
	Servicesetting interface{} `json:"ServiceSetting,omitempty"`
}

// GetInventoryRequest represents the GetInventoryRequest schema from the OpenAPI specification
type GetInventoryRequest struct {
	Aggregators interface{} `json:"Aggregators,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Resultattributes interface{} `json:"ResultAttributes,omitempty"`
}

// OpsItemRelatedItemsFilter represents the OpsItemRelatedItemsFilter schema from the OpenAPI specification
type OpsItemRelatedItemsFilter struct {
	Key interface{} `json:"Key"`
	Operator interface{} `json:"Operator"`
	Values interface{} `json:"Values"`
}

// DisassociateOpsItemRelatedItemResponse represents the DisassociateOpsItemRelatedItemResponse schema from the OpenAPI specification
type DisassociateOpsItemRelatedItemResponse struct {
}

// SessionFilter represents the SessionFilter schema from the OpenAPI specification
type SessionFilter struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// UpdateMaintenanceWindowTaskResult represents the UpdateMaintenanceWindowTaskResult schema from the OpenAPI specification
type UpdateMaintenanceWindowTaskResult struct {
	Targets interface{} `json:"Targets,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Logginginfo interface{} `json:"LoggingInfo,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Windowtaskid interface{} `json:"WindowTaskId,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
	Cutoffbehavior interface{} `json:"CutoffBehavior,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Taskinvocationparameters interface{} `json:"TaskInvocationParameters,omitempty"`
}

// DescribeMaintenanceWindowTasksResult represents the DescribeMaintenanceWindowTasksResult schema from the OpenAPI specification
type DescribeMaintenanceWindowTasksResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tasks interface{} `json:"Tasks,omitempty"`
}

// InventoryItemAttribute represents the InventoryItemAttribute schema from the OpenAPI specification
type InventoryItemAttribute struct {
	Name interface{} `json:"Name"`
	Datatype interface{} `json:"DataType"`
}

// SendAutomationSignalRequest represents the SendAutomationSignalRequest schema from the OpenAPI specification
type SendAutomationSignalRequest struct {
	Automationexecutionid interface{} `json:"AutomationExecutionId"`
	Payload interface{} `json:"Payload,omitempty"`
	Signaltype interface{} `json:"SignalType"`
}

// InstanceInformation represents the InstanceInformation schema from the OpenAPI specification
type InstanceInformation struct {
	Sourcetype interface{} `json:"SourceType,omitempty"`
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Computername interface{} `json:"ComputerName,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Lastpingdatetime interface{} `json:"LastPingDateTime,omitempty"`
	Pingstatus interface{} `json:"PingStatus,omitempty"`
	Platformversion interface{} `json:"PlatformVersion,omitempty"`
	Ipaddress interface{} `json:"IPAddress,omitempty"`
	Lastsuccessfulassociationexecutiondate interface{} `json:"LastSuccessfulAssociationExecutionDate,omitempty"`
	Activationid interface{} `json:"ActivationId,omitempty"`
	Sourceid interface{} `json:"SourceId,omitempty"`
	Iamrole interface{} `json:"IamRole,omitempty"`
	Agentversion interface{} `json:"AgentVersion,omitempty"`
	Associationoverview interface{} `json:"AssociationOverview,omitempty"`
	Islatestversion interface{} `json:"IsLatestVersion,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Lastassociationexecutiondate interface{} `json:"LastAssociationExecutionDate,omitempty"`
	Platformname interface{} `json:"PlatformName,omitempty"`
	Registrationdate interface{} `json:"RegistrationDate,omitempty"`
	Associationstatus interface{} `json:"AssociationStatus,omitempty"`
	Platformtype interface{} `json:"PlatformType,omitempty"`
}

// DocumentDescription represents the DocumentDescription schema from the OpenAPI specification
type DocumentDescription struct {
	Reviewstatus interface{} `json:"ReviewStatus,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Schemaversion interface{} `json:"SchemaVersion,omitempty"`
	Targettype interface{} `json:"TargetType,omitempty"`
	Category interface{} `json:"Category,omitempty"`
	Reviewinformation interface{} `json:"ReviewInformation,omitempty"`
	Approvedversion interface{} `json:"ApprovedVersion,omitempty"`
	Versionname interface{} `json:"VersionName,omitempty"`
	Displayname interface{} `json:"DisplayName,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Owner interface{} `json:"Owner,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Statusinformation interface{} `json:"StatusInformation,omitempty"`
	Pendingreviewversion interface{} `json:"PendingReviewVersion,omitempty"`
	Hash interface{} `json:"Hash,omitempty"`
	Categoryenum interface{} `json:"CategoryEnum,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Sha1 interface{} `json:"Sha1,omitempty"`
	Author interface{} `json:"Author,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Attachmentsinformation interface{} `json:"AttachmentsInformation,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Platformtypes interface{} `json:"PlatformTypes,omitempty"`
	Requires interface{} `json:"Requires,omitempty"`
	Documentformat interface{} `json:"DocumentFormat,omitempty"`
	Latestversion interface{} `json:"LatestVersion,omitempty"`
	Documenttype interface{} `json:"DocumentType,omitempty"`
	Defaultversion interface{} `json:"DefaultVersion,omitempty"`
	Hashtype interface{} `json:"HashType,omitempty"`
}

// ListDocumentVersionsResult represents the ListDocumentVersionsResult schema from the OpenAPI specification
type ListDocumentVersionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Documentversions interface{} `json:"DocumentVersions,omitempty"`
}

// Target represents the Target schema from the OpenAPI specification
type Target struct {
	Key interface{} `json:"Key,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// GetParameterHistoryResult represents the GetParameterHistoryResult schema from the OpenAPI specification
type GetParameterHistoryResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
}

// OpsItemEventFilter represents the OpsItemEventFilter schema from the OpenAPI specification
type OpsItemEventFilter struct {
	Operator interface{} `json:"Operator"`
	Values interface{} `json:"Values"`
	Key interface{} `json:"Key"`
}

// RemoveTagsFromResourceRequest represents the RemoveTagsFromResourceRequest schema from the OpenAPI specification
type RemoveTagsFromResourceRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Resourcetype interface{} `json:"ResourceType"`
	Tagkeys interface{} `json:"TagKeys"`
}

// GetDeployablePatchSnapshotForInstanceRequest represents the GetDeployablePatchSnapshotForInstanceRequest schema from the OpenAPI specification
type GetDeployablePatchSnapshotForInstanceRequest struct {
	Baselineoverride interface{} `json:"BaselineOverride,omitempty"`
	Instanceid interface{} `json:"InstanceId"`
	Snapshotid interface{} `json:"SnapshotId"`
}

// GetPatchBaselineForPatchGroupRequest represents the GetPatchBaselineForPatchGroupRequest schema from the OpenAPI specification
type GetPatchBaselineForPatchGroupRequest struct {
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Patchgroup interface{} `json:"PatchGroup"`
}

// ListComplianceItemsResult represents the ListComplianceItemsResult schema from the OpenAPI specification
type ListComplianceItemsResult struct {
	Complianceitems interface{} `json:"ComplianceItems,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetMaintenanceWindowTaskResult represents the GetMaintenanceWindowTaskResult schema from the OpenAPI specification
type GetMaintenanceWindowTaskResult struct {
	Windowid interface{} `json:"WindowId,omitempty"`
	Logginginfo interface{} `json:"LoggingInfo,omitempty"`
	Windowtaskid interface{} `json:"WindowTaskId,omitempty"`
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Taskparameters interface{} `json:"TaskParameters,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Taskarn interface{} `json:"TaskArn,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Cutoffbehavior interface{} `json:"CutoffBehavior,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Tasktype interface{} `json:"TaskType,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Taskinvocationparameters interface{} `json:"TaskInvocationParameters,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// ListOpsItemEventsResponse represents the ListOpsItemEventsResponse schema from the OpenAPI specification
type ListOpsItemEventsResponse struct {
	Summaries interface{} `json:"Summaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RelatedOpsItem represents the RelatedOpsItem schema from the OpenAPI specification
type RelatedOpsItem struct {
	Opsitemid interface{} `json:"OpsItemId"`
}

// DescribeAvailablePatchesRequest represents the DescribeAvailablePatchesRequest schema from the OpenAPI specification
type DescribeAvailablePatchesRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// Activation represents the Activation schema from the OpenAPI specification
type Activation struct {
	Description interface{} `json:"Description,omitempty"`
	Expirationdate interface{} `json:"ExpirationDate,omitempty"`
	Expired interface{} `json:"Expired,omitempty"`
	Defaultinstancename interface{} `json:"DefaultInstanceName,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Activationid interface{} `json:"ActivationId,omitempty"`
	Iamrole interface{} `json:"IamRole,omitempty"`
	Registrationlimit interface{} `json:"RegistrationLimit,omitempty"`
	Registrationscount interface{} `json:"RegistrationsCount,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
}

// AutomationExecutionFilter represents the AutomationExecutionFilter schema from the OpenAPI specification
type AutomationExecutionFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// GetCommandInvocationResult represents the GetCommandInvocationResult schema from the OpenAPI specification
type GetCommandInvocationResult struct {
	Executionelapsedtime interface{} `json:"ExecutionElapsedTime,omitempty"`
	Standarderrorcontent interface{} `json:"StandardErrorContent,omitempty"`
	Executionstartdatetime interface{} `json:"ExecutionStartDateTime,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Executionenddatetime interface{} `json:"ExecutionEndDateTime,omitempty"`
	Pluginname interface{} `json:"PluginName,omitempty"`
	Standarderrorurl interface{} `json:"StandardErrorUrl,omitempty"`
	Commandid interface{} `json:"CommandId,omitempty"`
	Responsecode interface{} `json:"ResponseCode,omitempty"`
	Instanceid interface{} `json:"InstanceId,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Cloudwatchoutputconfig interface{} `json:"CloudWatchOutputConfig,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Standardoutputcontent interface{} `json:"StandardOutputContent,omitempty"`
	Standardoutputurl interface{} `json:"StandardOutputUrl,omitempty"`
}

// GetOpsItemRequest represents the GetOpsItemRequest schema from the OpenAPI specification
type GetOpsItemRequest struct {
	Opsitemarn interface{} `json:"OpsItemArn,omitempty"`
	Opsitemid interface{} `json:"OpsItemId"`
}

// DocumentKeyValuesFilter represents the DocumentKeyValuesFilter schema from the OpenAPI specification
type DocumentKeyValuesFilter struct {
	Key interface{} `json:"Key,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// DescribeAutomationExecutionsResult represents the DescribeAutomationExecutionsResult schema from the OpenAPI specification
type DescribeAutomationExecutionsResult struct {
	Automationexecutionmetadatalist interface{} `json:"AutomationExecutionMetadataList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ComplianceItemDetails represents the ComplianceItemDetails schema from the OpenAPI specification
type ComplianceItemDetails struct {
}

// GetPatchBaselineRequest represents the GetPatchBaselineRequest schema from the OpenAPI specification
type GetPatchBaselineRequest struct {
	Baselineid interface{} `json:"BaselineId"`
}

// MaintenanceWindowRunCommandParameters represents the MaintenanceWindowRunCommandParameters schema from the OpenAPI specification
type MaintenanceWindowRunCommandParameters struct {
	Documenthash interface{} `json:"DocumentHash,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Notificationconfig interface{} `json:"NotificationConfig,omitempty"`
	Outputs3bucketname interface{} `json:"OutputS3BucketName,omitempty"`
	Outputs3keyprefix interface{} `json:"OutputS3KeyPrefix,omitempty"`
	Servicerolearn interface{} `json:"ServiceRoleArn,omitempty"`
	Cloudwatchoutputconfig CloudWatchOutputConfig `json:"CloudWatchOutputConfig,omitempty"` // Configuration options for sending command output to Amazon CloudWatch Logs.
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Comment interface{} `json:"Comment,omitempty"`
	Documenthashtype interface{} `json:"DocumentHashType,omitempty"`
	Timeoutseconds interface{} `json:"TimeoutSeconds,omitempty"`
}

// PutResourcePolicyRequest represents the PutResourcePolicyRequest schema from the OpenAPI specification
type PutResourcePolicyRequest struct {
	Policyid interface{} `json:"PolicyId,omitempty"`
	Resourcearn interface{} `json:"ResourceArn"`
	Policy interface{} `json:"Policy"`
	Policyhash interface{} `json:"PolicyHash,omitempty"`
}

// GetMaintenanceWindowExecutionTaskInvocationResult represents the GetMaintenanceWindowExecutionTaskInvocationResult schema from the OpenAPI specification
type GetMaintenanceWindowExecutionTaskInvocationResult struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Taskexecutionid interface{} `json:"TaskExecutionId,omitempty"`
	Executionid interface{} `json:"ExecutionId,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Windowexecutionid interface{} `json:"WindowExecutionId,omitempty"`
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
	Invocationid interface{} `json:"InvocationId,omitempty"`
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	Tasktype interface{} `json:"TaskType,omitempty"`
}

// UpdatePatchBaselineRequest represents the UpdatePatchBaselineRequest schema from the OpenAPI specification
type UpdatePatchBaselineRequest struct {
	Sources interface{} `json:"Sources,omitempty"`
	Approvedpatchescompliancelevel interface{} `json:"ApprovedPatchesComplianceLevel,omitempty"`
	Baselineid interface{} `json:"BaselineId"`
	Description interface{} `json:"Description,omitempty"`
	Rejectedpatches interface{} `json:"RejectedPatches,omitempty"`
	Replace interface{} `json:"Replace,omitempty"`
	Approvalrules interface{} `json:"ApprovalRules,omitempty"`
	Globalfilters interface{} `json:"GlobalFilters,omitempty"`
	Rejectedpatchesaction interface{} `json:"RejectedPatchesAction,omitempty"`
	Approvedpatchesenablenonsecurity interface{} `json:"ApprovedPatchesEnableNonSecurity,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Approvedpatches interface{} `json:"ApprovedPatches,omitempty"`
}

// MaintenanceWindowTarget represents the MaintenanceWindowTarget schema from the OpenAPI specification
type MaintenanceWindowTarget struct {
	Resourcetype interface{} `json:"ResourceType,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Windowid interface{} `json:"WindowId,omitempty"`
	Windowtargetid interface{} `json:"WindowTargetId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Ownerinformation interface{} `json:"OwnerInformation,omitempty"`
}

// CreateAssociationResult represents the CreateAssociationResult schema from the OpenAPI specification
type CreateAssociationResult struct {
	Associationdescription interface{} `json:"AssociationDescription,omitempty"`
}

// GetParametersByPathRequest represents the GetParametersByPathRequest schema from the OpenAPI specification
type GetParametersByPathRequest struct {
	Recursive interface{} `json:"Recursive,omitempty"`
	Withdecryption interface{} `json:"WithDecryption,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Parameterfilters interface{} `json:"ParameterFilters,omitempty"`
	Path interface{} `json:"Path"`
}

// AssociationExecutionFilter represents the AssociationExecutionFilter schema from the OpenAPI specification
type AssociationExecutionFilter struct {
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// GetOpsItemResponse represents the GetOpsItemResponse schema from the OpenAPI specification
type GetOpsItemResponse struct {
	Opsitem interface{} `json:"OpsItem,omitempty"`
}

// UpdateServiceSettingRequest represents the UpdateServiceSettingRequest schema from the OpenAPI specification
type UpdateServiceSettingRequest struct {
	Settingid interface{} `json:"SettingId"`
	Settingvalue interface{} `json:"SettingValue"`
}

// GetDefaultPatchBaselineResult represents the GetDefaultPatchBaselineResult schema from the OpenAPI specification
type GetDefaultPatchBaselineResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
}

// ParametersFilter represents the ParametersFilter schema from the OpenAPI specification
type ParametersFilter struct {
	Key interface{} `json:"Key"`
	Values interface{} `json:"Values"`
}

// DeletePatchBaselineRequest represents the DeletePatchBaselineRequest schema from the OpenAPI specification
type DeletePatchBaselineRequest struct {
	Baselineid interface{} `json:"BaselineId"`
}

// OpsItemOperationalData represents the OpsItemOperationalData schema from the OpenAPI specification
type OpsItemOperationalData struct {
}

// DocumentMetadataResponseInfo represents the DocumentMetadataResponseInfo schema from the OpenAPI specification
type DocumentMetadataResponseInfo struct {
	Reviewerresponse interface{} `json:"ReviewerResponse,omitempty"`
}

// OpsEntityItem represents the OpsEntityItem schema from the OpenAPI specification
type OpsEntityItem struct {
	Content interface{} `json:"Content,omitempty"`
	Capturetime interface{} `json:"CaptureTime,omitempty"`
}

// CommandPlugin represents the CommandPlugin schema from the OpenAPI specification
type CommandPlugin struct {
	Responsestartdatetime interface{} `json:"ResponseStartDateTime,omitempty"`
	Standardoutputurl interface{} `json:"StandardOutputUrl,omitempty"`
	Output interface{} `json:"Output,omitempty"`
	Outputs3bucketname interface{} `json:"OutputS3BucketName,omitempty"`
	Standarderrorurl interface{} `json:"StandardErrorUrl,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Outputs3keyprefix interface{} `json:"OutputS3KeyPrefix,omitempty"`
	Statusdetails interface{} `json:"StatusDetails,omitempty"`
	Outputs3region interface{} `json:"OutputS3Region,omitempty"`
	Responsecode interface{} `json:"ResponseCode,omitempty"`
	Responsefinishdatetime interface{} `json:"ResponseFinishDateTime,omitempty"`
	Status interface{} `json:"Status,omitempty"`
}

// DeregisterTaskFromMaintenanceWindowRequest represents the DeregisterTaskFromMaintenanceWindowRequest schema from the OpenAPI specification
type DeregisterTaskFromMaintenanceWindowRequest struct {
	Windowtaskid interface{} `json:"WindowTaskId"`
	Windowid interface{} `json:"WindowId"`
}

// GetMaintenanceWindowTaskRequest represents the GetMaintenanceWindowTaskRequest schema from the OpenAPI specification
type GetMaintenanceWindowTaskRequest struct {
	Windowid interface{} `json:"WindowId"`
	Windowtaskid interface{} `json:"WindowTaskId"`
}

// ListAssociationVersionsResult represents the ListAssociationVersionsResult schema from the OpenAPI specification
type ListAssociationVersionsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Associationversions interface{} `json:"AssociationVersions,omitempty"`
}

// DescribeMaintenanceWindowsResult represents the DescribeMaintenanceWindowsResult schema from the OpenAPI specification
type DescribeMaintenanceWindowsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Windowidentities interface{} `json:"WindowIdentities,omitempty"`
}

// InventoryFilter represents the InventoryFilter schema from the OpenAPI specification
type InventoryFilter struct {
	Values interface{} `json:"Values"`
	Key interface{} `json:"Key"`
	TypeField interface{} `json:"Type,omitempty"`
}

// AddTagsToResourceResult represents the AddTagsToResourceResult schema from the OpenAPI specification
type AddTagsToResourceResult struct {
}

// DescribeInstanceInformationResult represents the DescribeInstanceInformationResult schema from the OpenAPI specification
type DescribeInstanceInformationResult struct {
	Instanceinformationlist interface{} `json:"InstanceInformationList,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// OpsAggregator represents the OpsAggregator schema from the OpenAPI specification
type OpsAggregator struct {
	Aggregatortype interface{} `json:"AggregatorType,omitempty"`
	Aggregators interface{} `json:"Aggregators,omitempty"`
	Attributename interface{} `json:"AttributeName,omitempty"`
	Filters interface{} `json:"Filters,omitempty"`
	Typename interface{} `json:"TypeName,omitempty"`
	Values interface{} `json:"Values,omitempty"`
}

// DescribeParametersResult represents the DescribeParametersResult schema from the OpenAPI specification
type DescribeParametersResult struct {
	Parameters interface{} `json:"Parameters,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateAssociationStatusRequest represents the UpdateAssociationStatusRequest schema from the OpenAPI specification
type UpdateAssociationStatusRequest struct {
	Associationstatus interface{} `json:"AssociationStatus"`
	Instanceid interface{} `json:"InstanceId"`
	Name interface{} `json:"Name"`
}

// SendAutomationSignalResult represents the SendAutomationSignalResult schema from the OpenAPI specification
type SendAutomationSignalResult struct {
}

// ListOpsItemEventsRequest represents the ListOpsItemEventsRequest schema from the OpenAPI specification
type ListOpsItemEventsRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RegisterPatchBaselineForPatchGroupResult represents the RegisterPatchBaselineForPatchGroupResult schema from the OpenAPI specification
type RegisterPatchBaselineForPatchGroupResult struct {
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Patchgroup interface{} `json:"PatchGroup,omitempty"`
}

// PutComplianceItemsResult represents the PutComplianceItemsResult schema from the OpenAPI specification
type PutComplianceItemsResult struct {
}

// CreateOpsMetadataRequest represents the CreateOpsMetadataRequest schema from the OpenAPI specification
type CreateOpsMetadataRequest struct {
	Metadata interface{} `json:"Metadata,omitempty"`
	Resourceid interface{} `json:"ResourceId"`
	Tags interface{} `json:"Tags,omitempty"`
}

// DescribeMaintenanceWindowExecutionTaskInvocationsRequest represents the DescribeMaintenanceWindowExecutionTaskInvocationsRequest schema from the OpenAPI specification
type DescribeMaintenanceWindowExecutionTaskInvocationsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Taskid interface{} `json:"TaskId"`
	Windowexecutionid interface{} `json:"WindowExecutionId"`
	Filters interface{} `json:"Filters,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// DescribeInstanceInformationRequest represents the DescribeInstanceInformationRequest schema from the OpenAPI specification
type DescribeInstanceInformationRequest struct {
	Filters interface{} `json:"Filters,omitempty"`
	Instanceinformationfilterlist interface{} `json:"InstanceInformationFilterList,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DeleteActivationResult represents the DeleteActivationResult schema from the OpenAPI specification
type DeleteActivationResult struct {
}

// InventoryDeletionStatusItem represents the InventoryDeletionStatusItem schema from the OpenAPI specification
type InventoryDeletionStatusItem struct {
	Deletionstarttime interface{} `json:"DeletionStartTime,omitempty"`
	Deletionsummary interface{} `json:"DeletionSummary,omitempty"`
	Laststatus interface{} `json:"LastStatus,omitempty"`
	Laststatusmessage interface{} `json:"LastStatusMessage,omitempty"`
	Laststatusupdatetime interface{} `json:"LastStatusUpdateTime,omitempty"`
	Typename interface{} `json:"TypeName,omitempty"`
	Deletionid interface{} `json:"DeletionId,omitempty"`
}

// AccountSharingInfo represents the AccountSharingInfo schema from the OpenAPI specification
type AccountSharingInfo struct {
	Shareddocumentversion interface{} `json:"SharedDocumentVersion,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
}

// PutParameterResult represents the PutParameterResult schema from the OpenAPI specification
type PutParameterResult struct {
	Tier interface{} `json:"Tier,omitempty"`
	Version interface{} `json:"Version,omitempty"`
}

// GetPatchBaselineResult represents the GetPatchBaselineResult schema from the OpenAPI specification
type GetPatchBaselineResult struct {
	Approvedpatches interface{} `json:"ApprovedPatches,omitempty"`
	Approvedpatchescompliancelevel interface{} `json:"ApprovedPatchesComplianceLevel,omitempty"`
	Baselineid interface{} `json:"BaselineId,omitempty"`
	Description interface{} `json:"Description,omitempty"`
	Modifieddate interface{} `json:"ModifiedDate,omitempty"`
	Rejectedpatches interface{} `json:"RejectedPatches,omitempty"`
	Sources interface{} `json:"Sources,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Approvalrules interface{} `json:"ApprovalRules,omitempty"`
	Patchgroups interface{} `json:"PatchGroups,omitempty"`
	Createddate interface{} `json:"CreatedDate,omitempty"`
	Globalfilters interface{} `json:"GlobalFilters,omitempty"`
	Rejectedpatchesaction interface{} `json:"RejectedPatchesAction,omitempty"`
	Operatingsystem interface{} `json:"OperatingSystem,omitempty"`
	Approvedpatchesenablenonsecurity interface{} `json:"ApprovedPatchesEnableNonSecurity,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourceid interface{} `json:"ResourceId"`
	Resourcetype interface{} `json:"ResourceType"`
}

// StartSessionResponse represents the StartSessionResponse schema from the OpenAPI specification
type StartSessionResponse struct {
	Sessionid interface{} `json:"SessionId,omitempty"`
	Streamurl interface{} `json:"StreamUrl,omitempty"`
	Tokenvalue interface{} `json:"TokenValue,omitempty"`
}

// AutomationExecution represents the AutomationExecution schema from the OpenAPI specification
type AutomationExecution struct {
	Maxerrors interface{} `json:"MaxErrors,omitempty"`
	Mode interface{} `json:"Mode,omitempty"`
	Targets interface{} `json:"Targets,omitempty"`
	Stepexecutionstruncated interface{} `json:"StepExecutionsTruncated,omitempty"`
	Opsitemid interface{} `json:"OpsItemId,omitempty"`
	Associationid interface{} `json:"AssociationId,omitempty"`
	Parentautomationexecutionid interface{} `json:"ParentAutomationExecutionId,omitempty"`
	Progresscounters interface{} `json:"ProgressCounters,omitempty"`
	Documentname interface{} `json:"DocumentName,omitempty"`
	Currentstepname interface{} `json:"CurrentStepName,omitempty"`
	Scheduledtime interface{} `json:"ScheduledTime,omitempty"`
	Maxconcurrency interface{} `json:"MaxConcurrency,omitempty"`
	Triggeredalarms interface{} `json:"TriggeredAlarms,omitempty"`
	Changerequestname interface{} `json:"ChangeRequestName,omitempty"`
	Executedby interface{} `json:"ExecutedBy,omitempty"`
	Automationexecutionstatus interface{} `json:"AutomationExecutionStatus,omitempty"`
	Currentaction interface{} `json:"CurrentAction,omitempty"`
	Resolvedtargets interface{} `json:"ResolvedTargets,omitempty"`
	Stepexecutions interface{} `json:"StepExecutions,omitempty"`
	Executionendtime interface{} `json:"ExecutionEndTime,omitempty"`
	Parameters interface{} `json:"Parameters,omitempty"`
	Failuremessage interface{} `json:"FailureMessage,omitempty"`
	Alarmconfiguration interface{} `json:"AlarmConfiguration,omitempty"`
	Targetparametername interface{} `json:"TargetParameterName,omitempty"`
	Executionstarttime interface{} `json:"ExecutionStartTime,omitempty"`
	Targetlocations interface{} `json:"TargetLocations,omitempty"`
	Runbooks interface{} `json:"Runbooks,omitempty"`
	Targetmaps interface{} `json:"TargetMaps,omitempty"`
	Outputs interface{} `json:"Outputs,omitempty"`
	Automationexecutionid interface{} `json:"AutomationExecutionId,omitempty"`
	Automationsubtype interface{} `json:"AutomationSubtype,omitempty"`
	Documentversion interface{} `json:"DocumentVersion,omitempty"`
	Target interface{} `json:"Target,omitempty"`
}

// ResourceDataSyncSourceWithState represents the ResourceDataSyncSourceWithState schema from the OpenAPI specification
type ResourceDataSyncSourceWithState struct {
	Includefutureregions interface{} `json:"IncludeFutureRegions,omitempty"`
	Sourceregions interface{} `json:"SourceRegions,omitempty"`
	Sourcetype interface{} `json:"SourceType,omitempty"`
	State interface{} `json:"State,omitempty"`
	Awsorganizationssource interface{} `json:"AwsOrganizationsSource,omitempty"`
	Enableallopsdatasources interface{} `json:"EnableAllOpsDataSources,omitempty"`
}

// DeregisterTargetFromMaintenanceWindowRequest represents the DeregisterTargetFromMaintenanceWindowRequest schema from the OpenAPI specification
type DeregisterTargetFromMaintenanceWindowRequest struct {
	Windowtargetid interface{} `json:"WindowTargetId"`
	Safe interface{} `json:"Safe,omitempty"`
	Windowid interface{} `json:"WindowId"`
}

// CreateResourceDataSyncResult represents the CreateResourceDataSyncResult schema from the OpenAPI specification
type CreateResourceDataSyncResult struct {
}
