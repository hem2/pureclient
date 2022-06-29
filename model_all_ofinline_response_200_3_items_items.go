/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type AllOfinlineResponse2003ItemsItems struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// The name of the computer account in the Active Directory domain.
	ComputerName string `json:"computer_name,omitempty"`
	// A list of directory servers used for lookups related to user authorization. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array and are only communicated with over the secure LDAP (LDAPS) protocol. If this field is `null`, the servers are resolved for the domain in DNS.
	DirectoryServers []string `json:"directory_servers,omitempty"`
	// The Active Directory domain joined.
	Domain string `json:"domain,omitempty"`
	// A list of key distribution servers to use for Kerberos protocol. Servers must be specified in FQDN format. All specified servers must be registered to the domain appropriately in the configured DNS of the array. If this field is `null`, the servers are resolved for the domain in DNS.
	KerberosServers []string `json:"kerberos_servers,omitempty"`
}
