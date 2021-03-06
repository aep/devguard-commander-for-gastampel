/*
 * Devguard IoT Commander
 *
 * Although devguard does offer extremely efficient direct end to end encrypted connections, some users may prefer to use familiar http REST services. The Iot Commander is both a graphical user interface and a rest service that runs carrier connections on your behalf.  The public instance of IoT Commander is a demo that runs a conduit in the cloud and keeps secrets in a shared database. For anything business-critical please please contact sales and get an on-premise instance instead.  We still do not collect any data, however due to the nature of webservices, interception by third parties is possible and the commander service comes with no waranty whatsoever. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type RawRequest struct {

	Target  string                  `json:"target,omitempty"`
	Path    string                  `json:"path,omitempty"`
	Headers map[string]interface{}  `json:"headers,omitempty"`
}
