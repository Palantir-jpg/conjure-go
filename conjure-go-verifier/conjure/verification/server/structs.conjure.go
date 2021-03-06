// This file was generated by Conjure and should not be manually edited.

package server

import (
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safeyaml"
)

type ClientTestCases struct {
	AutoDeserialize         map[EndpointName]PositiveAndNegativeTestCases `json:"autoDeserialize"`
	SingleHeaderService     map[EndpointName][]string                     `json:"singleHeaderService"`
	SinglePathParamService  map[EndpointName][]string                     `json:"singlePathParamService"`
	SingleQueryParamService map[EndpointName][]string                     `json:"singleQueryParamService"`
}

func (o ClientTestCases) MarshalJSON() ([]byte, error) {
	if o.AutoDeserialize == nil {
		o.AutoDeserialize = make(map[EndpointName]PositiveAndNegativeTestCases, 0)
	}
	if o.SingleHeaderService == nil {
		o.SingleHeaderService = make(map[EndpointName][]string, 0)
	}
	if o.SinglePathParamService == nil {
		o.SinglePathParamService = make(map[EndpointName][]string, 0)
	}
	if o.SingleQueryParamService == nil {
		o.SingleQueryParamService = make(map[EndpointName][]string, 0)
	}
	type ClientTestCasesAlias ClientTestCases
	return safejson.Marshal(ClientTestCasesAlias(o))
}

func (o *ClientTestCases) UnmarshalJSON(data []byte) error {
	type ClientTestCasesAlias ClientTestCases
	var rawClientTestCases ClientTestCasesAlias
	if err := safejson.Unmarshal(data, &rawClientTestCases); err != nil {
		return err
	}
	if rawClientTestCases.AutoDeserialize == nil {
		rawClientTestCases.AutoDeserialize = make(map[EndpointName]PositiveAndNegativeTestCases, 0)
	}
	if rawClientTestCases.SingleHeaderService == nil {
		rawClientTestCases.SingleHeaderService = make(map[EndpointName][]string, 0)
	}
	if rawClientTestCases.SinglePathParamService == nil {
		rawClientTestCases.SinglePathParamService = make(map[EndpointName][]string, 0)
	}
	if rawClientTestCases.SingleQueryParamService == nil {
		rawClientTestCases.SingleQueryParamService = make(map[EndpointName][]string, 0)
	}
	*o = ClientTestCases(rawClientTestCases)
	return nil
}

func (o ClientTestCases) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ClientTestCases) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.UnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type IgnoredClientTestCases struct {
	AutoDeserialize         map[EndpointName][]string `json:"autoDeserialize"`
	SingleHeaderService     map[EndpointName][]string `json:"singleHeaderService"`
	SinglePathParamService  map[EndpointName][]string `json:"singlePathParamService"`
	SingleQueryParamService map[EndpointName][]string `json:"singleQueryParamService"`
}

func (o IgnoredClientTestCases) MarshalJSON() ([]byte, error) {
	if o.AutoDeserialize == nil {
		o.AutoDeserialize = make(map[EndpointName][]string, 0)
	}
	if o.SingleHeaderService == nil {
		o.SingleHeaderService = make(map[EndpointName][]string, 0)
	}
	if o.SinglePathParamService == nil {
		o.SinglePathParamService = make(map[EndpointName][]string, 0)
	}
	if o.SingleQueryParamService == nil {
		o.SingleQueryParamService = make(map[EndpointName][]string, 0)
	}
	type IgnoredClientTestCasesAlias IgnoredClientTestCases
	return safejson.Marshal(IgnoredClientTestCasesAlias(o))
}

func (o *IgnoredClientTestCases) UnmarshalJSON(data []byte) error {
	type IgnoredClientTestCasesAlias IgnoredClientTestCases
	var rawIgnoredClientTestCases IgnoredClientTestCasesAlias
	if err := safejson.Unmarshal(data, &rawIgnoredClientTestCases); err != nil {
		return err
	}
	if rawIgnoredClientTestCases.AutoDeserialize == nil {
		rawIgnoredClientTestCases.AutoDeserialize = make(map[EndpointName][]string, 0)
	}
	if rawIgnoredClientTestCases.SingleHeaderService == nil {
		rawIgnoredClientTestCases.SingleHeaderService = make(map[EndpointName][]string, 0)
	}
	if rawIgnoredClientTestCases.SinglePathParamService == nil {
		rawIgnoredClientTestCases.SinglePathParamService = make(map[EndpointName][]string, 0)
	}
	if rawIgnoredClientTestCases.SingleQueryParamService == nil {
		rawIgnoredClientTestCases.SingleQueryParamService = make(map[EndpointName][]string, 0)
	}
	*o = IgnoredClientTestCases(rawIgnoredClientTestCases)
	return nil
}

func (o IgnoredClientTestCases) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *IgnoredClientTestCases) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.UnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type IgnoredTestCases struct {
	Client IgnoredClientTestCases `json:"client"`
}

func (o IgnoredTestCases) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *IgnoredTestCases) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.UnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type PositiveAndNegativeTestCases struct {
	Positive []string `json:"positive"`
	Negative []string `json:"negative"`
}

func (o PositiveAndNegativeTestCases) MarshalJSON() ([]byte, error) {
	if o.Positive == nil {
		o.Positive = make([]string, 0)
	}
	if o.Negative == nil {
		o.Negative = make([]string, 0)
	}
	type PositiveAndNegativeTestCasesAlias PositiveAndNegativeTestCases
	return safejson.Marshal(PositiveAndNegativeTestCasesAlias(o))
}

func (o *PositiveAndNegativeTestCases) UnmarshalJSON(data []byte) error {
	type PositiveAndNegativeTestCasesAlias PositiveAndNegativeTestCases
	var rawPositiveAndNegativeTestCases PositiveAndNegativeTestCasesAlias
	if err := safejson.Unmarshal(data, &rawPositiveAndNegativeTestCases); err != nil {
		return err
	}
	if rawPositiveAndNegativeTestCases.Positive == nil {
		rawPositiveAndNegativeTestCases.Positive = make([]string, 0)
	}
	if rawPositiveAndNegativeTestCases.Negative == nil {
		rawPositiveAndNegativeTestCases.Negative = make([]string, 0)
	}
	*o = PositiveAndNegativeTestCases(rawPositiveAndNegativeTestCases)
	return nil
}

func (o PositiveAndNegativeTestCases) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *PositiveAndNegativeTestCases) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.UnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type TestCases struct {
	Client ClientTestCases `json:"client"`
}

func (o TestCases) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *TestCases) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.UnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}
