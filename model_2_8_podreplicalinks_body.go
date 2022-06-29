/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28PodreplicalinksBody struct {
	// Returns a value of `true` if the replica link is to be created in a `paused` state. Returns a value of `false` if the replica link is to be created not in a `paused` state. If not specified, defaults to `false`.
	Paused bool `json:"paused,omitempty"`
}
