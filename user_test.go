package deezer

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetUser(t *testing.T) {
	result, err := GetUser(374224155)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserAlbums(t *testing.T) {
	result, err := GetUserAlbums(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserArtists(t *testing.T) {
	result, err := GetUserArtists(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserCharts(t *testing.T) {
	result, err := GetUserCharts(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserFlow(t *testing.T) {
	result, err := GetUserFlow(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserFollowings(t *testing.T) {
	result, err := GetUserFollowings(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserPlaylists(t *testing.T) {
	result, err := GetUserPlaylists(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserRadios(t *testing.T) {
	result, err := GetUserRadios(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}

func TestGetUserTracks(t *testing.T) {
	result, err := GetUserTracks(374224155, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	json, _ := json.Marshal(result)
	fmt.Println(string(json))
}
