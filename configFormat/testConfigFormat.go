package configFormat

type TestConfig struct {
	BaseUrl string `json:"base_url"`
	TestCases []TestCase `json:"test_cases"`
}

type TestCase struct {
	Headers []Header `json:"headers"`
	Path string `json:"path"`
	Method string `json:"method"`
	RequestBody string `json:"request_body"`
	ExpectResponse ExpectResponse `json:"expect_response"`
}

type Header struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type ExpectResponse struct {
	Code int `json:"code"`
	Body string `json:"body"`
	Headers []Header `json:"headers"`
}
