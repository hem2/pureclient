/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type VolumeGroupSpace struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	Space *DirectorySpaceSpace `json:"space,omitempty"`
	// The timestamp of when the data was taken. Measured in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
}
