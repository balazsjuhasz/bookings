package main

import (
	"testing"

	"github.com/balazsjuhasz/bookings/internal/config"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing, test is passed
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
