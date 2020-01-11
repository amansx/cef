// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package cef

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

import (
	// #include <stdlib.h>
	// #include "include/internal/cef_types.h"
	"C"
)

// NewWindowInfo creates a new default WindowInfo instance.
func NewWindowInfo(parent unsafe.Pointer, bounds geom.Rect) *WindowInfo {
	d := &WindowInfo{
		X:      int32(bounds.X),
		Y:      int32(bounds.Y),
		Width:  int32(bounds.Width),
		Height: int32(bounds.Height),
	}
	d.platformInit(parent)
	return d
}
