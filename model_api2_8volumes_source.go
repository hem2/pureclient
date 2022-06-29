/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The source volume of a volume copy.
type Api28volumesSource struct {
	// A globally unique, system-generated ID. The ID cannot be modified.
	Id string `json:"id,omitempty"`
	// The resource name, such as volume name, pod name, snapshot name, and so on.
	Name string `json:"name,omitempty"`
}
