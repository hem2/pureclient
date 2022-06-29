/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Model28SnmpmanagersBody1 struct {
	// A user-specified name. The name must be locally unique and can be changed.
	Name string `json:"name,omitempty"`
	// DNS hostname or IP address of a computer that hosts an SNMP manager to which Purity//FA is to send trap messages when it generates alerts.
	Host string `json:"host,omitempty"`
	// The type of notification the agent will send. Valid values are `inform` and `trap`.
	Notification string `json:"notification,omitempty"`
	V2c *Api28snmpagentsV2c `json:"v2c,omitempty"`
	V3 *Api28snmpagentsV3 `json:"v3,omitempty"`
	// Version of the SNMP protocol to be used by Purity//FA to communicate with the specified manager. Valid values are `v2c` and `v3`.
	Version string `json:"version,omitempty"`
}
