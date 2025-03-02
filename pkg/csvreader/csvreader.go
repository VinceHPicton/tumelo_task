package csvreader

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

// ReadIgnoringHeader returns Read for a filepath but drops the first row, use this if the csv file has a header row you wish to ignore.
func ReadIgnoringHeader(filepath string) ([][]string, error) {
	data, err := Read(filepath)

	if len(data) == 0 {
		return [][]string{}, err
	}

	return data[1:], err
}

// Read takes a file path (including file extension) to a csv file and returns the data in the file
func Read(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data := [][]string{}

	for {
		record, err := reader.Read()

		if err != nil {			
			if errors.Is(err, io.EOF) {
				break
			}
			return [][]string{}, err
		}

		data = append(data, record)
	}

	return data, nil
}
