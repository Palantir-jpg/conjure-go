// This file was generated by Conjure and should not be manually edited.

package api

import (
	"regexp"
	"strings"

	"github.com/palantir/conjure-go-runtime/v2/conjure-go-contract/errors"
	werror "github.com/palantir/witchcraft-go-error"
	wparams "github.com/palantir/witchcraft-go-params"
)

var enumValuePattern = regexp.MustCompile("^[A-Z][A-Z0-9]*(_[A-Z0-9]+)*$")

type Enum string

const (
	EnumValue1 Enum = "VALUE1"
	// Docs for an enum value
	EnumValue2 Enum = "VALUE2"
)

func (e *Enum) UnmarshalText(data []byte) error {
	switch v := strings.ToUpper(string(data)); v {
	default:
		if !enumValuePattern.MatchString(v) {
			return werror.Convert(errors.NewInvalidArgument(wparams.NewSafeAndUnsafeParamStorer(map[string]interface{}{"enumType": "Enum", "message": "enum value must match pattern ^[A-Z][A-Z0-9]*(_[A-Z0-9]+)*$"}, map[string]interface{}{"enumValue": string(data)})))
		}
		*e = Enum(v)
	case "VALUE1":
		*e = EnumValue1
	case "VALUE2":
		*e = EnumValue2
	}
	return nil
}
