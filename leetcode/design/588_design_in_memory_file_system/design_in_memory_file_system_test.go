package design_in_memory_file_system

import "testing"

func equalStr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFileSystem(t *testing.T) {
	t.Run("mkdir and ls", func(t *testing.T) {
		fs := Constructor()
		fs.Mkdir("/a/b/c")
		got := fs.Ls("/a")
		want := []string{"b"}
		if !equalStr(got, want) {
			t.Errorf("Ls('/a') = %v, want %v", got, want)
		}
	})

	t.Run("addContent and readContent", func(t *testing.T) {
		fs := Constructor()
		fs.Mkdir("/a/b")
		fs.AddContentToFile("/a/b/file.txt", "hello")
		if got := fs.ReadContentFromFile("/a/b/file.txt"); got != "hello" {
			t.Errorf("ReadContentFromFile = %q, want %q", got, "hello")
		}
	})

	t.Run("ls root", func(t *testing.T) {
		fs := Constructor()
		got := fs.Ls("/")
		if len(got) != 0 {
			t.Errorf("Ls('/') = %v, want empty", got)
		}
		fs.Mkdir("/z")
		fs.Mkdir("/a")
		got = fs.Ls("/")
		want := []string{"a", "z"}
		if !equalStr(got, want) {
			t.Errorf("Ls('/') = %v, want %v", got, want)
		}
	})

	t.Run("nested dirs", func(t *testing.T) {
		fs := Constructor()
		fs.Mkdir("/a/b/c/d")
		got := fs.Ls("/a/b/c")
		want := []string{"d"}
		if !equalStr(got, want) {
			t.Errorf("Ls('/a/b/c') = %v, want %v", got, want)
		}
	})

	t.Run("append content", func(t *testing.T) {
		fs := Constructor()
		fs.AddContentToFile("/file", "hello")
		fs.AddContentToFile("/file", " world")
		if got := fs.ReadContentFromFile("/file"); got != "hello world" {
			t.Errorf("ReadContentFromFile = %q, want %q", got, "hello world")
		}
	})
}
