package main

import (
	"coletor-SIPPAG/coletor"
	"coletor-SIPPAG/entidades"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Erro ao ler arquivo .env")
	}
}

func main() {
	canal := make(chan entidades.URLELocalDeArmazenamento)

	go coletor.ColetarInformacoesPortarias(canal)

	for c := range canal {
		go coletor.BaixarArquivo(c.LocalArmazenamento, c.URL)
	}
}
