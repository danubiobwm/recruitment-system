# Recruitment System API

Sistema de recrutamento simples com autenticação, gestão de vagas e candidatos, feito em Go (Golang), Gin, GORM e PostgreSQL. Containerizado com Docker.

---

## 🧱 Tecnologias

- **Go** + **Gin** (framework web)
- **GORM** (ORM para Go)
- **JWT** para autenticação
- **PostgreSQL**
- **Docker / Docker Compose**

---

## 🚀 Como rodar o projeto com Docker

### Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Passo a passo

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/recruitment-system.git
cd recruitment-system

# Suba os containers
docker-compose up --build

# O backend estará disponível em:
📍 http://localhost:8080
```

## 🛠️ Estrutura
```
/models — Definição das entidades (User, Job, Candidate)

/controllers — Lógica dos endpoints da API

/routes — Rotas públicas e protegidas

/database — Conexão com o banco via GORM

main.go — Entry point da aplicação
```

## 🔐 Autenticação
Utiliza JWT. Primeiro é necessário se registrar e logar para obter um token, que deve ser usado nos endpoints protegidos.

```
Fluxo:
Registrar (/register)

Login (/login) → recebe o token

Usar o token no header:
Authorization: Bearer {token}

```

## 📮 Endpoints
### Auth

Método	Rota	Descrição

POST	/register	Cadastrar usuário

POST	/login	Fazer login

GET	/api/me	Ver dados do usuário

### Jobs (Vagas)

Método	Rota	Descrição

GET	/api/jobs	Listar todas as vagas

POST	/api/jobs	Criar nova vaga

GET	/api/jobs/:id	Obter vaga por ID

PUT	/api/jobs/:id	Atualizar vaga

DELETE	/api/jobs/:id	Deletar vaga (soft delete)

### Exemplo de criação de vaga
```
{
  "title": "Desenvolvedor Backend Go",
  "description": "Trabalhar com microsserviços em Golang",
  "status": "open"
}
```
## Candidates (Candidatos)

| Método | Rota                 | Descrição                  |
| ------ | -------------------- | -------------------------- |
| GET    | /api/candidates      | Listar todos os candidatos |
| POST   | /api/candidates      | Criar novo candidato       |
| GET    | /api/candidates/\:id | Obter candidato por ID     |
| PUT    | /api/candidates/\:id | Atualizar candidato        |
| DELETE | /api/candidates/\:id | Deletar candidato          |
```
{
  "name": "João Silva",
  "email": "joao@example.com",
  "phone": "11999999999",
  "resume": "https://linkedin.com/in/joao",
  "status": "applied",
  "jobId": 1
}
```

## 🧪 Testes manuais

Você pode testar com:

Insomnia

Postman

Arquivo .http no VSCode com plugin REST Client

## 👤 Autor
J. Danubio

GitHub: [danubiobwm](https://github.com/danubiobwm/recruitment-system)

## 📄 Licença
Este projeto está licenciado sob a MIT License.