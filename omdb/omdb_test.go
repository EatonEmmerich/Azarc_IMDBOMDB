package omdb_test

import (
    "fmt"
    "math/rand"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/require"

    "Azarc/omdb"
)

func TestInfo(t *testing.T){
    apikey := fmt.Sprintf("apiKey%d", rand.Int())
    id := fmt.Sprintf("id%d", rand.Int())

    testServer := httptest.NewServer(
        http.HandlerFunc(
            func( writer http.ResponseWriter, req *http.Request){
                err := req.ParseForm()
                require.NoError(t, err)
                require.Equal(t, id, req.Form.Get("i"))
                require.Equal(t, apikey, req.Form.Get("apikey"))

                _,err = writer.Write([]byte(`{"Title":"Blacksmith Scene","Year":"1893","Rated":"Unrated","Released":"09 May 1893","Runtime":"1 min","Genre":"Short, Comedy","Director":"William K.L. Dickson","Writer":"N/A","Actors":"Charles Kayser, John Ott","Info":"Three men hammer on an anvil and pass a bottle of beer around.","Language":"None","Country":"USA","Awards":"1 win.","Poster":"https://m.media-amazon.com/images/M/MV5BNDg0ZDg0YWYtYzMwYi00ZjVlLWI5YzUtNzBkNjlhZWM5ODk5XkEyXkFqcGdeQXVyNDk0MDg4NDk@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"6.1/10"}],"Metascore":"N/A","imdbRating":"6.1","imdbVotes":"2,223","imdbID":"tt0000005","Type":"movie","DVD":"N/A","BoxOffice":"N/A","Production":"N/A","Website":"N/A","Response":"True"}`))
                require.NoError(t, err)
            },
        ),
    )

    client := omdb.NewWithURL(testServer.URL, apikey)

    item, err := client.Info(id)
    require.NoError(t, err)
    require.Equal(t,omdb.Item{
        Title: "Blacksmith Scene",
        Year: "1893",
        Rated: "Unrated",
        Released: "09 May 1893",
        Runtime: "1 min",
        Genre: "Short, Comedy",
        Director: "William K.L. Dickson",
        Writer: "N/A",
        Actors: "Charles Kayser, John Ott",
        Plot: "Three men hammer on an anvil and pass a bottle of beer around.",
        Language: "None",
        Country: "USA",
        Awards: "1 win.",
        Poster: "https://m.media-amazon.com/images/M/MV5BNDg0ZDg0YWYtYzMwYi00ZjVlLWI5YzUtNzBkNjlhZWM5ODk5XkEyXkFqcGdeQXVyNDk0MDg4NDk@._V1_SX300.jpg",
    }, item)
}
