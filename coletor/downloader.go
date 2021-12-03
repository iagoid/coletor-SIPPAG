package coletor

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func BaixarArquivo(localArmazenamento string, url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Cria o arquivo
	out, err := os.Create(localArmazenamento)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	// Escreve no corpo do arquivo
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(url)

}
