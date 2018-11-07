package query

import (
	"testing"
	"net/http"
	"fmt"
	"io/ioutil"
	"bufio"
	"bytes"
	"strings"
)

func TestHttpGet(t *testing.T)  {
	fileName := "/Users/fuasahi/coolq/data/image/F7EB61A22794FAC40DD057A6B0B14A86.jpg.cqimg"
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Error("Read File Error")
	}
	//fmt.Println(string(file))
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
			url = strings.TrimSpace(strings.Join(strs[1:], "="))
			fmt.Println("right is:\n" + url)
			break
		}
	}

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {

		fmt.Println("fill")
		url = "https://gchat.qpic.cn/gchatpic_new/583036751/2060440534-2361584284-F7EB61A22794FAC40DD057A6B0B14A86/0?vuin=3187545268&term=2"
		_, err2 := http.Get(url)
		if err2 != nil {
			t.Error("re-error" + err2.Error())
		} else  {
			t.Error("just one error" + err.Error())
		}
	}
	fmt.Println(resp.Body)
}