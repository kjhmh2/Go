package server

import(
    "net/http"
	"github.com/go-martini/martini"
	"github.com/unrolled/render"
)

func Run(port string) {
    // default setting
    m := martini.Classic()
    // set path
    m.Use(martini.Static("static"));
    // Get request
    m.Get("/", func(res http.ResponseWriter, req *http.Request) {
        r := render.New()
        r.HTML(res, http.StatusOK, "jigsaw", "World");
    })
    // run
    m.RunOnAddr(":" + port)
}