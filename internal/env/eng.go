package env

import (
	"fmt"
	"os"
)

func GetEnv(k string, fallback string ) string {
	v := os.Getenv(k)
	if v == "" {
		return fallback
	}
	fmt.Println("k ", k, " v ", v)
	return v
}