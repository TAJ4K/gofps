// Copyright 2022 TAJ4K
package gofps

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/PuerkitoBio/goquery"
)

// Search takes a name, city, and state then returns a People object.
// City is optional, but state isn't.
func Search(name, city, state string) (People, error) {
	var people People

	if name == "" {
		return People{}, fmt.Errorf("name is required")
	}
	if state == "" {
		return People{}, fmt.Errorf("state is required")
	}
	if len(state) > 2 {
		return People{}, fmt.Errorf("state is invalid, use 2 character abbreviation")
	}

	url, err := createLink(name, city, state)
	if err != nil {
		return People{}, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return People{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return People{}, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

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

		if person.Addresses.Current[0] == ' ' {
			person.Addresses.Current = person.Addresses.Current[1:]
		}

		s.Find(".row > div > a").Each(func(i int, s *goquery.Selection) {
			person.Addresses.Past = append(person.Addresses.Past, strings.Replace(s.Text(), "\n", " ", -1))
		})

		person.Phones.Current = s.Find("strong > .nowrap").Text()
		s.Find("a.nowrap").Each(func(i int, s *goquery.Selection) {
			if s.Text()[0] == '(' {
				person.Phones.Past = append(person.Phones.Past, s.Text())
			} else {
				person.Relatives = append(person.Relatives, s.Text())
			}
		})

		people.People = append(people.People, person)
	})

	return people, nil
}

// Internal use for creating the links
func createLink(name, city, state string) (*url.URL, error) {
	var data string

	if city == "" {
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

// Takes in a AgeGroup to filter a People object
func (people People) SearchPeopleByAgeGroup(ageGroup AgeGroup) People {
	var result People

	for _, person := range people.People {
		if person.Age == "" {
			continue
		}
		if person.Age == "0" {
			continue
		}

		age, err := strconv.ParseInt(person.Age, 10, 64)
		if err != nil {
			continue
		}

		if age >= ageGroup.Min && age <= ageGroup.Max {
			result.People = append(result.People, person)
		}
	}

	return result
}

// Takes in a specific Age and returns all people that match
func (people People) SearchPeopleByAge(givenAge int64) People {
	var result People

	for _, person := range people.People {
		if person.Age == "" {
			continue
		}
		if person.Age == "0" {
			continue
		}

		age, err := strconv.ParseInt(person.Age, 10, 64)
		if err != nil {
			continue
		}

		if age == givenAge {
			result.People = append(result.People, person)
		}
	}

	return result
}

// Converts People struct to JSON
func (people People) ConvertToJson() (string, error) {
	json, err := json.Marshal(people)
	if err != nil {
		return "", err
	}

	return string(json), nil
}

// Converts Person struct to JSON
func (person Person) ConvertToJson() (string, error) {
	json, err := json.Marshal(person)
	if err != nil {
		return "", err
	}

	return string(json), nil
}