package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/selfharm-enj/otus_m2_hw9/internal/model"
)

var (
	Files  []model.File
	Images []model.Image
	mu     sync.Mutex
)

const (
	filesFilePath  = "files.json"
	imagesFilePath = "images.json"
)

func InitData() error {
	if err := loadData(&Files, filesFilePath); err != nil {
		return err
	}
	if err := loadData(&Images, imagesFilePath); err != nil {
		return err
	}
	return nil
}

func loadData(slice interface{}, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(slice); err != nil {
		return err
	}
	return nil
}

func AddData(data model.IDReader) {
	mu.Lock()
	defer mu.Unlock()

	switch v := data.(type) {
	case *model.File:
		Files = append(Files, *v)
	case *model.Image:
		Images = append(Images, *v)
	}

	saveData(Files, filesFilePath)
	saveData(Images, imagesFilePath)
}

func saveData(slice interface{}, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(slice); err != nil {
		fmt.Printf("Error encoding data to %s: %v\n", filename, err)
	}
}

func FilesImagesLen() (int, int) {
	mu.Lock()
	defer mu.Unlock()
	return len(Files), len(Images)
}

func FilesImagesItems() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("Files: %v\t", Files)
	fmt.Printf("Images: %v\n\n", Images)
}

func LastNFiles(n int) []model.File {
	mu.Lock()
	defer mu.Unlock()

	if n <= 0 {
		return nil
	}
	if n > len(Files) {
		n = len(Files)
	}
	result := make([]model.File, n)
	copy(result, Files[len(Files)-n:])
	return result
}

func LastNImages(n int) []model.Image {
	mu.Lock()
	defer mu.Unlock()

	if n <= 0 {
		return nil
	}
	if n > len(Images) {
		n = len(Images)
	}
	result := make([]model.Image, n)
	copy(result, Images[len(Images)-n:])
	return result
}
