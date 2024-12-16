# NP

> cria uma pasta no local especificado para ser usado pelo `tmux-sessionizer`

# Instalação
```
go install github.com/matherique/np@latest
```

## Uso 

Para criar uma pasta informe o local com um prefixo de `-`. A pasta `personal` é a padrão caso não informe
```
np nome_pasta -personal
np nome_pasta -work
np nome_pasta -ideias
np nome_pasta
```
Caso a pasta não exista, o sistema ira criar.



