// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"errors"
	"regexp"
)

// Match returns a validation rule that checks if a value matches the specified regular expression.
// This rule should only be used for validating strings and byte slices.
// An empty value is considered valid. Use the NotEmpty rule to make sure a value is not empty.
func Match(re *regexp.Regexp) *matchRule {
	return &matchRule{
		re:      re,
		message: "must be in a valid format",
	}
}

type matchRule struct {
	re      *regexp.Regexp
	message string
}

// Validate checks if the given value is valid or not.
func (v *matchRule) Validate(value interface{}, context interface{}) error {
	value, isNil := Indirect(value)
	if isNil {
		return nil
	}

	isString, str, isBytes, bs := StringOrBytes(value)
	if isString && (str == "" || v.re.MatchString(str)) {
		return nil
	} else if isBytes && (len(bs) == 0 || v.re.Match(bs)) {
		return nil
	}
	return errors.New(v.message)
}

// Error sets the error message for the rule.
func (v *matchRule) Error(message string) *matchRule {
	v.message = message
	return v
}
