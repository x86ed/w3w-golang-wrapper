package w3w

import (
	"fmt"
	"responses"
	"testing"
)

func TestForward(t *testing.T) {
	t.Log("Testing the forward API wrapper")
	w := W3W{"2OY1ZVS1", "fartFunction", nil, false, "en", "", true, false}
	fwd, err := w.Forward("index.home.raft")
	f1 := responses.Forward{
		Crs: responses.Crs{
			Properties: responses.Properties{
				Type: "ogcwkt",
				Href: "http://spatialreference.org/ref/epsg/4326/ogcwkt/",
			},
			Type: "link",
		},
		Bounds: responses.Bounds{
			Southwest: responses.Coord{
				Lng: -0.203607,
				Lat: 51.521236,
			},
			Northeast: responses.Coord{
				Lng: -0.203564,
				Lat: 51.521263,
			},
		},
		Words:    "index.home.raft",
		Map:      "http://w3w.co/index.home.raft",
		Language: "en",
		Geometry: responses.Coord{
			Lng: -0.203586,
			Lat: 51.52125,
		},
		Status: responses.Status{
			Code:    0,
			Reason:  "OK",
			Status:  0,
			Message: "",
		},
		Thanks: "Thanks from all of us at index.home.raft for using a what3words API",
	}
	e1 := responses.Error{
		Status: responses.Status{
			Code:    404,
			Reason:  "OK",
			Status:  200,
			Message: "Not found",
		},
		Thanks: "",
	}

	errBlank := responses.Error{}
	if err != errBlank {
		t.Errorf("error found")
	}
	if f1 != fwd {
		t.Errorf("%+v doesn't equal %+v.", f1, fwd)
	}
	w.APIKey = ""
	fwd2, err2 := w.Forward("i")
	if w.Error == nil {
		t.Error("no error was found")
	}
	if fwd2.Words != "" {
		t.Error("no words should have been returned")
	}
	if err2 != e1 {
		t.Error("errors not returned")
	}
}

func TestReverse(t *testing.T) {
	t.Log("Testing the reverse API wrapper")
	w := W3W{"2OY1ZVS1", "fartFunction", nil, false, "en", "", true, false}
	rev, err := w.Reverse("51.521251,-0.203586")
	errBlank := responses.Error{}
	if err != errBlank {
		fmt.Printf("%+v", rev)
		fmt.Printf("%+v", err)
		t.Errorf("error found")

	}
}

func TestAutoSuggest(t *testing.T) {
	t.Log("Testing the reverse API wrapper")
}

func TestStandardBlend(t *testing.T) {
	t.Log("Testing the standard blend API wrapper")
}

func TestGrid(t *testing.T) {
	t.Log("Testing the grid API wrapper")
}

func TestGetLanguages(t *testing.T) {
	t.Log("Testing the languages API wrapper")
}
