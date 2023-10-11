package filehandler

import (
	"context"
	"os"
)

type service struct {
}

type Service interface {
	OpenFile(ctx context.Context, filePath string) (*os.File, error)
	CloseFile(file *os.File) error
}

func New() Service {
	return &service{}
}

func (s service) OpenFile(ctx context.Context, filePath string) (*os.File, error) {
	return os.Open(filePath)
}

func (s service) CloseFile(file *os.File) error {
	return file.Close()
}
