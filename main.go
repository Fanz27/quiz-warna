package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var x bool
	var userGuess string
	var score int
	randomWarna := []string{"merah", "kuning", "hijau", "biru"}
	maxAttempts := 5
	attempts := 0

	fmt.Println("selamat datang digame tebak warna")
	fmt.Println("pilih warna yang benar")
	fmt.Println("selamat mencoba")

	for attempts < maxAttempts {
		random := rand.Intn(len(randomWarna))
		targetWarna := randomWarna[random]

		fmt.Println("warna apa ini", targetWarna)
		fmt.Println("Jawaban Anda")
		fmt.Scan(&userGuess)

		if userGuess == targetWarna {
			fmt.Println("benar. Lanjut")
			attempts++
			score++
		} else {
			fmt.Println("salah. Gameover")
			attempts += maxAttempts
		}
	}

	if score == maxAttempts {
		fmt.Println("Score anda", score, "/", maxAttempts)
		fmt.Println("Selamat Anda Memenangkan Game ini")
	} else {
		fmt.Println("Yahahah kalah")
	}

}
