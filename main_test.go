package main

import (
	"testing"
	app "example.com/src"
)

func BenchmarkSimpleMain(b *testing.B) {
	app.MainSingle()
}
