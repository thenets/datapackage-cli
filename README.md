# Brasil.io CLI Dockerfiles

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
