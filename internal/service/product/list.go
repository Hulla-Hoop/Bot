package product

import (
	"encoding/json"

	"io/ioutil"
	"log"
)

func (s *Service) List() []Product {
	data, err := ioutil.ReadFile(FILE)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var group allProduct

	jsonErr := json.Unmarshal(data, &group)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return group.Products
}
