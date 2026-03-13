# 588. Design In-Memory File System

**Difficulty:** Hard

**Topics:** Hash Table, String, Design, Trie, Sorting

## Description

Design a data structure that simulates an in-memory file system.

Implement the `FileSystem` class:

- `FileSystem()` Initializes the object of the system.
- `List<String> ls(String path)` If `path` is a file path, returns a list that only contains this file's name. If `path` is a directory path, returns the list of file and directory names in this directory, sorted in lexicographic order.
- `void mkdir(String path)` Makes a new directory according to the given path. The given directory path does not exist. If the middle directories in the path do not exist, you should create them as well.
- `void addContentToFile(String filePath, String content)` If `filePath` does not exist, creates that file containing given `content`. If `filePath` already exists, appends the given `content` to original content.
- `String readContentFromFile(String filePath)` Returns the content in the file at `filePath`.

## Examples

**Example 1:**

```
Input: ["FileSystem", "ls", "mkdir", "addContentToFile", "ls", "readContentFromFile"]
       [[], ["/"], ["/a/b/c"], ["/a/b/c/d", "hello"], ["/"], ["/a/b/c/d"]]
Output: [null, [], null, null, ["a"], "hello"]
```

## Constraints

- `1 <= path.length, filePath.length <= 100`
- `path` and `filePath` are absolute paths which always begin with `'/'`.
- The names of directories and files only contain lowercase letters.
- The paths to create directories are guaranteed not to exist.
- At most `300` calls will be made to `ls`, `mkdir`, `addContentToFile`, and `readContentFromFile`.

## Approach Hints

1. **Trie (prefix tree):** Each node represents a directory or file. Children stored in a map. Files have a content string.
2. **Nested hash maps:** Map of path string to directory/file metadata.
3. **Split path approach:** Split path by `/` and traverse/create nodes as needed.

## Related Problems

- [604. Design Compressed String Iterator](https://leetcode.com/problems/design-compressed-string-iterator/)
- [1166. Design File System](https://leetcode.com/problems/design-file-system/)

## Google Follow-ups

- How would you implement file permissions (read/write/execute)?
- How would you handle symbolic links?
- How would you add support for file deletion and moving files?
- How would you make this thread-safe for concurrent access?
