

# QQGroupBot

By mwish

## Overview

This application allow user to manage QQ Group by QQ robot. Before using this application, you should have a robot QQ number. This application can help you search image/pull message or chating with robot. The application use web technology. It uses [coolq-http-api](https://github.com/richardchien/coolq-http-api) to intereact with QQ, uses golang http package [gin-gogic](https://github.com/gin-gonic/gin) to intereact with the coolq-http. The application uses [vue.js](http://cn.vuejs.org) as front-end frameworkand [Vuetify](https://vuetifyjs.com/zh-Hans/) library as UI framework.

## QuickStart

The program use:

* Node.js 10.5.0
* coolq-http-api 4.5(latest)
* Golang 1.10.3

The program use docker-compose to manage our application. But you should do some extra operations.

```bash
git clone https://github.com/mapleFU/QQGroupBot

cd QQGroupBot
```

In config directory, you can create `qq-number.ini` file, like:

```ini
[3187545268]
post_url = http://httpserver:8085
serve_data_files = yes
port=5700
post_message_format=array
```

And you can build your service

```
docker-compose up
```

After building, enter `localhost:9005`, 

![F8DA9798-7B82-4DB9-8EAA-7D7AE1332C61](/Users/fuasahi/Desktop/F8DA9798-7B82-4DB9-8EAA-7D7AE1332C61.png)

Enter it with `MAX8char` and login your qq number. Now the program is started.

![84CD1196-DF97-443E-B65F-DAEED33F8DED](/Users/fuasahi/Desktop/84CD1196-DF97-443E-B65F-DAEED33F8DED.png)

Now you can visit our http server at http://localhost:8083/

![669862C6-0CD2-45FD-B515-689E72EB1411](/Users/fuasahi/Desktop/669862C6-0CD2-45FD-B515-689E72EB1411.png)

Now you can manage your group.

## Web Apis

### [Coolq-http](https://cqhttp.cc/docs/4.5/#/)

#### Receiving

Coolq-http is a plugin for coolq, it report the event in coolq via http or websocket.

When we get a message in QQ, we will receive that:

```json
{
    "message": [
        {
            "type": "text",
            "data": {"text": "这是第一段"}
        },
        {
            "type": "face",
            "data": {"id": "111"}
        },
        {
            "type": "text",
            "data": {"text": "这是表情之后的一段"}
        }
    ]
}
```

The data is just like this:

| 上报类型  | 说明                       |
| --------- | -------------------------- |
| `message` | 收到消息                   |
| `notice`  | 群、讨论组变动等通知类事件 |
| `request` | 加好友请求、加群请求／邀请 |

The most import data is message, we can see the definition of it:

#### 上报数据

| 字段名         | 数据类型       | 可能的值                              | 说明                                                         |
| -------------- | -------------- | ------------------------------------- | ------------------------------------------------------------ |
| `post_type`    | string         | `message`                             | 上报类型                                                     |
| `message_type` | string         | `private`                             | 消息类型                                                     |
| `sub_type`     | string         | `friend`、`group`、`discuss`、`other` | 消息子类型，如果是好友则是 `friend`，如果从群或讨论组来的临时会话则分别是 `group`、`discuss` |
| `message_id`   | number (int32) | -                                     | 消息 ID                                                      |
| `user_id`      | number (int64) | -                                     | 发送者 QQ 号                                                 |
| `message`      | message        | -                                     | 消息内容                                                     |
| `raw_message`  | string         | -                                     | 原始消息内容                                                 |
| `font`         | number (int32) | -                                     | 字体                                                         |
| `sender`       | object         | -                                     | 发送人信息                                                   |

其中 `sender` 字段的内容如下：

| 字段名     | 数据类型       | 说明                                  |
| ---------- | -------------- | ------------------------------------- |
| `user_id`  | number (int64) | 发送者 QQ 号                          |
| `nickname` | string         | 昵称                                  |
| `sex`      | string         | 性别，`male` 或 `female` 或 `unknown` |
| `age`      | number (int32) | 年龄                                  |

So, when we receive a private chat like this, you may receive:

```json
{
    "time": 1515204254,
    "post_type": "message",
    "message_type": "private",
    "sub_type": "friend",
    "message_id": 12,
    "user_id": 12345678,
    "message": "你好～",
    "raw_message": "你好～",
    "font": 456,
    "sender": {
        "nickname": "小不点",
        "sex": "male",
        "age": 18
    }
}
```

And there are many different data types in coolq, There are “at”(when you are ‘@‘), “text”, “image”. “icon”.

If you receive an image, the message will like:

```
[CQ:image,file=1.jpg]
```

The image will be store in dir `/Users/fuasahi/coolq/data/image` as cqimg file like 

```
FC819E24213817F10778A29F5DBA3542.jpg.cqimg
```

```ini
[image]
md5=FC819E24213817F10778A29F5DBA3542
width=350
height=344
size=17456
url=https://gchat.qpic.cn/gchatpic_new/1961118034/2060440534-2815190621-FC819E24213817F10778A29F5DBA3542/0?vuin=3187545268&term=2
addtime=1541615696
```

so, you get get the image in `url` field.

### Send Message

You can send message by yourself via http.

![A1C4D209-DF78-41B7-B59F-A1C8DA3963B8](/Users/fuasahi/Desktop/A1C4D209-DF78-41B7-B59F-A1C8DA3963B8.png)



![E8D4F37A-FD7E-4935-B533-5F1E270974FD](/Users/fuasahi/Desktop/E8D4F37A-FD7E-4935-B533-5F1E270974FD.png)

The api is as:

### `/send_private_msg` 发送私聊消息

#### 参数

| 字段名        | 数据类型 | 默认值  | 说明                                                         |
| ------------- | -------- | ------- | ------------------------------------------------------------ |
| `user_id`     | number   | -       | 对方 QQ 号                                                   |
| `message`     | message  | -       | 要发送的内容                                                 |
| `auto_escape` | boolean  | `false` | 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 `message` 字段是字符串时有效 |

#### 响应数据

| 字段名       | 数据类型       | 说明    |
| ------------ | -------------- | ------- |
| `message_id` | number (int32) | 消息 ID |

### `/send_group_msg` 发送群消息

#### 参数

| 字段名        | 数据类型 | 默认值  | 说明                                                         |
| ------------- | -------- | ------- | ------------------------------------------------------------ |
| `group_id`    | number   | -       | 群号                                                         |
| `message`     | message  | -       | 要发送的内容                                                 |
| `auto_escape` | boolean  | `false` | 消息内容是否作为纯文本发送（即不解析 CQ 码），只在 `message` 字段是字符串时有效 |

#### 响应数据

| 字段名       | 数据类型       | 说明    |
| ------------ | -------------- | ------- |
| `message_id` | number (int32) | 消息 ID |

### 

So, we can use it api to intereact with qq via http.

### [trace.moe](https://github.com/soruly/trace.moe)

Image Reverse Search for Anime Scenes

Use anime screenshots to search where this scene is taken from.

It tells you which anime, which episode, and exactly which moment this scene appears in Japanese Anime.

## Search

Seach request should be POST as JSON or FORM

```json
POST https://trace.moe/api/search?token=your_api_token
Content-Type: application/json

{
  "image" : "data:image/jpeg;base64,/9j/4AAQSkZJ......"
}
```

| Fields | Value             | Notes                                |
| ------ | ----------------- | ------------------------------------ |
| image  | String (Required) | Base64 Encoded Image                 |
| filter | Number (Optional) | Limit search to specific anilist ID. |

Example Response

```json
{
  "RawDocsCount": 3555648,
  "RawDocsSearchTime": 14056,
  "ReRankSearchTime": 1182,
  "CacheHit": false,
  "trial": 1,
  "limit": 9,
  "limit_ttl": 60,
  "quota": 148,
  "quota_ttl": 85899,
  "docs": [
    {
      "from": 663.17,
      "to": 665.42,
      "anilist_id": 98444,
      "at": 665.08,
      "season": "2018-01",
      "anime": "搖曳露營",
      "filename": "[Ohys-Raws] Yuru Camp - 05 (AT-X 1280x720 x264 AAC).mp4",
      "episode": 5,
      "tokenthumb": "bB-8KQuoc6u-1SfzuVnDMw",
      "similarity": 0.9563952960290518,
      "title": "ゆるキャン△",
      "title_native": "ゆるキャン△",
      "title_chinese": "搖曳露營",
      "title_english": "Laid-Back Camp",
      "title_romaji": "Yuru Camp△",
      "mal_id": 34798,
      "synonyms": [
        "Yurucamp",
        "Yurukyan△"
      ],
      "synonyms_chinese": [],
      "is_adult": false
    }
  ]
}
```

| Fields            | Meaning                                                      | Value            |
| ----------------- | ------------------------------------------------------------ | ---------------- |
| RawDocsCount      | Total number of frames searched                              | Number           |
| RawDocsSearchTime | Time taken to retrieve the frames from database (sum of all cores) | Number           |
| ReRankSearchTime  | Time taken to compare the frames (sum of all cores)          | Number           |
| CacheHit          | Whether the search result is cached. (Results are cached by extraced image feature) | Boolean          |
| trial             | Number of times searched                                     | Number           |
| limit             | Number of search limit remaining                             | Number           |
| limit_ttl         | Time until limit resets (seconds)                            | Number           |
| quota             | Number of search quota remaining                             | Number           |
| quota_ttl         | Time until quota resets (seconds)                            | Number           |
| docs              | Search results (see table below)                             | Array of Objects |

| Fields           | Meaning                                                 | Value                                 |
| ---------------- | ------------------------------------------------------- | ------------------------------------- |
| from             | Starting time of the matching scene                     | Number (seconds, in 2 decimal places) |
| to               | Ending time of the matching scene                       | Number (seconds, in 2 decimal places) |
| at               | Exact time of the matching scene                        | Number (seconds, in 2 decimal places) |
| episode          | The extracted episode number from filename              | Number, "OVA/OAD", "Special", ""      |
| similarity       | Similarity compared to the search image                 | Number (float between 0-1)            |
| anilist_id       | The matching [AniList](https://anilist.co/) ID          | Number                                |
| mal_id           | The matching [MyAnimeList](https://myanimelist.net/) ID | Number or null                        |
| is_adult         | Whether the anime is hentai                             | Boolean                               |
| title_native     | Native (Japanese) title                                 | String or null (Can be empty string)  |
| title_chinese    | Chinese title                                           | String or null (Can be empty string)  |
| title_english    | English title                                           | String or null (Can be empty string)  |
| title_romaji     | Title in romaji                                         | String                                |
| synonyms         | Alternate english titles                                | Array of String or []                 |
| synonyms_chinese | Alternate chinese titles                                | Array of String or []                 |
| filename         | The filename of file where the match is found           | String                                |
| tokenthumb       | A token for generating preview                          | String                                |

![狗屎2](/Users/fuasahi/Downloads/狗屎2.png)

### rss

**RSS** (originally **RDF Site Summary**; later, two competing approaches emerged, which used the [backronyms](https://en.wikipedia.org/wiki/Backronym) **Rich Site Summary** and **Really Simple Syndication** respectively)[[2\]](https://en.wikipedia.org/wiki/RSS#cite_note-powers-2003-1-2) is a type of [web feed](https://en.wikipedia.org/wiki/Web_feed)[[3\]](https://en.wikipedia.org/wiki/RSS#cite_note-Netsc99-3) which allows users and applications to access updates to [online content](https://en.wikipedia.org/wiki/Online_content) in a standardized, computer-readable format. These feeds can, for example, allow a user to keep track of many different websites in a single [news aggregator](https://en.wikipedia.org/wiki/News_aggregator). The news aggregator will automatically check the RSS feed for new content, allowing the content to be automatically passed from website to website or from website to user. This passing of content is called [web syndication](https://en.wikipedia.org/wiki/Web_syndication). Websites usually use RSS feeds to publish frequently updated information, such as [blog](https://en.wikipedia.org/wiki/Blog) entries, news headlines, audio, video. RSS is also used to distribute [podcasts](https://en.wikipedia.org/wiki/Podcast). An RSS document (called "feed", "web feed",[[4\]](https://en.wikipedia.org/wiki/RSS#cite_note-GuardWF-4) or "channel") includes full or summarized text, and [metadata](https://en.wikipedia.org/wiki/Metadata), like publishing date and author's name.

In our program, I build a rsshub app http://rsshub.app. So that we can get message from weibo/pixiv/wechat/douban...

![狗屎1](/Users/fuasahi/Downloads/狗屎1.png)

### hitokoto

Hitokoto api is very easy...

| **参数名称**   | **类型**                                      | **描述**                                          |
| -------------- | --------------------------------------------- | ------------------------------------------------- |
| c              | 可选                                          | Cat，即类型。提交不同的参数代表不同的类别，具体： |
| a              | Anime - 动画                                  |                                                   |
| b              | Comic – 漫画                                  |                                                   |
| c              | Game – 游戏                                   |                                                   |
| d              | Novel – 小说                                  |                                                   |
| e              | Myself – 原创                                 |                                                   |
| f              | Internet – 来自网络                           |                                                   |
| g              | Other – 其他                                  |                                                   |
| 其他不存在参数 | 任意类型随机取得                              |                                                   |
| encode         | 可选                                          |                                                   |
| text           | 返回纯净文本                                  |                                                   |
| json           | 返回不进行unicode转码的json文本               |                                                   |
| js             | 返回指定选择器(默认.hitokoto)的同步执行函数。 |                                                   |
| 其他不存在参数 | 返回unicode转码的json文本                     |                                                   |
| charset        | 可选                                          |                                                   |
| utf-8          | 返回 UTF-8 编码的内容，支持与异步函数同用。   |                                                   |
| gbk            | 返回 GBK 编码的内容，不支持与异步函数同用。   |                                                   |
| callback       | 可选                                          |                                                   |
| 回调函数       | 将返回的内容传参给指定的异步函数。            |                                                   |

![狗屎](/Users/fuasahi/Downloads/狗屎.png)

#### Response

| **返回参数名称**                                     | **描述**                                                     |
| ---------------------------------------------------- | ------------------------------------------------------------ |
| id                                                   | 本条一言的id。 可以链接到https://hitokoto.cn?id=[id]查看这个一言的完整信息。 |
| hitokoto                                             | 一言正文。编码方式unicode。使用utf-8。                       |
| type                                                 | 类型。请参考第三节参数的表格。                               |
| from                                                 | 一言的出处。                                                 |
| creator                                              | 添加者。                                                     |
| created_at                                           | 添加时间。                                                   |
| 注意：如果encode参数为text，那么输出的只有一言正文。 |                                                              |

## Design and implementation

### 0. dockerlize

``` yaml
version: '3'

services:
  cqhttp:
    image: richardchien/cqhttp:latest
#    links:
#      - httpserver
    ports:
      - "9005:9000"
      - "5700:5700"
      - "5911:5911"
    environment:
      - COOLQ_ACCOUNT=3187545268
      - CQHTTP_POST_URL=http://httpserver:8085
      - CQHTTP_SERVE_DATA_FILES=yes
      - CQHTTP_USE_WS=yes
      - CQHTTP_USE_HTTP=yes
      - CQHTTP_POST_MESSAGE_FORMAT=array
    volumes:
#    you should set your config in this dir
      - ~/coolq:/home/user/coolq
      = ./config:/home/user/coolq/app/io.github.richardchien.coolqhttpapi/config
#      debug with host mode
#    network_mode: "host"
#    extra_hosts:
#      - "httpserver:101.132.121.41"
  httpserver:
    volumes:
      - ~/coolq:/home/user/coolq
      - ./log:/home/user/log
    links:
      - cqhttp
#    environment:
#      - GIN_MODE=release
    build: ./qqbot
    ports:
      - "8085:8085"
    depends_on:
      - cqhttp

  front-page:
    build:
      ./qqbot-frontend
    depends_on:
      - httpserver
    ports:
      - "8083:80"

```

We use docker-compose and docker to simplize our deployment. `cqhttp` is coolq http service, it provides service for QQ data.  `httpserver` is our golang logic server. It intereact with `cqhttp` to handle qq robot logic, and intereact with front-end to manage state of robot. `Front-page` is frontend server. It compile `vue` project and use nginx to show the page.

### 1. httpserver

#### Servicer and Manager

When we add functions like “image search”/“hitokoto”, we add a “servicer” in servicer manager. We have a manager to manage these services: 

```go
type Manager struct {
	serviceMap map[string]Servicer
	requester Requester.Requester
	receiver chan group.ChatResponseData
	strReceiver chan group.StringRespMessage
	// 管理的群组
	managedGroups []string
	// 同步的 lock
	serviceLock sync.Mutex
}
```

It can add and remove service. And you can believe that the `Manager` is thread-save.

```GO
unc (manager *Manager) AddService(servicer Servicer, name string)  {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()

	manager.serviceMap[name] = servicer
	servicer.SetOutchan(&manager.strReceiver)
	go servicer.Run()
}

func (manager *Manager) RemoveService(name string) bool {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()
	_, ok := manager.serviceMap[name]
	delete(manager.serviceMap, name)
	return ok
}
```

Also, we have different servicers, and the servicers are like:

```go
type Servicer interface {
	// Generate Message Channel
	GetChan() (out chan<- *group.ChatRequestData)
	IfAcceptMessage(Request *group.ChatRequestData) bool
	PutRequest(Request *group.ChatRequestData)
	SetOutchan(respChan* chan group.StringRespMessage)
	SendData(data *group.StringRespMessage)

//	logic
	Run()
    Stop()
}

type BaseServicer struct {
	InChan chan *group.ChatRequestData
	// it might be nil
	OutChan *chan group.StringRespMessage
}
```

So, the servicer can:

1. Decide whether this service receive the message
2. Accept the message.
3. Send data to QQ

So, when the program receive an request, it will ask every `servicer` in its list whether it can receive the request.

```go
func (manager *Manager) RecvRequest(request *group.ChatRequestData) {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()

	for k, v := range manager.serviceMap {
		if v.IfAcceptMessage(request) {
			fmt.Println("Call service " + k)
			v.PutRequest(request)
		}
	}
}
```

And when we add a servicer, we will start a goroutine to run it:

```go
func (manager *Manager) AddService(servicer Servicer, name string)  {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()

	manager.serviceMap[name] = servicer
	servicer.SetOutchan(&manager.strReceiver)
	// RUN THE PROGRAM
	go servicer.Run()
}
```

And when we remove it, the service will be closed.

```go
func (manager *Manager) RemoveService(name string) bool {
	manager.serviceLock.Lock()
	defer manager.serviceLock.Unlock()
	servicer, ok := manager.serviceMap[name]
	delete(manager.serviceMap, name)
	servicer.Stop()
	return ok
}
```

#### Servicer

As we talked before, we have “image search” “subscirbe rss” service. The service will be classified as:

1. query
2. subscribe

Subscribe like rss will refuse all the input, just run and output.

```go
func (*Subscribe) IfAcceptMessage(Request *group.ChatRequestData) bool {
	return false
}

func (self *WeiboService) Run() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(self.ServiceUrl)
	if err != nil {
		panic(err.Error())

	}
	newest := feed.Items[0]
	title := feed.Title
	
	// InChan 被 close 会终止
	for true {
		// 10 分钟一次
		time.Sleep(time.Minute * 10)
		feed, _ := fp.ParseURL(self.ServiceUrl)
		if feed.Items == nil {
			fmt.Println("Feed.Items is nil!")
		}

		for _, item := range feed.Items {
			if item == nil {
				fmt.Println("item is nil here")
			}
			if item.Title == newest.Title {
				newest = item
				break
			} else {
				Resp := buildService(item, title)
				*self.OutChan <- Resp
			}
		}
	}
}

```

It will run and fetch news.

#### Data

Golang use json and tag to parsing data, for example, as for the message we talked above, we can define the data structure like that:

```json
type ArrayRespMessage struct {
	GroupID string `json:"group_id"`
	Message message.Message `json:"message"`
	AutoEscape bool `json:"auto_escape"`
}

// 定义的 QQ 消息的数据结构, 表示单条的消息
type MessageData struct {
//	field for file
	File string `json:"file,omitempty"`
	Url string `json:"url,omitempty"`
//	field for text
	Text string `json:"text,omitempty"`
//	field for face
	Id string `json:"id,omitempty"`
//	field for at
	QQ string `json:"qq"`
}

type MessageSegment struct {
	Type string `json:"type,omitempty"`
	// may be message type ?
	Data MessageData `json:"data,omitempty"`
}

type Message []MessageSegment
```

Golang will automatically parsing the data for you.