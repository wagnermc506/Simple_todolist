# Simple Todolist

Este projeto, como o nome indica, é uma aplicação simples de uma lista de tarefas.

Mais especificamente, um backend que fornece uma API Rest escrita em golang.

# Instruções de uso

Obs.: O programa está sendo testado apenas em ambiente linux, neste caso, no Ubuntu 20.04.

# Docker

O arquivo Dockerfile contém o script para a criação de uma imagem no docker que faz o build do código fonte.

Para usar este método, o sistema deve ter o Docker instalado.

## Build

Na raiz do projeto, execute o comando:

``` sh
$ docker build -t go-todolist .
```

Caso queira nomear a imagem de outra forma, modifique o nome _go-todolist_ para um nome de sua preferência.
Vale ressaltar que alguns sistemas precisam das permissões de root para executar o docker.

## Run

Para iniciar um container, execute:

```sh
$ docker run -p 8090:8090 go-todolist
```

A porta padrão do sistema é a 8090. Mude o segundo "8090" caso queira que o host receba as requisições por outra porta.

Obs.: O banco de dados está acessando o sistema de arquivos interno do container, logo os dados não estão persistentes. Isso será ajustado futuramente.

# A partir do código fonte

Para este método, deve-se ter a linguagem Go instalada no sistema. 

## Build

Vá para o diretório /src/todolist/ e execute:

```sh
$ go build -o ../../build/todolist
```

## Run

Para iniciar o server, volte para o diretório raiz do projeto e execute:

```sh
$ ./build/todolist
```
Obs.: O arquivo do banco de dados está atualmente sendo salvo na pasta home do usuário.