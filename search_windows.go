package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	var folder, kw string
	fmt.Print("Specify the domain: ")
	fmt.Scanln(&folder)

	fmt.Print("Specify your search string: ")
	fmt.Scanln(&kw)

	localPath, _ := os.Getwd()
	files, _ := os.Open(localPath + "/" + folder)
	defer files.Close()

	fileInfos, _ := files.Readdir(-1)

	fmt.Println("\nMatches: \n")
	for _, fileInfo := range fileInfos {
		file, _ := os.Open(localPath + "/" + folder + "/" + fileInfo.Name())
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if match, _ := regexp.MatchString(kw, scanner.Text()); match {
				if !strings.HasSuffix(fileInfo.Name(), "robots.txt") {
					fmt.Println(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fileInfo.Name(), ".txt", ""), "!!!", ":"), "§§", "?"), "£", "/"))
				} else {
					fmt.Println(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(fileInfo.Name(), "!!!", ":"), "§§", "?"), "£", "/"))
				}
				break
			}
		}
	}

	time.Sleep(10 * time.Second)
}
