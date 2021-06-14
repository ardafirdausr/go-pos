package storage

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type FileSystemStorage struct {
	Root string
	Host string
	Port string
}

func (fss FileSystemStorage) Save(file *multipart.FileHeader, dir string, filename string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	path := filepath.Join(fss.Root, "web", "storage", dir, filename)
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		log.Println(err)
		return "", err
	}

	// Create file
	dstFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dstFile.Close()

	// Copy File
	if _, err = io.Copy(dstFile, src); err != nil {
		return "", err
	}

	// Generate path (local)
	dirToPath := strings.ReplaceAll(dir, "\\", "/")
	paths := []string{"", "storage", dirToPath, filename}
	sotrageFilepath := strings.Join(paths, "/")
	return sotrageFilepath, nil
}