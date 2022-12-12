package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	TypeRegularFile = iota
	TypeDirectory
)

const TotalDiskSpace = 70000000
const UpgradeSpaceRequirement = 30000000

type File struct {
	Name     string
	Path     string
	FileType int
	Size     int
	Parent   *File
	Children []File
}

func getFileInDirectory(directory *File, fileName string) *File {
	for c := range directory.Children {
		if directory.Children[c].Name == fileName {
			return &directory.Children[c]
		}
	}

	return nil
}

func setDirectorySize(directory *File) {
	directory.Size = 0

	for c := range directory.Children {
		if directory.Children[c].FileType == TypeDirectory {
			setDirectorySize(&directory.Children[c])
		}

		directory.Size += directory.Children[c].Size
	}
}

func flattenTree(file *File) map[string]File {
	files := make(map[string]File)

	// Add the current file
	files[file.Path] = File{
		Name:     file.Name,
		Path:     file.Path,
		FileType: file.FileType,
		Size:     file.Size,
	}

	// Add all the children, recursively
	// Note: We may 'add' a file multiple times, however as the path is used
	// as the map key, they will only appear once
	for c := range file.Children {
		childFiles := flattenTree(&file.Children[c])

		for key, value := range childFiles {
			files[childFiles[key].Path] = File{
				Name:     value.Name,
				Path:     value.Path,
				FileType: value.FileType,
				Size:     value.Size,
			}
		}
	}

	return files
}

func main() {
	regularFileRegex := regexp.MustCompile(`^\d+\s+[a-z\.]+$`)

	// Define the root file as our filesystem must have one
	root := &File{
		Name:     "",
		Path:     "/",
		FileType: TypeDirectory,
		Size:     0,
	}

	currentFile := root

	scanner := bufio.NewScanner(os.Stdin)

	// Read in each line and convert into tree
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.Index(line, "$") == 0 {
			// Command
			command := strings.Fields(line)

			if command[1] == "cd" {
				// Change into a directory
				if command[2] == ".." {
					// Go up one level if we are not already at the top level
					if currentFile.Parent != nil {
						currentFile = currentFile.Parent
					}
				} else {
					// Change into a directory - create it if it does not already exist
					directoryName := command[2]

					if strings.Index(directoryName, "/") == 0 {
						// Absolute change of directory, find path
					} else {
						directory := getFileInDirectory(currentFile, directoryName)

						if directory != nil {
							currentFile = directory
						} else {
							filePath := currentFile.Path + "/" + directoryName

							// If there are two leading forward slashes, remove one
							if strings.Index(filePath, "//") == 0 {
								filePath = filePath[1:]
							}

							newDirectory := &File{
								Name:     directoryName,
								Path:     filePath,
								FileType: TypeDirectory,
								Size:     0,
								Parent:   currentFile,
							}

							currentFile.Children = append(currentFile.Children, *newDirectory)
							currentFile = newDirectory
						}
					}
				}
			} else if command[1] == "ls" {
				// Listing directory - we can skip this as we will process the contents
				// when we read the next lines
			} else {
				fmt.Println("Unexpected command: " + command[1])
				os.Exit(1)
			}
		} else {
			isDirectory := strings.Index(line, "dir") == 0
			isRegularFile := regularFileRegex.MatchString(line)

			if isDirectory || isRegularFile {
				fileParts := strings.Fields(line)

				// Assume we are working with a directory, then override
				// if we have a regular file
				fileSize := 0
				fileName := fileParts[1]
				fileType := TypeDirectory

				if isRegularFile {
					fileSize, _ = strconv.Atoi(fileParts[0])
					fileType = TypeRegularFile
				}

				// Check if file exists
				// This is necessary because we may have seen this file already,
				// for example if we have changed into a directory and run ls in its parent
				existingFile := getFileInDirectory(currentFile, fileName)

				if existingFile == nil {
					// Add file to this level of the tree
					filePath := currentFile.Path + "/" + fileName

					// If there are two leading forward slashes, remove one
					if strings.Index(filePath, "//") == 0 {
						filePath = filePath[1:]
					}

					newFile := File{
						Name:     fileName,
						Path:     filePath,
						FileType: fileType,
						Size:     fileSize,
						Parent:   currentFile,
					}

					currentFile.Children = append(currentFile.Children, newFile)
				}
			}
		}
	}

	// Recursively walk the tree and populate all directory sizes
	// Go all the way up to the top
	for ; currentFile.Parent != nil; currentFile = currentFile.Parent {
		// Intentionally empty body
	}

	// Set the directory sizes from the root down
	setDirectorySize(currentFile)

	currentFreeSpace := TotalDiskSpace - currentFile.Size
	extraSpaceRequired := UpgradeSpaceRequirement - currentFreeSpace

	// Flatten the tree into a map of path -> file so we can process it iteratively
	flattenedTree := flattenTree(currentFile)

	// Deleting the root directory will be definition free up enough space
	deleteDirectory := flattenedTree["/"]

	for _, file := range flattenedTree {
		if file.FileType == TypeDirectory && file.Size >= extraSpaceRequired && file.Size < deleteDirectory.Size {
			deleteDirectory = file
		}
	}

	fmt.Println(deleteDirectory.Size)
}
