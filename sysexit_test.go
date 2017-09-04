// Copyright 2017 The Mellium Contributors.
// Use of this source code is governed by the BSD 2-clause license that can be
// found in the LICENSE file.

package sysexit_test

import (
	"testing"

	"mellium.im/mel/internal/sysexit"
)

var _ error = sysexit.Code(0)

func TestCode(t *testing.T) {
	if es := sysexit.USAGE.Error(); es != "" {
		t.Errorf("Expected no error output, got %s", es)
	}
}
