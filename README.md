# Gestão de Livros API

Esta é uma API simples para gestão de livros, desenvolvida em Go utilizando o framework Gin. A API permite listar livros, obter detalhes de um livro específico, adicionar novos livros, e gerir a quantidade de livros disponíveis (check-out e devolução).

## Funcionalidades

- **Listar todos os livros**: Retorna uma lista de todos os livros disponíveis.
- **Obter livro por ID**: Retorna os detalhes de um livro específico com base no seu ID.
- **Adicionar novo livro**: Permite adicionar um novo livro à lista.
- **Check-out de livro**: Reduz a quantidade disponível de um livro (simula o empréstimo de um livro).
- **Devolução de livro**: Aumenta a quantidade disponível de um livro (simula a devolução de um livro).

## Como Utilizar

### Pré-requisitos

- Go instalado na tua máquina.
- Git instalado (opcional, para clonar o repositório).

### Instalação

1. Clona o repositório:

   ```bash
   git clone https://github.com/tuusuario/gestao-livros-api.git
