/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Port struct {
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// The iSCSI Qualified Name (or `null` if target is not iSCSI).
	Iqn string `json:"iqn,omitempty"`
	// NVMe Qualified Name (or `null` if target is not NVMeoF).
	Nqn string `json:"nqn,omitempty"`
	// IP and port number (or `null` if target is not iSCSI).
	Portal string `json:"portal,omitempty"`
	// Fibre Channel World Wide Name (or `null` if target is not Fibre Channel).
	Wwn string `json:"wwn,omitempty"`
	// If the array port has failed over, returns the name of the port to which this port has failed over.
	Failover string `json:"failover,omitempty"`
}
