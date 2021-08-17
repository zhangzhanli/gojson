package gojson

import (
	"fmt"
	"runtime"
)

const (
	logo = `
   
  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄        ▄ 
 ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░▌      ▐░▌
 ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀▀█░█▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌▐░▌░▌     ▐░▌
 ▐░▌          ▐░▌       ▐░▌      ▐░▌    ▐░▌          ▐░▌       ▐░▌▐░▌▐░▌    ▐░▌
 ▐░▌ ▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌      ▐░▌    ▐░█▄▄▄▄▄▄▄▄▄ ▐░▌       ▐░▌▐░▌ ▐░▌   ▐░▌
 ▐░▌▐░░░░░░░░▌▐░▌       ▐░▌      ▐░▌    ▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░▌  ▐░▌  ▐░▌
 ▐░▌ ▀▀▀▀▀▀█░▌▐░▌       ▐░▌      ▐░▌     ▀▀▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌▐░▌   ▐░▌ ▐░▌
 ▐░▌       ▐░▌▐░▌       ▐░▌      ▐░▌              ▐░▌▐░▌       ▐░▌▐░▌    ▐░▌▐░▌
 ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌ ▄▄▄▄▄█░▌     ▄▄▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌▐░▌     ▐░▐░▌
 ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░▌    ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░▌      ▐░░▌
  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀      ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀        ▀▀ 																	  
`
	app   = "gojson"
	major = 2
	minor = 0
)

func Version() string {
	return fmt.Sprintf("%s  ver %d.%d (Go runtime %s).", logo, major, minor, runtime.Version())
}
