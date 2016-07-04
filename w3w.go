package w3w

import (
	"encoding/json"
	"errors"
	"github.com/x86ed/w3w-golang-wrapper/responses"
	"net/http"
	"net/url"
)

/**/
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
		w.Error = errors.New(w.Error + " " + err.Status.Message)
	} else {
		w.Error = errors.New(err.Status.Message)
	}
}

func (w *W3W) makeErrorObject(body io.ReadCloser, Err *responses.Error) {
	if err := json.NewDecoder(body).Decode(Err); err != nil {
		if errr := xml.NewDecoder(body).Decode(Err); errr != nil {
			Err.Status = "Not found"
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
	fwdUrl := fmt.Sprintf("https://api.what3words.com/v2/forward?addr=%s&key=%s", words, w.APIKey)
	//optional  param handling
	if w.Lang && w.Lang != "en" {
		fwdUrl = fwdUrl + "&lang=" + w.Lang
	}
	if w.Display {
		fwdUrl = fwdUrl + "&display=terse"
	}
	if w.Format && w.Format != "json" {
		fwdUrl = fwdUrl + "&format=" + w.Format
	}
	if w.CallBack && (w.Format == "json" || w.Format == "geojson") {
		fwdUrl = fwdUrl + "&callback=" + w.CallBack
	}

	req, err := http.NewRequest("GET", fwdUrl, nil)

	if err != nil {
		w.makeW3WError(err)
		return Fwd, Err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		w.makeW3WError(err)
		return Fwd, Err
	}

	if err := json.NewDecoder(resp.Body).Decode(&Fwd); err != nil {
		Err = w.makeErrorObject(resp.Body)
	}

	resp.Body.Close()

	return Fwd, Err
}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/reverse?coords=51.521251%2C-0.203586&key=MY-API-KEY&lang=en&format=json&display=full'
*/
func (w *W3W) Reverse(coords string) (responses.Reverse, responses.Error) {

}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/autosuggest?addr=index.home.r&key=MY-API-KEY&focus=51.521251%2C-0.203586&clip=51.521251%2C-0.203586%2C5&lang=en&format=json&display=full'
*/
func (w *W3W) AutoSuggest(addr string) (responses.AutoSuggest, responses.Error) {

}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/standardblend?addr=index.home.r&key=MY-API-KEY&focus=51.521251%2C-0.203586&lang=en&format=json'
*/
func (w *W3W) StandardBlend(addr string) (responses.StandardBlend, responses.Error) {

}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/grid?bbox=52.208867%2C0.117540%2C52.207988%2C0.116126&key=MY-API-KEY&format=json'
*/
func (w *W3W) Grid(bbox string) (responses.Grid, responses.Error) {

}

/*
curl --request GET \
  --url 'https://api.what3words.com/v2/languages?key=MY-API-KEY'
*/
func (w *W3W) GetLanguages() (responses.Grid, responses.Error) {

}
