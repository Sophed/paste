package util

import (
	"os"

	"github.com/google/uuid"
)

const STORAGE_DIR = "storage"

func Store(data []byte) (string, error) {
	id := uuid.New().String()
	err := os.WriteFile(STORAGE_DIR+"/"+id, data, 0644)
	return id, err
}
