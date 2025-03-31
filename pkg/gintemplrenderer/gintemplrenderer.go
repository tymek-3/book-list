package gintemplrenderer

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

var Default = &HTMLTemplRenderer{}

type HTMLTemplRenderer struct {
}

func (r *HTMLTemplRenderer) Instance(s string, d any) render.Render {
	templData := d.(templ.Component)
	return &Renderer{
		Ctx:       context.Background(),
		Status:    -1,
		Component: templData,
	}
}

func New(ctx context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{
		Ctx:       ctx,
		Status:    status,
		Component: component,
	}
}

type Renderer struct {
	Ctx       context.Context
	Status    int
	Component templ.Component
}

func (t Renderer) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	if t.Status != -1 {
		w.WriteHeader(t.Status)
	}
	if t.Component != nil {
		return t.Component.Render(t.Ctx, w)
	}
	return nil
}

func (t Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
