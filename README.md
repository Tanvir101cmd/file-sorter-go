# Simple File Sorter written in GO (WIP)
A CLI tool that organizes files in a folder based on their extensions (e.g., `jpg`, `png` etc to Images, `mp3`, `wav` to Music etc).

## Features (for now)
- Groups files into images, documents, videos, audio, archives and code
- Unknown extensions are placed inside Others directory
- Skips directories 

## Installation
1. Clone the repo:
```bash
git clone https://github.com/tanvir101cmd/file-sorter-go.git
cd file-sorter
```

2. Build the binary:
```bash
go build -o organizer main.go
```

## Usage
Run the executable and enter the path of the directory you want to organize:
```bash
./organizer
``` 

### Supported Extensions
| Category | Extensions |
| :--- | :--- |
| **Images** | jpg, png, svg, webp, etc. |
| **Documents** | pdf, docx, txt, md, xlsx, etc. |
| **Videos** | mp4, mov, mkv, etc. |
| **Audio** | mp3, wav, flac, etc. |
| **Archives** | zip, rar, 7z, tar, etc. |
| **Code** | go, js, py, html, css, etc. |

### Features yet to implement
- [ ] dry-run mode to preview what changes it’ll make
- [ ] Undo to revert back changes
- [ ] Deep scan to sort the files inside the subdirectories
- [ ] A progress bar
- [ ] Flags/arguments for more options
- [ ] Config file