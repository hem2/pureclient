/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type InlineResponse200166 struct {
	// Returns a list of all items after filtering. The values are displayed for each name where meaningful. If `total_only=true`, the `items` list will be empty.
	Items []VolumeSnapshot `json:"items,omitempty"`
}
