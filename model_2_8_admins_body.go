/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28AdminsBody struct {
	// Password associated with the account.
	Password string `json:"password,omitempty"`
	Role *AllOf28AdminsBodyRole `json:"role,omitempty"`
}
