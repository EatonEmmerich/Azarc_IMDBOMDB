package omdb

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// http://www.omdbapi.com/?apikey=43a201a
type Item struct {
    Title string `json:"Title"`
    Year string `json:"Year"` //toint
    Rated string `json:"Rated"`
    Released string `json:"Released"` //totime
    Runtime string `json:"Runtime"`
    Genre string `json:"Genre"` //toslice
    Director string `json:"Director"`
    Writer string `json:"Writer"`
    Actors string `json:"Actors"` //toslice
    Plot string `json:"Info"`
    Language string `json:"Language"`
    Country string `json:"Country"`
    Awards string `json:"Awards"`
    Poster string `json:"Poster"`
    //"Ratings":[{"Source":"Internet Movie Database","Value":"6.1/10"}],"Metascore":"N/A","imdbRating":"6.1","imdbVotes":"2,223","imdbID":"tt0000005","Type":"movie","DVD":"N/A","BoxOffice":"N/A","Production":"N/A","Website":"N/A","Response":"True"}
}

type Client struct{
    url string
    apikey string
    httpClient *http.Client
}

func New(apikey string) Client {
    return Client{
        url: "https://www.omdbapi.com",
        apikey: apikey,
        httpClient: &http.Client{},
    }
}

func NewWithURL(url string, apikey string) Client {
    return Client{
        url: url,
        apikey: apikey,
        httpClient: &http.Client{},
    }
}

func (c Client) Info(id string) (Item, error){
    res, err := c.httpClient.Get(fmt.Sprintf("%s/?apikey=%s&i=%s&", c.url ,c.apikey, id))
    if err != nil {
        return  Item{}, err
    }

    var item Item
    err = json.NewDecoder(res.Body).Decode(&item)
    return item, err
}