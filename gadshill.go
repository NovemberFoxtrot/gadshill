package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
	"unicode"
)

type Config struct {
	ConnectTimeout   time.Duration
	ReadWriteTimeout time.Duration
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)

		if err != nil {
			return nil, err
		}

		conn.SetDeadline(time.Now().Add(rwTimeout))

		return conn, nil
	}
}

func FetchURL(theurl string) string {
	var client *http.Client

	if proxy := os.Getenv("http_proxy"); proxy != `` {
		proxyUrl, err := url.Parse(proxy)

		if err != nil {
			fmt.Println(err)
		}

		transport := http.Transport{
			Dial:  TimeoutDialer(5*time.Second, 5*time.Second),
			Proxy: http.ProxyURL(proxyUrl),
		}

		client = &http.Client{Transport: &transport}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest(`GET`, theurl, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func main() {
	counter := make(map[string]int, 0)

	data := FetchURL(os.Args[1])

	var buffer []rune

	printed := false

	for _, r := range data {
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) || r == 'ー' {
			buffer = append(buffer, r)

			printed = false
		} else if printed != true {
			printed = true
			counter[string(buffer)] += 1
			buffer = make([]rune, 0)
		}
	}

	for k, v := range counter {
		fmt.Println(v, k)
	}
}
