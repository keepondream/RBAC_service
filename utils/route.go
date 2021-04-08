package utils

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RouteList struct {
	Name   string `json:"route_name"`
	Uri    string `json:"route_uri"`
	Method string `json:"route_method"`
}

var RouteNum = 0

var AllRoutes = make(map[int]RouteList)

func AddRoute(router *gin.RouterGroup, uri, method, name string, handle ...gin.HandlerFunc) {
	prefix := router.BasePath()

	method = strings.ToUpper(method)

	AllRoutes[RouteNum] = RouteList{
		Name:   name,
		Uri:    fmt.Sprintf("%s%s", prefix, uri),
		Method: method,
	}
	RouteNum++

	switch method {
	case http.MethodGet:
		router.GET(uri, handle...)
	case http.MethodPost:
		router.POST(uri, handle...)
	case http.MethodPatch:
		router.PATCH(uri, handle...)
	case http.MethodPut:
		router.PUT(uri, handle...)
	case http.MethodDelete:
		router.DELETE(uri, handle...)
	default:
		log.Fatalf("The route type is not recognized method : %s , uri : %s , name : %s \n", method, uri, name)
	}
}
