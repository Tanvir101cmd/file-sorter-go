package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var categories = map[string]string {
		// Images
		"jpg": "Images", "jpeg": "Images", "png": "Images", "dng": "Images",
		"gif": "Images", "bmp": "Images", "svg": "Images",
		"webp": "Images", "ico": "Images", "tiff": "Images",

		// Documents
		"pdf": "Documents", "doc": "Documents", "docx": "Documents",
		"txt": "Documents", "rtf": "Documents", "odt": "Documents",
		"md": "Documents", "ppt": "Documents", "pptx": "Documents",
		"xls": "Documents", "xlsx": "Documents",

		// Videos
		"mp4": "Videos", "avi": "Videos", "mov": "Videos",
		"mkv": "Videos", "wmv": "Videos", "flv": "Videos",
		"webm": "Videos", "m4v": "Videos", "mpg": "Videos",

		// Audio
		"mp3": "Audio", "wav": "Audio", "flac": "Audio",
		"aac": "Audio", "ogg": "Audio", "m4a": "Audio",
		"wma": "Audio", "mid": "Audio", "midi": "Audio",

		// Archives
		"zip": "Archives", "rar": "Archives", "7z": "Archives",
		"tar": "Archives", "gz": "Archives", "bz2": "Archives",
		"xz": "Archives", "iso": "Archives",

		// Code
		"go": "Code", "js": "Code", "py": "Code", "c": "Code",
		"java": "Code", "cpp": "Code", "html": "Code",
		"css": "Code", "php": "Code", "rb": "Code",
		"json": "Code", "xml": "Code", "yaml": "Code",
		"yml": "Code", "csv": "Code", "sql": "Code",

		// Applications
		"dmg": "Apps", "app": "Apps", "exe": "Apps", "sh": "Apps", "msi": "Apps",
		"pkg": "Installer", "deb": "Installer", "rpm": "Installer",
		}

func get_cur_dir() string {
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
	ext := filepath.Ext(filename)

	if len(ext) > 0 {
		ext = strings.ToLower(ext[1:]) // slicing out the extension name jpg rather than .jpg
	}

	if ext == "" {
		return "Other"
	}
	if category, found := categories[ext]; found {
		return category
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

func moveFile(source, dest string) bool {
	err := os.Rename(source, dest)
	if err != nil {
		fmt.Println("Error while moving file", err)
		return false
	}
	return true
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

		if moveFile(source, dest) {
			fmt.Printf("Moved %s to %s\n", filename, fileType)
			movedCount++
		}
	}
	return movedCount
}

func main() {
	dir := get_cur_dir()
	fmt.Println("Currently at", dir)

	files := getAllFiles(dir)
	fmt.Println(files)
	fmt.Println("Number of files: ", len(files))

	count := moveFilesToDir(dir, files)

	fmt.Printf("\nDone! Moved %d files.\n", count)
}
