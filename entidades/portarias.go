package entidade

type Portarias struct {
	Data struct {
		Results []struct {
			Portaria struct {
				ID          int    `json:"id"`
				Interessado string `json:"interessado"`
				DataEfetiva struct {
					Date struct {
						Month string `json:"month"`
						Year  int    `json:"year"`
						Day   string `json:"day"`
					} `json:"date"`
				} `json:"dataEfetiva"`
			} `json:"portaria"`
			Assinatura struct {
				Hash string `json:"hash"`
			} `json:"assinatura"`
		} `json:"results"`
		Page  int `json:"page"`
		Size  int `json:"size"`
		Total int `json:"total"`
	} `json:"data"`
}
