package fullyear

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"sync"
	"time"
)

type FullYearClientResponse struct {
	Year        int    `json:"year"`
	YearString  string `json:"year_string"`
	SponsoredBy string `json:"sponsored_by"`
}

type GetFullYearClient struct {
	lock *sync.Mutex
}

func NewGetFullYearClient() *GetFullYearClient {
	return &GetFullYearClient{
		lock: &sync.Mutex{},
	}
}

func (c *GetFullYearClient) GetFullYear() (*FullYearClientResponse, error) {
	c.lock.Lock()

	result := FullYearClientResponse{}
	client := resty.New()
	_, err := client.R().SetResult(&result).Get("https://getfullyear.com/api/year")

	if err != nil {
		return nil, err
	}

	fmt.Println("Get full year is sponsored by: " + result.SponsoredBy)
	go func() {
		time.Sleep(time.Second * 1)
		c.lock.Unlock()
	}()

	return &result, nil
}
