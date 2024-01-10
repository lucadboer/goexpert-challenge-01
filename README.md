# Dollar exchange rate challenge

## Descrição

Este prójeto consiste em um aplicativo de servidor HTTP feita em Go que busca e armazena a cotação atual do dólar em relação ao real brasileiro (USD-BRL) usando uma API externa. Além disso, ele oferece uma funcionalidade para salvar o valor de "bid" (o valor pelo qual o mercado está disposto a comprar) em um arquivo local.

## Dependências

- **Go (Golang)**: A linguagem de programação em que o código está escrito.
- **Bibliotecas Go**:
  - `database/sql`: Para interação com banco de dados.
  - `github.com/mattn/go-sqlite3`: Driver SQLite para Go.
  - `net/http`: Para criar o servidor HTTP e fazer requisições web.
  - `encoding/json`: Para codificação e decodificação de dados JSON.

## Funcionalidades

### 1. Servidor HTTP
- Inicia um servidor HTTP na porta 8080.
- Responde a requisições na rota `/cotacao` com a cotação atual do dólar em relação ao real brasileiro em formato JSON.

### 2. Busca e Armazenamento da Cotação
- A função `getCotation` faz uma requisição à API `https://economia.awesomeapi.com.br/json/last/USD-BRL` para obter a cotação.
- A cotação é então armazenada em um banco de dados SQLite local usando a função `saveCotationToDB`.

### 3. Salvamento em Arquivo Local
- A função `saveCotation` salva o valor de "bid" em um arquivo chamado `cotacao.txt`.

## Como Rodar

1. **Pré-requisitos**:
    - Certifique-se de ter o Go instalado em sua máquina. Caso contrário, você pode baixá-lo [aqui](https://golang.org/dl/).

2. **Clonar o Repositório**:
    ```bash
    git clone https://github.com/lucadboer/goexpert-challenge-01.git || git@github.com:lucadboer/goexpert-challenge-01.git
    cd goexpert-challenge-01
    ```

3. **Executar a Aplicação**:
    ```bash
    go run main.go
    ```

## Notas Importantes

- O código contém funções para lidar com erros, usando `panic` para interromper a execução em caso de erro.
- O banco de dados SQLite é utilizado para armazenar a cotação, garantindo que ela possa ser acessada novamente em futuras execuções do programa.
- O valor de "bid" também é salvo em um arquivo local para fácil acesso e leitura.

## Contribuição

Se você encontrar algum problema ou tiver sugestões para melhorar este código, sinta-se à vontade para abrir uma issue ou enviar uma pull request.
