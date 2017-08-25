package common

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"strings"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func LoadUrlsFromFile(filePath string) []string {
	var urls = []string{}
	file, _ := os.Open(filePath)
	defer file.Close()

	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		urls = append(urls, strings.TrimSpace(line))
	}
	return urls
}

func getRandom() int {
	for {
		ri := r.Intn(9)
		if ri > 0 {
			if ri == 1 {
				ri++
			}
			return ri
		} else {
			continue
		}
	}
}
