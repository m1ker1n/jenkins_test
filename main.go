package main

import (
	"fmt"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Secret string `env:"SECRET"`
}

var (
	cfg = Config{}
)

func echoQueryHandle(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	data := fmt.Sprintf("%+v", queryParams)
	w.Write([]byte(data))
}

func secretHandle(w http.ResponseWriter, _ *http.Request) {
	data := fmt.Sprintf("%+v", cfg)
	w.Write([]byte(data))
}

func main() {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoQueryHandle)
	mux.HandleFunc("/secret", secretHandle)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
