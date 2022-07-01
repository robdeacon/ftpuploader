package main


import (
	"flag"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var (

	// Scs ...
	Scs = spew.ConfigState{
		Indent:   "    ",
		MaxDepth: 7,
	}
)

func main(){

	fmt.Println("Hello world!")
	
	var username string
	// var password string
	// var host string
	// var src string
	// var dest string
	flag.StringVar(&username, "username", "", "SFTP username")
	flag.StringVar(&username, "u", "", "SFTP username")
	
	// flag.StringVar(&password, "password", "", "SFTP password")
	// flag.StringVar(&password, "p", "", "SFTP password")

	// flag.StringVar(&host, "host", "", "SFTP hostname server")
	// flag.StringVar(&host, "h", "", "SFTP hostname server")

	// flag.StringVar(&src, "source", "", "source file path")
	// flag.StringVar(&src, "s", "", "source file path")

	// flag.StringVar(&dest, "destination", "", "destination file path")
	// flag.StringVar(&dest, "d", "", "destination file path")

	flag.Parse()
	
	fmt.Println("username",username)
	
	Scs.Dump(username)
	
	
}