package service

import (
	"fmt"
	"time"

	"github.com/selfharm-enj/otus_m2_hw9/internal/repository"
)

func LogChanges() {
	var (
		lastFiles  int
		lastImages int
	)

	for {
		time.Sleep(200 * time.Millisecond)
		currentFiles, currentImages := repository.FilesImagesLen()

		if currentFiles != lastFiles {
			fmt.Printf("Added %v Files\t", currentFiles-lastFiles)
			lastFiles = currentFiles
		}
		if currentImages != lastImages {
			fmt.Printf("Added %d Images\n", currentImages-lastImages)
			lastImages = currentImages
		}
		time.Sleep(1 * time.Second)
		repository.FilesImagesItems()
	}
}
