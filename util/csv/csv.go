package csv

import (
	"encoding/csv"
	"reflect"
	"strings"
)

func Parse(body string, v interface{}, f func(result interface{})) error {

	reader := csv.NewReader(strings.NewReader(strings.TrimSuffix(body, "\n")))
	reader.FieldsPerRecord = -1

	headers, err := reader.Read()
	if err != nil {
		return err
	}

	dataType := reflect.TypeOf(v)
	newData := reflect.New(dataType).Elem()

	const tag = "csv"
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		for i := 0; i < dataType.NumField(); i++ {
			f := dataType.Field(i)
			index := 0
			fieldName := f.Tag.Get(tag)
			for k, v := range headers {
				if v == fieldName {
					index = k
				}
			}
			newField := newData.FieldByName(f.Name)
			if !newField.IsValid() || !newField.CanSet() {
				continue
			}
			newField.Set(reflect.ValueOf(row[index]))
		}
		f(newData.Interface())
	}
	return nil
}
