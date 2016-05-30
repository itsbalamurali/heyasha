package database

import (
	"net/http"
	"fmt"
	"github.com/dghubble/sling"
	"encoding/base64"
	"log"
)

const baseURL = "https://api-eu.clusterpoint.com/v4/"

type Database struct {
	AccountID string
	DBName string
	sling *sling.Sling

}

func Client(username string, password string, accID string,dbName string,httpClient *http.Client) *Database {
	return &Database{
		AccountID:accID,
		DBName:dbName,
		sling: sling.New().Client(httpClient).Base(baseURL).Set("Authorization", "Basic "+basicAuth(username, password)),
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

type ClusterpointError struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Field    string `json:"field"`
		Code     string `json:"code"`
	} `json:"errors"`
	DocumentationURL string `json:"documentation_url"`
}

func (s *Database) Query(collection string, query string) (*http.Response, error) {
	//clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf(s.AccountID+"/"+s.DBName+".%s/_query", collection)
	req, err := s.sling.New().Post(path).Request()
	if err != nil {
		//err = clusterpointError
		log.Println(err.Error())
	}

	//defer resp.Body.Close()
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		//return resp, err
		log.Println(err.Error())

	}
	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()
	return resp, err
}

/*
func (s *Database) Insert(collection string, document interface{}) (*http.Response, error) {
	issues := new([]Issue)
	clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	resp, err := s.sling.New().Post(path).QueryStruct(params).Receive(issues, clusterpointError)
	if err == nil {
		err = clusterpointError
	}
	return *issues, resp, err
}

func (s *Database) InsertByID(owner, repo string, params *IssueListParams) ([]Issue, *http.Response, error) {
	issues := new([]Issue)
	clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	resp, err := s.sling.New().Post(path).QueryStruct(params).Receive(issues, clusterpointError)
	if err == nil {
		err = clusterpointError
	}
	return *issues, resp, err
}

func (s *Database) Delete(owner, repo string, params *IssueListParams) ([]Issue, *http.Response, error) {
	issues := new([]Issue)
	clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	resp, err := s.sling.New().Delete(path).QueryStruct(params).Receive(issues, clusterpointError)
	if err == nil {
		err = clusterpointError
	}
	return *issues, resp, err
}

func (s *Database) Update(owner, repo string, params *IssueListParams) ([]Issue, *http.Response, error) {
	issues := new([]Issue)
	clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	resp, err := s.sling.New().Patch(path).QueryStruct(params).Receive(issues, clusterpointError)
	if err == nil {
		err = clusterpointError
	}
	return *issues, resp, err
}

func (s *Database) Replace(owner, repo string, params *IssueListParams) ([]Issue, *http.Response, error) {
	issues := new([]Issue)
	clusterpointError := new(ClusterpointError)
	path := fmt.Sprintf("repos/%s/%s/issues", owner, repo)
	resp, err := s.sling.New().Put(path).QueryStruct(params).Receive(issues, clusterpointError)
	if err == nil {
		err = clusterpointError
	}
	return *issues, resp, err
}
*/