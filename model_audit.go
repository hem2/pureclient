/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Audit struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// The arguments provided to the command.
	Arguments string `json:"arguments,omitempty"`
	// The top level command that starts with the string \"pure\" as a convention.
	Command string `json:"command,omitempty"`
	Origin *AuditOrigin `json:"origin,omitempty"`
	// The `command` and `subcommand` combination determines which action the user attempted to perform.
	Subcommand string `json:"subcommand,omitempty"`
	// The time at which the command was run in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
	// The user who ran the command.
	User string `json:"user,omitempty"`
	// The user interface through which the user session event was performed. Valid values are `CLI`, `GUI`, and `REST`.
	UserInterface string `json:"user_interface,omitempty"`
}
