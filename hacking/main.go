package main

import (
	"fmt"
	"time"
)

func main() {
	Find4digit()
}

func Find4digit() {
	abjad := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hasil := ""
	find := "as10"
	// panjangCharacter := 4
	timeNow := time.Now()
	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(abjad); j++ {
			for k := 0; k < len(abjad); k++ {
				for l := 0; l < len(abjad); l++ {

					hasil = string(abjad[i]) + string(abjad[j]) + string(abjad[k]) + string(abjad[l])
					// fmt.Println(hasil)
					fmt.Println(hasil)
					if hasil == find {

						fmt.Println("KETEMU", hasil)
						fmt.Println(time.Since(timeNow))
						return
					}

				}
			}
		}
	}
}
func Find10digit() {
	abjad := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	hasil := ""
	find := "9s10aa8343"
	// panjangCharacter := 4
	timeNow := time.Now()
	for i := 0; i < len(abjad); i++ {
		for j := 0; j < len(abjad); j++ {
			for k := 0; k < len(abjad); k++ {
				for l := 0; l < len(abjad); l++ {
					for m := 0; m < len(abjad); m++ {
						for n := 0; n < len(abjad); n++ {
							for o := 0; o < len(abjad); o++ {
								for p := 0; p < len(abjad); p++ {
									for q := 0; q < len(abjad); q++ {
										for r := 0; r < len(abjad); r++ {
											hasil = string(abjad[i]) + string(abjad[j]) + string(abjad[k]) + string(abjad[l]) + string(abjad[m]) + string(abjad[n]) + string(abjad[o]) + string(abjad[p]) + string(abjad[q]) + string(abjad[r])
											// fmt.Println(hasil)
											fmt.Println(hasil)
											if hasil == find {
												fmt.Println("KETEMU", hasil)
												fmt.Println(time.Since(timeNow))
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
