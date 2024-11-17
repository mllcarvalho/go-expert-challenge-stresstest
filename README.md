
# Go Expert Challenge - Stress Test CLI

Este projeto é um CLI (Command-Line Interface) escrito em Go, projetado para realizar testes de stress em endpoints web, permitindo avaliar o desempenho de serviços HTTP.

---

## 🛠 Arquitetura

O CLI utiliza um modelo concorrente para disparar requisições HTTP em alta escala. A arquitetura segue os seguintes princípios:
- **Workers e Pool**: Um pool de workers gerencia as requisições simultâneas, onde cada worker executa uma requisição HTTP.
- **Comunicação via canais**: Os resultados das requisições são armazenados em canais e processados em tempo real.
- **Processamento**: Após a execução, os resultados são agregados e exibidos, fornecendo métricas sobre o teste.

---

## 🚀 Como usar

O CLI pode ser executado **via Docker** ou **localmente**, dependendo do ambiente e das preferências do usuário.

### 🔹 Usando Docker

A imagem Docker está disponível no Docker Hub. Para rodar o teste, execute o seguinte comando, substituindo os valores de `--url`, `--requests` e `--concurrency` conforme a necessidade:

```bash
docker run mllcarvalho/go-expert-challenge-stresstest:latest \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

### 🔹 Rodando localmente

Para executar o CLI diretamente no ambiente local, certifique-se de ter o **Go** instalado. Siga os passos abaixo:
1. Clone o repositório.
2. No diretório do projeto, execute o comando:

```bash
go run cmd/cli/main.go \
    --url https://google.com.br \
    --requests 100 \
    --concurrency 10
```

Substitua os parâmetros `--url`, `--requests` e `--concurrency` pelos valores desejados:
- **`--url`**: O endpoint HTTP a ser testado.
- **`--requests`**: Número total de requisições a serem feitas.
- **`--concurrency`**: Número de requisições simultâneas.

---

## 🧪 Testes

O projeto inclui testes de unidade para garantir a qualidade do código. Para executá-los, use o comando abaixo:

```bash
make test
```

Esse comando executa todos os testes automatizados definidos no projeto.

---

## 📋 Exemplos de uso

1. Teste simples com 100 requisições e 10 workers:

```bash
go run cmd/cli/main.go \
    --url https://example.com \
    --requests 100 \
    --concurrency 10
```

2. Rodando via Docker com parâmetros personalizados:

```bash
docker run mllcarvalho/go-expert-challenge-stresstest:latest \
    --url https://your-api-endpoint.com \
    --requests 200 \
    --concurrency 20
```

---

Com essa ferramenta, você pode realizar testes de stress simples e eficazes em APIs e servidores web para validar o desempenho em condições de carga. 🚀
