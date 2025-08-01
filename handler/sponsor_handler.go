package handler

import (
	"encoding/json"
	"html/template"
	"net/http"
	"uptimemonitor"
	"uptimemonitor/html"
)

func (*Handler) ListSponsors() http.HandlerFunc {
	layout := template.Must(template.ParseFS(html.FS, "layout.html"))
	sponsor := template.Must(template.ParseFS(html.FS, "sponsor.html"))

	type data struct {
		Sponsors []uptimemonitor.Sponsor
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("HX-Request") != "true" {
			layout.ExecuteTemplate(w, "sponsors", nil)
			return
		}

		d := data{
			Sponsors: []uptimemonitor.Sponsor{
				{
					Name:  "AIR Labs",
					Url:   "https://airlabs.pl",
					Image: "/static/img/airlabs.svg",
				},
			},
		}

		req, err := http.NewRequest(http.MethodGet, "https://sponsors.uptimemonitor.dev", nil)
		if err != nil {
			sponsor.Execute(w, d)
			return
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			sponsor.Execute(w, d)
			return
		}

		var sponsors []uptimemonitor.Sponsor
		err = json.NewDecoder(res.Body).Decode(&sponsors)
		if err != nil {
			sponsor.Execute(w, d)
			return
		}

		sponsor.Execute(w, data{
			Sponsors: sponsors,
		})
	}
}
