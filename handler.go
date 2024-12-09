package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ListHandler(path string) {
	err := filepath.Walk(path, func(filePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(filePath)
		return nil
	})
	if err != nil {
		fmt.Printf("Error listing files : %v\n", err)
	}
}

func SearchHandler(searchPath, searchQuery string) {
	if searchQuery == "" {
		fmt.Println("Query empty")
		return
	}
	filepath.Walk(searchPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(info.Name(), searchQuery) {
			fmt.Println(path)
		}
		return nil
	})

}

func CopyHandler(src, dest, action string) {
	if src == "" || dest == "" {
		fmt.Println("Source and destination path cannot be empty")
		return
	}
	info, err := os.Stat(src)
	if err != nil {
		fmt.Printf("Not Found : %v\n", err)
	}

	if info.IsDir() {
		filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			relativePath, _ := filepath.Rel(src, path)
			targetPath := filepath.Join(dest, relativePath)
			if info.IsDir() {
				return os.MkdirAll(targetPath, os.ModePerm)
			} else if action == "copy" {
				return CopyFile(path, targetPath)
			} else if action == "move" {
				return os.Rename(path, targetPath)
			}
			return nil
		})
	} else {
		if action == "copy" {
			err = CopyFile(src, dest)
		} else if action == "move" {
			err = os.Rename(src, dest)
		}
	}
	if err != nil {
		fmt.Printf("Error during %s operation: %v\n", action, err)
	} else {
		fmt.Printf("%s operation completed successfully\n", strings.Title(action))
	}
}

func CopyFile(src, dest string) error {
	inputFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	err = os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	if err != nil {
		return err
	}

	output, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, inputFile)
	if err != nil {
		return err
	}
	return nil
}

func DeleteHandler(src string) {
	if src == "" {
		fmt.Println("Path cannot be empty")
		return
	}
	var res string
	fmt.Printf("sure? you want to delete '%s'[Y/N]:", src)
	fmt.Scan(&res)

	if strings.ToLower(res) == "y" {
		err := os.RemoveAll(src)
		if err != nil {
			fmt.Printf("Err while deleting file : %v\n", err)
		} else {
			fmt.Println("Hurray file deleted")
		}

	} else if strings.ToLower(res) == "n" {
		fmt.Println("deletion cancelled")
	} else {
		fmt.Println("invalid input only response in Y/N")
	}
}
