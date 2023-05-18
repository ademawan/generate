package main

import (
	"testing"
)

// func BenchmarkLooping(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Looping(test)
// 	}
// }

// func BenchmarkEnvTest(b *testing.B) {

// 	for n := 0; n < b.N; n++ {
// 		LoopingWithRange(test)
// 	}
// }

// func BenchmarkTes(b *testing.B) {
// 	err := godotenv.Load(".env")
// 	handleError(err)
// 	for n := 0; n < b.N; n++ {
// 		EnvTest()
// 	}
// }

func BenchmarkEnvTestOneGetEnv(b *testing.B) {
	c := NewEnvTestWithFunc()
	for n := 0; n < b.N; n++ {
		EnvTestOneGetEnv(c)
	}
}
