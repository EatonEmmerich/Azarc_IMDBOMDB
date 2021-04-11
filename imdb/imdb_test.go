package imdb_test

import (
    "io/ioutil"
    "os"
    "testing"

    "github.com/stretchr/testify/require"

    "Azarc/imdb"
)

func TestList(t *testing.T){
    file, err := ioutil.TempFile(os.TempDir(),"testIMDBFile-")
    require.NoError(t, err)
    file.Write([]byte("tconst\ttitleType\tprimaryTitle\toriginalTitle\tisAdult\tstartYear\tendYear\truntimeMinutes\tgenres\n"+
        "tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t\\N\t1\tDocumentary,Short\n"+
        "tt0000002\tshort\tLe clown et ses chiens\tLe clown et ses chiens\t0\t1892\t\\N\t\\N\tAnimation,Short\n"+
        "tt0033122\tmovie\t\"Swing it\" magistern\t\"Swing it\" magistern\t0\t1940\t\\N\t92\tComedy,Music\n"))
    file.Close()

    imdbClient, err := imdb.New(file.Name())
    require.NoError(t, err)

    resp, err := imdbClient.List()
    require.NoError(t, err)
    require.ElementsMatch(t, []imdb.Item{
        {
            TConst:         "tt0000001",
            TitleType:      "short",
            PrimaryTitle:   "Carmencita",
            OriginalTitle:  "Carmencita",
            StartYear:      1894,
            RuntimeMinutes: 1,
            Genres:         []string{"Documentary","Short"},
        },
        {
            TConst:         "tt0000002",
            TitleType:      "short",
            PrimaryTitle:   "Le clown et ses chiens",
            OriginalTitle:  "Le clown et ses chiens",
            StartYear:      1892,
            Genres:         []string{"Animation","Short"},
        },
        {
            TConst:         "tt0033122",
            TitleType:      "movie",
            PrimaryTitle:   "\"Swing it\" magistern",
            OriginalTitle:  "\"Swing it\" magistern",
            StartYear:      1940,
            RuntimeMinutes: 92,
            Genres:         []string{"Comedy","Music"},
        },
    }, resp)
}
