package utilities

type TestingCodeAPI struct {
	Description  string // description of the test case
	Route        string // route path to test
	ExpectedCode int    // expected HTTP status code
	Body         []byte //Body
	ExpectedBody []byte //Expected body
}
