package main

type File struct {
	Name     string
	Size     int
	IsDir    bool
	Parent   *File
	Children map[string]*File
}

func (f *File) addNew(file *File) {
	file.Parent = f
	f.Children[file.Name] = file
}

func (f *File) getParent() *File {
	return f.Parent
}

func (f *File) getChild(name string) *File {
	return f.Children[name]
}

func (f *File) getChildren() []*File {
	ch := []*File{}
	for _, v := range f.Children {
		ch = append(ch, v)
	}
	return ch
}

func newFile(name string, isDir bool, size int) *File {
	return &File{
		Name:     name,
		Size:     size,
		IsDir:    isDir,
		Parent:   nil,
		Children: make(map[string]*File),
	}
}
