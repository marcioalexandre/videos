package main

type produto struct {
	ID    int    `json:"id"`
	NOME  string `json:"nome"`
	ATIVO bool   `json:"ativo"`
}
