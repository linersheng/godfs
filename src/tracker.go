package main

import (
	"app"
	"flag"
	"libtracker"
	"os"
	"path/filepath"
	"runtime"
	"util/file"
	"util/logger"
	"validate"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	abs, _ := filepath.Abs(os.Args[0])
	s, _ := filepath.Split(abs)
	s = file.FixPath(s)
	var confPath = flag.String("c", s+string(filepath.Separator)+".."+string(filepath.Separator)+"conf"+string(filepath.Separator)+"tracker.conf", "custom config file")
	flag.Parse()
	logger.Info("using config file:", *confPath)
	m, e := file.ReadPropFile(*confPath)
	if e == nil {
		validate.Check(m, 2)
		for k, v := range m {
			logger.Debug(k, "=", v)
		}
		app.RUN_WITH = 2
		libtracker.StartService(m)
	} else {
		logger.Fatal("error read file:", e)
	}
}
