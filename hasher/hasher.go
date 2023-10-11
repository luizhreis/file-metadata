package hasher

import (
	"context"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"io"
	"os"
)

type service struct {
}

type Service interface {
	GetHashes(ctx context.Context, fileHandle *os.File) (string, error)
}

func New() Service {
	return &service{}
}

func (s service) GetHashes(ctx context.Context, fileHandle *os.File) (string, error) {
	sha512, err := getSha512(fileHandle)
	if err != nil {
		return "", err
	}
	hashes := fmt.Sprintf("[SHA512] %s", sha512)

	sha256, err := getSha256(fileHandle)
	if err != nil {
		return "", err
	}
	hashes += fmt.Sprintf("\n[SHA256] %s", sha256)

	return hashes, nil
}

func getSha512(fileHandle *os.File) (string, error) {
	sha512 := sha512.New()
	hash, err := getHash(sha512, fileHandle)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func getSha256(fileHandle *os.File) (string, error) {
	sha256 := sha256.New()
	hash, err := getHash(sha256, fileHandle)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func getHash(shaDigest hash.Hash, fileHandle *os.File) (string, error) {
	if _, err := io.Copy(shaDigest, fileHandle); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", shaDigest.Sum(nil)), nil
}
