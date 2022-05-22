package pkg

import "os"

type Env struct {
	DbHost         string
	DbToken        string
	DbOrganization string
	DbBucket       string
}

func GetEnv() Env {
	return Env{
		DbHost:         os.Getenv("DB_HOST"),
		DbToken:        os.Getenv("DB_TOKEN"),
		DbOrganization: os.Getenv("DB_ORGANIZATION"),
		DbBucket:       os.Getenv("DB_BUCKET"),
	}
}
