package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func remove(deleteID string) error {
	buf, err := ioutil.ReadFile("fake.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}
	str := string(buf)
	Lines := strings.Split(str, "\n")
	var result []string
	for i, v := range Lines {
		linePieces := strings.Split(v, ",")
		if strings.Compare(deleteID, linePieces[0]) == 0 {
			result = append(Lines[0:i], Lines[i+1:]...)
			break
		}

	}
	str = strings.Join(result, "\n")
	err = ioutil.WriteFile("fake.csv", []byte(str), 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func main() {
	remove("1")
}
