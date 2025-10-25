package service

import (
	"math/rand/v2"
	"time"

	"github.com/selfharm-enj/otus_m2_hw9/internal/model"
)

func GenerateData(ch chan<- model.IDReader) {
	for {
		time.Sleep(1 * time.Second)
		if rand.IntN(2) == 0 {
			ch <- &model.File{
				ID:   rand.IntN(10),
				Path: randPath(15),
			}
		} else {
			ch <- &model.Image{
				ID:   rand.IntN(10),
				Path: randPath(15),
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

func StartService(ch chan<- model.IDReader) {
	go GenerateData(ch)
}
