package main

import api "github.com/omidhaqi/clean-web-api/api"


// @SecurityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main()  {
	api.InitServer()
}