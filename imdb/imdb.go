package imdb

import (
    "bufio"
    "context"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
    "sync"
)

type Item struct {
    TConst         string
    TitleType      string
    PrimaryTitle   string
    OriginalTitle  string
    IsAdult        int
    StartYear      int
    EndYear        int
    RuntimeMinutes int
    Genres         []string
}

type Client struct {
    path       string
    goroutines int
}

func New(path string, goroutines int) (Client, error) {
    file, err := os.Open(path)
    if err != nil {
        return Client{}, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    ok := scanner.Scan()
    if !ok {
        return Client{}, scanner.Err()
    }

    header := strings.Split(scanner.Text(), "\t")
    if len(header) != 9 {
        errorString := fmt.Sprintf("invalid headers for file, expecting 9, got %d", len(header))
        return Client{}, errors.New(errorString)
    }

    return Client{
        path:       path,
        goroutines: goroutines,
    }, nil
}

// interface filter is used to filter out Items when doing a List
type Filter interface {
    // filter returns true when the given item should be included.
    filter(i Item) bool
}

func (c *Client) List(ctx context.Context, filters ...Filter) ([]Item, error) {
    errorChan := make(chan error)
    wgDone := make(chan bool)
    wg := new(sync.WaitGroup)
    wg.Add(c.goroutines)

    listTotal := make([][]Item, c.goroutines)
    for i := 0; i < c.goroutines; i++ {
        iCopy := i
        go func() {
            err := c.list(ctx, iCopy, &listTotal[iCopy], filters)
            if err != nil {
                errorChan <- err
                return
            }
            wg.Done()
        }()
    }

    go func() {
        wg.Wait()
        wgDone <- true
    }()

    select {
    case err := <-errorChan:
        return nil, err
    case <-wgDone:
    case <-ctx.Done():
        return nil, ctx.Err()
    }

    var list []Item
    for i := 0; i < c.goroutines; i++ {
        list = append(list, listTotal[i]...)
    }

    return list, nil
}

func (c *Client) list(ctx context.Context, offset int, out *[]Item, filters []Filter) error {
    file, err := os.Open(c.path)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for i := 0; i < offset+1; i++ {
        if !scanner.Scan(){
            return scanner.Err()
        }
    }

    for {
        select {
        case <-ctx.Done():
            return nil
        default:
        }

        done, err := scan(scanner, out, filters)
        if err != nil {
            return err
        }
        if done {
            return nil
        }

        for i := 0; i < c.goroutines-1; i++ {
            if !scanner.Scan() {
                return scanner.Err()
            }
        }
    }

    return nil
}

type scanner interface {
    Scan() bool
    Err() error
    Text() string
}

func scan(scanner scanner, out *[]Item, filters []Filter) (bool, error) {
    if !scanner.Scan() {
        if scanner.Err() != nil {
            return false, scanner.Err()
        }
        return true, nil
    }
    res := strings.Split(scanner.Text(), "\t")

    isAdult, err := parseInt(res[4])
    if err != nil {
        return false, err
    }

    startYear, err := parseInt(res[5])
    if err != nil {
        return false, err
    }

    endYear, err := parseInt(res[6])
    if err != nil {
        return false, err
    }

    runtimeMinutes, err := parseInt(res[7])
    if err != nil {
        return false, err
    }

    genres := strings.Split(res[8], ",")
    item := Item{
        TConst:         res[0],
        TitleType:      res[1],
        PrimaryTitle:   res[2],
        OriginalTitle:  res[3],
        IsAdult:        isAdult,
        StartYear:      startYear,
        EndYear:        endYear,
        RuntimeMinutes: runtimeMinutes,
        Genres:         genres,
    }

    for _, filter := range filters {
        if !filter.filter(item) {
            return  false, nil
        }
    }
    *out = append(*out, item)
    return false, nil
}

func parseInt(input string) (int, error) {
    if input != `\N` {
        return strconv.Atoi(input)
    }
    return 0, nil
}
