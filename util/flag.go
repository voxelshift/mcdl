package util

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func UnwrapFlag[F interface{}](f F, err error) F {
	if err != nil {
		log.Fatalf("failed to unwrap flag: %v", err)
		os.Exit(1)
	}

	return f
}
