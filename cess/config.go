package cess

type Config struct {
    Databases [] Database `json:"db"`
    Services  [] Api `json:"api"`
}
