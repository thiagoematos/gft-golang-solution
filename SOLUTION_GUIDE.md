# Guia de Solução - Para o Avaliador

## Objetivo do Teste

Este teste avalia:
1. **Leitura de código**: Identificar problemas
2. **Boas práticas Go**: Naming, estrutura, idiomas
3. **Design**: Separação de responsabilidades
4. **Qualidade**: Código limpo e manutenível
5. **API HTTP**: net/http nativo, JSON, handlers idiomáticos

---

## Análise da Solução

### ⭐ Nível PLENO - Melhorias Superficiais

Um desenvolvedor **pleno** faz mudanças básicas:

```go
// User ao invés de user (exportado)
type User struct {
    Name  string
    Age   int
    Email string
}

func Process(u User) error {
    if u.Name == "" {
        return errors.New("name is empty")
    }
    // ... mesmas validações
    
    // Usa Printf ao invés de concatenação
    fmt.Printf("User %s is %d years old\n", u.Name, u.Age)
    fmt.Printf("Email: %s\n", u.Email)
    
    return nil
}
```

✅ **Melhorias feitas**:
- Capitalização correta (User, Name, Age, Email)
- Usa `fmt.Printf` ao invés de concatenação de strings
- Talvez adicione alguns comentários

⚠️ **O que falta**:
- Validação ainda está misturada com processamento
- Não extrai constantes (150, 0)
- Erros não são descritivos
- Não usa métodos
- Código ainda é difícil de testar

---

### ⭐⭐ Nível SÊNIOR - Refatoração Completa

Um desenvolvedor **sênior** faz refatoração estrutural:

```go
const (
    MinAge = 0
    MaxAge = 150
)

type User struct {
    Name  string
    Age   int
    Email string
}

// Validate separa lógica de validação
func (u User) Validate() error {
    if u.Name == "" {
        return fmt.Errorf("validation failed: name cannot be empty")
    }
    
    if u.Age < MinAge {
        return fmt.Errorf("validation failed: age cannot be negative (got %d)", u.Age)
    }
    
    if u.Age > MaxAge {
        return fmt.Errorf("validation failed: age %d exceeds maximum (%d)", u.Age, MaxAge)
    }
    
    if u.Email == "" {
        return fmt.Errorf("validation failed: email cannot be empty")
    }
    
    return nil
}

// String implementa Stringer para formatação
func (u User) String() string {
    return fmt.Sprintf("User %s is %d years old (Email: %s)", 
        u.Name, u.Age, u.Email)
}

// ProcessUser agora só foca no processamento
func ProcessUser(u User) error {
    if err := u.Validate(); err != nil {
        return err
    }
    
    fmt.Println(u)  // Usa String() automaticamente
    return nil
}
```

✅ **Por que é sênior**:
- ✅ **Método Validate()**: Separação de responsabilidades
- ✅ **Constantes**: MinAge/MaxAge extraídos (Single Source of Truth)
- ✅ **Método String()**: Implementa interface Stringer
- ✅ **Erros descritivos**: Usa `fmt.Errorf` com contexto
- ✅ **Testabilidade**: Validate() testável isoladamente
- ✅ **Documentação**: Comentários godoc
- ✅ **Extensibilidade**: Fácil adicionar validações

**Diferenciais bônus** (sênior forte):
- Validação de email com regex
- Custom error types
- Validação de múltiplos campos com slice de erros
- Builder pattern para User

---

## Checklist de Avaliação (100 pontos)

### Parte 1: Refatoração (50 pts)

**Melhorias Básicas (15 pts)**
- [ ] Capitaliza User e campos **(5 pts)**
- [ ] Usa `fmt.Printf` ao invés de concatenação **(5 pts)**
- [ ] Código funciona corretamente **(5 pts)**

**Estrutura e Design (20 pts)**
- [ ] Método `Validate()` separado **(10 pts)**
- [ ] Constantes MinAge/MaxAge **(5 pts)**
- [ ] Erros descritivos com `fmt.Errorf` **(5 pts)**

