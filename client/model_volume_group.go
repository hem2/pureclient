/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type VolumeGroup struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// Returns a value of `true` if the volume group has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed volume group is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed volume group can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the volume group is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`
	Qos *VolumeGroupQos `json:"qos,omitempty"`
	Space *OffloadSpace `json:"space,omitempty"`
	// The amount of time left until the destroyed volume group is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed volume group can be recovered by setting `destroyed=false`.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	// The number of volumes in the volume group.
	VolumeCount int64 `json:"volume_count,omitempty"`
}