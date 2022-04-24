package gofps

// Person contains all the information of a single person that is publically available.
type Person struct {
	Name   string
	GoesBy string
	Age    string

	Addresses struct {
		Current string
		Past    []string
	}

	Phones struct {
		Current string
		Past    []string
	}

	Relatives []string
}
// People is a collection of Person objects that gets returned from the search function and all sorts
type People struct {
	People []Person
}

// AgeGroup is a struct that contains the min and max age of a group of people used for sorting struct People
// You can create your own instead of using premades by `gofps.AgeGroup{Min: 0, Max: 100}`
type AgeGroup struct {
	Min int64
	Max int64
}

// Age Group estimates are provided to help narrow down the search results.
var (
	AgeGroupYoung  = AgeGroup{Min: 18, Max: 28}
	AgeGroupMiddle = AgeGroup{Min: 28, Max: 45}
	AgeGroupOld    = AgeGroup{Min: 45, Max: 100}
)

// Generations are also provided, incase it helps more.
var (
	AgeGroupGenZ       = AgeGroup{Min: 18, Max: 25}
	AgeGroupMillenials = AgeGroup{Min: 25, Max: 41}
	AgeGroupGenX       = AgeGroup{Min: 41, Max: 57}
	AgeGroupBoomers    = AgeGroup{Min: 57, Max: 76}
	AgeGroupSilent     = AgeGroup{Min: 76, Max: 100}
)
