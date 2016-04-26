//Project : Hey Asha!
//Copyright (C) Balamurali Pandranki - All Rights Reserved
//Unauthorized copying of this file, via any medium is strictly prohibited
//Proprietary and confidential
//Written by Balamurali Pandranki <balamurali@live.com>, 26/4/2016 2:16 PM
package datatypes

var (
	ErrNotAcceptable        = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
	ErrInternalServer       = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
)

type ApiDefaultResponse struct {
	Version string `json:"version"`
	Message	string `json:"message"`
}

type Errors struct {
	Errors []*Error `json:"errors"` // We use an array of *Error to set common errors in variables. More on that later.
}

type Error struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

