<div align="right">
<img height="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
</div>

# RH Api GOlang com PostgreSQL e Redis #

É uma API que controla os funcionários de uma epresa e emite seus contracheques

**O usuário será capaz de**

- endpoint GET /funcionarios/
    * Lista todos os funcionários
    * Se não tiver funcionários cadastrados, deve retornar o Status 200 com o body vazio

Exemplo de retorno:
   ```json
  [
  {
    "ïd": 1,
    "nome": "Priscila",
    "sobrenome": "Tavares",
    "documento": "26008495721",
    "setor": "backend",
    "salario_bruto": 7000,
    "data_admissao": "2022-09-20T00:00:00Z",
    "saude": true,
    "dental": true,
    "transporte": true,
    "dependentes": 0,
    "pensao_alimenticia": 0
  },
  {
    "ïd": 2,
    "nome": "Alberto",
    "sobrenome": "Ouverney Paz",
    "documento": "90003218759",
    "setor": "backend",
    "salario_bruto": 7000,
    "data_admissao": "2022-10-01T00:00:00Z",
    "saude": true,
    "dental": true,
    "transporte": true,
    "dependentes": 0,
    "pensao_alimenticia": 0
  }
]
```

- endpoint GET /funcionarios/:id
    * Retora o funcionário com o ID passado na URL com status 200


    Exemplo de retorno:

 ```json
   {
	   "ïd": 1,
	   "nome": "Alberto",
	   "sobrenome": "Ouverney Paz",
	   "documento": "90003218759",
	   "setor": "backend",
	   "salario_bruto": 7000,
	   "data_admissao": "2022-09-20T00:00:00Z",
	   "saude": true,
	   "dental": true,
	   "transporte": true,
       "dependentes": 0,
	   "pensao_alimenticia": 0
  }
  ```
  
   * Se o funcionário não existir, retorna o Status 404 com a mensagem

- ```json
  {
	     "message": "funcionário não encontrado"
  }
  ```

    - endpoint POST /funcionarios/
    * Adiciona um funcionário ao banco de dados

    Exemplo de body:

    ```json
     { 
        "nome": "Priscila",
        "sobrenome": "Tavares",
        "setor": "backend",
        "documento": "260.084.957-21",
        "salario_bruto": 7000.00,
	    "data_admissao": "20/09/2022",
	    "saude": true,
	    "dental": true,
	    "transporte": true,
	    "depedente":0,
        "pensao_alimenticia": 0
     }
    ````
  * O endpoint retorna o Status 201 com o ID do funcionário criado conforme exemplo abaixo:
  
- ```json
    {
        "id": 1
    }
    ``` 


- Validações do endpoint

* O campo nome é obrigatório
* O campo sobrenome é obrigatório
* O campo setor é obrigatório
* O campo documento é obrigatório e dever ser um cpf válido
* O campo salario_bruto é obrigatório
* O campo data_admissao é obrigatório e deve ser no formato DD/MM/YYYY

- Caso a validação do campo nome falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	     "message": "nome não pode ser vazio"
   }
  ```
- Caso a validação do campo sobrenome falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	    "message": "sobrenome não pode ser vazio"
   }
  ```
- Caso a validação do campo setor falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	    "message": "setor não pode ser vazio"
   }
  ```
- Caso a validação do campo documento falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	    "message": "CPF inválido"
   }
  ```
- Caso a validação do campo salario_bruto falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	  "message": "salário bruto deve ser maior que zero"
   }
  ```  
  - Caso a validação do campo data_admissao falhe o endpoint retorna o Status 400 com a mensagem:
  ```json
   {
	  "message": "salário bruto deve ser maior que zero"
   }
  ```  
    - endpoint GET /contracheque/:id
        * Retorna o contracheque de acordo com o ID do funcionario passado na URL
        * O endpoint retorna o Status 200 com os dados do contracheque no formato JSON
    
    Exemplo de retorno: 
  ```json
    {
        "nome": "Priscila",
        "sobrenome": "Tavares",
        "documento": "26008495721",
        "setor": "backend",
        "salario_bruto": 7000,
        "data_admissao": "2022-09-20T00:00:00Z",
        "saude": true,
        "dental": true,
        "transporte": true,
        "dependentes": 0,
        "pensao_alimenticia": 0,
        "desconto_saude": 0,
        "desconto_dental": 0,
        "desconto_transporte": 0,
        "desconto_dependentes": 0,
        "desconto_pensao_alimenticia": 0,
        "desconto_inss": 0,
        "desconto_irrf": 0,
        "salario_liquido": 0
    }
    ```
  * O endpoint retorna o Status 404 com a mensagem "funcionário não encontrado" caso o funcionário não exista
   
    Exemplo de retorno:
  
    ```json
        {
            "message": "funcionário não encontrado"
        }
    ```

## O Desenvolvimento

Foi utilizado o método TDD, para o desenvolvimento. Foi escolhido chi como router por
este possuir o melhor benchmark de performance, e o banco de dados escolhido foi o Postgres e Redis para o cache. 
Todo ele dockerizado.

### Ferramentas usadas

- [Golang](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)
- [Chi](https://github.com/go-chi/chi)
- [Docker](https://www.docker.com/)

## Uso

### Docker

Este projeto 

- Clone o repositório
    ```bash
    git@github.com:beto-ouverney/rh-api-go.git
  ```
- Entre na pasta do projeto 

  ```bash
  cd rh-api-go
  ```

- Utilize o comando abaixo para subir os containers da API, REDIS e POSTGRES

   ```bash
   docker-compose -f docker-compose.dev.yml up -d --build
   ```

- Utilize o comando abaixo para retirar os containers da API, REDIS e POSTGRES

  ```bash
  docker-compose -f docker-compose.dev.yml down --remove-orphans
  ```

- exitem dois docker-compose para subir a aplicação, um para testes e outro para produção
para que não haja conflito entre os dados de teste e os dados de uso continuo. Altere o .env.test caso
seja para testes e .env para produção

- Para rodar o projeto

  ```bash
  docker exec -it rh_api bash
  ```
Após isso rode o comando abaixo para rodar a API no ambiente de produção:

   ```bash
   go run main.go
   ```
## Teste

* O projeto possui testes unitários e de integração, os testes devem rodar preferencialmente no 
seus respectivos containers para que não haja conflito entre os dados de teste e os dados de uso continuo.
Para rodar subir os containers de teste utilize o comando abaixo:

- Não esqueça de alterar o .env.test para .env caso seja para testes e .env para produção

   ```bash
   docker-compose -f docker-compose.test.yml up -d --build
   ```

- Para rodar os testes
     ```bash
     go test -v ./... 
     ```

- Para rodar os testes com cores para diferenciar os erros mais facilmente
  ```bash
  go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
  ```

- Caso queira rodar a API no ambiente de teste execute o comando abaixo:

   ```bash
   docker exec -it rh_api_test bash
   ```

Após isso rode o comando abaixo para rodar a API no ambiente de teste:

   ```bash
   go run main.go
   ```

- Utilize o comando abaixo para retirar os containers de testes da API, REDIS e POSTGRES

  ```bash
  docker-compose -f docker-compose.test.yml down --remove-orphans
  ```
## Author

- LinkedIn - [Alberto Ouverney Paz](https://www.linkedin.com/in/beto-ouverney-paz/)
