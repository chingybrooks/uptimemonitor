package handler

import (
	"html/template"
	"net/http"
	"uptimemonitor"
	"uptimemonitor/html"
	"uptimemonitor/pkg/version"
)

func (h *Handler) HomePage() http.HandlerFunc {
	tmpl := template.Must(template.ParseFS(html.FS, "layout.html", "app.html", "home.html"))

	type data struct {
		Version string
		User    uptimemonitor.User
	}

	return func(w http.ResponseWriter, r *http.Request) {
		count := h.Store.CountMonitors(r.Context())
		if count == 0 {
			http.Redirect(w, r, "/new", http.StatusSeeOther)
			return
		}

		tmpl.Execute(w, data{
			Version: version.Version,
			User:    getUserFromRequest(r),
		})
	}
}
