/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type Tag struct {
	// Specifies whether or not to include the tag when copying the parent resource. If set to `true`, the tag is included in resource copying. If set to `false`, the tag is not included. If not specified, defaults to `true`.
	Copyable bool `json:"copyable,omitempty"`
	// Key of the tag. Supports up to 64 Unicode characters.
	Key string `json:"key,omitempty"`
	// Optional namespace of the tag. Namespace identifies the category of the tag. Omitting the namespace defaults to the namespace `default`. The `pure&#42;` namespaces are reserved for plugins and integration partners. It is recommended that customers avoid using reserved namespaces.
	Namespace string `json:"namespace,omitempty"`
	Resource *AlertEventAlert `json:"resource,omitempty"`
	// Value of the tag. Supports up to 256 Unicode characters.
	Value string `json:"value,omitempty"`
}
