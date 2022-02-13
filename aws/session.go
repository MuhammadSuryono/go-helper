package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
)

var Session *session.Session

type sessionHandler struct {
	Session *session.Session
}

func CreateAwsSession(accessKeyId, secreteKey, region string) {
	logrus.WithField("REGION", region).WithField(
		"ACCESS_KEY", accessKeyId).WithField(
		"SECRET_KEY", secreteKey).Info(
		"AWS Configuration")
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyId,
			secreteKey,
			""),
	})
	if err != nil {
		fmt.Printf("Unable to list buckets, %v", err)
	}

	Session = sess
	_ = sessionHandler{Session: sess}
}
