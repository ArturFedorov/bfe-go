package design_in_memory_file_system

type FileSystem struct{}

func Constructor() FileSystem                                           { return FileSystem{} }
func (fs *FileSystem) Ls(path string) []string                          { return nil }
func (fs *FileSystem) Mkdir(path string)                                {}
func (fs *FileSystem) AddContentToFile(filePath string, content string) {}
func (fs *FileSystem) ReadContentFromFile(filePath string) string       { return "" }
