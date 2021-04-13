package imdb

import (
    "testing"

    "github.com/stretchr/testify/require"
)

type mockScanner struct {
    scans       []bool
    err         error
    texts       []string
    scanCounter int
    textCounter int
}

func (m *mockScanner) Scan() bool {
    return m.scans[m.scanCounter]
}

func (m *mockScanner) Err() error {
    return m.err
}

func (m *mockScanner) Text() string {
    return m.texts[m.textCounter]
}

type mockFilter struct {
    t          *testing.T
    expectItem Item
    pass       bool
}

func (m mockFilter) filter(i Item) bool {
    require.Equal(m.t, m.expectItem, i)
    return m.pass
}

func TestScan(t *testing.T) {
    baseScanner := mockScanner{
        scans: []bool{true},
        texts: []string{"tt0000001\tshort\tCarmencita\tCarmencita\t0\t1894\t2020\t1\tDocumentary,Short"},
    }

    itemA := Item{
        TConst:         "tt0000001",
        TitleType:      "short",
        PrimaryTitle:   "Carmencita",
        OriginalTitle:  "Carmencita",
        StartYear:      1894,
        EndYear:        2020,
        RuntimeMinutes: 1,
        Genres:         []string{"Documentary", "Short"},
    }

    itemB := Item{
        TConst:        "tt0000002",
        TitleType:     "short",
        PrimaryTitle:  "Le clown et ses chiens",
        OriginalTitle: "Le clown et ses chiens",
        Genres:        []string{"Animation", "Short"},
    }

    tests := []struct {
        name         string
        scanner      mockScanner
        out          []Item
        filters      []Filter
        expectedOut  []Item
        expectedDone bool
    }{
        {
            name:    "scan done",
            scanner: mockScanner{scans: []bool{false}},
            out: []Item{
                itemB,
            },
            filters: []Filter{mockFilter{t: t, expectItem: itemA}},
            expectedOut: []Item{
                itemB,
            },
            expectedDone: true,
        },
        {
            name:    "with filter",
            scanner: baseScanner,
            out: []Item{
                itemB,
            },
            filters: []Filter{mockFilter{t: t, expectItem: itemA}},
            expectedOut: []Item{
                itemB,
            },
        },
        {
            name:    "with filter passed",
            scanner: baseScanner,
            out: []Item{
                itemB,
            },
            filters: []Filter{mockFilter{t: t, expectItem: itemA, pass: true}},
            expectedOut: []Item{
                itemB,
                itemA,
            },
        },
        {
            name:    "no filter",
            scanner: baseScanner,
            out: []Item{
                {
                    TConst:        "tt0000002",
                    TitleType:     "short",
                    PrimaryTitle:  "Le clown et ses chiens",
                    OriginalTitle: "Le clown et ses chiens",
                    Genres:        []string{"Animation", "Short"},
                },
            },
            expectedOut: []Item{
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
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            var out []Item
            out = append(out, test.out...)
            resp, err := scan(&test.scanner, &out, test.filters)
            require.NoError(t, err)
            require.ElementsMatch(t, test.expectedOut, out)
            require.Equal(t, test.expectedDone, resp)
        })
    }
}
