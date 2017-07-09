package cess

type Api struct {
    Name    string `json:"name"`
    Url     string `json:"url"`
    Headers map[string]string `json:"headers"`
    Data    map[string]string `json:"data"`
    Action  string `json:"test_action"`
    Method  string `json:"test_method"`
}