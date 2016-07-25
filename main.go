/*
Fresh is a command line tool that builds and (re)starts your web application everytime you save a go or template file.

If the web framework you are using supports the Fresh runner, it will show build errors on your browser.

It currently works with Traffic (https://github.com/pilu/traffic), Martini (https://github.com/codegangsta/martini) and gocraft/web (https://github.com/gocraft/web).

Fresh will watch for file events, and every time you create/modifiy/delete a file it will build and restart the application.
If `go build` returns an error, it will logs it in the tmp folder.

Traffic (https://github.com/pilu/traffic) already has a middleware that shows the content of that file if it is present. This middleware is automatically added if you run a Traffic web app in dev mode with Fresh.
*/
package main

import (
	"flag"

	"github.com/c2h5oh/fresh/runner"
)

func main() {
	var watchList, excludeList runner.Multiflag

	configPath := flag.String("c", "", "config file path")
	buildArgs := flag.String("b", "", "build command line arguments")
	runArgs := flag.String("r", "", "run command line arguments")
	buildPath := flag.String("p", "", "build at path")
	outputBuildPath := flag.String("o", "runner-build", "output build path")
	flag.Var(&watchList, "w", "watch path (recursive), repeat multiple times to watch multiple paths")
	flag.Var(&excludeList, "e", "exclude path (recursive), repeat multiple times to exclude multiple paths")
	flag.Parse()

	runner.Start(configPath, buildArgs, runArgs, buildPath, outputBuildPath, watchList, excludeList)
}
