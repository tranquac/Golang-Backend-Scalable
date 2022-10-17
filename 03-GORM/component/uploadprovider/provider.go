package uploadprovider

import (
	"03-GORM/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
