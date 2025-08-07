package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GoogleCloudMemcacheV1LocationMetadata represents the GoogleCloudMemcacheV1LocationMetadata schema from the OpenAPI specification
type GoogleCloudMemcacheV1LocationMetadata struct {
	Availablezones map[string]interface{} `json:"availableZones,omitempty"` // Output only. The set of available zones in the location. The map is keyed by the lowercase ID of each zone, as defined by GCE. These keys can be specified in the `zones` field when creating a Memcached instance.
}

// ZoneMetadata represents the ZoneMetadata schema from the OpenAPI specification
type ZoneMetadata struct {
}

// LocationMetadata represents the LocationMetadata schema from the OpenAPI specification
type LocationMetadata struct {
	Availablezones map[string]interface{} `json:"availableZones,omitempty"` // Output only. The set of available zones in the location. The map is keyed by the lowercase ID of each zone, as defined by GCE. These keys can be specified in the `zones` field when creating a Memcached instance.
}

// MaintenancePolicy represents the MaintenancePolicy schema from the OpenAPI specification
type MaintenancePolicy struct {
	Updatepolicy UpdatePolicy `json:"updatePolicy,omitempty"` // Maintenance policy applicable to instance updates.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the resource was updated.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the resource was created.
	Description string `json:"description,omitempty"` // Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	Labels map[string]interface{} `json:"labels,omitempty"` // Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user.
	Name string `json:"name,omitempty"` // Required. MaintenancePolicy name using the form: `projects/{project_id}/locations/{location_id}/maintenancePolicies/{maintenance_policy_id}` where {project_id} refers to a GCP consumer project ID, {location_id} refers to a GCP region/zone, {maintenance_policy_id} must be 1-63 characters long and match the regular expression `[a-z0-9]([-a-z0-9]*[a-z0-9])?`.
	State string `json:"state,omitempty"` // Optional. The state of the policy.
}

// MemcacheParameters represents the MemcacheParameters schema from the OpenAPI specification
type MemcacheParameters struct {
	Id string `json:"id,omitempty"` // Output only. The unique ID associated with this set of parameters. Users can use this id to determine if the parameters associated with the instance differ from the parameters associated with the nodes. A discrepancy between parameter ids can inform users that they may need to take action to apply parameters on nodes.
	Params map[string]interface{} `json:"params,omitempty"` // User defined set of parameters to use in the memcached process.
}

// WeeklyMaintenanceWindow represents the WeeklyMaintenanceWindow schema from the OpenAPI specification
type WeeklyMaintenanceWindow struct {
	Day string `json:"day,omitempty"` // Required. Allows to define schedule that runs specified day of the week.
	Duration string `json:"duration,omitempty"` // Required. Duration of the time window.
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
}

// GoogleCloudSaasacceleratorManagementProvidersV1PerSliSloEligibility represents the GoogleCloudSaasacceleratorManagementProvidersV1PerSliSloEligibility schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1PerSliSloEligibility struct {
	Eligibilities map[string]interface{} `json:"eligibilities,omitempty"` // An entry in the eligibilities map specifies an eligibility for a particular SLI for the given instance. The SLI key in the name must be a valid SLI name specified in the Eligibility Exporter binary flags otherwise an error will be emitted by Eligibility Exporter and the oncaller will be alerted. If an SLI has been defined in the binary flags but the eligibilities map does not contain it, the corresponding SLI time series will not be emitted by the Eligibility Exporter. This ensures a smooth rollout and compatibility between the data produced by different versions of the Eligibility Exporters. If eligibilities map contains a key for an SLI which has not been declared in the binary flags, there will be an error message emitted in the Eligibility Exporter log and the metric for the SLI in question will not be emitted.
}

// GoogleCloudMemcacheV1ZoneMetadata represents the GoogleCloudMemcacheV1ZoneMetadata schema from the OpenAPI specification
type GoogleCloudMemcacheV1ZoneMetadata struct {
}

// GoogleCloudMemcacheV1OperationMetadata represents the GoogleCloudMemcacheV1OperationMetadata schema from the OpenAPI specification
type GoogleCloudMemcacheV1OperationMetadata struct {
	Target string `json:"target,omitempty"` // Output only. Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Output only. Name of the verb executed by the operation.
	Apiversion string `json:"apiVersion,omitempty"` // Output only. API version used to start the operation.
	Cancelrequested bool `json:"cancelRequested,omitempty"` // Output only. Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
	Createtime string `json:"createTime,omitempty"` // Output only. Time when the operation was created.
	Endtime string `json:"endTime,omitempty"` // Output only. Time when the operation finished running.
	Statusdetail string `json:"statusDetail,omitempty"` // Output only. Human-readable status of the operation, if any.
}

