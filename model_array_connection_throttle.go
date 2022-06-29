/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

// The bandwidth throttling for an array connection. Configurable on PATCH only.
type ArrayConnectionThrottle struct {
	// Default maximum bandwidth threshold for outbound traffic in bytes. Once exceeded, bandwidth throttling occurs.
	DefaultLimit int64 `json:"default_limit,omitempty"`
	Window *Api28arrayconnectionsWindow `json:"window,omitempty"`
	// Maximum bandwidth threshold for outbound traffic during the specified `window_limit` time range in bytes. Once exceeded, bandwidth throttling occurs.
	WindowLimit int64 `json:"window_limit,omitempty"`
}
