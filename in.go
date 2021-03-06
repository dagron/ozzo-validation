// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import "errors"

// In returns a validation rule that checks if a value can be found in the given list of values.
// Note that the value being checked and the possible range of values must be of the same type.
// An empty value is considered valid. Use the NotEmpty rule to make sure a value is not empty.
func In(values ...interface{}) *inRule {
	return &inRule{
		elements: values,
		message:  "must be a valid value",
	}
}

type inRule struct {
	elements []interface{}
	message  string
}

// Validate checks if the given value is valid or not.
func (r *inRule) Validate(value interface{}, context interface{}) error {
	value, isNil := Indirect(value)
	if isNil {
		return nil
	}

	for _, e := range r.elements {
		if e == value {
			return nil
		}
	}
	return errors.New(r.message)
}

// Error sets the error message for the rule.
func (r *inRule) Error(message string) *inRule {
	r.message = message
	return r
}