// GoogleCloudSaasacceleratorManagementProvidersV1ProvisionedResource represents the GoogleCloudSaasacceleratorManagementProvidersV1ProvisionedResource schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1ProvisionedResource struct {
	Resourcetype string `json:"resourceType,omitempty"` // Type of the resource. This can be either a GCP resource or a custom one (e.g. another cloud provider's VM). For GCP compute resources use singular form of the names listed in GCP compute API documentation (https://cloud.google.com/compute/docs/reference/rest/v1/), prefixed with 'compute-', for example: 'compute-instance', 'compute-disk', 'compute-autoscaler'.
	Resourceurl string `json:"resourceUrl,omitempty"` // URL identifying the resource, e.g. "https://www.googleapis.com/compute/v1/projects/...)".
}

// NodeConfig represents the NodeConfig schema from the OpenAPI specification
type NodeConfig struct {
	Cpucount int `json:"cpuCount,omitempty"` // Required. Number of cpus per Memcached node.
	Memorysizemb int `json:"memorySizeMb,omitempty"` // Required. Memory size in MiB for each Memcached node.
}

// GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSchedule represents the GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSchedule schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSchedule struct {
	Endtime string `json:"endTime,omitempty"` // The scheduled end time for the maintenance.
	Rolloutmanagementpolicy string `json:"rolloutManagementPolicy,omitempty"` // The rollout management policy this maintenance schedule is associated with. When doing reschedule update request, the reschedule should be against this given policy.
	Scheduledeadlinetime string `json:"scheduleDeadlineTime,omitempty"` // schedule_deadline_time is the time deadline any schedule start time cannot go beyond, including reschedule. It's normally the initial schedule start time plus maintenance window length (1 day or 1 week). Maintenance cannot be scheduled to start beyond this deadline.
	Starttime string `json:"startTime,omitempty"` // The scheduled start time for the maintenance.
	Canreschedule bool `json:"canReschedule,omitempty"` // This field is deprecated, and will be always set to true since reschedule can happen multiple times now. This field should not be removed until all service producers remove this for their customers.
}

// DenyMaintenancePeriod represents the DenyMaintenancePeriod schema from the OpenAPI specification
type DenyMaintenancePeriod struct {
	Startdate Date `json:"startDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
	Time TimeOfDay `json:"time,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Enddate Date `json:"endDate,omitempty"` // Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
}

// Date represents the Date schema from the OpenAPI specification
type Date struct {
	Day int `json:"day,omitempty"` // Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Month int `json:"month,omitempty"` // Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Year int `json:"year,omitempty"` // Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
}

// GoogleCloudSaasacceleratorManagementProvidersV1SloMetadata represents the GoogleCloudSaasacceleratorManagementProvidersV1SloMetadata schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1SloMetadata struct {
	Nodes []GoogleCloudSaasacceleratorManagementProvidersV1NodeSloMetadata `json:"nodes,omitempty"` // Optional. List of nodes. Some producers need to use per-node metadata to calculate SLO. This field allows such producers to publish per-node SLO meta data, which will be consumed by SSA Eligibility Exporter and published in the form of per node metric to Monarch.
	Perslieligibility GoogleCloudSaasacceleratorManagementProvidersV1PerSliSloEligibility `json:"perSliEligibility,omitempty"` // PerSliSloEligibility is a mapping from an SLI name to eligibility.
	Tier string `json:"tier,omitempty"` // Name of the SLO tier the Instance belongs to. This name will be expected to match the tiers specified in the service SLO configuration. Field is mandatory and must not be empty.
}

// MaintenanceSchedule represents the MaintenanceSchedule schema from the OpenAPI specification
type MaintenanceSchedule struct {
	Endtime string `json:"endTime,omitempty"` // Output only. The end time of any upcoming scheduled maintenance for this instance.
	Scheduledeadlinetime string `json:"scheduleDeadlineTime,omitempty"` // Output only. The deadline that the maintenance schedule start time can not go beyond, including reschedule.
	Starttime string `json:"startTime,omitempty"` // Output only. The start time of any upcoming scheduled maintenance for this instance.
}

