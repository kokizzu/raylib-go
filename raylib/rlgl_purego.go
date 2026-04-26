//go:build !cgo
// +build !cgo

package rl

import "github.com/jupiterrider/ffi"

var (
	// Matrix operations

	rlMatrixMode   = dll.MustPrep("rlMatrixMode", &ffi.TypeVoid, &ffi.TypeSint32)
	rlPushMatrix   = dll.MustPrep("rlPushMatrix", &ffi.TypeVoid)
	rlPopMatrix    = dll.MustPrep("rlPopMatrix", &ffi.TypeVoid)
	rlLoadIdentity = dll.MustPrep("rlLoadIdentity", &ffi.TypeVoid)
)

// MatrixMode - Choose the current matrix to be transformed
func MatrixMode(mode int32) {
	rlMatrixMode.Call(nil, &mode)
}

// PushMatrix - Push the current matrix to stack
func PushMatrix() {
	rlPushMatrix.Call(nil)
}

// PopMatrix - Pop lattest inserted matrix from stack
func PopMatrix() {
	rlPopMatrix.Call(nil)
}

// LoadIdentity - Reset current matrix to identity matrix
func LoadIdentity() {
	rlLoadIdentity.Call(nil)
}
