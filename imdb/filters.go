package imdb


type titleType string

func (f titleType) filter(i Item) bool {
    return i.TitleType == string(f)
}

func NewTitleTypeFilter(s string) Filter {
    return titleType(s)
}

type primaryTitle string

func (f primaryTitle) filter(i Item) bool {
    return i.PrimaryTitle == string(f)
}

func NewPrimaryTitleFilter(s string) Filter {
    return primaryTitle(s)
}

type originalTitle string

func (f originalTitle) filter(i Item) bool {
    return i.OriginalTitle == string(f)
}

func NewOriginalTitleFilter(s string) Filter {
    return originalTitle(s)
}

type genre string

func (f genre) filter(i Item) bool {
    for _, genre := range i.Genres {
        if genre == string(f){
            return true
        }
    }
    return false
}

func NewGenreFilter(s string) Filter {
    return genre(s)
}

type startYear int

func (f startYear) filter(i Item) bool {
    return i.StartYear == int(f)
}

func NewStartYearFilter(i int) Filter {
    return startYear(i)
}

type endYear int

func (f endYear) filter(i Item) bool {
    return i.EndYear == int(f)
}

func NewEndYearFilter(i int) Filter {
    return endYear(i)
}

type runtimeMinutes int

func (f runtimeMinutes) filter(i Item) bool {
    return i.RuntimeMinutes == int(f)
}

func NewRuntimeMinutesFilter(i int) Filter {
    return runtimeMinutes(i)
}

type genres []string

// todo test
func (f genres) filter(i Item) bool {
    if len([]string(f)) != len(i.Genres){
        return false
    }

    for _, genreIn := range i.Genres{
        for _, genreFilter := range []string(f){
            if genreIn == genreFilter{
                continue
            }
        }
    }
    return true
}

func NewGenresFilter(s []string) Filter {
    return genres(s)
}
