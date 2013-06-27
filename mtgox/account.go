package mtgox

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	BaseUrl = "https://data.mtgox.com/api/2/"
	InfoUrl = "BTCUSD/money/info"
)

type InfoFeed struct {
	Result string                 `json:"result"`
	Data   map[string]interface{} `json:"data"`
}

func GetInfo(apikey, secret string) {
	bytes := postRequest(apikey, secret, InfoUrl)

	feed := &InfoFeed{}

	err := json.Unmarshal(bytes, feed)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(feed)
}

func postRequest(apikey, secret, url string) []byte {
	request := getRequest(apikey, secret, url)
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err.Error())
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err.Error())
	}

	return bytes
}

func getRequest(apiKey, secretKey, path string) *http.Request {
	secret := decode64(secretKey)
	hash := hmac.New(sha512.New, secret)

	body := getNonce()
	message := []byte(path + "\000" + body)
	hash.Write(message)
	hmac := hash.Sum(nil)
	reader := strings.NewReader(body)

	request, err := http.NewRequest("POST", BaseUrl+path, reader)

	if err != nil {
		panic(err.Error())
	}

	request.Header.Add("User-Agent", "My-First-Trade-Bot")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Rest-Key", apiKey)
	request.Header.Add("Rest-Sign", encode64(hmac))

	return request
}

func decode64(data string) []byte {
	result, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		panic(err.Error())
	}

	return result
}

func encode64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func getNonce() string {
	return "nonce=" + strconv.Itoa(int(time.Now().Unix()*1000))
}
