package random

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func GenerateRandomDates() error {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create("dates.txt")

	if err != nil {
		return err
	}
	defer file.Close()

	randLength := rand.Intn(100) + 1

	for i := 0; i < randLength; i++ {
		randDay := rand.Intn(28) + 1
		randMonth := rand.Intn(12) + 1
		randYear := rand.Intn(25) + 2000

		file.WriteString(fmt.Sprintf("%02d/%02d/%04d\n", randDay, randMonth, randYear))
	}

	return nil
}