// MaintenanceWindow represents the MaintenanceWindow schema from the OpenAPI specification
type MaintenanceWindow struct {
	Dailycycle DailyCycle `json:"dailyCycle,omitempty"` // Time window specified for daily operations.
	Weeklycycle WeeklyCycle `json:"weeklyCycle,omitempty"` // Time window specified for weekly operations.
}

// RescheduleMaintenanceRequest represents the RescheduleMaintenanceRequest schema from the OpenAPI specification
type RescheduleMaintenanceRequest struct {
	Rescheduletype string `json:"rescheduleType,omitempty"` // Required. If reschedule type is SPECIFIC_TIME, must set up schedule_time as well.
	Scheduletime string `json:"scheduleTime,omitempty"` // Timestamp when the maintenance shall be rescheduled to if reschedule_type=SPECIFIC_TIME, in RFC 3339 format, for example `2012-11-15T16:19:00.094Z`.
}

// ApplyParametersRequest represents the ApplyParametersRequest schema from the OpenAPI specification
type ApplyParametersRequest struct {
	Applyall bool `json:"applyAll,omitempty"` // Whether to apply instance-level parameter group to all nodes. If set to true, users are restricted from specifying individual nodes, and `ApplyParameters` updates all nodes within the instance.
	Nodeids []string `json:"nodeIds,omitempty"` // Nodes to which the instance-level parameter group is applied.
}

// GoogleCloudSaasacceleratorManagementProvidersV1Instance represents the GoogleCloudSaasacceleratorManagementProvidersV1Instance schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1Instance struct {
	Instancetype string `json:"instanceType,omitempty"` // Optional. The instance_type of this instance of format: projects/{project_number}/locations/{location_id}/instanceTypes/{instance_type_id}. Instance Type represents a high-level tier or SKU of the service that this instance belong to. When enabled(eg: Maintenance Rollout), Rollout uses 'instance_type' along with 'software_versions' to determine whether instance needs an update or not.
	Maintenanceschedules map[string]interface{} `json:"maintenanceSchedules,omitempty"` // The MaintenanceSchedule contains the scheduling information of published maintenance schedule with same key as software_versions.
	Slminstancetemplate string `json:"slmInstanceTemplate,omitempty"` // Link to the SLM instance template. Only populated when updating SLM instances via SSA's Actuation service adaptor. Service producers with custom control plane (e.g. Cloud SQL) doesn't need to populate this field. Instead they should use software_versions.
	State string `json:"state,omitempty"` // Output only. Current lifecycle state of the resource (e.g. if it's being created or ready to use).
	Notificationparameters map[string]interface{} `json:"notificationParameters,omitempty"` // Optional. notification_parameter are information that service producers may like to include that is not relevant to Rollout. This parameter will only be passed to Gamma and Cloud Logging for notification/logging purpose.
	Tenantprojectid string `json:"tenantProjectId,omitempty"` // Output only. ID of the associated GCP tenant project. See go/get-instance-metadata.
	Consumerdefinedname string `json:"consumerDefinedName,omitempty"` // consumer_defined_name is the name of the instance set by the service consumers. Generally this is different from the `name` field which reperesents the system-assigned id of the instance which the service consumers do not recognize. This is a required field for tenants onboarding to Maintenance Window notifications (go/slm-rollout-maintenance-policies#prerequisites).
	Labels map[string]interface{} `json:"labels,omitempty"` // Optional. Resource labels to represent user provided metadata. Each label is a key-value pair, where both the key and the value are arbitrary strings provided by the user.
	Maintenancesettings GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSettings `json:"maintenanceSettings,omitempty"` // Maintenance settings associated with instance. Allows service producers and end users to assign settings that controls maintenance on this instance.
	Provisionedresources []GoogleCloudSaasacceleratorManagementProvidersV1ProvisionedResource `json:"provisionedResources,omitempty"` // Output only. The list of data plane resources provisioned for this instance, e.g. compute VMs. See go/get-instance-metadata.
	Name string `json:"name,omitempty"` // Unique name of the resource. It uses the form: `projects/{project_number}/locations/{location_id}/instances/{instance_id}` Note: This name is passed, stored and logged across the rollout system. So use of consumer project_id or any other consumer PII in the name is strongly discouraged for wipeout (go/wipeout) compliance. See go/elysium/project_ids#storage-guidance for more details.
	Softwareversions map[string]interface{} `json:"softwareVersions,omitempty"` // Software versions that are used to deploy this instance. This can be mutated by rollout services.
	Slometadata GoogleCloudSaasacceleratorManagementProvidersV1SloMetadata `json:"sloMetadata,omitempty"` // SloMetadata contains resources required for proper SLO classification of the instance.
	Maintenancepolicynames map[string]interface{} `json:"maintenancePolicyNames,omitempty"` // Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the referenced policy must define the same policy type. For details, please refer to go/cloud-saas-mw-ug. Should not be set if maintenance_settings.maintenance_policies is set.
	Producermetadata map[string]interface{} `json:"producerMetadata,omitempty"` // Output only. Custom string attributes used primarily to expose producer-specific information in monitoring dashboards. See go/get-instance-metadata.
	Createtime string `json:"createTime,omitempty"` // Output only. Timestamp when the resource was created.
	Updatetime string `json:"updateTime,omitempty"` // Output only. Timestamp when the resource was last modified.
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata. For example the available capacity at the given location.
	Name string `json:"name,omitempty"` // Resource name for the location, which may vary between implementations. For example: `"projects/example-project/locations/us-east1"`
	Displayname string `json:"displayName,omitempty"` // The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	Labels map[string]interface{} `json:"labels,omitempty"` // Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Locationid string `json:"locationId,omitempty"` // The canonical id for this location. For example: `"us-east1"`.
}

