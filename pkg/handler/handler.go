package handler

import (
	"net/http"

	"github.com/shan251197/bookings/pkg/config"
	"github.com/shan251197/bookings/pkg/models"
	"github.com/shan251197/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}

}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	name := "suresh"
	m.App.Session.Put(r.Context(), "Name", name)

	render.RenderTemplate(rw, "home.page.html", &models.TemplateData{})
}
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, world"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	namefromhome := m.App.Session.GetString(r.Context(), "Name")

	stringMap["remote_ip"] = remoteIP
	stringMap["Name"] = namefromhome

	render.RenderTemplate(rw, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})

}