/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type VolumePerformance struct {
	// The average I/O size per mirrored write. Measured in bytes.
	BytesPerMirroredWrite int64 `json:"bytes_per_mirrored_write,omitempty"`
	// The average I/O size for both read and write (all) operations.
	BytesPerOp int64 `json:"bytes_per_op,omitempty"`
	// The average I/O size per read. Measured in bytes.
	BytesPerRead int64 `json:"bytes_per_read,omitempty"`
	// The average I/O size per write. Measured in bytes.
	BytesPerWrite int64 `json:"bytes_per_write,omitempty"`
	// The number of mirrored bytes written per second.
	MirroredWriteBytesPerSec int64 `json:"mirrored_write_bytes_per_sec,omitempty"`
	// The number of mirrored writes per second.
	MirroredWritesPerSec int64 `json:"mirrored_writes_per_sec,omitempty"`
	// The average time it takes the array to process a mirrored I/O write request. Measured in microseconds.
	QosRateLimitUsecPerMirroredWriteOp int64 `json:"qos_rate_limit_usec_per_mirrored_write_op,omitempty"`
	// The average time spent waiting due to QoS rate limiting for a read request. Measured in microseconds.
	QosRateLimitUsecPerReadOp int64 `json:"qos_rate_limit_usec_per_read_op,omitempty"`
	// The average time that a write I/O request spends waiting as a result of the volume reaching its QoS bandwidth limit. Measured in microseconds.
	QosRateLimitUsecPerWriteOp int64 `json:"qos_rate_limit_usec_per_write_op,omitempty"`
	// The average time that a mirrored write I/O request spends in the array waiting to be served. Measured in microseconds.
	QueueUsecPerMirroredWriteOp int64 `json:"queue_usec_per_mirrored_write_op,omitempty"`
	// The average time that a read I/O request spends in the array waiting to be served. Measured in microseconds.
	QueueUsecPerReadOp int64 `json:"queue_usec_per_read_op,omitempty"`
	// The average time that a write I/O request spends in the array waiting to be served. Measured in microseconds.
	QueueUsecPerWriteOp int64 `json:"queue_usec_per_write_op,omitempty"`
	// The number of bytes read per second.
	ReadBytesPerSec int64 `json:"read_bytes_per_sec,omitempty"`
	// The number of read requests processed per second.
	ReadsPerSec int64 `json:"reads_per_sec,omitempty"`
	// The average time required to transfer data from the initiator to the array for a mirrored write request. Measured in microseconds.
	SanUsecPerMirroredWriteOp int64 `json:"san_usec_per_mirrored_write_op,omitempty"`
	// The average time required to transfer data from the array to the initiator for a read request. Measured in microseconds.
	SanUsecPerReadOp int64 `json:"san_usec_per_read_op,omitempty"`
	// The average time required to transfer data from the initiator to the array for a write request. Measured in microseconds.
	SanUsecPerWriteOp int64 `json:"san_usec_per_write_op,omitempty"`
	// The average time required for the array to service a mirrored write request. Measured in microseconds.
	ServiceUsecPerMirroredWriteOp int64 `json:"service_usec_per_mirrored_write_op,omitempty"`
	// The average time required for the array to service a read request. Measured in microseconds.
	ServiceUsecPerReadOp int64 `json:"service_usec_per_read_op,omitempty"`
	// The average time required for the array to service a write request. Measured in microseconds.
	ServiceUsecPerWriteOp int64 `json:"service_usec_per_write_op,omitempty"`
	// The time when the sample performance data was taken. Measured in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
	// The average time it takes the array to process a mirrored I/O write request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time.
	UsecPerMirroredWriteOp int64 `json:"usec_per_mirrored_write_op,omitempty"`
	// The average time it takes the array to process an I/O read request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time.
	UsecPerReadOp int64 `json:"usec_per_read_op,omitempty"`
	// The average time it takes the array to process an I/O write request. Measured in microseconds. The average time does not include SAN time, queue time, or QoS rate limit time.
	UsecPerWriteOp int64 `json:"usec_per_write_op,omitempty"`
	// The number of bytes written per second.
	WriteBytesPerSec int64 `json:"write_bytes_per_sec,omitempty"`
	// The number of write requests processed per second.
	WritesPerSec int64 `json:"writes_per_sec,omitempty"`
	// The percentage reduction in `service_usec_per_read_op` due to data cache hits. For example, a value of 0.25 indicates that the value of `service_usec_per_read_op` is 25&#37; lower than it would have been without any data cache hits.
	ServiceUsecPerReadOpCacheReduction float32 `json:"service_usec_per_read_op_cache_reduction,omitempty"`
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
}
