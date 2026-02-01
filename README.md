# Go Pizza API

API REST para gerenciamento de pedidos de pizza, desenvolvida em Go como projeto de estudo e portfólio.

## Tech Stack

- **Go 1.21+** com standard library (`net/http`, `log/slog`)
- **PostgreSQL** como banco de dados
- **Docker** para containerização
- **go-playground/validator** para validação de input

## Arquitetura

O projeto segue **Clean Architecture** com separação em camadas:

```
cmd/api/          → Entry point
internal/
  ├── entity/     → Domain models
  ├── repository/ → Data access (interface + implementations)
  ├── service/    → Business logic
  └── handler/    → HTTP handlers
```

### Princípios Aplicados

| Conceito | Aplicação |
|----------|-----------|
| **Dependency Injection** | Services e handlers recebem dependências via construtor |
| **Interface Segregation** | Interfaces definidas onde são consumidas (Go idiomático) |
| **Type Safety** | Custom types para `OrderStatus` e `PizzaSize` com constantes |
| **Structured Logging** | `log/slog` com key-value pairs para observabilidade |
| **Input Validation** | Validação declarativa com struct tags |

## Como Executar

### Docker (Recomendado)

```bash
docker-compose up
```

### Local

```bash
cp .env.example .env
# Configure DATABASE_URL
go run cmd/api/main.go
```

## API

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| `POST` | `/orders` | Criar pedido |
| `GET` | `/orders/status?id={id}` | Consultar status |
| `GET` | `/orders/all` | Listar pedidos |

### Exemplo

```bash
curl -X POST http://localhost:3333/orders \
  -H "Content-Type: application/json" \
  -d '{"flavor_id": "margherita", "size": "G", "client_id": "cliente-1"}'
```

**Tamanhos:** `P`, `M`, `G`, `F`

**Status:** `PENDING` → `COOKING` → `READY` → `DELIVERED`

## Variáveis de Ambiente

| Variável | Descrição | Default |
|----------|-----------|---------|
| `PORT` | Porta do servidor | `3333` |
| `DATABASE_URL` | Connection string PostgreSQL | - |

## Roadmap

- [ ] Testes unitários e de integração
- [ ] Graceful shutdown
- [ ] Middleware de logging/metrics
- [ ] Autenticação JWT

## Licença

MIT
