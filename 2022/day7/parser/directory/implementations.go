package directory

/*
 * Directory
 */

type directoryImpl struct {
	name string

	files          []File
	parent         Directory
	subdirectories []Directory
}

func newDirectory(name string, parent Directory) *directoryImpl {
	return &directoryImpl{
		name:           name,
		files:          []File{},
		parent:         parent,
		subdirectories: []Directory{},
	}
}

func (d *directoryImpl) Name() string {
	return d.name
}

func (d *directoryImpl) Files() []File {
	return d.files
}

func (d *directoryImpl) Parent() Directory {
	return d.parent
}

func (d *directoryImpl) Subdirectories() []Directory {
	return d.subdirectories
}

/*
 * File
 */

type fileImpl struct {
	name string
	size int
}

func newFile(name string, size int) File {
	return &fileImpl{
		name: name,
		size: size,
	}
}

func (f *fileImpl) Name() string {
	return f.name
}

func (f *fileImpl) Size() int {
	return f.size
}
