package image

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetImage(t *testing.T) {
	fileName := "/Users/fuasahi/coolq/data/image/F7EB61A22794FAC40DD057A6B0B14A86.jpg.cqimg"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error("Read File Error")
	}
	fmt.Println(string(file))
	var url string

	w := bufio.NewReader(bytes.NewBuffer(file))
	for true {

		s, err := w.ReadString('\n')
		if err != nil {
			break
		}

		strs := strings.Split(s, "=")
		if len(strs) < 2 {
			continue
		}
		if strs[0] == "url" {
			url = strings.Join(strs[1:], "=")
			fmt.Println("right is:\n" + url)
			break
		}
	}

}
