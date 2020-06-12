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

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

    SetupAuth(router);
    SetupUi(router);
    InitConduitManager();

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
    c.Redirect(http.StatusFound, "/_ui/")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"SystemDiscoveryGet",
		http.MethodGet,
		"/system/discovery",
		SystemDiscoveryGet,
	},

	{
		"SystemSensorsGet",
		http.MethodGet,
		"/system/sensors",
		SystemSensorsGet,
	},

	{
		"SystemSysinfoGet",
		http.MethodGet,
		"/system/sysinfo",
		SystemSysinfoGet,
	},

	{
		"HiClaimPost",
		http.MethodPost,
		"/hi/claim",
		HiClaimPost,
	},

	{
		"HiMeGet",
		http.MethodGet,
		"/hi/me",
		HiMeGet,
	},

	{
		"StreamsGet",
		http.MethodGet,
		"/streams",
		StreamsGet,
	},

	{
		"StreamsIdDelete",
		http.MethodDelete,
		"/streams/:id",
		StreamsIdDelete,
	},

	{
		"StreamsIdPut",
		http.MethodPut,
		"/streams/:id",
		StreamsIdPut,
	},

	{
		"StreamsPost",
		http.MethodPost,
		"/streams",
		StreamsPost,
	},

	{
		"VaultIdentitiesGet",
		http.MethodGet,
		"/vault/identities",
		VaultIdentitiesGet,
	},

	{
		"VaultNetworksGet",
		http.MethodGet,
		"/vault/networks",
		VaultNetworksGet,
	},


    {
		"DevicesGet",
		http.MethodGet,
		"/devices",
		DevicesGet,
    },

	{
		"ConduitAgent",
		http.MethodGet,
        "/conduits/:conduit/agent/:agent",
		ConduitAgentWs,
	},

	{
		"RawPost",
		http.MethodPost,
        "/raw",
		RawPost,
	},
}