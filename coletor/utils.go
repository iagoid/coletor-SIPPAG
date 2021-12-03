package coletor

import (
	"errors"
	"fmt"
	"os"
)

func verificaSeDiretorioExiste(diretorio string) {
	if _, err := os.Stat(diretorio); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(diretorio, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	err := os.MkdirAll(diretorio, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}
