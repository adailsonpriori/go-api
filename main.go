package main

import (
	"fmt"

	"github.com/adailsonpriori/go-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
