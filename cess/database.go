package cess

type Database struct {
    Engine   string `json:"engine"`
    Host     string `json:"host"`
    Name     string `json:"name"`
    Username string `json:"username"`
    Password string `json:"password"`
    Port     string `json:"port"`
}