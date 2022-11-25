package env

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// LoadEnv load env variables
func Env(key string) string {
	err := godotenv.Load(RootDir()+"\\.env")
	// err := godotenv.Load(".env")
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
