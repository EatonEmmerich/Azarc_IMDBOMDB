package imdb_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"Azarc/imdb"
)

func TestList(t *testing.T) {
	ctx := context.Background()

	testTable := []struct {
	    name string
	    routines int
	    fileData []byte
		expectedItems []imdb.Item
	}{
	   /* {
            name :"one",
            routines: 1,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
                "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n"),
            expectedItems:  []imdb.Item{{
                TConst:         "tt0000001",
                TitleType:      "short",
                PrimaryTitle:   "Carmencita",
                OriginalTitle:  "Carmencita",
                StartYear:      1894,
                EndYear:        2020,
                RuntimeMinutes: 1,
                Genres:         []string{"Documentary", "Short"},
            }},
        },
        {
            name :"no value for integer",
            routines: 1,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
                "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n"),
            expectedItems:  []imdb.Item{{
                TConst:        "tt0000002",
                TitleType:     "short",
                PrimaryTitle:  "Le clown et ses chiens",
                OriginalTitle: "Le clown et ses chiens",
                Genres:        []string{"Animation", "Short"},
            }},
        },
        {
	        name: "quotes in tabbed field",
            routines: 1,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
	            "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n"),
            expectedItems: []imdb.Item{{
                    TConst:         "tt0033122",
                    TitleType:      "movie",
                    PrimaryTitle:   "\"Swing it\" magistern",
                    OriginalTitle:  "\"Swing it\" magistern",
                    StartYear:      1940,
                    RuntimeMinutes: 92,
                    Genres:         []string{"Comedy", "Music"},
            }},
        },
        {
            name: "many entries",
            routines: 1,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
                "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n" +
                "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n" +
                "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n"),
            expectedItems: []imdb.Item{
                {
                    TConst:         "tt0033122",
                    TitleType:      "movie",
                    PrimaryTitle:   "\"Swing it\" magistern",
                    OriginalTitle:  "\"Swing it\" magistern",
                    StartYear:      1940,
                    RuntimeMinutes: 92,
                    Genres:         []string{"Comedy", "Music"},
                },
                {
                    TConst:        "tt0000002",
                    TitleType:     "short",
                    PrimaryTitle:  "Le clown et ses chiens",
                    OriginalTitle: "Le clown et ses chiens",
                    Genres:        []string{"Animation", "Short"},
                },
                {
                    TConst:         "tt0000001",
                    TitleType:      "short",
                    PrimaryTitle:   "Carmencita",
                    OriginalTitle:  "Carmencita",
                    StartYear:      1894,
                    EndYear:        2020,
                    RuntimeMinutes: 1,
                    Genres:         []string{"Documentary", "Short"},
                },
            },
        },*/
        {
            name: "more entries",
            routines: 1,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
                "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n" +
                "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n" +
                "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n" +
                "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n" +
                "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n"),
            expectedItems: []imdb.Item{
                {
                    TConst:         "tt0033122",
                    TitleType:      "movie",
                    PrimaryTitle:   "\"Swing it\" magistern",
                    OriginalTitle:  "\"Swing it\" magistern",
                    StartYear:      1940,
                    RuntimeMinutes: 92,
                    Genres:         []string{"Comedy", "Music"},
                },
                {
                    TConst:        "tt0000002",
                    TitleType:     "short",
                    PrimaryTitle:  "Le clown et ses chiens",
                    OriginalTitle: "Le clown et ses chiens",
                    Genres:        []string{"Animation", "Short"},
                },
                {
                    TConst:         "tt0000001",
                    TitleType:      "short",
                    PrimaryTitle:   "Carmencita",
                    OriginalTitle:  "Carmencita",
                    StartYear:      1894,
                    EndYear:        2020,
                    RuntimeMinutes: 1,
                    Genres:         []string{"Documentary", "Short"},
                },
                {
                    TConst:        "tt0000002",
                    TitleType:     "short",
                    PrimaryTitle:  "Le clown et ses chiens",
                    OriginalTitle: "Le clown et ses chiens",
                    Genres:        []string{"Animation", "Short"},
                },
                {
                    TConst:         "tt0000001",
                    TitleType:      "short",
                    PrimaryTitle:   "Carmencita",
                    OriginalTitle:  "Carmencita",
                    StartYear:      1894,
                    EndYear:        2020,
                    RuntimeMinutes: 1,
                    Genres:         []string{"Documentary", "Short"},
                },
            },
        },
       /* {
            name:     "many routines",
            routines: 3,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n" +
                "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n" +
                "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n" +
                "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n"),
            expectedItems: []imdb.Item{
                {
                    TConst:         "tt0033122",
                    TitleType:      "movie",
                    PrimaryTitle:   "\"Swing it\" magistern",
                    OriginalTitle:  "\"Swing it\" magistern",
                    StartYear:      1940,
                    RuntimeMinutes: 92,
                    Genres:         []string{"Comedy", "Music"},
                },
                {
                    TConst:        "tt0000002",
                    TitleType:     "short",
                    PrimaryTitle:  "Le clown et ses chiens",
                    OriginalTitle: "Le clown et ses chiens",
                    Genres:        []string{"Animation", "Short"},
                },
                {
                    TConst:         "tt0000001",
                    TitleType:      "short",
                    PrimaryTitle:   "Carmencita",
                    OriginalTitle:  "Carmencita",
                    StartYear:      1894,
                    EndYear:        2020,
                    RuntimeMinutes: 1,
                    Genres:         []string{"Documentary", "Short"},
                },
            },
        },
        {
            name:     "more routines than data",
            routines: 10,
            fileData: []byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n" +
                "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n"),
            expectedItems: []imdb.Item{
                {
                    TConst:         "tt0033122",
                    TitleType:      "movie",
                    PrimaryTitle:   "\"Swing it\" magistern",
                    OriginalTitle:  "\"Swing it\" magistern",
                    StartYear:      1940,
                    RuntimeMinutes: 92,
                    Genres:         []string{"Comedy", "Music"},
                },
            },
        },*/
    }

    for _, test := range testTable {
        t.Run(test.name, func(t *testing.T) {
            file, err := ioutil.TempFile(os.TempDir(), "testIMDBFile-")
            require.NoError(t, err)
            file.Write(test.fileData)
            file.Close()

            imdbClient, err := imdb.New(file.Name(), test.routines)
            require.NoError(t, err)

            resp, err := imdbClient.List(ctx)
            require.NoError(t, err)
            require.ElementsMatch(t, test.expectedItems, resp)
        })
    }
}

func TestList_ContextCancelled(t *testing.T){
    ctx, cancel := context.WithCancel(context.Background())
    file, err := ioutil.TempFile(os.TempDir(), "testIMDBFile-")
    require.NoError(t, err)
    file.Write([]byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
        "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n" +
        "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t\\N\t\\N\t\\N\t\\N\tAnimation,Short\n" +
        "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short\n"))
    file.Close()

    done := make(chan bool)

    imdbClient, err := imdb.New(file.Name(), 1)
    require.NoError(t, err)

    cancel()

    go func () {
        _, err = imdbClient.List(ctx)
        done <- true
    }()

    <- done

    require.EqualError(t, err, "context canceled")
}