package env

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
)

func Env(key string) string {
	err := godotenv.Load(RootDir() + ".env")
	if err != nil {
		return ""
	}
	return os.Getenv(key)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	pathSplited := strings.Split(d, "infraestructure")
	return pathSplited[0]
}
