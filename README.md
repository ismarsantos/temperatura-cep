README.md:

# Sistema de temperatura por CEP

Este projeto é uma aplicação Go que recebe um CEP brasileiro, recupera a cidade correspondente e retorna o clima atual em Celsius, Fahrenheit e Kelvin. A aplicação utiliza a [API ViaCEP](https://viacep.com.br/) para obter informações de localização e a [WeatherAPI](https://www.weatherapi.com/) para buscar os dados do clima.


### Google Cloud Run Demo

```sh
curl "https://temperatura-cep-675408290066.us-central1.run.app/weather?cep=01001000"
```

## Estrutura do Projeto

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go
├── tests.go
└── README.md
```

### Arquivos

- **Dockerfile**: Contém as instruções para construir uma imagem Docker para a aplicação Go.
- **docker-compose.yml**: Define os serviços para executar a aplicação de clima, incluindo o servidor mock e o executor de testes.
- **main.go**: Código principal da aplicação, incluindo o servidor HTTP e as funções para lidar com requisições de clima.
- **tests.go**: Testes automatizados para validar o formato do CEP e as conversões de temperatura.
- **README.md**: Este arquivo, que contém informações sobre o projeto e instruções de configuração e uso.

## Pré-requisitos

- Docker
- Docker Compose

## Começando

Siga estas etapas para configurar e executar a aplicação em um novo ambiente.

### 1. Clonar o Repositório

```sh
git clone <repository_url>
cd temperatura-cep
```

### 2. Configurar Variáveis de Ambiente

Você precisa declarar a variável de ambiente `WEATHERAPI_KEY` com sua chave do WeatherAPI. Você pode exportá-la no terminal ou adicioná-la em um arquivo `.env`.

#### Exemplo de Arquivo `.env`

```
WEATHERAPI_KEY=sua_chave_weatherapi_aqui
```

### 3. Construir e Executar com Docker Compose

Use o Docker Compose para construir e iniciar todos os serviços.

```sh
docker-compose up --build
```

Este comando iniciará os seguintes serviços:

- **weather-service**: A aplicação principal que lida com as requisições para obter dados de clima.
- **test**: Um serviço para executar testes automatizados.

O serviço de clima estará disponível em [http://localhost:8080/weather](http://localhost:8080/weather).

### 4. Testando a Aplicação

Você pode executar os testes usando o seguinte comando:

```sh
docker-compose run test
```

Alternativamente, você pode executar os testes localmente se o Go estiver instalado:

```sh
go test -v
```

### 5. Fazendo uma Requisição

Para testar a aplicação, você pode fazer uma requisição GET para o endpoint `/weather` com um CEP válido. Por exemplo:

```sh
curl "http://localhost:8080/weather?cep=01001000"
```

Se bem-sucedido, você receberá uma resposta similar à seguinte:

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

## Notas

- Certifique-se de que o `WEATHERAPI_KEY` é válido e que sua máquina tem acesso à internet para consultar a WeatherAPI.
