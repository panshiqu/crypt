package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	n := make(map[string]int)

	for i := 0x4E00; i <= 0x9FFF; i++ {
		n[fmt.Sprintf("%c", i)] = i
	}

	f, err := os.Open("3500.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	m := make(map[string]int)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		k := scanner.Text()
		if v, ok := n[k]; ok {
			m[k] = v
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	data, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(data))
}
