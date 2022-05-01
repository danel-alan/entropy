package forms

import "mime/multipart"

type EntropyReport struct {
	File      *multipart.FileHeader `form:"file" binding:"required"`
	BlockSize uint64                `form:"block_size"`
}
