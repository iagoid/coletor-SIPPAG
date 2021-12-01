package entidades

type Portarias struct {
	Data struct {
		Results []struct {
			Portaria struct {
				Numero      uint   `json:"numero"`
				Interessado string `json:"interessado"`
				DataEfetiva struct {
					Date struct {
						Month string `json:"month"`
						Year  uint   `json:"year"`
						Day   string `json:"day"`
					} `json:"date"`
				} `json:"dataEfetiva"`
			} `json:"portaria"`
			Assinatura struct {
				Hash string `json:"hash"`
			} `json:"assinatura"`
		} `json:"results"`
		Page  uint `json:"page"`
		Size  uint `json:"size"`
		Total uint `json:"total"`
	} `json:"data"`
}
