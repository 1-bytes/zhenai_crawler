package model

import "encoding/json"

//
// Profile
// @Description: 要爬取记录的用户基础信息
//
type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xingzuo    string
	House      string
	Car        string
}

// FromJsonObj .
func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
