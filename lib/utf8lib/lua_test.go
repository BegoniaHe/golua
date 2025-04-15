package utf8lib_test

import (
	"testing"

	"github.com/BegoniaHe/golua/lib"
	"github.com/BegoniaHe/golua/luatesting"
)

func TestUtf8Lib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
