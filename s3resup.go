package s3resup

import (
	"errors"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// Uploader stores configurations for uploading
type Uploader struct {
	S3          s3iface.S3API
	Concurrency int
	PartSize    int64
}

// NewUploader creates an uploader
func NewUploader(c client.ConfigProvider, options ...func(*Uploader)) *Uploader {
	uploader := Uploader{
		Concurrency: 5,
		PartSize:    1024 * 1024 * 5,
		S3: s3.New(c),
	}
	for _, opt := range options {
		opt(&uploader)
	}	
	return &uploader
}

// Upload uploads the contents of reader to the s3 location specified by input.
// The upload is multi-part and will resume, preserving it's progress if the
// upload is interrupted for any reason.
func (u *Uploader) Upload(reader io.Reader, input s3.PutObjectInput) error {
	return errors.New("unimplemented")
}

// UploadFile uploads the contents of a file to the s3 location specified by input.
// The upload is multi-part and will resume, preserving it's progress if the
// upload is interrupted for any reason.
func (u *Uploader) UploadFile(name string, input s3.PutObjectInput) error {
	file, err := os.Open(name)
	defer file.Close()
	if err != nil {
		return err
	}
	return u.Upload(file, input)
}
