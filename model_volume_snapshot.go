/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type VolumeSnapshot struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// The snapshot creation time. Measured in milliseconds since the UNIX epoch.
	Created int64 `json:"created,omitempty"`
	// Returns a value of `true` if the snapshot has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed volume snapshot is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed volume snapshot can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the volume snapshot is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`
	Pod *RemoteVolumeSnapshotPod `json:"pod,omitempty"`
	// The provisioned space of the snapshot. Measured in bytes.
	Provisioned int64 `json:"provisioned,omitempty"`
	Source *RemoteVolumeSnapshotSource `json:"source,omitempty"`
	// The suffix that is appended to the `source_name` value to generate the full volume snapshot name in the form `VOL.SUFFIX`. If the suffix is not specified, the system constructs the snapshot name in the form `VOL.NNN`, where `VOL` is the volume name, and `NNN` is a monotonically increasing number.
	Suffix string `json:"suffix,omitempty"`
	// The amount of time left until the destroyed snapshot is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed snapshot can be recovered by setting `destroyed=false`.
	TimeRemaining int64 `json:"time_remaining,omitempty"`
	// A serial number generated by the system when the snapshot is created. The serial number is unique across all arrays.
	Serial string `json:"serial,omitempty"`
	Space *DirectorySpaceSpace `json:"space,omitempty"`
	VolumeGroup *VolumeSnapshotVolumeGroup `json:"volume_group,omitempty"`
}
