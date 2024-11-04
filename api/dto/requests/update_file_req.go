package requests

import "mime/multipart"

type UpdateFileReq struct {
	AccessToken string
	ServiceKey  *string               `form:"serviceKey" validate:"required,max=255"`
	BucketName  *string               `form:"bucketName" validate:"required,max=255"`
	Directory   *string               `form:"directory" validate:"required,directory,max=255"`
	File        *multipart.FileHeader `form:"file" validate:"required"`
}
