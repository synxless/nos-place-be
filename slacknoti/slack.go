package slacknoti

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func send(text string, endpoint string) {
	content := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}
	contentBytes, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
		return
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Post(endpoint, "application/json", bytes.NewReader(contentBytes))
	if resp.Status != "200" || err != nil {
		log.Println(err)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body), resp.Status)
	}
	defer resp.Body.Close()
}

type SlackNoti struct {
	Endpoint  string
	notiChan  chan string
	notiArray []string
	notiLock  sync.Mutex
}

func (sl *SlackNoti) StartSlackChan() {
	sl.notiChan = make(chan string)
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case noti := <-sl.notiChan:
			sl.notiLock.Lock()
			sl.notiArray = append(sl.notiArray, noti)
			sl.notiLock.Unlock()
		case <-t.C:
			sl.notiLock.Lock()
			if len(sl.notiArray) > 0 {
				texts := strings.Join(sl.notiArray, "\n")
				send(texts, sl.Endpoint)
				sl.notiArray = []string{}
			}
			sl.notiLock.Unlock()
		}
	}
}

func (sl *SlackNoti) SendSlackNoti(msg string) {
	sl.notiChan <- msg
}

func SendWithCustomChannel(text string, channel string) {
	content := struct {
		Text string `json:"text"`
	}{
		Text: text,
	}
	contentBytes, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
		return
	}
	httpClient := http.DefaultClient
	resp, err := httpClient.Post(channel, "application/json", bytes.NewReader(contentBytes))
	if resp.Status != "200" || err != nil {
		log.Println(err)
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	}
	defer resp.Body.Close()
}
