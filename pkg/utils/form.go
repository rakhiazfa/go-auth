package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
)

func StructToForm(w io.Writer, payload interface{}) (*multipart.Writer, error) {
	writer := multipart.NewWriter(w)

	v := reflect.ValueOf(payload)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("form")
		value := v.Field(i)

		if !value.IsValid() || (value.Kind() == reflect.Ptr && value.IsNil()) {
			continue
		}

		if fileHeader, ok := value.Interface().(*multipart.FileHeader); ok {
			if err := writeFileHeader(writer, tag, fileHeader); err != nil {
				return nil, err
			}
		} else {
			if err := writeUnknownValue(writer, tag, value.Interface()); err != nil {
				return nil, err
			}
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	return writer, nil
}

func writeFileHeader(writer *multipart.Writer, tag string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer func(file multipart.File) {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}(file)

	part, err := writer.CreateFormFile(tag, fileHeader.Filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(part, file); err != nil {
		return err
	}

	return err
}

func writeUnknownValue(writer *multipart.Writer, tag string, value interface{}) error {
	switch v := value.(type) {
	case string:
		if err := writer.WriteField(tag, v); err != nil {
			return err
		}
	case int, int8, int16, int32, int64:
		if err := writer.WriteField(tag, fmt.Sprintf("%d", v)); err != nil {
			return err
		}
	case float32, float64:
		if err := writer.WriteField(tag, fmt.Sprintf("%f", v)); err != nil {
			return err
		}
	case bool:
		if err := writer.WriteField(tag, fmt.Sprintf("%d", btoi(v))); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
