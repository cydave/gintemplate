package server

import (
	"html/template"
	"net/http"

	"github.com/cydave/staticfs"
	"github.com/gin-gonic/gin"

	"github.com/cydave/gintemplate/internal/assets"
	"github.com/cydave/gintemplate/internal/config"
	"github.com/cydave/gintemplate/internal/controllers"
)

func configureStaticFS(r *gin.Engine) error {
	s := staticfs.New(assets.Static).WithRootAliases()
	s.Configure(r)
	return nil
}

func configureTemplating(r *gin.Engine) error {
	funcMaps := template.FuncMap{}
	templ := template.New("").Funcs(funcMaps)
	templ, err := templ.ParseFS(assets.Templates, "templates/*.tmpl")
	if err != nil {
		return err
	}
	r.SetHTMLTemplate(templ)
	return nil
}

func Init() (*gin.Engine, error) {
	cfg := config.Get()
	if env := cfg.GetString("environment"); env == "" || env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	if err := configureStaticFS(r); err != nil {
		return nil, err
	}
	if err := configureTemplating(r); err != nil {
		return nil, err
	}

	// Register controllers / routes here.
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Title": "Hello World",
		})
	})

	x := new(controllers.ExampleController)
	r.GET("/example", x.Get)
	r.POST("/example", x.Post)
	r.PUT("/example", x.Put)

	return r, nil
}
