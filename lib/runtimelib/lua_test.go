//go:build !noquotas

package runtimelib_test

import (
	"testing"

	"github.com/BegoniaHe/golua/lib"
	"github.com/BegoniaHe/golua/luatesting"
)

func TestRuntimeLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
