package open

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Save Struct
type Save struct {
	Word    string // word to find
	Try     int    // number of try used
	Letters []rune // letters found
}

// Open the os.Args[1] file and return the contains in an array string
func Open() []string {
	files := os.Args[1]
	file, err := os.Open(files)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()
	return lines
}

// open a file and return the contains in an array string
func OpenPath(Path string) []string {
	file, err := os.Open(Path)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()
	return lines
}

// check if a file exist
func CheckPath(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// create a save file
func CreateSAveFile() {
	f, err := os.Create("./ressources/save.json")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	fmt.Println("Save file created")
}

// unmarshal json and return the string struct
func UnmarshalJson(path string) Save {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var save Save
	err = decoder.Decode(&save)
	if err != nil {
		log.Fatal(err)
	}
	return save
}

// marshal json with the struct
func MarshalJson(path string, save Save) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(save)
	if err != nil {
		log.Fatal(err)
	}
}
