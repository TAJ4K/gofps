# gofps
## Golang [FastPeopleSearch](https://fastpeoplesearch.com) Wrapper

## Usage
    
```go
package main

import(
	"fmt"
	"github.com/TAJ4K/gofps"
)

func main() {
    people, err := gofps.Search("John Smith", "", "NY")
    if err != nil {
        fmt.Println(err)
    }

    genz := people.SearchPeopleByAgegroup(gofps.AgeGroupBoomers)
    for _, person := range genz.People {
        fmt.Println(person.Addresses.Current)
    }
}
```