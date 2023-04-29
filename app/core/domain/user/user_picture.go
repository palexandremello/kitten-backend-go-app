package users

import "errors"

type UserPicture struct {
	URL      string
	Data     []byte
	MimeType string
}

const maxFileSize = 10 << 20 // 10MB

var supportedMimeTypes = []string{
	"image/jpeg",
	"image/png",
}

func NewUserPicture(data []byte, mimeType string, filePath string) (*UserPicture, error) {

	if len(data) == 0 {
		return nil, errors.New("user picture data is empty")
	}

	if len(data) > maxFileSize {
		return nil, errors.New("Picture is too large: Max 10MB")
	}

	if !isValidMimeType(mimeType) {
		return nil, errors.New("unsupported file type")
	}

	return &UserPicture{
		URL:      filePath,
		Data:     data,
		MimeType: mimeType,
	}, nil
}

func isValidMimeType(mimeType string) bool {
	for _, t := range supportedMimeTypes {
		if t == mimeType {
			return true
		}
	}
	return false
}
