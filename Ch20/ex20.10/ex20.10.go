package main

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
}

func ReadFile(reader Reader) {
	c := reader.(Closer)
	c.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
