# gofps - Golang [FastPeopleSearch](https://fastpeoplesearch.com) Wrapper

## Why?
Almost every single library for people searching is paid and charges a ridiculous fee. This wrapper is free and open source.
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

    boomers := people.SearchPeopleByAgegroup(gofps.AgeGroupBoomers)
    for _, person := range boomers.People {
        fmt.Println(person.Addresses.Current)
    }
}
```
