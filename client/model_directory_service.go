/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type DirectoryService struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// Base of the Distinguished Name (DN) of the directory service groups.
	BaseDn string `json:"base_dn,omitempty"`
	// Masked password used to query the directory.
	BindPassword string `json:"bind_password,omitempty"`
	// Username used to query the directory.
	BindUser string `json:"bind_user,omitempty"`
	// The certificate of the Certificate Authority (CA) that signed the certificates of the directory servers, which is used to validate the authenticity of the configured servers.
	CaCertificate string `json:"ca_certificate,omitempty"`
	// Whether or not server authenticity is enforced when a certificate is provided.
	CheckPeer bool `json:"check_peer,omitempty"`
	// Whether or not the directory service is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// Services for which the directory service configuration is used.
	Services []string `json:"services,omitempty"`
	// List of URIs for the configured directory servers.
	Uris []string `json:"uris,omitempty"`
	Management *DirectoryServiceManagement `json:"management,omitempty"`
}