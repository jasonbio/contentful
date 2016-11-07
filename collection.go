package contentful

import "github.com/jasonbio/contentful/entities"

type Collection struct {
	Sys      entities.Sys             `json:"sys"`
	Total    int                      `json:"total"`
	Skip     int                      `json:"skip"`
	Limit    int                      `json:"limit"`
	Items    []interface{}            `json:"items"`
	Includes map[string][]interface{} `json:"includes"`
}
