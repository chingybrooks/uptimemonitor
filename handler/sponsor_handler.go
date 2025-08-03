package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"slices"
	"time"
	"uptimemonitor"
	"uptimemonitor/html"
	"uptimemonitor/pkg/version"
)

var initialSponsors = []uptimemonitor.Sponsor{
	{
		Name:  "AIR Labs",
		Url:   "https://airlabs.pl",
		Image: "/static/img/airlabs.png",
	},
}
var sponsors = []uptimemonitor.Sponsor{}

func init() {
	reloadSponsors()

	go func() {
		t := time.NewTicker(time.Minute * 60 * 12)

		for range t.C {
			reloadSponsors()
		}
	}()
}

func reloadSponsors() {
	req, err := http.NewRequest(http.MethodGet, "https://sponsors.uptimemonitor.dev", nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	var data []uptimemonitor.Sponsor
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return
	}

	sponsors = slices.Concat(initialSponsors, data)
}

func (*Handler) ListSponsors() http.HandlerFunc {
	layout := template.Must(template.ParseFS(html.FS, "layout.html"))
	sponsor := template.Must(template.ParseFS(html.FS, "sponsor.html"))

	type data struct {
		Sponsors []uptimemonitor.Sponsor
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			layout.ExecuteTemplate(w, "sponsors", struct {
				Version string
			}{
				Version: version.Version,
			})
			return
		}

		d := data{
			Sponsors: sponsors,
		}

		sponsor.Execute(w, d)
	}
}
