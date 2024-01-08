package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ymsodev/braindump/editor"
	"github.com/ymsodev/braindump/store"
)

const dumpFileName = ".bd"

func printHelp() {
	fmt.Println("usage: bd <command>")
	fmt.Println("Commands:")
	fmt.Println("  add		Add a new entry")
	fmt.Println("  log		Show entry logs")
	fmt.Println("  help		Print help message")
}

func addEntry(path string) {
	s, err := editor.Run()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Write(path, s); err != nil {
		log.Fatal(err)
	}
}

func showLogs(path string) {
	entries, err := store.Read(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Content)
	}
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dumpPath := filepath.Join(homedir, dumpFileName)

	switch cmd := os.Args[1]; cmd {
	case "add":
		addEntry(dumpPath)
	case "log":
		showLogs(dumpPath)
	case "help":
		printHelp()
	default:
		log.Fatalf("invalid command: '%s'\n", cmd)
	}
}
