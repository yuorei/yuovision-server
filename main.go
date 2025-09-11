package main

import (
	"fmt"

	"github.com/yuorei/video-server/app/driver/router"
)

func main() {
	router.NewHTTPRouter()
}

func init() {
	yuorei :=
		`
██╗   ██╗██╗   ██╗ ██████╗ ██████╗ ███████╗██╗
╚██╗ ██╔╝██║   ██║██╔═══██╗██╔══██╗██╔════╝██║
 ╚████╔╝ ██║   ██║██║   ██║██████╔╝█████╗  ██║
  ╚██╔╝  ██║   ██║██║   ██║██╔══██╗██╔══╝  ██║
   ██║   ╚██████╔╝╚██████╔╝██║  ██║███████╗██║
   ╚═╝    ╚═════╝  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚═╝
`
	fmt.Println(yuorei)
}
