/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Alert struct {
	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	Id string `json:"id,omitempty"`
	// A locally unique, system-generated name. The name cannot be modified.
	Name string `json:"name,omitempty"`
	// Actual condition at the time the alert is created.
	Actual string `json:"actual,omitempty"`
	// The category of the alert. Valid values include `array`, `hardware` and `software`.
	Category string `json:"category,omitempty"`
	// The time the alert was closed in milliseconds since the UNIX epoch.
	Closed int64 `json:"closed,omitempty"`
	// The code number of the alert.
	Code int64 `json:"code,omitempty"`
	// The name of the component that generated the alert.
	ComponentName string `json:"component_name,omitempty"`
	// The type of component that generated the alert.
	ComponentType string `json:"component_type,omitempty"`
	// The time the alert was created in milliseconds since the UNIX epoch.
	Created int64 `json:"created,omitempty"`
	// A short description of the alert.
	Description string `json:"description,omitempty"`
	// Expected state or threshold under normal conditions.
	Expected string `json:"expected,omitempty"`
	// If set to `true`, the message is flagged. Important messages can can be flagged and listed separately.
	Flagged bool `json:"flagged,omitempty"`
	// Information about the alert cause.
	Issue string `json:"issue,omitempty"`
	// The URL of the relevant knowledge base page.
	KnowledgeBaseUrl string `json:"knowledge_base_url,omitempty"`
	// The time the most recent alert notification was sent in milliseconds since the UNIX epoch.
	Notified int64 `json:"notified,omitempty"`
	// The severity level of the alert. Valid values include `info`, `warning`, `critical`, and `hidden`.
	Severity string `json:"severity,omitempty"`
	// The current state of the alert. Valid values include `open`, `closing`, and `closed`.
	State string `json:"state,omitempty"`
	// A summary of the alert.
	Summary string `json:"summary,omitempty"`
	// The time the alert was last updated in milliseconds since the UNIX epoch.
	Updated int64 `json:"updated,omitempty"`
}
