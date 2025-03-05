package main

import (
	"mycareerapp/internal/bootstrap"
)

func main() {

	// cfg, err := env.New()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(cfg.AppPort)

	if err := bootstrap.Start(); err != nil {
		panic(err)
	}
}
