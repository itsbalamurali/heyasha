package database

import (
	"testing"
	"net/http"
)

//db := database.Cleint

func TestClient(t *testing.T) {
	httpClient := http.DefaultClient
	client := Client("balamurali@live.com","yourass1994","5161","heyasha",httpClient)
	res,err := client.Query("users","SELECT * FROM users LIMIT 0, 20")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDatabase_Query(t *testing.T) {
	httpClient := http.DefaultClient
	client := Client("balamurali@live.com","yourass1994","5161","heyasha",httpClient)
	res,err := client.Query("users","SELECT * FROM users WHERE FirstName = \"Balamurali\"")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}