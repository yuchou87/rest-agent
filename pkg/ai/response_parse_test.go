package ai

import (
	_ "embed"
	"github.com/yuchou87/rest-agent/pkg/models"
	"testing"
)

var (
	//go:embed test_data/response_json.txt
	responseJsonFile string

	//go:embed test_data/response_yaml.txt
	responseYamlFile string
)

func TestStructuredResponse_Parse(t *testing.T) {
	type args struct {
		file     string
		codeType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "json file test",
			args: args{
				file:     responseJsonFile,
				codeType: StructuredResponseCodeTypeJSON,
			},
			want: true,
		},
		{
			name: "yaml file test",
			args: args{
				file:     responseYamlFile,
				codeType: StructuredResponseCodeTypeYAML,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				testCases models.TestCases
			)
			s := NewStructuredResponse(tt.args.file, tt.args.codeType, &testCases)
			if err := s.Parse(); err != nil {
				t.Fatalf("failed to parse response file: %v", err)
			}
			t.Logf("test cases: %v", testCases)
		})
	}

}
