package image

import (
	"github.com/mapleFU/QQBot/qqbot/data/group/message"
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
	"bufio"
)

const IMAGE_DIR  = "/home/user/coolq/data/image"
// return the real url for the image

func GetImage(segment *message.MessageSegment) (string, bool) {
	if segment == nil {
		return "", false
	}
	if segment.Type != "image" {
		return "", false
	}
	fmt.Println("Hey")
	// 如果是 image，需要有这个字段
	fileLink := segment.Data.File
	fileName := IMAGE_DIR + "/" + fileLink + ".cqimg"
	fmt.Println("fileName " + fileName)
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read File Error")
		fmt.Println(err.Error())
		return "", false
	}

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
			// you should trim space
			url = strings.TrimSpace(strings.Join(strs[1:], "="))

			//fmt.Println("right is:\n" + url)
			break
		}
	}
	if url != "" {
		return url, true
	} else {
		return url, false
	}


}