package main

import (
	"coletor-SIPPAG/coletor"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Erro ao ler arquivo .env")
	}
}

func main() {
	coletor.ColetarInformacoesPortarias()
}
