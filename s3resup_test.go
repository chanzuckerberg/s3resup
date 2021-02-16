package s3resup

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestAbs(t *testing.T) {
    sess := session.Must(session.NewSession())
	uploader := NewUploader(sess)
    input := s3.PutObjectInput{}
    result := uploader.Upload(nil, input)
    if result.Error() != "unimplemented" {
        t.Error("Upload should not be implemented")
    }
}