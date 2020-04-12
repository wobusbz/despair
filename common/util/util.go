package util

import "os"

func Exist(d string) bool {
	_, err := os.Stat(d)
	return err == nil && os.IsExist(err)
}
