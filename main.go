package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/siraiwaqarali/go-pexels-api/models"
)

func NewClient(token string) *models.Client {
	c := http.Client{}
	return &models.Client{Token: token, HC: c}
}

func main() {
	os.Setenv("PexelsToken", "Your-Api-Key")
	var TOKEN = os.Getenv("PexelsToken")

	var c = NewClient(TOKEN)

	// result, err := c.SearchPhotos("waves", 15, 1)
	// result, err := c.CuratedPhotos(15, 1)
	result, err := c.GetPhoto(3573351)
	// result, err := c.GetRandomPhoto()
	// result, err := c.SearchVideos("nature", 15, 1)
	// result, err := c.PopularVideos(15, 1)
	// result, err := c.GetRandomVideo()
	remainingRequestsInMonth := c.GetRemainingRequestsInMonth()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(result)
	fmt.Println("Remaining Requests in Month:", remainingRequestsInMonth)
}
