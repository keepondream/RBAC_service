package utils

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/render"
	"github.com/go-pg/pg/v10"
)

// Render 原生http统一输出,适用于chi框架
func Render(w http.ResponseWriter, r *http.Request, code int, modResp ...ModResponse) {
	resp := Response{}
	render.Status(r, code)
	resp.Msg = http.StatusText(code)

	for _, fn := range modResp {
		fn(&resp)
	}

	if resp.Data != nil {
		render.JSON(w, r, resp.Data)
		return
	}

	if resp.Err != nil {
		log.Printf("render response err: %v \n", resp.Err)
		resp.Msg = resp.Err.Error()
		if e, ok := resp.Err.(net.Error); ok && e.Timeout() {
			render.Status(r, http.StatusGatewayTimeout)
		} else if resp.Err == sql.ErrNoRows {
			render.Status(r, http.StatusNotFound)
		} else if e, ok := resp.Err.(pg.Error); ok && e.IntegrityViolation() {
			render.Status(r, http.StatusBadRequest)
		}
	}

	render.JSON(w, r, resp)
}
