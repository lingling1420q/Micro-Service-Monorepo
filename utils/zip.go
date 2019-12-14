package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"monaco/config"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// UnZip 解压压缩包中文件到tmp目录并返回文件绝对路径
func UnZip(zipFile string) ([]string, error) {
	var fileNames []string
	tmpFolder := config.Config.TmpFolder

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return fileNames, err
	}
	defer r.Close()

	for _, f := range r.File {
		// Store filename/path for returning and using later on
		fpath := filepath.Join(tmpFolder, f.Name)
		if macOsxDir(fpath) || f.FileInfo().IsDir() {
			continue
		}

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(tmpFolder)+string(os.PathSeparator)) {
			return fileNames, fmt.Errorf("%s: illegal file path", fpath)
		}

		fileUUID := AbsoluteFileName(path.Base(f.Name))
		fileNames = append(fileNames, fileUUID)

		outFile, e := os.OpenFile(fileUUID, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if e != nil {
			return fileNames, e
		}

		src, e := f.Open()
		if e != nil {
			return fileNames, e
		}

		_, e = io.Copy(outFile, src)
		if e != nil {
			return fileNames, e
		}

		// Close the file without defer to close before next iteration of loop
		e = outFile.Close()
		if e != nil {
			return fileNames, e
		}

		e = src.Close()

		if e != nil {
			return fileNames, e
		}
	}
	return fileNames, nil
}
