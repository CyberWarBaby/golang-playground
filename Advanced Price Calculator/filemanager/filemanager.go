package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed to read line in file")
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
		return errors.New("could not create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil

}
