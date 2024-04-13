package requests

import "mime/multipart"

type ParseMenu struct {
	MenuType string               `form:"menuType" binding:"required"`
	File     multipart.FileHeader `form:"file" binding:"required"`
}
