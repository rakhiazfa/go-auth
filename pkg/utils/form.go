package utils

import (
	"io"
	"mime/multipart"
	"reflect"
)

func CreateFormFromStruct(w io.Writer, payload interface{}) (*multipart.Writer, error) {
	writer := multipart.NewWriter(w)

	v := reflect.ValueOf(payload)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("form")
		value := v.Field(i).Interface()

		if fileHeader, ok := value.(*multipart.FileHeader); ok {
			if err := writeFileHeader(writer, tag, fileHeader); err != nil {
				return nil, err
			}
		} else {
			if err := writer.WriteField(tag, value.(string)); err != nil {
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
