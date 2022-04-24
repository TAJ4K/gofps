package gofps

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
}

type People struct {
	People []Person
}

type AgeGroup struct {
	Min int64
	Max int64
}

var (
	AgeGroupYoung  = AgeGroup{Min: 18, Max: 28}
	AgeGroupMiddle = AgeGroup{Min: 28, Max: 45}
	AgeGroupOld    = AgeGroup{Min: 45, Max: 100}
)

var (
	AgeGroupGenZ       = AgeGroup{Min: 18, Max: 25}
	AgeGroupMillenials = AgeGroup{Min: 25, Max: 41}
	AgeGroupGenX       = AgeGroup{Min: 41, Max: 57}
	AgeGroupBoomers    = AgeGroup{Min: 57, Max: 76}
	AgeGroupSilent     = AgeGroup{Min: 76, Max: 100}
)
