package rogerdev_generator

import (
	"fmt"
	"math/rand"
)

type Generator struct {
}

var (
	alfabet_ABC = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	alfabet_abc = []rune("abcdefghijklmnopqrstuvwxyz")
	number      = "0123456789"
	space       = []rune(" ")
	price       = []int{10000, 20000, 30000, 40000, 50000, 60000, 70000, 80000, 90000, 100000, 15000, 25000, 35000, 45000, 55000, 65000, 75000, 85000, 95000, 150000}
)

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateName(lengthCharacter int, useSpace bool) string {
	if lengthCharacter < 5 {
		if lengthCharacter < 3 {
			lengthCharacter = 3
		}
		useSpace = false
	}
	// for i := 0; i < len(lengthCharacter); i++ {

	// }
	var divide int
	if useSpace {
		if lengthCharacter == 5 {
			divide = 3
		} else {
			divide = lengthCharacter / 2

		}

	}
	fmt.Println(divide)
	b := make([]rune, lengthCharacter)

	for i := range b {
		if i == 0 {
			b[i] = alfabet_ABC[rand.Intn(len(alfabet_ABC))]
		} else {
			if useSpace {
				if i == divide {
					b[i] = space[0]
				} else {
					if i == divide+1 {
						b[i] = alfabet_ABC[rand.Intn(len(alfabet_ABC))]

					} else {
						b[i] = alfabet_abc[rand.Intn(len(alfabet_abc))]

					}

				}
			} else {
				b[i] = alfabet_abc[rand.Intn(len(alfabet_abc))]

			}
		}

	}

	return string(b)
}

// func (g *Generator) GenerateEmail()  string {

// 	email:=
// return email
// }

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func (g *Generator) GenerateEmail() string {
	// ademawan1210@gmail.com
	b := make([]rune, 6)
	for i := range b {
		b[i] = alfabet_abc[rand.Intn(len(alfabet_abc))]
	}
	email := string(b)
	email += "@gmail.com"
	return email
}
func (g *Generator) GenerateString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func (g *Generator) RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func (g *Generator) GeneratePrice() int {

	priceIndex := rand.Intn(len(price))

	return price[priceIndex]
}
