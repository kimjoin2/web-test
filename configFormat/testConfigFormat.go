package configFormat

type TestConfig struct {
	BaseUrl string `json:"base_url"`
	TestCases []TestCase `json:"test_cases"`
}

type TestCase struct {
	RequestData      RequestData    `json:"request_data"`
	ExpectedResponse ExpectResponse `json:"expected_response"`
}

type RequestData struct {
	Headers []Header `json:"headers"`
	Path string `json:"path"`
	Method string `json:"method"`
	Body string `json:"body"`
	ContentType string `json:"content_type"`
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
