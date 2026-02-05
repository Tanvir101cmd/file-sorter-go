package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var categories = map[string][]string {
		// Images
		"Images": {"jpg", "jpeg", "png", "dng", "gif", "bmp", "svg", "webp", "ico", "tiff"},

		// Documents
		"Documents": {"pdf", "doc", "docx", "txt", "rtf", "odt", "md", "ppt", "pptx", "xls", "xlsx"},

		// Videos
		"Videos": {"mp4", "avi", "mov", "mkv", "wmv", "flv", "webm", "m4v", "mpg"},

		// Audio
		"Audio": {"mp3", "wav", "flac", "aac", "ogg", "m4a", "wma", "mid", "midi"},

		// Archives
		"Archives": {"zip", "rar", "7z", "tar", "gz", "bz2", "xz", "iso"},

		// Code
		"Code": {"go", "js", "py", "c", "java", "cpp", "html", "css", "php", "rb", "json", "xml", "yaml", "yml", "csv", "sql"},

		// Applications
		"Apps": {"dmg", "app", "exe", "sh", "msi"},
		"Installer": {"pkg", "deb", "rpm"},
		"Other": {"other"},
	}

func getCurrentDir() string {
	var dir string
	fmt.Print("Enter the directory to organize: ")
	fmt.Scanln(&dir)

	if dir == "" {
		dir = "."
	}

	dir = strings.TrimSpace(dir)
	return dir
}

func getAllFiles(dir string) []string {
	files_with_dir, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error while reading the directory", err)
		return []string{}
	}
	var files []string
	for _, file := range files_with_dir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	return files
}

func getFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))

	if ext == "" {
		return "Other"
	}
	for category, extensions := range categories {
		for _, extension := range extensions {
			if ext == "."+extension {
				return category
			}
		}
	}

	return ext + "_Files"
}

func createDir(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true
	}

	err = os.Mkdir(dirPath, 0755)
	if err != nil {
		fmt.Println("Error while creating a folder", err)
		return false
	}
	fmt.Println("Created folder:", filepath.Base(dirPath))
	return true
}

func moveFile(source, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return fmt.Errorf("destination file %s already exists", dest)
	}
	return os.Rename(source, dest)
}

func moveFilesToDir(dir string, files []string) int {
	movedCount := 0

	for _, filename := range files {
		fileType := getFileType(filename)

		typeDir := filepath.Join(dir, fileType)

		if !createDir(typeDir) {
			continue
		}

		source := filepath.Join(dir, filename)
		dest := filepath.Join(typeDir, filename)

		if err := moveFile(source, dest); err == nil {
			fmt.Printf("Moved %s to %s\n", filename, fileType)
			movedCount++
		} else {
			fmt.Printf("Failed to move %s: %v\n", filename, err)
		}
	}
	return movedCount
}

func main() {
	dir := getCurrentDir()
	fmt.Println("Currently at", dir)

	files := getAllFiles(dir)
	fmt.Println(files)
	fmt.Println("Number of files: ", len(files))

	count := moveFilesToDir(dir, files)

	fmt.Printf("\nDone! Moved %d files.\n", count)
}
