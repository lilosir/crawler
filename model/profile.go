package model

import "encoding/json"

// Profile denotes user info
type Profile struct {
	Name     string
	Gender   string
	Area     string
	Age      int
	Marriage string
	Height   int
	Income   string
	Zodiac   string
}

// FromJSONObj will parse the json file to string, then converte to json again
func FromJSONObj(o interface{}) (Profile, error) {
	var profile Profile
	str, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(str, &profile)
	return profile, err
}
