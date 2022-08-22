/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type OauthTokenResponse struct {
	// The serialized OAuth 2.0 Bearer token used to perform authenticated requests. The access token must be added to the Authorization header of all API calls.
	AccessToken string `json:"access_token,omitempty"`
	// The type of token that is issued. The Pure Storage REST API supports OAuth 2.0 access tokens.
	IssuedTokenType string `json:"issued_token_type,omitempty"`
	// Indicates how the API client can use the access token issued. The Pure Storage REST API supports the `Bearer` token.
	TokenType string `json:"token_type,omitempty"`
	// The duration after which the access token will expire. Measured in seconds. This differs from other duration fields that are expressed in milliseconds.
	ExpiresIn int32 `json:"expires_in,omitempty"`
}