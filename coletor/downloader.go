package coletor

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func baixarArquivos(localArmazenamento string, url string) error {
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Cria o arquivo
	out, err := os.Create(localArmazenamento)
	if err != nil {
		return err
	}
	defer out.Close()

	// Escreve no corpo do arquivo
	_, err = io.Copy(out, resp.Body)
	return err
}
