package main

import "runtime"

func compile() {
	defer func() {
		recover()
	}()

	runtime.GOMAXPROCS()
}