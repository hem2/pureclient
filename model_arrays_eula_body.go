/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type ArraysEulaBody struct {
	// End User Agreement text.
	Agreement string `json:"agreement,omitempty"`
	Signature *Api28arrayseulaSignature `json:"signature,omitempty"`
}
