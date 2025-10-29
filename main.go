package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binaryToDecimal(binStr string) (int64, error) {

	for _, ch := range binStr {
		if ch != '0' && ch != '1' {
			return 0, fmt.Errorf("yanliz 0 ve 1 daxil edilmelidir")
		}
	}
	dec, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return 0, fmt.Errorf("cevirerken xeta bas verdi: %v", err)
	}

	if dec > 1000 {
		return 0, fmt.Errorf("netice 1000-den boyuk ola bilmez: %d", dec)
	}
	return dec, nil
}

func decimalToBinary(dec int64) (string, error) {
	if dec > 1000 {
		return "", fmt.Errorf("daxil edilen eded 1000-den boyuk ola bilmez: %d", dec)
	}

	if dec < 0 {
		return "", fmt.Errorf("menfi ededler qebul edilmir: %d", dec)
	}

	return strconv.FormatInt(dec, 2), nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Say Sistemleri Cevirici (Binary <-> Decimal) ===")
	fmt.Println("Secimler:")
	fmt.Println("1 - Binary->Decimal")
	fmt.Println("2 - Decimal ->Binary")
	fmt.Print("Seciminizi edin (1 veya 2): ")

	choiceInput, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(choiceInput)

	switch choice {
	case "1":
		fmt.Print("Binary ededi daxil edin (yanliz 0 ve 1): ")
		binInput, _ := reader.ReadString('\n')
		bin := strings.TrimSpace(binInput)

		result, err := binaryToDecimal(bin)
		if err != nil {
			fmt.Printf("Xeta: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Binary %s  ->  Decimal %d\n", bin, result)

	case "2":
		fmt.Print("Decimal ededi daxil edin (maksimum 1000): ")
		decInput, _ := reader.ReadString('\n')
		decStr := strings.TrimSpace(decInput)

		decVal, err := strconv.ParseInt(decStr, 10, 64)
		if err != nil {
			fmt.Printf("Xeta: duzgun decimal eded daxil edilmeyib: %v\n", err)
			os.Exit(1)
		}
		result, err := decimalToBinary(decVal)
		if err != nil {
			fmt.Printf("Xeta: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Decimal %d -> Binary %s\n", decVal, result)

	default:
		fmt.Println("Yanlis secim. Yanliz 1 veya 2 daxil edile biler.")
		os.Exit(1)
	}
}
