package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

// Post sends a post request to the URL
func (c *Client) GetStudentClasses(studentId int) (*[]Class, error) {
	classesUrl := fmt.Sprintf("%s/classes/%d", c.url, studentId)
	var classes *[]Class

	response, err := http.Get(classesUrl)
	if err != nil {
		return classes, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return classes, err
	}

	err = json.NewDecoder(response.Body).Decode(&classes)
	if err != nil {
		return classes, err
	}

	return classes, nil
}
