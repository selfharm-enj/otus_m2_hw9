package model

type IDReader interface {
	GetID() int
}

type Image struct {
	ID   int
	Path string
}

func (i *Image) GetID() int {
	return i.ID
}

type File struct {
	ID   int
	Path string
}

func (f *File) GetID() int {
	return f.ID
}
