package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// flags

	// list
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listPath := listCmd.String("path", ".", "path to list files and directories")

	//search
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchPath := searchCmd.String("path", ".", "path to search files")
	searchQuery := searchCmd.String("query", ".", "Name/extension to search")

	// copy
	copyCmd := flag.NewFlagSet("copy", flag.ExitOnError)
	copySrc := copyCmd.String("src", "", "Source file/directory")
	copyDest := copyCmd.String("dest", "", "Destination path")

	// delete
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deletePath := deleteCmd.String("path", "", "File to delete")

	// parsing CLI argruments
	if len(os.Args) < 2 {
		fmt.Println("need 'list', 'search', 'copy', 'move', or 'delete' command")
		os.Exit(1)
	}

	// switching to the command
	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
		ListHandler(*listPath)
	case "search":
		searchCmd.Parse(os.Args[2:])
		SearchHandler(*searchPath, *searchQuery)
	case "copy":
		copyCmd.Parse(os.Args[2:])
		CopyHandler(*copySrc, *copyDest, "copy")
	case "move":
		copyCmd.Parse(os.Args[2:])
		CopyHandler(*copySrc, *copyDest, "move")
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		DeleteHandler(*deletePath)
	default:
		fmt.Println("no such command found. use 'list', 'search', 'copy', 'move', or 'delete'")
		os.Exit(1)
	}
}
