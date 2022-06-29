/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Volume struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// The total number of hosts and host groups connected to the volume.
	ConnectionCount int64 `json:"connection_count,omitempty"`
	// The volume creation time. Measured in milliseconds since the UNIX epoch.
	Created int64 `json:"created,omitempty"`
	// Returns a value of `true` if the volume has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed volume is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed volume can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the volume is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`
	// The host encryption key status for this volume. Possible values include `none`, `detected`, and `fetched`.
	HostEncryptionKeyStatus string `json:"host_encryption_key_status,omitempty"`
	// The virtual size of the volume. Measured in bytes and must be a multiple of 512.
	Provisioned int64 `json:"provisioned,omitempty"`
	// Displays QoS limit information.
	Qos *interface{} `json:"qos,omitempty"`
	// A globally unique serial number generated by the system when the volume is created.
	Serial string `json:"serial,omitempty"`
	// Displays size and space consumption information.
	Space *interface{} `json:"space,omitempty"`
	// The amount of time left until the destroyed volume is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed volume can be recovered by setting `destroyed=false`.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	// A reference to the pod.
	Pod *interface{} `json:"pod,omitempty"`
	// A reference to the originating volume as a result of a volume copy.
	Source *interface{} `json:"source,omitempty"`
	// The type of volume. Valid values are `protocol_endpoint` and `regular`.
	Subtype string `json:"subtype,omitempty"`
	// A reference to the volume group.
	VolumeGroup *interface{} `json:"volume_group,omitempty"`
	// Valid values are `promoted` and `demoted`. Patch `requested_promotion_state` to `demoted` to demote the volume so that the volume stops accepting write requests. Patch `requested_promotion_state` to `promoted` to promote the volume so that the volume starts accepting write requests.
	RequestedPromotionState string `json:"requested_promotion_state,omitempty"`
	// Current promotion status of a volume. Valid values are `promoted` and `demoted`. A status of `promoted` indicates that the volume has been promoted and can accept write requests from hosts. This is the default status for a volume when it is created. A status of `demoted` indicates that the volume has been demoted and no longer accepts write requests.
	PromotionStatus string `json:"promotion_status,omitempty"`
}
