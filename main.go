package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

// SOLUÇÃO SÊNIOR - Código refatorado com API HTTP

const (
	// Constantes para limites de validação
	MinAge = 0
	MaxAge = 150
)

var (
	// Regex simples para validação de email
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

// User representa um usuário do sistema
type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// Validate valida os campos do usuário
func (u User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("validation failed: name cannot be empty")
	}

	if u.Age < MinAge {
		return fmt.Errorf("validation failed: age cannot be negative (got %d)", u.Age)
	}

	if u.Age > MaxAge {
		return fmt.Errorf("validation failed: age %d exceeds maximum allowed (%d)", u.Age, MaxAge)
	}

	if u.Email == "" {
		return fmt.Errorf("validation failed: email cannot be empty")
	}

	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("validation failed: invalid email format '%s'", u.Email)
	}

	return nil
}

// String implementa a interface Stringer para formatação legível
func (u User) String() string {
	return fmt.Sprintf("User %s is %d years old (Email: %s)", u.Name, u.Age, u.Email)
}

// CreateUserHandler processa requisições POST para criar usuário
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Valida método HTTP
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Parse JSON do body
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid JSON format")
		return
	}
	defer r.Body.Close()

	// Valida dados do usuário
	if err := user.Validate(); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Sucesso - retorna usuário criado
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	})

	// Log para acompanhamento
	log.Printf("User created: %s", user.String())
}

// respondWithJSON envia resposta JSON com status code
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

// respondWithError envia resposta de erro em JSON
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func main() {
	// Registra handler
	http.HandleFunc("/users", CreateUserHandler)

	// Inicia servidor
	addr := ":8080"
	fmt.Printf("🚀 Server starting on http://localhost%s\n", addr)
	fmt.Println("📝 POST /users - Create a new user")
	fmt.Println("\nExample:")
	fmt.Println(`  curl -X POST http://localhost:8080/users \`)
	fmt.Println(`    -H "Content-Type: application/json" \`)
	fmt.Println(`    -d '{"name":"John","age":30,"email":"john@example.com"}'`)
	fmt.Println()

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Principais características desta solução SÊNIOR:
//
// REFATORAÇÃO:
// ✅ Struct User exportada com json tags
// ✅ Constantes MinAge/MaxAge (configuráveis)
// ✅ Método Validate() separado e testável
// ✅ Método String() implementa Stringer
// ✅ Erros descritivos com contexto
// ✅ Validação de email com regex
//
// API HTTP:
// ✅ Handler idiomático com validação de método
// ✅ JSON encoding/decoding correto
// ✅ Status codes apropriados (200, 400, 405)
// ✅ Helper functions (respondWithJSON, respondWithError)
// ✅ Headers corretos (Content-Type)
// ✅ defer r.Body.Close() para cleanup
// ✅ Logging de operações
// ✅ Error handling robusto
// ✅ Separação de concerns (handlers vs lógica)
//
// EXTENSIBILIDADE:
// ✅ Fácil adicionar novos endpoints
// ✅ Fácil adicionar novas validações
// ✅ Código testável (Validate() isolado)
// ✅ Reutilização (helpers)
