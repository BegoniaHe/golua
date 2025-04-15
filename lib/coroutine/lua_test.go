package coroutine_test

import (
	"testing"

	"github.com/BegoniaHe/golua/lib"
	"github.com/BegoniaHe/golua/luatesting"
)

func TestCoroutineLib(t *testing.T) {
	luatesting.RunLuaTestsInDir(t, "lua", lib.LoadAll)
}
