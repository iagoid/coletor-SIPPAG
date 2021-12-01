package main

import (
	"baixador-SIPPAG/entidades"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://sippag-web.ifrs.edu.br/api/v1/portaria?ano=2021&page=0&size=1")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var portarias entidades.Portaria
	err = json.Unmarshal(body, &portarias)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(portarias)

}
