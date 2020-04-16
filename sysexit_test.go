// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause
// license that can be found in the LICENSE file.

package sysexit_test

import (
	"fmt"
	"testing"

	"mellium.im/sysexit"
)

var _ error = sysexit.Code(0)
var _ fmt.Stringer = sysexit.Code(0)

func TestCode(t *testing.T) {
	t.Run("Ok", func(t *testing.T) {
		s := sysexit.Ok.String()
		if es := sysexit.Ok.Error(); es != s {
			t.Errorf("Expected Error and String to have same output: want=%q, got=%q", s, es)
		}
	})
	t.Run("ErrUsage", func(t *testing.T) {
		s := sysexit.ErrUsage.String()
		if es := sysexit.ErrUsage.Error(); es != s {
			t.Errorf("Expected Error and String to have same output: want=%q, got=%q", s, es)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		c := sysexit.Code(1000)
		s := c.String()
		es := c.Error()
		if es != s {
			t.Errorf("Expected Error and String to have same output: want=%q, got=%q", s, es)
		}
		expected := fmt.Sprintf("Code(%d)", c)
		if es != expected {
			t.Errorf("wrong message for invalid code: want=%q, got=%q", expected, es)
		}
	})
}
