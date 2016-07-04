package responses

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Coord struct {
	Lng float32 `json:"lng" `
	Lat float32 `json:"lat"`
}

type Properties struct {
	Type string `json:"type"`
	Href string `json:"href"`
}

type Crs struct {
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}

type Bounds struct {
	Southwest Coord `json:"southwest"`
	Northeast Coord `json:"northeast"`
}

type Status struct {
	Code    int    `json:"code"`
	Reason  string `json:"reason,omitempty"`
	Status  int    `json:"code,omitempty"`
	Message string `json:"message"`
}

type Suggestion struct {
	Distance int     `json:"distance"`
	Rank     int     `json:"rank"`
	Words    string  `json:"words"`
	Score    float64 `json:"score"`
	Place    string  `json:"place"`
	Geometry Coord   `json:"geometry"`
	Country  String  `json:"country"`
}

type Blend struct {
	Distance int    `json:"distance"`
	Rank     int    `json:"rank"`
	Words    string `json:"words"`
	Language string `json:"languages"`
	Place    string `json:"place"`
	Geometry Coord  `json:"geometry"`
	Country  string `json:"country"`
}

type Line struct {
	Start Coord `json:"start"`
	End   Coord `json:"end"`
}

type Language struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
}

type Forward struct {
	XMLName  xml.Name `json:"-" xml:"response"`
	Crs      Crs      `json:"crs"`
	Bounds   Bounds   `json:"bounds"`
	Words    string   `json:"words"`
	Map      string   `json:"map"`
	Language string   `json:"language"`
	Geometry Coord    `json:"geometry"`
	Status   Status   `json:"status"`
	Thanks   string   `json:"thanks"`
}

func (f *Forward) toString() {
	return fmt.Printf("%+v\n", f)
}

type Reverse Forward

func (r *Reverse) toString() {
	return fmt.Printf("%+v\n", r)
}

type AutoSuggest struct {
	XMLName     xml.Name     `json:"-" xml:"response"`
	Suggestions []Suggestion `json:"suggestions"`
	Status      Status       `json:"status"`
	Thanks      string       `json:"thanks"`
}

func (a *AutoSuggest) toString() {
	return fmt.Printf("%+v\n", a)
}

type StandardBlend struct {
	XMLName xml.Name `json:"-" xml:"response"`
	Blends  []Blend  `json:"blends"`
	Status  Status   `json:"status"`
	Thanks  string   `json:"thanks"`
}

func (s *StandardBlend) toString() {
	return fmt.Printf("%+v\n", s)
}

type Grid struct {
	XMLName xml.Name `json:"-" xml:"response"`
	Lines   []Line   `json:"lines"`
	Status  Status   `json:"status"`
	Thanks  string   `json:"thanks"`
}

func (g *Grid) toString() {
	return fmt.Printf("%+v\n", g)
}

type Languages struct {
	XMLName   xml.Name   `json:"-" xml:"languages"`
	Languages []Language `json:"languages"`
	Status    Status     `json:"status"`
	Thanks    string     `json:"thanks"`
}

func (l *Languages) toString() {
	return fmt.Printf("%+v\n", l)
}

type Error struct {
	XMLName xml.Name `json:"-" xml:"error"`
	Status  Status
	Thanks  string
}

func (e *Error) toString() {
	return fmt.Printf("%+v\n", e)
}

type Response interface {
	ToGeoJSON() string
	ToJSON() string
	ToJSONP() string
	ToKML() string
	ToString() string
	ToXML() string
}
