/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// SSL Certificate managed by Purity//FA.
type CertificatesCertificatesigningrequestsBody struct {
	Certificate *Api28certificatescertificatesigningrequestsCertificate `json:"certificate,omitempty"`
	// The common name field listed in the certificate.
	CommonName string `json:"common_name,omitempty"`
	// Two-letter country (ISO) code listed in the certificate.
	Country string `json:"country,omitempty"`
	// The email field listed in the certificate.
	Email string `json:"email,omitempty"`
	// The locality field listed in the certificate.
	Locality string `json:"locality,omitempty"`
	// The organization field listed in the certificate.
	Organization string `json:"organization,omitempty"`
	// The organizational unit field listed in the certificate.
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	// The state/province field listed in the certificate.
	State string `json:"state,omitempty"`
}
