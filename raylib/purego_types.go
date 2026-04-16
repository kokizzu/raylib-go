//go:build !cgo
// +build !cgo

package rl

import "github.com/jupiterrider/ffi"

var typeImage = ffi.NewType(&ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
var typeVector2 = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat)
