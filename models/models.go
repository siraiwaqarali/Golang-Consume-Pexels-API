package models

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/siraiwaqarali/go-pexels-api/constants"
)

type Client struct {
	Token          string
	HC             http.Client
	RemainingTimes int32
}

type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

type CuratedResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

type Photo struct {
	ID              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	URL             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerURL string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type VideoSearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}

type Video struct {
	ID            int32           `json:"id"`
	Width         int32           `json:"width"`
	Height        int32           `json:"height"`
	URL           string          `json:"url"`
	Image         string          `json:"image"`
	FullRes       interface{}     `json:"full_res"`
	Duration      float64         `json:"duration"`
	VideoFiles    []VideoFiles    `json:"video_files"`
	VideoPictures []VideoPictures `json:"video_pictures"`
}

type VideoFiles struct {
	ID       int32  `json:"id"`
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
	Link     string `json:"link"`
}

type VideoPictures struct {
	ID      int32  `json:"id"`
	Picture string `json:"picture"`
	Nr      int32  `json:"nr"`
}

type PopularVideos struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	URL          string  `json:"url"`
	Videos       []Video `json:"videos"`
}

func (c *Client) RequestDoWithAuth(method string, url string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", c.Token)
	response, err := c.HC.Do(request)
	if err != nil {
		return nil, err
	}

	time, err := strconv.Atoi(response.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return response, nil
	}

	c.RemainingTimes = int32(time)

	return response, nil
}

func (c *Client) SearchPhotos(query string, perPage int, page int) (*SearchResult, error) {
	fmt.Println("xxx============= Search Photos =============xxx")
	url := fmt.Sprintf(constants.PhotoApi+"/search?query=%s&per_page=%d&page=%d",
		query, perPage, page)

	response, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) CuratedPhotos(perPage int, page int) (*CuratedResult, error) {
	fmt.Println("xxx============= Curated Photos =============xxx")
	url := fmt.Sprintf(constants.PhotoApi+"/curated?per_page=%d&page=%d", perPage, page)

	response, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result CuratedResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetPhoto(id int32) (*Photo, error) {
	fmt.Println("xxx============= Get Photo =============xxx")
	url := fmt.Sprintf(constants.PhotoApi+"/photos/%d", id)

	response, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result Photo
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRandomPhoto() (*Photo, error) {
	fmt.Println("xxx============= Get Random Photo =============xxx")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Intn(1001)

	result, err := c.CuratedPhotos(1, randNum)
	if err == nil && len(result.Photos) == 1 {
		return &result.Photos[0], nil
	}
	return nil, err
}

func (c *Client) SearchVideos(query string, perPage int, page int) (*VideoSearchResult, error) {
	fmt.Println("xxx============= Search Video =============xxx")
	url := fmt.Sprintf(constants.VideoApi+"/search?query=%s&per_page=%d&page=%d",
		query, perPage, page)

	response, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result VideoSearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) PopularVideos(perPage int, page int) (*PopularVideos, error) {
	fmt.Println("xxx============= Popular Videos =============xxx")
	url := fmt.Sprintf(constants.VideoApi+"/popular?per_page=%d&page=%d", perPage, page)

	response, err := c.RequestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var result PopularVideos
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetRandomVideo() (*Video, error) {
	fmt.Println("xxx============= Random Video =============xxx")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum := rand.Intn(1001)

	result, err := c.PopularVideos(1, randNum)
	if err == nil && len(result.Videos) == 1 {
		return &result.Videos[0], nil
	}
	return nil, err
}

func (c *Client) GetRemainingRequestsInMonth() int32 {
	fmt.Println("xxx============= Remaining Requests =============xxx")
	return c.RemainingTimes
}
