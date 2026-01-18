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
| **Images** | jpg, jpeg, png, dng, gif, bmp, svg, webp, ico, tiff |
| **Documents** | pdf, doc, docx, txt, rtf, odt, md, ppt, pptx, xls, xlsx |
| **Videos** | mp4, avi, mov, mkv, wmv, flv, webm, m4v, mpg |
| **Audio** | mp3, wav, flac, aac, ogg, m4a, wma, mid, midi |
| **Archives** | zip, rar, 7z, tar, gz, bz2, xz, iso |
| **Code** | go, py, java, c, html, css, php, rb, json, xml, yaml, yml, csv, sql |
| **Apps** | app, sh, exe, msi|
| **Installer** | pkg, deb, rpm |

### Features yet to implement
- [ ] dry-run mode to preview what changes it’ll make
- [ ] Undo to revert back changes
- [ ] Deep scan to sort the files inside the subdirectories
- [ ] A progress bar
- [ ] Flags/arguments for more options
- [ ] Config file
