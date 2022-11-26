package env

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(RootDir()+"\\.env") // for local
	// err := godotenv.Load(".env") // for docker
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
