package router


type Application struct {
	StaticDir map[string]string
}

func (app *Application) SetStaticPath(url string, path string) *Application {
	app.StaticDir[url] = path
	return app
}