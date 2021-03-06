// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import "errors"

// NotEmpty is a validation rule that checks if a value is not empty.
// A value is considered not empty if
// - integer, float: not zero
// - bool: true
// - string, array, slice, map: len() > 0
// - interface, pointer: not nil and the referenced value is not empty
// - any other types
var NotEmpty = &notEmptyRule{message: "cannot be blank"}

type notEmptyRule struct {
	message string
}

// Validate checks if the given value is valid or not.
func (v *notEmptyRule) Validate(value interface{}, context interface{}) error {
	if IsEmpty(value) {
		return errors.New(v.message)
	}
	return nil
}

// Error sets the error message for the rule.
func (v *notEmptyRule) Error(message string) *notEmptyRule {
	return &notEmptyRule{
		message: message,
	}
}
