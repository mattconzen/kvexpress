package kvexpress

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

func ReadFile(filepath string) string {
	dat, err := ioutil.ReadFile(filepath)
	check(err)
	return string(dat)
}

func SortFile(file string) string {
	lines := strings.Split(file, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

func WriteFile(data string, filepath string, perms int, direction string) {
	err := ioutil.WriteFile(filepath, []byte(data), os.FileMode(perms))
	check(err)
	log.Print(direction, ": file_wrote='true' location='", filepath, "' permissions='", perms, "'")
}

func CompareFilename(file string) string {
	compare := fmt.Sprintf("%s.compare", file)
	full_path := path.Join(path.Dir(file), compare)
	log.Print("in: file='compare' full_path='", full_path, "'")
	return full_path
}

func LastFilename(file string) string {
	last := fmt.Sprintf("%s.last", file)
	full_path := path.Join(path.Dir(file), last)
	log.Print("in: file='last' full_path='", full_path, "'")
	return full_path
}

func CheckLastFile(file string, perms int) {
	if _, err := os.Stat(file); err != nil {
		log.Print("in: Last File: ", file, " does not exist.")
		WriteFile("This is a blank file.\n", file, perms, "in")
	}
}