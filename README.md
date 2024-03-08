
# 🚀 Inicializando

1. Primeiro builde a imagem do Docker.

```sh
docker build -t cognito-hackagon .
```
2. Antes de subir a imagem na sua máquina, é necessário um arquivo `config.yaml`, com todas as configurações sensíveis solicitadas no arquivo `config.go`

- Informações do Servidor, Banco de Dados & Redis.  

3. Agora builde a imagem.

```sh
docker run -p 9738:9738 -v /path/config/config.yaml:/app/config/config.yaml cognito-hackagon:latest
```

## ⚙️ API Endpoints

Por padrão a API é executada na porta `9738` localmente.

A API usa o formato JSON no conteúdo do `body` seguindo os princípios REST.

### POST /user

Cria uma conta na API

**Parâmetros**

- Não são necessários parametros.

**Resposta**

```json
{"user":{"username":"WDLyjHmb","password":"CoGCMltond"}}
```
