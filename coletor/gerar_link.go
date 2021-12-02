package coletor

import (
	"coletor-SIPPAG/entidades"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func gerarURLELocalDeArmazenamento(portarias entidades.Portarias, siglaInstituicao string, canal chan entidades.URLELocalDeArmazenamento) {

	anoInicio := 2021 // Inicio 2017
	anoAtual := time.Now().Year()

	for ano := anoInicio; ano <= anoAtual; ano++ {

		diretorio := fmt.Sprintf("%s/%s/%d", os.Getenv("DIRETORIO"), siglaInstituicao, anoInicio)
		verificaSeDiretorioExiste(diretorio)

		for _, portaria := range portarias.Data.Results {
			url := fmt.Sprintf("https://sippag.%s.edu.br/portarias/visualizar/?ano=%d&numero=%d&hash=%s",
				siglaInstituicao, ano, portaria.Portaria.Numero, portaria.Assinatura.Hash)

			localArmazenamento := fmt.Sprintf("%s/https_DOISpont__baraduplas_sippag.%s.edu.brbarraportariasbarravisualizarinterrogacaoano=%d&numero=%d&hash=%s.pdf",
				diretorio, siglaInstituicao, ano, portaria.Portaria.Numero, portaria.Assinatura.Hash)

			linkElocal := entidades.URLELocalDeArmazenamento{URL: url, LocalArmazenamento: localArmazenamento}

			time.Sleep(time.Second * 5)
			canal <- linkElocal
		}
	}

}

func verificaSeDiretorioExiste(diretorio string) {
	if _, err := os.Stat(diretorio); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(diretorio, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	err := os.MkdirAll(diretorio, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}
