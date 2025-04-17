package handlers

import (
	"net/http"

	"github.com/dtg-lucifer/goth-stack-starter/utils"
	"github.com/dtg-lucifer/goth-stack-starter/views/home"
)

func Index(w http.ResponseWriter, r *http.Request) error {
	return utils.Render(w, r, home.Home())
}
