# Documentação da API Trilha

## Visão Geral

A API Trilha é uma solução de backend para um sistema de gerenciamento de projetos, projetada para ser robusta, escalável e fácil de manter. A aplicação é construída em Go, seguindo uma arquitetura modular e utilizando ferramentas modernas para garantir a qualidade e a eficiência do código.

## Módulos

A aplicação é dividida nos seguintes módulos:

*   **Account**: Responsável pelo gerenciamento de contas de usuário, incluindo criação, autenticação e autorização.
*   **Shared**: Contém componentes compartilhados por toda a aplicação, como configurações, manipulação de banco de dados e respostas de API.

## Estrutura de Diretórios

A estrutura de diretórios do projeto é organizada da seguinte forma:

*   `cmd/server`: Contém o ponto de entrada da aplicação.
*   `db`: Armazena os arquivos de schema e queries SQL.
*   `docs`: Contém a documentação do projeto.
*   `internal`: Abriga a lógica de negócios da aplicação, dividida por módulos.
*   `wire`: Contém os arquivos de injeção de dependência do Wire.

## Configurações

As configurações da aplicação são gerenciadas através de variáveis de ambiente, carregadas a partir de um arquivo `.env`. As principais configurações incluem:

*   `APP_PORT`: A porta em que a aplicação será executada.
*   `DB_USER`: O nome de usuário do banco de dados.
*   `DB_PASSWORD`: A senha do banco de dados.
*   `DB_NAME`: O nome do banco de dados.

## Dependências

A aplicação utiliza as seguintes dependências:

*   **Gin**: Um framework web para Go, utilizado para criar as rotas da API.
*   **pgx**: Um driver PostgreSQL para Go, utilizado para se comunicar com o banco de dados.
*   **sqlc**: Uma ferramenta para gerar código Go a partir de queries SQL.
*   **Wire**: Uma ferramenta de injeção de dependência para Go.
*   **godotenv**: Uma biblioteca para carregar variáveis de ambiente a partir de um arquivo `.env`.
*   **crypto**: Uma biblioteca para criptografar senhas de usuário.

## Ferramentas Auxiliares

As seguintes ferramentas auxiliares são utilizadas no desenvolvimento da aplicação:

*   **Docker**: Utilizado para criar um ambiente de desenvolvimento containerizado.
*   **Docker Compose**: Utilizado para orquestrar os containers da aplicação.

## Ideia da Aplicação

A API Trilha foi projetada para ser o backend de um sistema de gerenciamento de projetos completo. A ideia é que a aplicação possa ser utilizada para gerenciar não apenas projetos de software, mas também projetos em geral, como projetos de marketing, projetos de construção, etc.

## Metas Futuras

A seguir, algumas metas para o futuro desenvolvimento da aplicação:

*   **Implementar um sistema de autenticação e autorização mais robusto**: Atualmente, a aplicação utiliza um sistema de autenticação simples, baseado em senhas criptografadas. No futuro, pretendemos implementar um sistema mais robusto, utilizando tokens JWT e um sistema de permissões baseado em papéis.
*   **Adicionar suporte para múltiplos projetos**: Atualmente, a aplicação suporta apenas um projeto por vez. No futuro, pretendemos adicionar suporte para múltiplos projetos, permitindo que os usuários criem e gerenciem vários projetos simultaneamente.
*   **Implementar um sistema de tarefas**: Atualmente, a aplicação não possui um sistema de tarefas. No futuro, pretendemos implementar um sistema de tarefas completo, permitindo que os usuários criem, atribuam e gerenciem tarefas dentro de cada projeto.
*   **Adicionar suporte para equipes**: Atualmente, a aplicação não possui suporte para equipes. No futuro, pretendemos adicionar suporte para equipes, permitindo que os usuários convidem outros usuários para colaborar em seus projetos.
*   **Implementar um sistema de notificações**: Atualmente, a aplicação não possui um sistema de notificações. No futuro, pretendemos implementar um sistema de notificações, permitindo que os usuários recebam notificações sobre eventos importantes, como novas tarefas, comentários, etc.
*   **Adicionar suporte para anexos**: Atualmente, a aplicação não suporta o upload de anexos. No futuro, pretendemos adicionar suporte para anexos, permitindo que os usuários anexem arquivos a tarefas e projetos.
*   **Implementar um sistema de relatórios**: Atualmente, a aplicação não possui um sistema de relatórios. No futuro, pretendemos implementar um sistema de relatórios, permitindo que os usuários gerem relatórios sobre o andamento de seus projetos.
*   **Adicionar suporte para integrações**: Atualmente, a aplicação não suporta integrações com outras ferramentas. No futuro, pretendemos adicionar suporte para integrações com outras ferramentas, como o GitHub, o Slack, etc.
