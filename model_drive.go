/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Drive struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// Physical storage capacity of the module in bytes.
	Capacity int64 `json:"capacity,omitempty"`
	// Details about the status of the module if not healthy.
	Details string `json:"details,omitempty"`
	// Storage protocol of the module. Valid values are `NVMe` and `SAS`.
	Protocol string `json:"protocol,omitempty"`
	// Current status of the module. Valid values are `empty`, `failed`, `healthy`, `identifying`, `missing`, `recovering`, `unadmitted`, `unhealthy`, `unrecognized`, and `updating`.
	Status string `json:"status,omitempty"`
	// The type of the module. Valid values are `cache`, `NVRAM`, `SSD`, and `virtual`.
	Type_ string `json:"type,omitempty"`
}
