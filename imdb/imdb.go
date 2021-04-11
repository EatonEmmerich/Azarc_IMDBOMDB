package imdb

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Item struct {
    TConst string
    TitleType string
    PrimaryTitle string
    OriginalTitle string
    IsAdult int
    StartYear int
    EndYear int // \N
    RuntimeMinutes  int
    Genres []string
}

type Client struct {
    file *os.File
}

func New(path string) (Client, error){
    file, err := os.Open(path)
    if err != nil {
        return Client{}, err
    }

    return Client{
        file: file,
    }, nil
}

func (c Client) List() ([]Item, error) {
    var list []Item
    scanner := bufio.NewScanner(c.file)
    ok := scanner.Scan()
    if !ok {
        return nil, errors.New("could not scan")
    }


    header := strings.Split(scanner.Text(), "\t")
    if len(header) != 9 {
        errorString := fmt.Sprintf("invalid headers for file, expecting 9, got %d", len(header))
        return nil, errors.New(errorString)
    }

    res := make([]string,9)

    for {
        ok := scanner.Scan()
        if !ok {
            break
        }
        res = strings.Split(scanner.Text(), "\t")

        isAdult, err := parseInt(res[4])
        if err != nil {
            return nil, err
        }

        startYear, err := parseInt(res[5])
        if err != nil {
            return nil, err
        }

        endYear, err := parseInt(res[6])
        if err != nil {
            return nil, err
        }

        runtimeMinutes, err := parseInt(res[7])
        if err != nil {
            return nil, err
        }

        genres := strings.Split(res[8],",")

        list = append(list, Item{
            TConst: res[0],
            TitleType: res[1],
            PrimaryTitle: res[2],
            OriginalTitle: res[3],
            IsAdult: isAdult,
            StartYear: startYear,
            EndYear: endYear,
            RuntimeMinutes: runtimeMinutes,
            Genres : genres,
        })
    }

    return list, nil
}

func parseInt(input string) (int, error){
    if input != `\N` {
        return strconv.Atoi(input)
    }
    return 0, nil
}