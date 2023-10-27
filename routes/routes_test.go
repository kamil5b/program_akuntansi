package routes_test

import (
	"net/http/httptest"
	"program_akuntansi/routes"
	"program_akuntansi/utilities"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var tests []utilities.TestingCodeAPI = []utilities.TestingCodeAPI{
	// First test case
	{
		Description:  "get HTTP status 200",
		Route:        "/",
		ExpectedCode: 200,
	},
	// Second test case
	{
		Description:  "get HTTP status 404, when route is not exists",
		Route:        "/not-found",
		ExpectedCode: 404,
	},
}

func TestHelloRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case

	// Define Fiber app.
	app := fiber.New()

	routes.Setup(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.Route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.ExpectedCode, resp.StatusCode, test.Description)
	}
}
