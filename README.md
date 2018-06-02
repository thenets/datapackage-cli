# Brasil.io CLI (Command Line Interface)

ESTE PROJETO AINDA ESTÁ EM DESENVOLVIMENTO!

Escreve o objetivo do projeto aqui! TODO

## Como instalar

### Requisitos

Linux x64 com suporte para Docker.

É necessário ter o Docker instalado:

```bash
# Instala do Docker automaticamente
curl -sSL https://get.docker.io | sudo sh

# Adiciona o usuário atual ao grupo do Docker
usermod -aG docker $(getent passwd "$(id -u)" | cut -d: -f1)

# Reinicia o sistema (necessário para alguns Linux)
sudo shutdown -rf 0
```

### Instalação

Por enquanto não existe uma maneira simples de instalar. É necessário clonar o repositório e construir o projeto com o Go:

```bash
go get github.com/thenets/brasilio-cli
cd $GOPATH/src/github.com/thenets/brasilio-cli/
go install
```

## Gerenciar projetos de conjuntos de dados

### Criar projeto

Para criar um novo projeto do Brasil.io:

```bash
brasilio-cli new
```

### Execute projeto

Execute o projeto e teste se a saída está no padrão do [datapackage](https://frictionlessdata.io/data-packages/).

```bash
brasilio-cli run
```

### Teste

Verifique se o código está no padrão para o Brasil.io.

```bash
brasilio-cli test
```

### Execute e teste

Execute e teste ao mesmo tempo.

```bash
brasilio-cli run-test
```
