package lib

import (
	"github.com/BegoniaHe/golua/lib/base"
	"github.com/BegoniaHe/golua/lib/coroutine"
	"github.com/BegoniaHe/golua/lib/debuglib"
	"github.com/BegoniaHe/golua/lib/golib"
	"github.com/BegoniaHe/golua/lib/iolib"
	"github.com/BegoniaHe/golua/lib/mathlib"
	"github.com/BegoniaHe/golua/lib/oslib"
	"github.com/BegoniaHe/golua/lib/packagelib"
	"github.com/BegoniaHe/golua/lib/runtimelib"
	"github.com/BegoniaHe/golua/lib/stringlib"
	"github.com/BegoniaHe/golua/lib/tablelib"
	"github.com/BegoniaHe/golua/lib/utf8lib"
	rt "github.com/BegoniaHe/golua/runtime"
)

func LoadLibs(r *rt.Runtime, loaders ...packagelib.Loader) func() {
	var cleanups []func()
	for _, loader := range loaders {
		cleanup := loader.Run(r)
		if cleanup != nil {
			cleanups = append(cleanups, cleanup)
		}
	}
	return func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}
}

func LoadAll(r *rt.Runtime) func() {
	return LoadLibs(
		r,
		base.LibLoader,
		packagelib.LibLoader,
		coroutine.LibLoader,
		stringlib.LibLoader,
		tablelib.LibLoader,
		mathlib.LibLoader,
		iolib.LibLoader,
		utf8lib.LibLoader,
		oslib.LibLoader,
		debuglib.LibLoader,
		golib.LibLoader,
		runtimelib.LibLoader,
	)
}
