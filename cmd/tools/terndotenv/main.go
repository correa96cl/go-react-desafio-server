package main

import (
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar variáveis de ambiente:", err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations", // Replace with your migrations directory
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	if err := cmd.Run(); err != nil {
		log.Printf("Erro ao executar o comando: %s\n", err)
		log.Printf("Comando executado: %s\n", cmd.String())
	} else {
		log.Println("Variáveis de ambiente carregadas com sucesso!")

	}

}
