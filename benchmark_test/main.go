package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Benchmark struct {
	Month       string `json:"month"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

var test = []int{1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 2, 3, 1, 43, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}

func main() {
	data1 := &Benchmark{}

	data2 := &Benchmark{}
	data2.Description = "halo"
	fmt.Println(&data1, &data2, data1.Description, data2.Description)
	// file, _ := json.MarshalIndent(result, "", " ")

	// _ = ioutil.WriteFile("test.json", file, 0644)

	err := godotenv.Load(".env")
	handleError(err)

}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func Looping(x []int) {
	count := 0
	for i := 0; i < len(x); i++ {
		count = x[i] + count
	}
}
func LoopingWithRange(x []int) {
	count := 0
	for _, v := range x {
		count += v + count
	}
}

// func BenchmarkConcat(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		Concat("a", "b")
// 	}
// }

// func BenchmarkStringsBuilder(b *testing.B) {
// 	var builder strings.Builder

// 	for n := 0; n < b.N; n++ {
// 		builder.WriteString("a")
// 	}
// }

type TestEnv struct {
	Variable1  string
	Variable2  string
	Variable3  string
	Variable4  string
	Variable5  string
	Variable6  string
	Variable7  string
	Variable8  string
	Variable9  string
	Variable11 string
	Variable13 string
	Variable14 string
	Variable15 string
	Variable16 string
	Variable17 string
	Variable18 string
	Variable19 string
	Variable20 string
	Variable21 string
	Variable22 string
	Variable23 string
	Variable24 string
	Variable25 string
	Variable26 string
	Variable27 string
}

func EnvTest() {

	Variable1 := os.Getenv("TESTING1")
	Variable2 := os.Getenv("TESTING2")
	Variable3 := os.Getenv("TESTING3")
	Variable4 := os.Getenv("TESTING4")
	Variable5 := os.Getenv("TESTING5")
	Variable6 := os.Getenv("TESTING6")
	Variable7 := os.Getenv("TESTING7")
	Variable8 := os.Getenv("TESTING8")
	Variable9 := os.Getenv("TESTING9")
	fmt.Println(Variable1, Variable2, Variable3, Variable4, Variable5, Variable6, Variable7, Variable8, Variable9)
	// Variable11 := os.Getenv("TESTING10")
	// Variable13 := os.Getenv("TESTING11")
	// Variable14 := os.Getenv("TESTING12")
	// Variable15 := os.Getenv("TESTING13")
	// Variable16 := os.Getenv("TESTING14")
	// Variable17 := os.Getenv("TESTING15")
	// Variable18 := os.Getenv("TESTING16")
	// Variable19 := os.Getenv("TESTING17")
	// Variable20 := os.Getenv("TESTING18")
	// Variable21 := os.Getenv("TESTING19")
	// Variable22 := os.Getenv("TESTING20")
	// Variable23 := os.Getenv("TESTING21")
	// Variable24 := os.Getenv("TESTING22")
	// Variable25 := os.Getenv("TESTING23")
	// Variable26 := os.Getenv("TESTING24")
	// Variable27 := os.Getenv("TESTING25")
	// Variable28 := os.Getenv("TESTING26")
	// Variable29 := os.Getenv("TESTING27")
	// Variable30 := os.Getenv("TESTING28")

}
func EnvTestOneGetEnv(t *TestEnv) {
	fmt.Println(t.Variable1, t.Variable2, t.Variable3, t.Variable4, t.Variable5, t.Variable6, t.Variable7, t.Variable8, t.Variable9)

}
func NewEnvTestWithFunc() *TestEnv {
	testing := &TestEnv{}
	testing.Variable1 = "testing"
	testing.Variable2 = "testing"
	testing.Variable3 = "testing"
	testing.Variable4 = "testing"
	testing.Variable5 = "testing"
	testing.Variable6 = "testing"
	testing.Variable7 = "testing"
	testing.Variable8 = "testing"
	testing.Variable9 = "testing"
	// testing.Variable11 = "testing"
	// testing.Variable13 = "testing"
	// testing.Variable14 = "testing"
	// testing.Variable15 = "testing"
	// testing.Variable16 = "testing"
	// testing.Variable17 = "testing"
	// testing.Variable18 = "testing"
	// testing.Variable19 = "testing"
	// testing.Variable20 = "testing"
	// testing.Variable21 = "testing"
	// testing.Variable22 = "testing"
	// testing.Variable23 = "testing"
	// testing.Variable24 = "testing"
	// testing.Variable25 = "testing"
	// testing.Variable26 = "testing"
	// testing.Variable27 = "testing"
	return testing

}
