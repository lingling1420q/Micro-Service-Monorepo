package utils

import "testing"

func TestUnZip(t *testing.T) {
	zipFile := "/Users/luxuze/xxx.zip"
	files, err := UnZip(zipFile)
	t.Log(files, err)
}
