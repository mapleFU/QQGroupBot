package image

import (
	"github.com/mapleFU/QQBot/qqbot/data/group/message"
	"fmt"
	"io/ioutil"
	"github.com/BurntSushi/toml"
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
	// 如果是 image，需要有这个字段
	fileLink := segment.Data.File

	file, err := ioutil.ReadFile(IMAGE_DIR + "/" + fileLink + ".cqimg")
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}
	var image Image
	if _, err := toml.Decode(string(file), &image); err != nil {
		fmt.Println(err.Error())
		return "", false
	}

	return image.url, true

}