// OperationMetadata represents the OperationMetadata schema from the OpenAPI specification
type OperationMetadata struct {
	Statusdetail string `json:"statusDetail,omitempty"` // Output only. Human-readable status of the operation, if any.
	Target string `json:"target,omitempty"` // Output only. Server-defined resource path for the target of the operation.
	Verb string `json:"verb,omitempty"` // Output only. Name of the verb executed by the operation.
	Apiversion string `json:"apiVersion,omitempty"` // Output only. API version used to start the operation.
	Cancelrequested bool `json:"cancelRequested,omitempty"` // Output only. Identifies whether the user has requested cancellation of the operation. Operations that have successfully been cancelled have Operation.error value with a google.rpc.Status.code of 1, corresponding to `Code.CANCELLED`.
	Createtime string `json:"createTime,omitempty"` // Output only. Time when the operation was created.
	Endtime string `json:"endTime,omitempty"` // Output only. Time when the operation finished running.
}

// Schedule represents the Schedule schema from the OpenAPI specification
type Schedule struct {
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
	Day string `json:"day,omitempty"` // Allows to define schedule that runs specified day of the week.
	Duration string `json:"duration,omitempty"` // Output only. Duration of the time window, set by service producer.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Details []map[string]interface{} `json:"details,omitempty"` // A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Message string `json:"message,omitempty"` // A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Code int `json:"code,omitempty"` // The status code, which should be an enum value of google.rpc.Code.
}

// TimeOfDay represents the TimeOfDay schema from the OpenAPI specification
type TimeOfDay struct {
	Hours int `json:"hours,omitempty"` // Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value "24:00:00" for scenarios like business closing time.
	Minutes int `json:"minutes,omitempty"` // Minutes of hour of day. Must be from 0 to 59.
	Nanos int `json:"nanos,omitempty"` // Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	Seconds int `json:"seconds,omitempty"` // Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.
}

// GoogleCloudMemcacheV1MaintenancePolicy represents the GoogleCloudMemcacheV1MaintenancePolicy schema from the OpenAPI specification
type GoogleCloudMemcacheV1MaintenancePolicy struct {
	Description string `json:"description,omitempty"` // Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time when the policy was updated.
	Weeklymaintenancewindow []WeeklyMaintenanceWindow `json:"weeklyMaintenanceWindow,omitempty"` // Required. Maintenance window that is applied to resources covered by this policy. Minimum 1. For the current version, the maximum number of weekly_maintenance_windows is expected to be one.
	Createtime string `json:"createTime,omitempty"` // Output only. The time when the policy was created.
}

// ListLocationsResponse represents the ListLocationsResponse schema from the OpenAPI specification
type ListLocationsResponse struct {
	Locations []Location `json:"locations,omitempty"` // A list of locations that matches the specified filter in the request.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
}

// DailyCycle represents the DailyCycle schema from the OpenAPI specification
type DailyCycle struct {
	Duration string `json:"duration,omitempty"` // Output only. Duration of the time window, set by service producer.
	Starttime TimeOfDay `json:"startTime,omitempty"` // Represents a time of day. The date and time zone are either not significant or are specified elsewhere. An API may choose to allow leap seconds. Related types are google.type.Date and `google.protobuf.Timestamp`.
}

