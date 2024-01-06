package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"time"
)

func printHelp() {
	fmt.Println("usage: bd <command>")
	fmt.Println("Commands:")
	fmt.Println("  add		Add a new entry")
	fmt.Println("  log		Show entry logs")
	fmt.Println("  help		Print help message")
}

func addEntry() {
	s, err := runEditor()
	if err != nil {
		log.Fatal(err)
	}
	if err := writeEntry("bd.txt", s); err != nil {
		log.Fatal(err)
	}
}

type entry struct {
	Time int64 `json:"time"`
	Content string `json:"content"`
}

func writeEntry(path string, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	e := &entry{
		Time: time.Now().Unix(),
		Content: content,
	}
	dat, err := json.MarshalIndent(e, "", "")
	if err != nil {
		return err
	}
	f.Write(dat)
	f.WriteString("\r\n")
	return nil	
}

func showLogs() {
	fmt.Println("TODO: show logs!")
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch cmd := os.Args[1]; cmd {
		case "add":
			addEntry()
		case "log":
			showLogs()
		case "help":
			printHelp()
		default:
			log.Fatalf("invalid command: '%s'\n", cmd)
	}
}
