package imdb

type Item struct {

}

type Client struct {
    path string
}

func New(path string) Client{
    return Client{
        path: path,
    }
}

func (c Client)List() ([]Item, error) {
    return []Item{}, nil
}