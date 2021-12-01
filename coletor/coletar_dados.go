package coletor

import (
	"coletor-SIPPAG/entidades"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ColetarInformacoesPortarias() {

	institutos := []string{
		// "ifrs",
		"ifce",
		"ifnmg",
		"IFBaiano",
		"ifal",
	}

	canal := make(chan entidades.URLELocalDeArmazenamento)

	for _, siglaInstituicao := range institutos {

		resp, err := http.Get("https://sippag-web." + siglaInstituicao + ".edu.br/api/v1/portaria?ano=2021&page=0")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}

		var portarias entidades.Portarias
		err = json.Unmarshal(body, &portarias)
		if err != nil {
			log.Fatalln(err)
		}

		go gerarURLELocalDeArmazenamento(portarias, siglaInstituicao, canal)

		for linkLocal := range canal {
			baixarArquivos(linkLocal.LocalArmazenamento, linkLocal.URL)
		}
	}
	close(canal)
}
