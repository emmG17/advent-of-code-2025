package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Boilerplate file reading
func check(e error) {
	if e != nil { panic(e) }
}

func getData(path , separator string) []string {
	cleanPath := filepath.Clean(path)
	bytes, err := os.ReadFile(cleanPath)
	check(err)
	return strings.Split(strings.TrimSpace(string(bytes)), separator)
}
