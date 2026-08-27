package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Define categories for common file extensions
var extensionMap = map[string]string{
	// Images
	".png":  "images",
	".jpg":  "images",
	".jpeg": "images",
	".gif":  "images",
	".webp": "images",
	".svg":  "images",

	// Videos
	".mp4":  "videos",
	".mov":  "videos",
	".mkv":  "videos",
	".avi":  "videos",
	".webm": "videos",

	// Documents
	".pdf":  "documents",
	".docx": "documents",
	".txt":  "documents",
	".xlsx": "documents",
	".pptx": "documents",

	// Audio
	".mp3":  "audio",
	".wav":  "audio",
	".flac": "audio",

	// Archives
	".zip": "archives",
	".tar": "archives",
	".gz":  "archives",
	".rar": "archives",
}

// MoveRecord keeps track of where a file started and where it went
type MoveRecord struct {
	OriginalPath string `json:"original_path"`
	NewPath      string `json:"new_path"`
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  Organize: go run main.go /path/to/folder")
		fmt.Println("  Revert:   go run main.go --revert /path/to/folder")
		return
	}

	// Check if the user is asking to revert
	if args[1] == "--revert" {
		if len(args) < 3 {
			fmt.Println("Please provide the target directory to revert. Example: go run main.go --revert /path/to/folder")
			return
		}
		revert(args[2])
		return
	}

	// Otherwise, organize normally
	organize(args[1])
}

func organize(targetDir string) {
	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	var moveRecords []MoveRecord

	for _, file := range files {
		// Skip directories, we only want to move files
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		// Do not accidentally organize our own undo log file!
		if fileName == ".ranger_undo.json" {
			continue
		}

		// Get the extension (e.g. ".png") and make it lowercase
		ext := strings.ToLower(filepath.Ext(fileName))
		if ext == "" {
			continue // Skip files without an extension
		}

		// Realize the kind of file it is
		category, known := extensionMap[ext]
		if !known {
			// If it's a dynamic/unknown file type, create a folder based on the extension name
			extName := strings.TrimPrefix(ext, ".")
			category = extName + "_files"
		}

		// Create category folder safely
		categoryDir := filepath.Join(targetDir, category)
		err = os.MkdirAll(categoryDir, 0755)
		if err != nil {
			fmt.Printf("Failed to create directory %s: %v\n", categoryDir, err)
			continue
		}

		oldPath := filepath.Join(targetDir, fileName)
		newPath := filepath.Join(categoryDir, fileName)

		// Move the file
		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Failed to move %s: %v\n", fileName, err)
		} else {
			fmt.Printf("Moved: %s -> %s/\n", fileName, category)
			
			// Record this successful move in our log
			moveRecords = append(moveRecords, MoveRecord{
				OriginalPath: oldPath,
				NewPath:      newPath,
			})
		}
	}

	// If we actually moved files, save the undo log
	if len(moveRecords) > 0 {
		saveUndoLog(targetDir, moveRecords)
		fmt.Println("Organization complete! (Undo log created)")
	} else {
		fmt.Println("No files needed organizing.")
	}
}

func revert(targetDir string) {
	undoPath := filepath.Join(targetDir, ".ranger_undo.json")
	
	// Read the undo log
	data, err := os.ReadFile(undoPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No undo log found in this directory. Nothing to revert.")
		} else {
			fmt.Printf("Error reading undo log: %v\n", err)
		}
		return
	}

	var moveRecords []MoveRecord
	err = json.Unmarshal(data, &moveRecords)
	if err != nil {
		fmt.Printf("Error parsing undo log: %v\n", err)
		return
	}

	successCount := 0
	for _, record := range moveRecords {
		// Try to move the file back to its original location
		err := os.Rename(record.NewPath, record.OriginalPath)
		if err != nil {
			fmt.Printf("Failed to revert %s: %v\n", record.NewPath, err)
		} else {
			fmt.Printf("Reverted: %s\n", filepath.Base(record.OriginalPath))
			successCount++
		}
	}

	fmt.Printf("\nRevert complete! Successfully reverted %d out of %d files.\n", successCount, len(moveRecords))
	
	// Delete the undo log so we don't accidentally revert twice
	os.Remove(undoPath)
}

func saveUndoLog(targetDir string, records []MoveRecord) {
	undoPath := filepath.Join(targetDir, ".ranger_undo.json")
	
	// Convert our records slice into nicely formatted JSON data
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Printf("Failed to create undo log data: %v\n", err)
		return
	}
	
	// Save it to a hidden file in the target directory
	err = os.WriteFile(undoPath, data, 0644)
	if err != nil {
		fmt.Printf("Failed to save undo log: %v\n", err)
	}
}