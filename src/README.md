> # Para configurar o POSTGRES/DOCKER(subir minha maquina) usar:
> 
> ```bash
> $ docker run -d --name mercadolivre -p 5432:5432 -e POSTGRES_PASSWORD=om315 postgres:13.5 
> ```
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```docker run```:
> 
> **Para Rodar o Docker**
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```-d```:
> 
> **Ambiente de Desenvolvimento**
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```--name mercadolivre```:
> 
> **Nome do meu Banco de Dados**
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```-p 5432:5432```:
> 
> **A Porta que estará Rodando**
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```-e POSTGRES_PASSWORD=om315```:
> 
> **Variável de Ambiente com a Senha**
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### ```postgres:13.5```:
> 
> **A versão o POSTGRES**
> 
> --------------------------------------------------------//--------------------------------------------------------
> --------------------------------------------------------//--------------------------------------------------------
> 
> <mark>[!MPORTANT]
> NÃO ESQUECER DE INICIAR O DOCKER NA MAQUINA</mark>
> --------------------------------------------------------//--------------------------------------------------------
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### Depois de subir o servidor usar: ```docker ps```
> 
> ```docker ps``` -> ***Verificar se realmente subiu***
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### Depois rodar: ```docker exec -it mercadolivre psql -U postgres```
> 
> ***Comando para realmente executar o docker***
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### Para criar o banco de dados: ```create database [NOME DO BANCO];``` -> api_mercadolivre
> 
> #### ***OBS: Tem que ser no terminal do docker, vai ficar:***
> 
> ```bash
> postgres=# create database [NOME DO BANCO]; 
> ```
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### Criar usuário: ```create user [NOME DO USUÁRIO];``` -> user_mercadolivre
> 
> ### Criar  senha: ```alter user [NOME DO USUÁRIO] with encrypted password ['SENHA'];``` -> 1122
> 
> --------------------------------------------------------//--------------------------------------------------------
> 
> ### Dar permissão para o usuário mexer no banco: ```grant all privileges on database [NOME DO BANCO] to [NOME DO USUÁRIO];```
> ### Dar permissão para o usuário mexer nas tabelas: ```grant all privileges on all tables in schema public to [NOME DO USUÁRIO];```
> 
> --------------------------------------------------------//--------------------------------------------------------
>
>### Limpar tudo: ```CRTL+L```
>
> --------------------------------------------------------//--------------------------------------------------------
>
>### Listar bancos de dados: ```\l```
>
> --------------------------------------------------------//--------------------------------------------------------
>
>### Entrar no banco(SELECIONAR O BANCO): ```\c [NOME DO BANCO];```
>
> --------------------------------------------------------//--------------------------------------------------------
>
>### Listar todas as tabelas: ```\dt```
>
> --------------------------------------------------------//--------------------------------------------------------
>
>### Criando a tabela: ```create table [NOME DA TABELA] ([ESPECIFICAÇÕES]);```
>#### No meu caso será: ```create table todos (id serial primary key, name varchar, descripton text, done bool);```
>
> --------------------------------------------------------//--------------------------------------------------------
> <mark>[!IMPORTANT]</mark>
> <mark>NÃO SE ESQUECER DO PONTO E VIRGULA NO FINAL</mark>
