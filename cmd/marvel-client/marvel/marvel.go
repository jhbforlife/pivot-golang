package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var CharBaseURL = "https://gateway.marvel.com/v1/public/characters"
var publicKey, privateKey = getKeys()
var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func getKeys() (public, private string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	pub := os.Getenv("MARVEL_PUBLIC_KEY")
	priv := os.Getenv("MARVEL_PRIVATE_KEY")
	return pub, priv
}

type Client struct {
	baseURL    string
	publicKey  string
	privateKey string
	httpClient *http.Client
}

func NewClient(url string) Client {
	return Client{url, publicKey, privateKey, httpClient}
}

func (c *Client) getHash(t int64) string {
	ts := strconv.FormatInt(t, 10)
	hash := md5.Sum([]byte(ts + c.privateKey + c.publicKey))
	return hex.EncodeToString(hash[:])
}

func (c *Client) signURL(url string) string {
	t := time.Now().Unix()
	hash := c.getHash(t)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, t, c.publicKey, hash)
}

func (c *Client) GetCharsWithLimit(l int) ([]CharResults, error) {
	url := c.baseURL + fmt.Sprintf("?limit=%d", l)
	url = c.signURL(url)

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var charResponse CharHTTPResponse
	if err := json.NewDecoder(res.Body).Decode(&charResponse); err != nil {
		return nil, err
	}

	return charResponse.Data.Results, nil
}

type CharHTTPResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int           `json:"offset"`
		Limit   int           `json:"limit"`
		Total   int           `json:"total"`
		Count   int           `json:"count"`
		Results []CharResults `json:"results"`
	} `json:"data"`
}

type CharResults struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
