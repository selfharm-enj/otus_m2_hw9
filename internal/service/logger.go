package service

import (
	"context"
	"fmt"
	"time"

	"github.com/selfharm-enj/otus_m2_hw9/internal/repository"
)

func LogChanges(ctx context.Context) {
	var (
		lastFiles  int
		lastImages int
	)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentFiles, currentImages := repository.FilesImagesLen()
			if currentFiles != lastFiles {
				diff := currentFiles - lastFiles
				if diff > 0 {
					fmt.Printf("Added %d Files", diff)
					newFiles := repository.LastNFiles(diff)
					fmt.Printf(" %v\n", newFiles)
				} else if diff < 0 {
					fmt.Printf("Removed %d Files\n", -diff)
				}
				lastFiles = currentFiles
			}
			if currentImages != lastImages {
				diff := currentImages - lastImages
				if diff > 0 {
					fmt.Printf("Added %d Images", diff)
					newImages := repository.LastNImages(diff)
					fmt.Printf(" %v\n", newImages)
				} else if diff < 0 {
					fmt.Printf("Removed %d Images\n", -diff)
				}
				lastImages = currentImages
			}
		}
	}
}
