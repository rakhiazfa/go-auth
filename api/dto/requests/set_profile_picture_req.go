package requests

import "mime/multipart"

type SetProfilePictureReq struct {
	ProfilePicture *multipart.FileHeader `form:"profilePicture" validate:"required"`
}
