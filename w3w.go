package w3w

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"responses"
	"strings"
)

/*2OY1ZVS1*/
type W3W struct {
	APIKey      string
	CallBack    string
	Error       error
	JSONP       bool
	Lang        string
	Format      string
	QueryString bool
	Display     bool
}

func (w *W3W) makeW3WError(err responses.Error) {
	if w.Error != nil {
		w.Error = errors.New(w.Error.Error() + " " + err.Status.Message)
	} else {
		w.Error = errors.New(err.Status.Message)
	}
}

func (w *W3W) makeErrorObject(body io.Reader, Err *responses.Error) {
	if err := json.NewDecoder(body).Decode(Err); err != nil {
		if errr := xml.NewDecoder(body).Decode(Err); errr != nil {
			Err.Status.Message = "Not found"
			Err.Status.Status = 200
			Err.Status.Code = 404
			Err.Status.Reason = "OK"

		}
	}
}

/*curl --request GET \
  --url 'https://api.what3words.com/v2/forward?addr=index.home.raft&key=MY-API-KEY&lang=en&format=json&display=full'
*/
func (w *W3W) Forward(words string) (responses.Forward, responses.Error) {
	var Fwd responses.Forward
	var Err responses.Error

	escaped := url.QueryEscape(words)
	fwdUrl := fmt.Sprintf("https://api.what3words.com/v2/forward?addr=%s&key=%s", escaped, w.APIKey)
	//optional  param handling
	if w.Lang != "" && w.Lang != "en" {
		fwdUrl = fwdUrl + "&lang=" + w.Lang
	}
	if w.Display {
		fwdUrl = fwdUrl + "&display=terse"
	}
	if w.Format != "" && w.Format != "json" {
		fwdUrl = fwdUrl + "&format=" + w.Format
	}
	if w.CallBack != "" && (w.Format == "json" || w.Format == "geojson") {
		fwdUrl = fwdUrl + "&callback=" + w.CallBack
	}

	req, err := http.NewRequest("GET", fwdUrl, nil)

	if err != nil {
		body := strings.NewReader(err.Error())
		w.makeErrorObject(body, &Err)
		w.makeW3WError(Err)
		return Fwd, Err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		body := strings.NewReader(err.Error())
		w.makeErrorObject(body, &Err)
		w.makeW3WError(Err)
		return Fwd, Err
	}

	if err := json.NewDecoder(resp.Body).Decode(&Fwd); err != nil {
		w.makeErrorObject(resp.Body, &Err)
	}

	if Fwd.Words == "" {
		w.makeErrorObject(resp.Body, &Err)
		w.makeW3WError(Err)
	}

	resp.Body.Close()

	return Fwd, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/reverse?coords=51.521251%2C-0.203586&key=MY-API-KEY&lang=en&format=json&display=full'
*/
func (w *W3W) Reverse(rawcoords string) (responses.Reverse, responses.Error) {
	var Rev responses.Reverse
	var Err responses.Error

	coords := url.QueryEscape(rawcoords)
	revUrl := fmt.Sprintf("https://api.what3words.com/v2/reverse?coords=%s&key=%s", coords, w.APIKey)
	//optional  param handling
	if w.Lang != "" && w.Lang != "en" {
		revUrl = revUrl + "&lang=" + w.Lang
	}
	if w.Display {
		revUrl = revUrl + "&display=terse"
	}
	if w.Format != "" && w.Format != "json" {
		revUrl = revUrl + "&format=" + w.Format
	}
	if w.CallBack != "" && (w.Format == "json" || w.Format == "geojson") {
		revUrl = revUrl + "&callback=" + w.CallBack
	}

	req, err := http.NewRequest("GET", revUrl, nil)

	if err != nil {
		body := strings.NewReader(err.Error())
		w.makeErrorObject(body, &Err)
		w.makeW3WError(Err)
		return Rev, Err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		body := strings.NewReader(err.Error())
		w.makeErrorObject(body, &Err)
		w.makeW3WError(Err)
		return Rev, Err
	}

	if err := json.NewDecoder(resp.Body).Decode(&Rev); err != nil {
		w.makeErrorObject(resp.Body, &Err)
	}

	if Rev.Words == "" {
		w.makeErrorObject(resp.Body, &Err)
		w.makeW3WError(Err)
	}

	resp.Body.Close()

	return Rev, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/autosuggest?addr=index.home.r&key=MY-API-KEY&focus=51.521251%2C-0.203586&clip=51.521251%2C-0.203586%2C5&lang=en&format=json&display=full'
*/
func (w *W3W) AutoSuggest(addr string) (responses.AutoSuggest, responses.Error) {
	var Auto responses.AutoSuggest
	var Err responses.Error
	return Auto, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/standardblend?addr=index.home.r&key=MY-API-KEY&focus=51.521251%2C-0.203586&lang=en&format=json'
*/
func (w *W3W) StandardBlend(addr string) (responses.StandardBlend, responses.Error) {
	var Fwd responses.StandardBlend
	var Err responses.Error
	return Fwd, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/grid?bbox=52.208867%2C0.117540%2C52.207988%2C0.116126&key=MY-API-KEY&format=json'
*/
func (w *W3W) Grid(bbox string) (responses.Grid, responses.Error) {
	var Fwd responses.Grid
	var Err responses.Error
	return Fwd, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/languages?key=MY-API-KEY'
*/
func (w *W3W) GetLanguages() (responses.Grid, responses.Error) {
	var Fwd responses.Grid
	var Err responses.Error
	return Fwd, Err
}
