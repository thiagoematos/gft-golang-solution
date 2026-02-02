# Teste de Senioridade - Desenvolvedor Golang
**Tempo: 10 minutos**

## Desafio: API RESTful com Validações

Você precisa criar uma **API HTTP simples** para gerenciar usuários, com validações robustas.

### Parte 1: Refatoração de Código

Você recebeu um código funcional mas mal escrito. Refatore-o seguindo boas práticas Go.

**Código Original (Ruim):**

```go
package main

import (
    "errors"
    "fmt"
)

type user struct {
    name string
    age int
    email string
}

func process(u user) error {
    if u.name == "" {
        return errors.New("name is empty")
    }
    if u.age < 0 {
        return errors.New("age is negative")
    }
    if u.age > 150 {
        return errors.New("age is too high")
    }
    if u.email == "" {
        return errors.New("email is empty")
    }
    
    fmt.Println("User " + u.name + " is " + fmt.Sprintf("%d", u.age) + " years old")
    fmt.Println("Email: " + u.email)
    
    return nil
}
```

### Parte 2: Criar Endpoint HTTP

Crie um endpoint **POST /users** usando **net/http nativo** (sem frameworks) que:

1. ✅ Recebe JSON com dados do usuário
2. ✅ Valida os dados usando suas funções de validação
3. ✅ Retorna resposta apropriada (200 ou 400)
4. ✅ Usa handlers idiomáticos

**Exemplo de Request:**
```bash
POST /users
Content-Type: application/json

{
    "name": "John Doe",
    "age": 30,
    "email": "john@example.com"
}
```

**Exemplo de Response (Sucesso):**
```json
{
    "message": "User created successfully",
    "user": {
        "name": "John Doe",
        "age": 30,
        "email": "john@example.com"
    }
}
```

**Exemplo de Response (Erro):**
```json
{
    "error": "validation failed: age must be between 0 and 150"
}
```

### Requisitos

- ✅ Código deve seguir **boas práticas Go**
- ✅ Usar **net/http nativo** (sem Gin, Echo, etc.)
- ✅ **Validações** reutilizáveis (métodos)
- ✅ **JSON** marshaling/unmarshaling correto
- ✅ **HTTP status codes** apropriados (200, 400, 405)
- ✅ Máximo **10 minutos**

## Avaliação

### ⭐ Desenvolvedor Pleno

**Refatoração básica:**
- Renomeia `user` para `User` (exportado)
- Usa `fmt.Printf` ao invés de concatenação
- Código funciona mas mudanças são superficiais

**Endpoint básico:**
- Cria handler que funciona
- Faz parse de JSON
- Retorna alguma resposta
- **Pode ter**: Error handling básico, código um pouco verboso

### ⭐⭐ Desenvolvedor Sênior

**Refatoração completa:**
- ✅ **Método Validate()** separado e testável
- ✅ **String() method** implementando Stringer
- ✅ **Constantes** para valores mágicos (MinAge/MaxAge)
- ✅ **Erros descritivos** com `fmt.Errorf`
- ✅ **Código limpo** e bem organizado
- ✅ **Comentários godoc**

**Endpoint profissional:**
- ✅ **Handler idiomático** com `http.HandlerFunc`
- ✅ **JSON encoding/decoding** correto
- ✅ **HTTP status codes** apropriados (200, 400, 405)
- ✅ **Validação de método** (apenas POST)
- ✅ **Error responses** em JSON
- ✅ **Separação de concerns** (handler vs lógica de negócio)

### Exemplos de Diferenciais

**Pleno:**
```go
// Handler básico funcional
func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    json.NewDecoder(r.Body).Decode(&u)
    
    if err := u.Validate(); err != nil {
        w.WriteHeader(400)
        json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
        return
    }
    
    w.WriteHeader(200)
    json.NewEncoder(w).Encode(u)
}
```

**Sênior:**
```go
// Handler robusto e idiomático
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Valida método HTTP
    if r.Method != http.MethodPost {
        respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
        return
    }
    
    // Parse JSON
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        respondWithError(w, http.StatusBadRequest, "invalid JSON")
        return
    }
    defer r.Body.Close()
    
    // Valida
    if err := user.Validate(); err != nil {
        respondWithError(w, http.StatusBadRequest, err.Error())
        return
    }
    
    // Sucesso
    respondWithJSON(w, http.StatusOK, map[string]interface{}{
        "message": "User created successfully",
        "user":    user,
    })
}

// Helper functions (separação de concerns)
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}
```

## Como Executar

```bash
# Rodar servidor
go run main.go

# Testar endpoint (em outro terminal)
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John","age":30,"email":"john@example.com"}'
```

## Entrega (10 minutos)

1. Arquivo com código refatorado + endpoint HTTP
2. Breve comentário explicando **principais decisões de design**

**Boa sorte! ⏱️**
