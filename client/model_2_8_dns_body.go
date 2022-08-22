/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// Configuration information for the domain name servers.
type Model28DnsBody struct {
	// Domain suffix to be appended by the appliance when performing DNS lookups.
	Domain string `json:"domain,omitempty"`
	// List of DNS server IP addresses.
	Nameservers []string `json:"nameservers,omitempty"`
}