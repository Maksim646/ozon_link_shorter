package main

import (
	"fmt"

	"github.com/your-username/ozon_link_shorter/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
