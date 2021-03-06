// Copyright 2015 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
    "fmt"
    uuid "github.com/odeke-em/go-uuid"
)

func main() {
    fmt.Println(uuid.NewRandom())
}