**Qualidade (15 pts)**
- [ ] Método `String()` implementado **(10 pts)**
- [ ] Comentários godoc **(5 pts)**

### Parte 2: Endpoint HTTP (50 pts)

**Funcionalidade Básica (20 pts)**
- [ ] Handler funciona e responde **(10 pts)**
- [ ] Parse de JSON correto **(5 pts)**
- [ ] Retorna JSON na resposta **(5 pts)**

**HTTP Idiomático (20 pts)**
- [ ] Valida método HTTP (apenas POST) **(5 pts)**
- [ ] Status codes corretos (200, 400, 405) **(10 pts)**
- [ ] Header Content-Type application/json **(5 pts)**

**Qualidade (10 pts)**
- [ ] Helper functions (respondWithJSON, etc) **(5 pts)**
- [ ] Error handling robusto **(5 pts)**

### Bônus Sênior
- [ ] Validação de email com regex **(+5 pts)**
- [ ] defer r.Body.Close() **(+3 pts)**
- [ ] Logging de operações **(+2 pts)**
- [ ] Testes unitários **(+5 pts)**

---

## Diferenciação Rápida

| Aspecto | Pleno | Sênior |
|---------|-------|--------|
| **Refatoração** | Superficial | ✅ Estrutural |
| **Validação** | ❌ Misturada | ✅ Método Validate() |
| **Constantes** | ❌ Hardcoded | ✅ Min/MaxAge |
| **String()** | ❌ Não tem | ✅ Implementa |
| **Handler** | ⚠️ Básico | ✅ Idiomático |
| **HTTP Status** | ⚠️ Só 200 | ✅ 200/400/405 |
| **JSON** | ⚠️ Funciona | ✅ + helpers |
| **Error handling** | ⚠️ Básico | ✅ Robusto |
| **Tempo** | ~10 min | ~8-9 min |

---

## Red Flags 🚩

### Críticos
- ❌ Código não funciona após refatoração
- ❌ Piora a legibilidade
- ❌ Remove funcionalidade

### Preocupantes
- ⚠️ Não capitaliza structs/campos
- ⚠️ Ainda concatena strings
- ⚠️ Não extrai validação
- ⚠️ Over-engineering (adiciona patterns desnecessários)

---

## O que Observar

### Pleno típico faz:
1. Renomeia `user` → `User`
2. Usa `fmt.Printf`
3. Talvez adicione comentários
4. **Para por aqui**

### Sênior típico faz:
1. Tudo que o pleno faz
2. **Cria método `Validate()`**
3. **Extrai constantes**
4. **Implementa `String()`**
5. **Melhora mensagens de erro**
6. Pensa em testabilidade

---

## Perguntas de Follow-up

### 1. "Por que você criou o método Validate()?"
- **Pleno**: "Para organizar melhor" (vago)
- **Sênior**: "Separação de responsabilidades, testabilidade, reutilização"

### 2. "Por que extrair MinAge/MaxAge em constantes?"
- **Pleno**: "É boa prática"
- **Sênior**: "Single Source of Truth, fácil mudar, auto-documentação"

### 3. "O que é a interface Stringer?"
- **Pleno**: Não sabe ou explica de forma vaga
- **Sênior**: "Interface com String() string, usada por fmt.Print automaticamente"

---

## Tempo Esperado

**10 minutos** é suficiente:

✅ **< 6 min** + todas as melhorias → **Sênior forte**  
✅ **6-8 min** + Validate() + constantes → **Sênior**  
✅ **8-10 min** + melhorias básicas → **Pleno sólido**  
⚠️ **~10 min** só naming → **Pleno júnior**  
❌ **Não completa** ou quebra código → **Júnior**

---

## Conclusão

Este teste é excelente porque:

✅ **Não requer conhecimento específico** (sem concorrência complexa)  
✅ **Avalia qualidade de código** diretamente  
✅ **Diferencia claramente** através de design  
✅ **Prático** - refatoração é trabalho real  
✅ **Rápido** - código pequeno, mudanças focadas  

**Ponto-chave**: Validate() + String() + constantes = Sênior

**Boa avaliação!** 🎯
