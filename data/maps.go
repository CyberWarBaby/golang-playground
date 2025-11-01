package main

import "fmt"

type floatmap map[string]float64

func (m floatmap) output() {
	fmt.Println(m)
}
func main() {
	websitesArray := []string{
		"https://google.com",
		"https://aws.com",
		"https://github.com",
		"https://stackoverflow.com",
		"https://reddit.com",
		"https://medium.com",
		"https://linkedin.com",
		"https://twitter.com",
		"https://youtube.com",
		"https://docker.com",
	}

	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon web services": "https://aws.com",
	}
	fmt.Println(websites["Google"])

	websites["Linkedin"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

	makeSlice := make([]string, 2, 5)
	makeSlice = append(makeSlice, "sosyyyslice")
	fmt.Println(makeSlice)

	// courseRatings := map[string]float64{}
	courseRatings := make(floatmap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5

	courseRatings.output()
	fmt.Println(courseRatings)

	for i, v := range websitesArray {
		// fmt.Println(i + 1)
		fmt.Println(i+1, v)
	}
}
