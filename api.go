package main

import (
	"encoding/xml"
	"net/http"
)

// Note that APIKEY is defined elsewhere and not uploaded.
const (
	APIURL = "https://www.dictionaryapi.com/api/v1/references/collegiate/xml/"
)

func Definition(word string) ([]string, error) {
	// Download the raw xml data
	resp, err := http.Get(APIURL + word + "?key=" + APIKEY)
	if err != nil {
		return nil, err
	}

	// API result's format
	type Response struct {
		Definitions []string `xml:"entry>def>dt>un"`
	}

	// Parse out the defintions from the xml
	v := &Response{}
	dec := xml.NewDecoder(resp.Body)
	err = dec.Decode(v)
	if err != nil {
		return nil, err
	}

	return v.Definitions, nil
}
