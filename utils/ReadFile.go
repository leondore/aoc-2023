package utils

import "os"

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	CheckError(err)

	return string(data)
}
