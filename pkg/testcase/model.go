package testcase

type TestCase struct {
	TestCaseType      string `json:"test_case_type" yaml:"test_case_type"`
	TestCaseName      string `json:"test_case_name" yaml:"test_case_name"`
	RequestUrl        string `json:"request_url" yaml:"request_url"`
	RequestMethod     string `json:"request_method" yaml:"request_method"`
	RequestBody       any    `json:"request_body" yaml:"request_body"`
	RequestHeaders    any    `json:"request_headers" yaml:"request_headers"`
	AssertResult      int    `json:"assert_result" yaml:"assert_result"`
	AssertDescription string `json:"assert_description" yaml:"assert_description"`
}

type TestCases struct {
	TestCases []TestCase `json:"test_cases" yaml:"test_cases"`
}
