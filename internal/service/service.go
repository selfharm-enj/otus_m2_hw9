package service

import (
	"context"
	"math/rand/v2"

	"github.com/selfharm-enj/otus_m2_hw9/internal/model"
)

func GenerateData(ctx context.Context, ch chan<- model.IDReader) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var v model.IDReader
			if rand.IntN(2) == 0 {
				v = &model.File{
					ID:   rand.IntN(1000),
					Path: randPath(15),
				}
			} else {
				v = &model.Image{
					ID:   rand.IntN(1000),
					Path: randPath(15),
				}
			}
			select {
			case ch <- v:
			case <-ctx.Done():
				return
			}
		}
	}

}

func randPath(n int) string {
	letters := []rune("ababcdefghijklmnopqrstuvwxyzc")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.IntN(len(letters))]
	}
	return string(b)
}

func StartService(ctx context.Context, ch chan<- model.IDReader) {
	go GenerateData(ctx, ch)
}
