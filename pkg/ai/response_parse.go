package ai

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

const (
	// StructuredResponseCodeTypeJSON is the code type for JSON responses.
	StructuredResponseCodeTypeJSON = "json"
	// StructuredResponseCodeTypeYAML is the code type for YAML responses.
	StructuredResponseCodeTypeYAML = "yaml"
)

type StructuredResponse struct {
	Text            string
	CodeType        string
	ResponseSchemas any
}

// ParseError is the error type returned by output parsers.
type ParseError struct {
	Text   string
	Reason string
}

func (e ParseError) Error() string {
	return fmt.Sprintf("parse text %s. %s", e.Text, e.Reason)
}

func NewStructuredResponse(text string, codeType string, responseSchemas any) StructuredResponse {
	return StructuredResponse{
		Text:            text,
		CodeType:        codeType,
		ResponseSchemas: responseSchemas,
	}
}

func (s *StructuredResponse) Parse() error {
	switch s.CodeType {
	case StructuredResponseCodeTypeJSON:
		return s.parseJSON()
	case StructuredResponseCodeTypeYAML:
		return s.parseYAML()
	default:
		return ParseError{Text: s.Text, Reason: fmt.Sprintf("unsupported code type %s", s.CodeType)}
	}
}

func (s *StructuredResponse) parseJSON() error {
	content, err := s.parse()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(content), s.ResponseSchemas); err != nil {
		return ParseError{Text: s.Text, Reason: fmt.Sprintf("failed to unmarshal JSON: %s", err)}
	}

	return nil
}

func (s *StructuredResponse) parseYAML() error {
	content, err := s.parse()
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal([]byte(content), s.ResponseSchemas); err != nil {
		return ParseError{Text: s.Text, Reason: fmt.Sprintf("failed to unmarshal YAML: %s", err)}

	}

	return nil
}

func (s *StructuredResponse) parse() (string, error) {
	withoutStart := strings.Split(s.Text, fmt.Sprintf("```%s", s.CodeType))
	if !(len(withoutStart) > 1) {
		return "", ParseError{Text: s.Text, Reason: fmt.Sprintf("no ```%s at start of output", s.CodeType)}
	}
	withoutEnd := strings.Split(withoutStart[1], "```")
	if len(withoutEnd) < 1 {
		return "", ParseError{Text: s.Text, Reason: "no ``` at end of output"}
	}
	content := withoutEnd[0]
	return content, nil
}
