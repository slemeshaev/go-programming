// Exercise 7.9: Use the html/template package (§4.6) to replace `printTracks` with a function
// that displays the tracks as an HTML table. Use the solution to the previous exercise to arrange
// that each click on a column head makes an HTTP request to sort the table.

package main

import (
	"log"
	"net/http"
	"sort"
	"text/template"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func tracks() []*Track {
	return []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type byTitle []*Track

func (x byTitle) Len() int {
	return len(x)
}

func (x byTitle) Less(i, j int) bool {
	return x[i].Title < x[j].Title
}

func (x byTitle) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type less func(x, y *Track) bool

type Column struct {
	name string
	f    less
}

func colTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func colArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func colAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func colYear(x, y *Track) bool {
	return x.Year < y.Year
}

func colLength(x, y *Track) bool {
	return x.Length < y.Length
}

type byColumns struct {
	tracks  []*Track
	columns []*Column
}

func sortByColumns(t []*Track, f ...less) *byColumns {
	bc := &byColumns{
		tracks: t,
	}

	for _, foo := range f {
		bc.columns = append(bc.columns, &Column{f: foo})
	}

	return bc
}

func (x byColumns) Len() int {
	return len(x.tracks)
}

func (x byColumns) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func (x byColumns) Less(i, j int) bool {
	if len(x.columns) == 0 {
		return false
	}

	a, b := x.tracks[i], x.tracks[j]
	var k int

	for k = 0; k < len(x.columns)-1; k++ {
		f := x.columns[k].f
		switch {
		case f(a, b):
			return true
		case f(b, a):
			return false
		}
	}

	return x.columns[k].f(a, b)
}

var tplt = template.Must(template.New("trackTable").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <title>ex7.9</title>
      <style>
        table {
	      border-collapse: collapse;
        }
        td, th {
	      border: solid 1px;
	      padding: 0.5em;
          text-align: right;
        }
      </style>
  </head>
  <body>
    <table>
      <tr>
	    <th><a href="./?by=title">Title</a></th>
	    <th><a href="./?by=artist">Artist</a></th>
	    <th><a href="./?by=album">Album</a></th>
	    <th><a href="./?by=year">Year</a></th>
	    <th><a href="./?by=length">Length</a></th>
	  </tr>
      {{range .}}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Artist}}</td>
        <td>{{.Album}}</td>
        <td>{{.Year}}</td>
        <td>{{.Length}}</td>
      </tr>
      {{end}}
    </table>
  </body>
</html>
`))

func (x *byColumns) doSort(w http.ResponseWriter, req *http.Request) {
	col := req.URL.Query().Get("by")
	if col != "" {
		x.selected(col)
		sort.Sort(x)
	}

	if err := tplt.Execute(w, x.tracks); err != nil {
		log.Fatal(err)
	}
}

func (x *byColumns) selected(s string) {
	var f less

	switch s {
	case "title":
		f = colTitle
	case "artist":
		f = colArtist
	case "album":
		f = colAlbum
	case "year":
		f = colYear
	case "length":
		f = colLength
	default:
		s = "title"
		f = colTitle
	}

	for i, c := range x.columns {
		if c.name == s {
			if i != 0 {
				x.columns[0], x.columns[i] = x.columns[i], x.columns[0]
			}

			return
		}
	}

	x.columns = append(x.columns, &Column{f: f, name: s})
	i := len(x.columns) - 1

	if i != 0 {
		x.columns[0], x.columns[i] = x.columns[i], x.columns[0]
	}
}

func main() {
	t := sortByColumns(tracks())
	http.HandleFunc("/", t.doSort)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
