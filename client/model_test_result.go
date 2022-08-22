/*
 * FlashArray REST API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package pureclient

type TestResult struct {
	// Address of the component running the test.
	ComponentAddress string `json:"component_address,omitempty"`
	// Name of the component running the test.
	ComponentName string `json:"component_name,omitempty"`
	// What the test is doing.
	Description string `json:"description,omitempty"`
	// The URI of the target server being tested.
	Destination string `json:"destination,omitempty"`
	// Whether the object being tested is enabled or not. Returns a value of `true` if the the service is enabled. Returns a value of `false` if the service is disabled.
	Enabled bool `json:"enabled,omitempty"`
	// Additional information about the test result.
	ResultDetails string `json:"result_details,omitempty"`
	// Whether the object being tested passed the test or not. Returns a value of `true` if the specified test has succeeded. Returns a value of `false` if the specified test has failed.
	Success bool `json:"success,omitempty"`
	// Displays the type of test being performed. The returned values are determined by the `resource` being tested and its configuration. Values include `array-admin-group-searching`, `binding`, `connecting`, `phonehome`, `phonehome-ping`, `remote-assist`, `rootdse-searching`, `read-only-group-searching`, `storage-admin-group-searching`, and `validate-ntp-configuration`.
	TestType string `json:"test_type,omitempty"`
}