package main

import (
	"testing"
)

var platform1, platform2, platform3, platform4, platform5, platform6 Platform

func init() {

	platform1 = Platform{
		ID:   1,
		Name: "Test-Platform-1",
	}

	platform2 = Platform{
		ID:   2,
		Name: "Test-Platform-2",
	}

	platform3 = Platform{
		ID:   3,
		Name: "Test-Platform-3",
	}

	platform4 = Platform{
		ID:   4,
		Name: "Test-Platform-4",
	}

	platform5 = Platform{
		ID:   0,
		Name: "Test-Platform-0",
	}

	platform6 = Platform{
		ID:   6,
		Name: "",
	}
}

func TestRegisterValidPlatform(t *testing.T) {
	err := RegisterPlatform(platform3)
	if err != nil {
		t.FailNow()
	}
}

func TestRegisterInvalidPlatformId(t *testing.T) {
	err := RegisterPlatform(platform5)
	if err == nil {
		t.FailNow()
	}
}

func TestRegisterInvalidPlatform(t *testing.T) {
	err := RegisterPlatform(platform6)
	if err == nil {
		t.FailNow()
	}
}

func TestRegisterExistingPlatform(t *testing.T) {
	err := RegisterPlatform(platform1)
	if err == nil {
		t.FailNow()
	}
}

func TestRegisterEmptyPlatform(t *testing.T) {
	err := RegisterPlatform(platform6)
	if err == nil {
		t.FailNow()
	}
}

func TestGetPlatform(t *testing.T) {
	p, err := GetPlatform(1)
	if err != nil {
		t.FailNow()
	}

	if p.ID != 1 {
		t.FailNow()
	}
}

func TestGetPlatformInvalid(t *testing.T) {
	_, err := GetPlatform(99999)
	if err == nil {
		t.FailNow()
	}
}

func TestRegisterGetPlatforms(t *testing.T) {
	var defaultPlatforms []Platform
	defaultPlatforms = append(defaultPlatforms, platform1)
	defaultPlatforms = append(defaultPlatforms, platform2)

	err := RegisterPlatform(platform4)
	if err != nil {
		t.FailNow()
	}

	currentListAfter := GetPlatforms()

	found := false
	for _, a := range currentListAfter {
		if a.ID == platform4.ID {
			found = true
		}
	}

	if !found {
		t.FailNow()
	}
}
