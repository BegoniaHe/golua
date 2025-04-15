package stringlib_test

import (
	"testing"

	"github.com/BegoniaHe/golua/lib"
	"github.com/BegoniaHe/golua/luatesting"
)

func TestStringLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
