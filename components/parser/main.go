package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type StringField string

func (s StringField) ToString() string {
	return string(s)
}

type ArrayField []any
type ObjectField []map[string]any

func unflattenJSON(data map[string]any) (map[string]any, error) {
	result := make(map[string]any)
	for key, value := range data {
		if err := setNestedValue(result, key, value); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func setNestedValue(result map[string]any, key string, value any) error {
	pattern := regexp.MustCompile(`(\w+)(\[(\d+)\])?(\/(\w+))?`)
	matches := pattern.FindAllStringSubmatch(key, -1)

	current := result
	for i, match := range matches {
		isLast := i == len(matches)-1
		field := match[1]

		if match[2] != "" { // Array
			index, _ := strconv.Atoi(match[3])
			array, exists := current[field].([]any)
			if !exists {
				array = make([]any, index+1)
				current[field] = array
			} else if index >= len(array) {
				newArray := make([]any, index+1)
				copy(newArray, array)
				current[field] = newArray
				array = newArray
			}

			if isLast {
				array[index] = value
			} else {
				nextMap, exists := array[index].(map[string]any)
				if !exists {
					nextMap = make(map[string]any)
					array[index] = nextMap
				}
				current = nextMap
			}
		} else if match[5] != "" { // Object in array
			if isLast {
				current[match[5]] = value
			} else {
				nextMap, exists := current[match[5]].(map[string]any)
				if !exists {
					nextMap = make(map[string]any)
					current[match[5]] = nextMap
				}
				current = nextMap
			}
		} else { // Regular field
			if isLast {
				current[field] = value
			} else {
				nextMap, exists := current[field].(map[string]any)
				if !exists {
					nextMap = make(map[string]any)
					current[field] = nextMap
				}
				current = nextMap
			}
		}
	}
	return nil
}

type Foo struct {
	Name    StringField `json:"name,omitempty"`
	Likes   ArrayField  `json:"likes,omitempty"`
	Feature ArrayField  `json:"feature,omitempty"`
	Owns    ObjectField `json:"owns,omitempty"`
}

func main() {
	flatJSON := `{
		"name": "the name",
		"likes[1]": "wings",
		"likes[0]": "beer",
		"likes[2]": "hikes",
		"feature[0]/name": "gamer",
		"feature[1]/name": "griller",
		"owns/[0]/type": "dog",
		"owns/[0]/amount": "1",
		"owns/[2]/type": "tv",
		"owns/[2]/amount": "2",
		"owns/[1]/type": "computers",
		"owns/[1]/amount": "2"
	}`

	var flatData map[string]any
	if err := json.Unmarshal([]byte(flatJSON), &flatData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	unflattenedData, err := unflattenJSON(flatData)
	if err != nil {
		fmt.Println("Error unflattening JSON:", err)
		return
	}

	unflattenedJSON, _ := json.MarshalIndent(unflattenedData, "", "  ")
	fmt.Println("Unflattened JSON", string(unflattenedJSON))
}
