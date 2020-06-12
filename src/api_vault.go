/*
 * Devguard Open Hardware API
 *
 * Although devguard does offer direct end to end encrypted connections, some users may prefer a third party to establish a connection for them, as it is done by data hungry cloud corporations.  The Open Hardware API can be called with any standard https client by passing an secretkit in the headers, and a devguard hosted webservice will establish a device connection on behalf of the user.  We still do not collect any data, however due to the nature of webservices, interception by third parties is possible and the OHA service comes with no waranty whatsoever. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VaultIdentitiesGet - List carrier identities
func VaultIdentitiesGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// VaultNetworksGet - List carrier networks
func VaultNetworksGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
