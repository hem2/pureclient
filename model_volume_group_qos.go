/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type VolumeGroupQos struct {
	// The maximum QoS bandwidth limit for the volume. Whenever throughput exceeds the bandwidth limit, throttling occurs. Measured in bytes per second. Maximum limit is 512 GB/s.
	BandwidthLimit int64 `json:"bandwidth_limit,omitempty"`
	// The QoS IOPs limit for the volume.
	IopsLimit int64 `json:"iops_limit,omitempty"`
}
