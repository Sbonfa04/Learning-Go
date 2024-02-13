package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	temperatureFrequenze := make(map[int]int)
	tempMin, tempMax := 0, 0

	freqMax := 0

	for scanner.Scan() {
		singleNumbers := strings.Split(scanner.Text(), " ")

		for i := 0; i < len(singleNumbers); i++ {
			temp, _ := strconv.Atoi(singleNumbers[i])

			//memorizzo le temperature nella mappa
			_, check := temperatureFrequenze[temp] 
			if !check {
				temperatureFrequenze[temp] = 1
			} else {
				temperatureFrequenze[temp]++
			}

			if temp < tempMin {
				tempMin = temp //trovo temperatura min
			}

			if temp > tempMax {
				tempMax = temp //trovo temperatura max
			}

			if temperatureFrequenze[temp] > freqMax {
				freqMax = temperatureFrequenze[temp] //trovo frequenza massima
			}
		}
	}

	fmt.Println("max:", tempMax, temperatureFrequenze[tempMax], "volte")
	fmt.Println("min:", tempMin, temperatureFrequenze[tempMin], "volte")

	for key, value := range temperatureFrequenze {
		fmt.Print(key, ":", value, " ")
	}

	fmt.Print("\ntemperature ")

	for key, value := range temperatureFrequenze {
		if value == freqMax {
			fmt.Print(key, " ")
		}
	}

	fmt.Println("con frequenza ", freqMax)
}
