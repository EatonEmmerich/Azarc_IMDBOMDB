package main

import (
    "context"
    "flag"
    "fmt"
    "os"
    "os/signal"
    "time"

    "Azarc/imdb"
    "Azarc/omdb"
)

var apiKey = flag.String("apiKey", "", "the omdb API key")
var filePath = flag.String("filePath", "title.basics.tsv", "Absolute path to the inflated `title.basics.tsv.gz` file")
var titleType = flag.String("titleType", "movie", "Filter on `titleType` column")
var primaryTitle = flag.String("primaryTitle", "The Mask", "filter on `primaryTitle` column")
var originalTitle = flag.String("originalTitle", "", "filter on `originalTitle` column")
var genre = flag.String("genre", "", "filter on `genre` column")
var startYear = flag.Int("startYear", 0, "filter on `startYear` column")
var endYear = flag.Int("endYear", 0, "filter on `endYear` column")
var runtimeMinutes = flag.Int("runtimeMinutes", 0, "filter on `runtimeMinutes` column")
var genres = flag.String("genres", "", "filter on `genres` column")
var maxApiRequests = flag.String("maxApiRequests", "", "maximum number of requests to be made to [omdbapi](https://www.omdbapi.com/)")
var maxRunTime = flag.Duration("maxRunTime", time.Minute, "maximum run time of the application. Format is a `time.Duration` string see [here](https://godoc.org/time#ParseDuration)")
var maxRequests = flag.String("maxRequests", "", "maximum number of requests to send to [omdbapi](https://www.omdbapi.com/)")
var plotFilter = flag.String("plotFilter", "", "regex pattern to apply to the plot of a film retrieved from [omdbapi](https://www.omdbapi.com/)")
var fileReadRoutines = flag.Int("fileReadRoutines", 4, "number of routines used to parse the input file")

func main() {
    flag.Parse()
    ctx, cancel := context.WithTimeout(context.Background(), *maxRunTime)
    defer cancel()

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)

    go func() {
        for {
            select {
                case <-c:
                    cancel()
                case <-ctx.Done():
                    maybeExitGracefully(ctx.Err())
                    os.Exit(0)
                default:
                    //fmt.Printf("number of goroutines %d\n", runtime.NumGoroutine())
                    time.Sleep(time.Second)
            }
        }
    }()

    filters := buildFilters()

    imdbClient, err := imdb.New(*filePath, *fileReadRoutines)
    list, err := imdbClient.List(ctx, filters...)
    if err != nil {
        maybeExitGracefully(err)
    }

    omdbClient := omdb.New(*apiKey)

    for _, imdbitem := range list {
        omdbItem, err := omdbClient.Info(ctx, imdbitem.TConst)
        if err != nil {
            maybeExitGracefully(err)
        }
        fmt.Printf("imdbItem: %#v\nomdbItem: %#v\n\n", imdbitem, omdbItem)
    }
}

func maybeExitGracefully(err error){
    if err == context.DeadlineExceeded || err == context.Canceled {
        fmt.Printf("Stopping execution\n")
        os.Exit(0)
    }
    panic(err)
}

func buildFilters() []imdb.Filter{
    var filters []imdb.Filter
    if *titleType != ""{
        filters = append(filters,imdb.NewTitleTypeFilter(*titleType))
    }
    if *primaryTitle != "" {
        filters = append(filters, imdb.NewPrimaryTitleFilter(*primaryTitle))
    }
    if *originalTitle != "" {
        filters = append(filters, imdb.NewOriginalTitleFilter(*originalTitle))
    }
    if *genre != "" {
        filters = append(filters, imdb.NewGenreFilter(*genre))
    }
    if *startYear != 0 {
        filters = append(filters, imdb.NewStartYearFilter(*startYear))
    }
    if *endYear != 0 {
        filters = append(filters, imdb.NewEndYearFilter(*endYear))
    }
    if *runtimeMinutes != 0 {
        filters = append(filters, imdb.NewRuntimeMinutesFilter(*runtimeMinutes))
    }
    return filters
}
