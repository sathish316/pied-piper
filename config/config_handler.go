package config

import (
	"fmt"
	"os"
	"path/filepath"
	"io"
)

const SAMPLE_CONFIG_FILE_PATH = "./config/config.sample.yml"
const DEFAULT_CONFIG_DIR = ".pied-piper"
const DEFAULT_CONFIG_FILE = "config.yml"

type Config struct {
	Path string
	File string
}

func (c *Config) GetFilePath() string {
	return filepath.Join(os.Getenv("HOME"), DEFAULT_CONFIG_DIR, c.File)
}

func (c *Config) Init() {
	fmt.Println("Initializing config file at ", c.GetFilePath())
	os.MkdirAll(os.Getenv("HOME") + "/" + DEFAULT_CONFIG_DIR, 0755)
	c.copyFileContents(SAMPLE_CONFIG_FILE_PATH, c.GetFilePath())
}

func (c *Config) copyFileContents(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening src file: %v", err)
		return err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating dst file: %v", err)
		return err
	}
	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
	return nil
}