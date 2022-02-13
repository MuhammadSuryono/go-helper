package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func ListBucket(sess *session.Session) {
	svc := s3.New(sess)
	result, err := svc.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
func DownloadObject(sess *session.Session, bucket string, path string) []byte {

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}
	buff := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(sess)
	_, err := downloader.Download(buff, input)

	if err != nil {
		logrus.WithField("message", err.Error()).Error("Failed to download")
	}

	return buff.Bytes()
}
func PresignUrl(sess *session.Session, bucket string, path string) string {

	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	})
	urlStr, err := req.Presign(1 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	return urlStr
}
func IsObjectExist(sess *session.Session, bucket string, path string) bool {

	svc := s3.New(sess)
	output, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String("bucket_name"),
		Key:    aws.String("object_key"),
	})

	if err != nil {
		return false
	} else {
		fmt.Println(output.LastModified)
		return true
	}

}
