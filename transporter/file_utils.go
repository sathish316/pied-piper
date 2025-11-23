package transporter

import (
	"path/filepath"
	cp "github.com/otiai10/copy"
)

type FileUtils struct {}

func (f *FileUtils) CopyFile(srcFile, dstDir string) error {
	return cp.Copy(srcFile, dstDir)
}

func (f *FileUtils) CopyFiles(srcPattern, dstDir string) error {
	// Get all files matching the pattern
	files, err := filepath.Glob(srcPattern)
	if err != nil {
		return err
	}
	// Copy each file to the destination directory
	for _, file := range files {
		dstFile := filepath.Join(dstDir, filepath.Base(file))
		err = f.CopyFile(file, dstFile)
		if err != nil {
			return err
		}
	}
	return nil
}