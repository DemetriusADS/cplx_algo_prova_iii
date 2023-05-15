# cplx_algo_prova_ii

# Prova II Unidade - Complexidade de Algoritmos

> Aluno: Demetrius Araujo <br>
> Tema: Monitoramento de Calderia <br>
> Resumo: O programa tem como objetivo simular o sensoriamento de algumas Caldeiras, tendo como principais métricas a temperatura em graus celsius e o volume em litros. Por fim, de forma ficticia, o programa realiza um ajuste nas métricas obtidas nas temperaturas, baseado no volume daquele momento.

## Requisitos

- Ter a versão 1.20+ ou superior instalado na máquina
  - https://go.dev/doc/install

## Instalando os pacotes

- Realizar o clone do projeto em questão.
- Instalar os pacotes do projeto. Para isso, acesse a pasta root do projeto, a partir do seu terminal, e execute o comando abaixo:

```sh
go mod download
```

## Executando o projeto

É possível executar o programa de 3 formas diferentes. **Mas atente-se à flag que será utilizada**. São duas possíveis:

- `cenario-1`
- `cenario-2`<br>

#### Comando para executar o projeto:

- Prmeira forma:

```sh
go run main.go <flag>
```

- Segunda, gerando o binário

```sh
  go build ./build/main .
```

```sh
  ./build/main <flag>
```

- A Terceira forma é através da opcao de debug do próprio Vscode.

## Respostas da prova

Todos os comentários necessários para responder à prova, estão nos arquivos:
`monitor/monitor.go`
`machine/machine.go`
