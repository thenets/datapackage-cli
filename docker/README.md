Olá.

Eu criei uma imagem do `Docker` baseada no `Alpine` para construir e executar o projeto.
Tentei tornar a utilização o mais simples possível. Basicamente só é necessário ter o `docker` instalado.

## Apenas executar

Se quiser utilizar o meu build, basta executar a linha abaixo:
```bash
docker run --rm -it -v $(pwd)/output:/app/output thenets/opendata-salarios-magistrados
```

## Construir a imagem

Se quiser construir os pacotes do zero, basta ter o `make` e o `docker` instalados:
```bash
# Constroi o pacote
make build

# Testa se consegue baixar os CSVs do governo
make test

# Apenas executa o script principal e coloca a saída em `output`
make run
```

## Travis build

Além disso, eu também criei o build para o Travis CI. A cada semana ele executará o seu script e testará se tudo ocorreu como o esperado. Ou seja, caso algum servidor do governo pare de funcionar ou o CSV tenha algum problema, o Travis enviará um report sobre o erro gerado.

Abraço ae, mano!