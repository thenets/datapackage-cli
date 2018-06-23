# Brasil.io CLI (Command Line Interface)

ESTE PROJETO AINDA ESTÁ EM DESENVOLVIMENTO!

Escreve o objetivo do projeto aqui! TODO

## Como instalar

### Requisitos

É necessário ter o Docker instalado:

```bash
# Se você estiver usando Linux, siga os comandos abaixo

# Instala do Docker automaticamente
curl -sSL https://get.docker.io | sudo sh

# Adiciona o usuário atual ao grupo do Docker
usermod -aG docker $(getent passwd "$(id -u)" | cut -d: -f1)

# Reinicia o sistema (necessário para alguns Linux)
sudo shutdown -rf 0
```

### Instalação

#### Linux

Como instalar e atualizar:

```
sudo ln -s $HOME/.brasilio/bin/brasilio /usr/bin/brasilio 
mkdir -p $HOME/.brasilio/bin/
curl https://brasilio.thenets.org/builds/linux/brasilio > ~/.brasilio/bin/brasilio && chmod +x ~/.brasilio/bin/brasilio
```

Como atualizar:

```
curl https://brasilio.thenets.org/builds/linux/brasilio > ~/.brasilio/bin/brasilio && chmod +x ~/.brasilio/bin/brasilio
```


#### MacOS X

Como instalar:

```
mkdir -p ~/.brasilio/bin/
echo 'export PATH=$PATH:$HOME/.brasilio/bin/' >> ~/.bash_profile
curl https://brasilio.thenets.org/builds/macosx/brasilio > ~/.brasilio/bin/brasilio
chmod +x ~/.brasilio/bin/brasilio
source ~/.bash_profile
```

Como atualizar:

```
curl https://brasilio.thenets.org/builds/macosx/brasilio > ~/.brasilio/bin/brasilio && chmod +x ~/.brasilio/bin/brasilio
```


## Gerenciar projetos de conjuntos de dados

### Criar projeto

Para criar um novo projeto do Brasil.io:

```bash
brasilio-cli init <nome_do_projeto>
```

### Execute projeto

Execute o projeto e teste se a saída está no padrão do [datapackage](https://frictionlessdata.io/data-packages/).

```bash
brasilio-cli run
```

### Teste (ainda não implementado)

Verifique se o código está no padrão para o Brasil.io.

```bash
brasilio-cli test
```

### Execute e teste (ainda não implementado)

Execute e teste ao mesmo tempo.

```bash
brasilio-cli run-test
```
