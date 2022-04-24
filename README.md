# gofps
## Golang [FastPeopleSearch](https://fastpeoplesearch.com) Wrapper

## Usage
    
```go
import (
    "github.com/TAJ4K/gofps"
    "fmt"
)

func main() {
    people, err := fps.Search("John Smith", "", "NY")
    if err != nil {
        fmt.Println(err)
    }

    genz := people.SearchPeopleByAgegroup(gofps.AgeGroupGenZ)
    for _, person := range genz {
        fmt.Println(person.Addresses.Current)
    }
}
```