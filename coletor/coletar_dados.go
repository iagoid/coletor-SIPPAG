package coletor

import (
	"coletor-SIPPAG/entidades"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func ColetarInformacoesPortarias(canal chan entidades.URLELocalDeArmazenamento) {

	institutos := []string{
		// "ifrs",
		// "ifce",
		// "ifnmg",
		// "IFBaiano",
		"ifal",
	}

	for _, siglaInstituicao := range institutos {

		resp, err := http.Get("https://sippag-web." + siglaInstituicao + ".edu.br/api/v1/portaria?ano=2021&page=0")
		if err != nil {
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}

		var portarias entidades.Portarias
		err = json.Unmarshal(body, &portarias)
		if err != nil {
			panic(err)
		}

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

				time.Sleep(time.Second * 3)
				canal <- linkElocal
			}
		}
	}
	close(canal)
}
