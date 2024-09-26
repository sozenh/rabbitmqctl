package rocketmq

import (
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func NewMqAdmin(hosts []string, accessKey, secretKey string) (admin.Admin, error) {
	return admin.NewAdmin(
		admin.WithResolver(primitive.NewPassthroughResolver(hosts)),
		admin.WithCredentials(primitive.Credentials{AccessKey: accessKey, SecretKey: secretKey}),
	)
}
