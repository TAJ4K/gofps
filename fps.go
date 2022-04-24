package gofps

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func Search(name, city, state string) (People, error) {
	var people People

	if name == "" {return People{}, fmt.Errorf("name is required")}
	if state == "" {return People{}, fmt.Errorf("state is required")}
	if len(state) > 2 {return People{}, fmt.Errorf("state is invalid, use 2 character abbreviation")}

	url, err := createLink(name, city, state)
	if err != nil {
		return People{}, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return People{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {return People{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return People{}, err
	}

	doc.Find(".people-list > .card > .card-block").Each(func(i int, s *goquery.Selection) {
		var person Person
		person.Name = s.Find(".card-title > a > .larger").Text()
		person.GoesBy = s.Find(".card-title > a > span:nth-child(3) > strong").Text()
		person.Age = strings.Replace(strings.Split(s.Text(), "\n")[4], "Age: ", "", -1)

		person.Addresses.Current = strings.Replace(s.Find("div > strong > a:not(.nowrap)").Text(), "\n", " ", -1)

		s.Find(".row > div > a").Each(func(i int, s *goquery.Selection) {
			person.Addresses.Past = append(person.Addresses.Past, strings.Replace(s.Text(), "\n", " ", -1))
		})

		person.Phones.Current = s.Find("strong > .nowrap").Text()
		s.Find(".nowrap").Each(func(i int, s *goquery.Selection) {
			person.Phones.Past = append(person.Phones.Past, strings.Replace(s.Text(), "\n", " ", -1))
		})

		people.People = append(people.People, person)
	})

	return people, nil
}

func createLink(name, city, state string) (*url.URL, error) {
	var data string

	if city == ""{
		data = strings.Replace(name, " ", "-", -1) + "_" + state
	} else {
		data = strings.Replace(name, " ", "-", -1) + "_" + city + "-" + state
	}
	
	url, err := url.Parse("https://www.fastpeoplesearch.com/name/" + data)
	if err != nil {
		return nil, err
	}

	return url, nil 
}

func (people People) SearchPeopleByAgegroup(ageGroup AgeGroup) People {
	var result People

	for _, person := range people.People {
		if person.Age == "" {continue}
		if person.Age == "0" {continue}

		age, err := strconv.ParseInt(person.Age, 10, 64)
		if err != nil {continue}

		if age >= ageGroup.Min && age <= ageGroup.Max {
			result.People = append(result.People, person)
		}
	}

	return result
}

func (people People) SearchPeopleByAge(givenAge int64) People {
	var result People

	for _, person := range people.People {
		if person.Age == "" {continue}
		if person.Age == "0" {continue}

		age, err := strconv.ParseInt(person.Age, 10, 64)
		if err != nil {continue}

		if age == givenAge {
			result.People = append(result.People, person)
		}
	}

	return result
}