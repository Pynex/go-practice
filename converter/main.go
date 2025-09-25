package main

import (
	"fmt"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		fmt.Println("Выход...")
		os.Exit(0)
	}()

	fmt.Println("Введите число с префиксом (0x, 0o, 0b) или 'stop' для выхода")
	var input string
	for {
		fmt.Scanln(&input)

		input = strings.ToLower(input)
		if input == "stop" {
			return
		}

		i := new(big.Int)
		res, system := proccessNumber(input)
		if system == 0 {
			fmt.Println("unsuccessful convertation")
			continue
		} else {
			_, success := i.SetString(res, system)
			if !success {
				fmt.Println("unsuccessful convertation")
				continue
			} else {
				fmt.Println(i)
			}
		}
	}
}

func proccessNumber(hexStr string) (string, int) {
	if len(hexStr) < 3 {
		return hexStr, 0
	}
	prefix := hexStr[0:2]
	withoutPrefix := hexStr[2:]
	switch prefix {
	case "0x":
		return withoutPrefix, 16
	case "0o":
		return withoutPrefix, 8
	case "0b":
		return withoutPrefix, 2
	default:
		return hexStr, 0
	}
}
