// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package name

import (
	"testing"
)

func TestName(t *testing.T) {
	var name string
	for id, realname := range devicenames {
		name = Name(id)
		if name != realname {
			t.Fatalf("Error TestName: different names (%s != %s).", realname, name)
		}
	}
	if Name(666) != "unknown hardware" {
		t.Fatalf("Error TestName: not unknown (%s).", Name(666))
	}
}
