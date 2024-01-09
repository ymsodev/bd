package store

import (
	"bufio"
	"encoding/json"
	"os"
	"time"
)

type Entry struct {
	Time    int64  `json:"time"`
	Content string `json:"content"`
}

func Read(path string) ([]Entry, error) {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		if atEOF {
			return 0, nil, nil
		}
		for i := 0; i < len(data)-1; i++ {
			if data[i] == '\r' && data[i+1] == '\n' {
				return i + 2, data[:i], nil
			}
		}
		return len(data), data, bufio.ErrFinalToken
	})

	entries := []Entry{}
	
	var e Entry
	for s.Scan() {
		if err := json.Unmarshal(s.Bytes(), &e); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	
	return entries, nil
}

func Write(path string, content string) error {
	if len(content) == 0 {
		return nil
	}
	
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	e := Entry{
		Time:    time.Now().Unix(),
		Content: content,
	}
	dat, err := json.Marshal(&e)
	if err != nil {
		return err
	}
	f.Write(dat)
	f.WriteString("\r\n")
	return nil
}
