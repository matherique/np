package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var (
	helpMenu = `Cria pasta para ser usada em novos projetos e ser usado pelo tmux-sessionizer
    
    np NOME
    np -work WORK_NOME 
    np -personal PERSONAL_PROJECT

-work: cria pasta em ~/work
-personal: cria pasta em ~/personal
`

	name, location string
)

func main() {
	if debug := os.Getenv("DEBUG"); len(debug) > 0 {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	Assert(len(os.Args) == 1, helpMenu)
	Assert(len(os.Args) == 2 && strings.HasPrefix(os.Args[1], "-"), "faltando nome da pasta")

	if len(os.Args) == 2 {
		name = os.Args[1]
	}

	for _, p := range os.Args[1:] {
		slog.Debug("line " + p)
		if strings.HasPrefix(p, "-") {
			slog.Debug("has prefix " + p)
			location, _ = strings.CutPrefix(p, "-")
			continue
		}

		if len(name) == 0 {
			slog.Debug("set name ", p)
			name = p
		}
	}

	if len(location) == 0 {
		location = "personal"
	}

	Assert(len(name) == 0, "faltando nome da pasta")

	home := os.Getenv("HOME")

	Assert(len(home) == 0, "faltando variável de ambiente HOME")

	p := filepath.Join(home, location, name)

	if err := os.MkdirAll(p, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao criar pasta %q: %v")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "pasta criada em %q", p)
}

func Assert(b bool, message string) {
	if b {
		fmt.Fprintln(os.Stderr, message)
		os.Exit(1)
	}
}
