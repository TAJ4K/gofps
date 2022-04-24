# gofps
## [FastPeopleSearch](https://fastpeoplesearch.com) Wrapper

## Usage
    
    ```go
    import "github.com/TAJ4K/gofps"
    
    people, err := fps.Search("John Smith", "", "NY")
    if err != nil {
		fmt.Println(err)
	}
    
    genz := people.SearchPeopleByAgegroup(AgeGroupGenZ)
    for _, person := range genz {
        fmt.Println(person.Addresses.Current)
    }
    ```