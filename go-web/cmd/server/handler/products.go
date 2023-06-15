package handler

type Product struct {
	ID          int     `json:"id"`
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     int     `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"dataCriacao"`
}

type ProductRequest struct {
	Nome        string  `json:"nome"`
	Cor         string  `json:"cor"`
	Preco       float64 `json:"preco"`
	Estoque     int     `json:"estoque"`
	Codigo      string  `json:"codigo"`
	Publicacao  bool    `json:"publicacao"`
	DataCriacao string  `json:"dataCriacao"`
}
