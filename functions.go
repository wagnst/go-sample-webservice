package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Platform struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var platform []Platform

func init() {
	RegisterPlatform(Platform{
		ID:   1,
		Name: "Test-Platform-1",
	})
	RegisterPlatform(Platform{
		ID:   2,
		Name: "Test-Platform-2",
	})
}

func RegisterPlatform(p Platform) error {
	if p.ID <= 0 {
		return NewErrIdInvalid("id must be larger 0")
	}

	_, err := GetPlatform(p.ID)
	if err == nil {
		return NewErrIdInvalid("platform with id " + strconv.Itoa(p.ID) + " is already existing")
	}

	if p.Name == "" {
		return NewErrPlatformInvalid("name of the platform must not be empty")
	}

	platform = append(platform, p)

	return nil
}

func GetPlatform(id int) (Platform, error) {
	if len(platform) >= id && id > 0 {
		for i, p := range platform {
			if p.ID == id {
				return platform[i], nil
			}
		}
	}
	return Platform{}, NewErrNotFound("no platform found")
}

func GetPlatforms() []Platform {
	return platform
}

func GetPlatformsJson(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(GetPlatforms())
}

func GetPlatformJson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	platform, err := GetPlatform(id)

	if err != nil {
		log.Printf("got an error during request: %s ", err)
		w.WriteHeader(http.StatusNotFound)

	} else {
		json.NewEncoder(w).Encode(platform)
	}
}