// GoogleCloudSaasacceleratorManagementProvidersV1NotificationParameter represents the GoogleCloudSaasacceleratorManagementProvidersV1NotificationParameter schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1NotificationParameter struct {
	Values []string `json:"values,omitempty"` // Optional. Array of string values. e.g. instance's replica information.
}

// Instance represents the Instance schema from the OpenAPI specification
type Instance struct {
	Memcacheversion string `json:"memcacheVersion,omitempty"` // The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
	Instancemessages []InstanceMessage `json:"instanceMessages,omitempty"` // List of messages that describe the current state of the Memcached instance.
	Memcachefullversion string `json:"memcacheFullVersion,omitempty"` // Output only. The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be "memcached-1.5.16".
	Nodeconfig NodeConfig `json:"nodeConfig,omitempty"` // Configuration for a Memcached Node.
	Createtime string `json:"createTime,omitempty"` // Output only. The time the instance was created.
	Memcachenodes []Node `json:"memcacheNodes,omitempty"` // Output only. List of Memcached nodes. Refer to Node message for more details.
	Authorizednetwork string `json:"authorizedNetwork,omitempty"` // The full name of the Google Compute Engine [network](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. If left unspecified, the `default` network will be used.
	Labels map[string]interface{} `json:"labels,omitempty"` // Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Name string `json:"name,omitempty"` // Required. Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
	Nodecount int `json:"nodeCount,omitempty"` // Required. Number of nodes in the Memcached instance.
	Discoveryendpoint string `json:"discoveryEndpoint,omitempty"` // Output only. Endpoint for the Discovery API.
	Updatetime string `json:"updateTime,omitempty"` // Output only. The time the instance was updated.
	Zones []string `json:"zones,omitempty"` // Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
	Maintenancepolicy GoogleCloudMemcacheV1MaintenancePolicy `json:"maintenancePolicy,omitempty"` // Maintenance policy per instance.
	Parameters MemcacheParameters `json:"parameters,omitempty"`
	State string `json:"state,omitempty"` // Output only. The state of this Memcached instance.
	Maintenanceschedule MaintenanceSchedule `json:"maintenanceSchedule,omitempty"` // Upcoming maintenance schedule.
	Displayname string `json:"displayName,omitempty"` // User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
}

// Operation represents the Operation schema from the OpenAPI specification
type Operation struct {
	Done bool `json:"done,omitempty"` // If the value is `false`, it means the operation is still in progress. If `true`, the operation is completed, and either `error` or `response` is available.
	ErrorField Status `json:"error,omitempty"` // The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Service-specific metadata associated with the operation. It typically contains progress information and common metadata such as create time. Some services might not provide such metadata. Any method that returns a long-running operation should document the metadata type, if any.
	Name string `json:"name,omitempty"` // The server-assigned name, which is only unique within the same service that originally returns it. If you use the default HTTP mapping, the `name` should be a resource name ending with `operations/{unique_id}`.
	Response map[string]interface{} `json:"response,omitempty"` // The normal response of the operation in case of success. If the original method returns no data on success, such as `Delete`, the response is `google.protobuf.Empty`. If the original method is standard `Get`/`Create`/`Update`, the response should be the resource. For other methods, the response should have the type `XxxResponse`, where `Xxx` is the original method name. For example, if the original method name is `TakeSnapshot()`, the inferred response type is `TakeSnapshotResponse`.
}

// Empty represents the Empty schema from the OpenAPI specification
type Empty struct {
}

// InstanceMessage represents the InstanceMessage schema from the OpenAPI specification
type InstanceMessage struct {
	Code string `json:"code,omitempty"` // A code that correspond to one type of user-facing message.
	Message string `json:"message,omitempty"` // Message on memcached instance which will be exposed to users.
}

// ListOperationsResponse represents the ListOperationsResponse schema from the OpenAPI specification
type ListOperationsResponse struct {
	Nextpagetoken string `json:"nextPageToken,omitempty"` // The standard List next-page token.
	Operations []Operation `json:"operations,omitempty"` // A list of operations that matches the specified filter in the request.
}

