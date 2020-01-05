package web

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
)

type App struct {
	*Config
	http *fasthttp.Server
	tls  *fasthttp.Server
}

func NewApp(c *Config, r phi.Router) (app App, err error) {
	h := &fasthttp.Server{
		Handler:      r.ServeFastHTTP,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		Name:         "rw3",
	}

	app = App{Config: c, http: h}

	return app, nil
}

func (a *App) ListenAndServe() (errs chan error) {
	go func() {
		errs <- a.http.ListenAndServe(a.Config.Addr)
	}()

	return errs
}

func (a *App) Shutdown() (err error) {
	if a.http != nil {
		err = a.http.Shutdown()
		if err != nil {
			return err
		}
	}
	return nil
}