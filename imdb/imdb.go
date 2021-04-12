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
        go c.list(ctx, i, &listTotal[i], filters, wg, errorChan)
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

func (c *Client) list(ctx context.Context, offset int, out *[]Item, filters []Filter, group *sync.WaitGroup, errorChan chan error) {
    // todo: rather add group and errorChan in the go func in calling function.
    file, err := os.Open(c.path)
    if err != nil {
        errorChan <- err
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for i := 0; i < offset+1; i++ {
        ok := scanner.Scan()
        if !ok {
            err := scanner.Err()
            if err != nil {
                errorChan <- err
                return
            }
            break
        }
    }

    res := make([]string, 9)
    var item Item
nextItem:
    for {
        select {
        case <-ctx.Done():
            return
        default:
        }
        ok := scanner.Scan()
        if !ok {
            err := scanner.Err()
            if err != nil {
                errorChan <- err
                return
            }
            break
        }
        res = strings.Split(scanner.Text(), "\t")

        isAdult, err := parseInt(res[4])
        if err != nil {
            errorChan <- err
            return
        }

        startYear, err := parseInt(res[5])
        if err != nil {
            errorChan <- err
            return
        }

        endYear, err := parseInt(res[6])
        if err != nil {
            errorChan <- err
            return
        }

        runtimeMinutes, err := parseInt(res[7])
        if err != nil {
            errorChan <- err
            return
        }

        genres := strings.Split(res[8], ",")
        item = Item{
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
                continue nextItem
            }
        }
        *out = append(*out, item)

        for i := 0; i < c.goroutines; i++ {
            ok := scanner.Scan()
            if !ok {
                err := scanner.Err()
                if err != nil {
                    errorChan <- err
                    return
                }
                break
            }
        }
    }
    group.Done()
    return
}

func parseInt(input string) (int, error) {
    if input != `\N` {
        return strconv.Atoi(input)
    }
    return 0, nil
}