// Node represents the Node schema from the OpenAPI specification
type Node struct {
	Nodeid string `json:"nodeId,omitempty"` // Output only. Identifier of the Memcached node. The node id does not include project or location like the Memcached instance name.
	Parameters MemcacheParameters `json:"parameters,omitempty"`
	Port int `json:"port,omitempty"` // Output only. The port number of the Memcached server on this node.
	State string `json:"state,omitempty"` // Output only. Current state of the Memcached node.
	Zone string `json:"zone,omitempty"` // Output only. Location (GCP Zone) for the Memcached node.
	Host string `json:"host,omitempty"` // Output only. Hostname or IP address of the Memcached node used by the clients to connect to the Memcached server on this node.
}

// GoogleCloudSaasacceleratorManagementProvidersV1NodeSloMetadata represents the GoogleCloudSaasacceleratorManagementProvidersV1NodeSloMetadata schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1NodeSloMetadata struct {
	Nodeid string `json:"nodeId,omitempty"` // The id of the node. This should be equal to SaasInstanceNode.node_id.
	Perslieligibility GoogleCloudSaasacceleratorManagementProvidersV1PerSliSloEligibility `json:"perSliEligibility,omitempty"` // PerSliSloEligibility is a mapping from an SLI name to eligibility.
	Location string `json:"location,omitempty"` // The location of the node, if different from instance location.
}

// GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSettings represents the GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSettings schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1MaintenanceSettings struct {
	Isrollback bool `json:"isRollback,omitempty"` // Optional. If the update call is triggered from rollback, set the value as true.
	Maintenancepolicies map[string]interface{} `json:"maintenancePolicies,omitempty"` // Optional. The MaintenancePolicies that have been attached to the instance. The key must be of the type name of the oneof policy name defined in MaintenancePolicy, and the embedded policy must define the same policy type. For details, please refer to go/cloud-saas-mw-ug. Should not be set if maintenance_policy_names is set. If only the name is needed, then only populate MaintenancePolicy.name.
	Exclude bool `json:"exclude,omitempty"` // Optional. Exclude instance from maintenance. When true, rollout service will not attempt maintenance on the instance. Rollout service will include the instance in reported rollout progress as not attempted.
}

// CancelOperationRequest represents the CancelOperationRequest schema from the OpenAPI specification
type CancelOperationRequest struct {
}

// UpdateParametersRequest represents the UpdateParametersRequest schema from the OpenAPI specification
type UpdateParametersRequest struct {
	Parameters MemcacheParameters `json:"parameters,omitempty"`
	Updatemask string `json:"updateMask,omitempty"` // Required. Mask of fields to update.
}

// ListInstancesResponse represents the ListInstancesResponse schema from the OpenAPI specification
type ListInstancesResponse struct {
	Instances []Instance `json:"instances,omitempty"` // A list of Memcached instances in the project in the specified location, or across all locations. If the `location_id` in the parent field of the request is "-", all regions available to the project are queried, and the results aggregated.
	Nextpagetoken string `json:"nextPageToken,omitempty"` // Token to retrieve the next page of results, or empty if there are no more results in the list.
	Unreachable []string `json:"unreachable,omitempty"` // Locations that could not be reached.
}

// UpdatePolicy represents the UpdatePolicy schema from the OpenAPI specification
type UpdatePolicy struct {
	Channel string `json:"channel,omitempty"` // Optional. Relative scheduling channel applied to resource.
	Denymaintenanceperiods []DenyMaintenancePeriod `json:"denyMaintenancePeriods,omitempty"` // Deny Maintenance Period that is applied to resource to indicate when maintenance is forbidden. User can specify zero or more non-overlapping deny periods. Maximum number of deny_maintenance_periods expected is one.
	Window MaintenanceWindow `json:"window,omitempty"` // MaintenanceWindow definition.
}

// WeeklyCycle represents the WeeklyCycle schema from the OpenAPI specification
type WeeklyCycle struct {
	Schedule []Schedule `json:"schedule,omitempty"` // User can specify multiple windows in a week. Minimum of 1 window.
}

// GoogleCloudSaasacceleratorManagementProvidersV1SloEligibility represents the GoogleCloudSaasacceleratorManagementProvidersV1SloEligibility schema from the OpenAPI specification
type GoogleCloudSaasacceleratorManagementProvidersV1SloEligibility struct {
	Eligible bool `json:"eligible,omitempty"` // Whether an instance is eligible or ineligible.
	Reason string `json:"reason,omitempty"` // User-defined reason for the current value of instance eligibility. Usually, this can be directly mapped to the internal state. An empty reason is allowed.
}
