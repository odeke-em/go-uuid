// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid_test

import (
	"fmt"
	"testing"

	"github.com/odeke-em/go-uuid"
)

func TestSetNodeInterfaceGoroutineSafety(t *testing.T) {
	n := 30
	doneChan := make(chan struct{})
	for i := 0; i < n; i++ {
		go func(id int) {
			defer func() { doneChan <- struct{}{} }()
			uuid.SetNodeInterface(fmt.Sprintf("%d", id))
		}(i)
	}

	for i := 0; i < n; i++ {
		_ = <-doneChan
	}
}
