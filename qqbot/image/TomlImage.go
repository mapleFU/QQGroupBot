package image

type image struct {
	md5 string	`toml:md5`
	width int64 `toml:width`
	height int64 `toml:height`
	size int64	`toml:size`
	url string	`toml:url`
	addtime string	`toml:addtime`
}

type Image struct {
	image image `toml:image`
}
