package main

type File struct {
	name     string
	size     int
	isDir    bool
	parent   *File
	children map[string]*File
}

func (f *File) addNew(file *File) {
	file.parent = f
	f.children[f.name] = file
}

func newFile(name string, isDir bool, size int) *File {
	return &File{
		name:     name,
		size:     size,
		isDir:    isDir,
		parent:   nil,
		children: make(map[string]*File),
	}
}
