package main

import (
	"github.com/codegangsta/martini"




)
func all()string{
    name:="chicco"
	return name

}


func main() {
	m := martini.Classic()
	m.Get("/", func() string {
			return all()
		})
	m.Run()

}

