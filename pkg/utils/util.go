package utils

import (
	"github.com/sirupsen/logrus"
)

func CheckErr(err error) {
	if err != nil {
		logrus.WithError(err).Fatal("exit with error")
	}
}

func BrokerRole(id string) string {
	if id == "0" {
		return "master"
	}
	return "slave"
}
