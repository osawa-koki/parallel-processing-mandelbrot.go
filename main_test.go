package main

import (
	app "example.com/src"
	"testing"
)

func BenchmarkSimpleMain(b *testing.B) {
	app.MainSingle()
}

func BenchmarkMultiMain(b *testing.B) {
	app.MainMulti()
}
