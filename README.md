# create-and-send-tx-go

Este programa é um exemplo de criação de transação BSV com golang.

- [create-and-send-tx-go](#create-and-send-tx-go)
  - [Configuração .env](#configuração-env)
  - [Instalação de dependências](#instalação-de-dependências)
  - [Utilização](#utilização)

## Configuração .env
Criar/Editar o arquivo .env:
```bash
RPC_HOST="hostname do bitcoind:porta"
RPC_USERNAME="username para acesso ao bitcoind"
RPC_PASSWORD="password para acesso ao bitcoind"
TX_ID="txid da transacao a ser pesquisada com gettx"
```

## Instalação de dependências
```bash
go get ./...
```

## Utilização

Gerar executável após editar o .env e instalar dependências:

```bash
$ go build
```

Comandos disponíveis:

```bash
$ ./create-and-send-tx-go create # para criar tx
$ ./create-and-send-tx-go get # para pesquisar tx informada no .env
```

