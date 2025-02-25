package csv_reader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
)

// ReadIgnoringHeader returns Read for a filepath but drops the first row, use this if the file has a header row you wish to ignore.
func ReadIgnoringHeader[T any](filepath string) ([]T, error) {
	data, err := Read[T](filepath)

	if len(data) == 0 {
		return data, fmt.Errorf("given file has no data")
	}

	return data[1:], err
}

// Read takes a file path (including file extension) and a pointer to the slice of structs you want the csv data to be loaded into
// each line of the CSV will be placed into a struct in the dataContainer slice
func Read[T any](filepath string) ([]T, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []T{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	data := []T{}

	for {
		record, err := reader.Read()

		if err != nil {			
			if errors.Is(err, io.EOF) {
				break
			}
			return []T{}, err
		}

		var entry T
		err = parseRecord(record, &entry)
		if err != nil {
			return nil, err
		}

		data = append(data, entry)
	}

	return data, nil
}

// parseRecord dynamically assigns CSV fields to struct fields based on struct order.
func parseRecord[T any](record []string, entry *T) error {
	v := reflect.ValueOf(entry).Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct, got %T", entry)
	}

	if len(record) != v.NumField() {
		return fmt.Errorf("expected %d fields, got %d", v.NumField(), len(record))
	}

	for i := 0; i < len(record); i++ {
		field := v.Field(i)
		if field.CanSet() {
			field.SetString(record[i])
		}
	}
	return nil
}