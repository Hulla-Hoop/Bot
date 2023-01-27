package product

import (
	"encoding/json"

	"io/ioutil"
	"log"
)

func (s *Service) Delete(idx int) {
	data, err := ioutil.ReadFile(FILE)
	if err != nil {
		log.Fatal("Cannot load settings:", err)
	}

	var group allProduct

	jsonErr := json.Unmarshal(data, &group)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	group.Products[idx] = Product{}

	data, err = json.MarshalIndent(&group, "", "  ")
	if err != nil {
		log.Fatal("Json marshaling failed:", err)
	}
	err = ioutil.WriteFile(FILE, data, 0)
	if err != nil {
		log.Fatal("cannot write updated settings file:", err)
	}
}
