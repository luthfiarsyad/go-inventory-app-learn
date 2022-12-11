package helper

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk kembali ke menu
func BackHandler() {
	fmt.Print("Tekan enter untuk kembali ke menu")
	var back int
	fmt.Scanln(&back)
}

// This func returns order number + type order combined in 1 variable
func OrderNumberCombiner(typeOrder string, number string) string {
	return typeOrder + number
}

// This func returns 2 variable consist of type order  and number
func OrderNumberSeparator(orderNum string) (string, string) {
	typeOrder := orderNum[:2]
	number := orderNum[2:]
	return typeOrder, number
}

func NumberOrderIncrement(number string) string {
	res := "00000"
	n, _ := strconv.Atoi(number)
	n++
	if n < 10 {
		number = strconv.Itoa(n)
		number = res[:5-1] + number
	} else if n < 100 {
		number = strconv.Itoa(n)
		number = res[:5-2] + number
	} else if n < 1000 {
		number = strconv.Itoa(n)
		number = res[:5-3] + number
	} else if n < 10000 {
		number = strconv.Itoa(n)
		number = res[:5-4] + number
	} else {
		number = strconv.Itoa(n)
	}
	return number
}
