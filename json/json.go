package json

import (
	"encoding/json"
	"log"
)

var B = []byte(`{ 
   "addons":{ 
      "localtest@techtouch.jp":{ 
         "updates":[ 
            { 
               "version":"0.0.0.2",
               "update_link":"version 2"
            },
            { 
               "version":"0.0.0.3",
               "update_link":"version 3"
            }
         ]
      }
   }
}`)

type UpdatesJson struct {
	Addons map[string] struct {
		Updates *[]Update`json:"updates"`
	} `json:"addons"`
}

type Update struct {
	Version    string `json:"version"`
	UpdateLink string `json:"update_link"`
}

// json → struct
func PrintJson(b []byte) {
	var device UpdatesJson

	err := json.Unmarshal(b, &device)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range device.Addons {
		log.Printf("%v \n", key)
		for _ , v := range *val.Updates {
			log.Printf("%v\n", v)
		}
	}
}

func UpdateJson(b []byte) ([]byte, error) {
	var updateJson UpdatesJson

	err := json.Unmarshal(b, &updateJson)
	if err != nil {
		log.Fatal(err)
	}

	var update Update
	update.Version = "version"
	update.UpdateLink = "link"

	for _, val := range updateJson.Addons {
		*val.Updates = append(*val.Updates, update)
	}

	return json.Marshal(updateJson)
}

func CreateJson() ([]byte, error){
	var update Update
	update.Version = "version"
	update.UpdateLink = "link"

	var updates []Update
	updates = append(updates, update)

	updatesMap := make(map[string][]Update)
	updatesMap["updates"] = updates

	gecko := make(map[string]map[string][]Update)
	gecko["localtest@techtouch.jp"] = updatesMap

	addons := make(map[string]map[string]map[string][]Update)
	addons["addons"] = gecko

	return json.Marshal(addons)
}
