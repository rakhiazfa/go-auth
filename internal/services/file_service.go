package services

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/rakhiazfa/vust-identity-service/api/dto/requests"
	"github.com/rakhiazfa/vust-identity-service/api/dto/responses"
	"github.com/rakhiazfa/vust-identity-service/pkg/utils"
	"github.com/spf13/viper"
	"net/http"
)

type FileService struct {
	client *resty.Client
}

func NewFileService() *FileService {
	client := resty.New()
	client.SetBaseURL(viper.GetString("services.file_service"))

	return &FileService{client}
}

func (s *FileService) UploadFile(payload requests.UploadFileReq) (*responses.FileRes, error) {
	var buffer bytes.Buffer
	var fileAPIRes responses.FileAPIRes

	writer, err := utils.StructToForm(&buffer, payload)
	if err != nil {
		return nil, err
	}

	res, err := s.client.R().
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetHeader("Authorization", "Bearer "+payload.AccessToken).
		SetBody(&buffer).
		SetResult(&fileAPIRes).
		Post("/files")
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusCreated {
		return nil, fmt.Errorf("upload failed with status: %s", res.Status())
	}

	return &fileAPIRes.File, nil
}

func (s *FileService) UpdateFile(id uuid.UUID, payload requests.UpdateFileReq) (*responses.FileRes, error) {
	var buffer bytes.Buffer
	var fileAPIRes responses.FileAPIRes

	writer, err := utils.StructToForm(&buffer, payload)
	if err != nil {
		return nil, err
	}

	res, err := s.client.R().
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetHeader("Authorization", "Bearer "+payload.AccessToken).
		SetBody(&buffer).
		SetResult(&fileAPIRes).
		Put("/files/" + id.String())
	if err != nil {
		return nil, err
	}

	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("upload failed with status: %s", res.Status())
	}

	return &fileAPIRes.File, nil
}
