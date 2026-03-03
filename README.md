# Comparei - Serviço de Logs 📝

O **Comparei - Serviço de Logs** é um microsserviço dedicado à centralização, processamento e armazenamento de registos (logs) e eventos gerados por todo o ecossistema "Comparei". 

Ao utilizar uma arquitetura orientada a eventos, este serviço garante que as atividades de outros microsserviços são registadas de forma assíncrona, sem impactar o desempenho da plataforma principal.

## 🛠️ Tecnologias Utilizadas

Este serviço foi desenvolvido com foco em desempenho e fiabilidade, utilizando:

* **Linguagem:** [Go 1.23](https://golang.org/)
* **Banco de Dados Relacional:** **MySQL 8** (armazenamento estruturado de logs, níveis de log, utilizadores e eventos).
* **Mensageria:** **Redis** (utilizado como *Message Broker* para capturar eventos de log em tempo real).
* **Infraestrutura:** Docker e Docker Compose.
* **Roteamento HTTP:** [Gorilla Mux](https://github.com/gorilla/mux)
* **Segurança e Autenticação:** Autenticação via JWT (JSON Web Tokens) e chaves de API.

## ⚙️ Arquitetura do Sistema

A aplicação executa dois processos fundamentais em simultâneo:
1. **Subscriber de Mensageria (Redis):** Uma *goroutine* que escuta continuamente o Redis à procura de novos eventos de log emitidos por outros serviços na rede.
2. **Servidor HTTP REST:** Disponibiliza *endpoints* para consulta de logs, gestão de eventos e níveis de log, permitindo a integração com *dashboards* de monitorização.

## 🚀 Como Executar o Projeto Localmente

### Pré-requisitos
* [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

### Passo a Passo

1. **Clonar o repositório:**
```bash
   git clone https://github.com/ViniciusPereira03/comparei-servico-logs
   cd comparei-servico-logs

```

2. **Configuração de Variáveis de Ambiente:**
Copia o ficheiro de exemplo `.env.example` para criar o teu ficheiro `.env`:
```bash
cp .env.example .env

```


Preenche o ficheiro `.env` com os valores adequados para o teu ambiente local. Exemplo:
```env
# Configurações da Base de Dados MySQL
MYSQL_HOST=db:3306
MYSQL_USER=root
MYSQL_PASSWORD=root
MYSQL_DB=logsdb

# Configurações do Servidor
PORT=8084

# Mensageria Redis (deve estar acessível na rede comparei_net)
REDIS_MESSAGING_HOST=redis
REDIS_MESSAGING_PORT=6379

# Segurança
API_KEY=tua_api_key_secreta
JWT_SECRET=teu_jwt_secret
VALIDATION_HASH=teu_hash_de_validacao

```


3. **Executar a Aplicação com `run.sh`:**
Dá permissão de execução ao script (caso ainda não tenha dado) e executa-o. Este ficheiro já está configurado para inicializar a aplicação corretamente com os devidos parâmetros:
```bash
chmod +x run.sh
./run.sh

```

> **Nota:** O script aguardará automaticamente (`wait-for-it.sh`) os bancos de dados estarem prontos antes de iniciar o servidor Go.

4. **Verificar os Logs:**
Após a execução do script, deves ver as seguintes mensagens no terminal confirmando que os serviços estão a rodar:
* `📡 Inicializando subscriber...`
* `🚀 Servidor iniciado na porta 8084`


## 📂 Estrutura de Diretórios (Resumo)

* `/cmd` ou Raiz: Ponto de entrada da aplicação (`main.go`).
* `/config`: Lógica de carregamento de variáveis de ambiente (`config.go`).
* `/internal`: Camadas de domínio e infraestrutura da aplicação.
    * `/app`: Serviços de domínio (`events_service.go`, `levels_service.go`, `logs_service.go`, `user_service.go`).
    * `/domain`: Definição de interfaces e entidades de negócio.
    * `/infrastructure`:
        * `/http`: *Handlers*, rotas REST e *middlewares* (como validação de `API_KEY`).
        * `/messaging`: Lógica de publicação e subscrição no Redis.
        * `/repository`: Implementação de persistência de dados no MySQL.
* `/migrations`: Scripts para criar as tabelas iniciais na base de dados (`init.sql`).
