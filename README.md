# Guia de Avaliação - Teste Golang (Pleno vs Sênior)

## 📋 Visão Geral do Desafio

O candidato recebe um código mal escrito para refatorar e deve criar um endpoint HTTP POST /users usando Go nativo (sem frameworks).

**Tempo limite**: 10 minutos  
**Objetivo**: Diferenciar desenvolvedores **Pleno** de **Sênior** através de qualidade de código, design e conhecimento de APIs HTTP.

---

## 🎯 Critérios de Avaliação (100 pontos)

### Parte 1: Refatoração de Código (50 pontos)

| Critério | Pleno | Sênior | Pontos |
|----------|-------|--------|--------|
| **Naming** | Capitaliza `User` | ✅ + campos exportados | 5 |
| **Validação** | Mantém na função | ✅ Método `Validate()` separado | 15 |
| **Constantes** | Valores hardcoded (150, 0) | ✅ `MinAge`, `MaxAge` | 10 |
| **Formatação** | Usa `fmt.Printf` | ✅ Método `String()` (Stringer) | 10 |
| **Erros** | `errors.New()` genérico | ✅ `fmt.Errorf()` descritivo | 5 |
| **Comentários** | Poucos ou nenhum | ✅ Godoc comments | 5 |

### Parte 2: Endpoint HTTP (50 pontos)

| Critério | Pleno | Sênior | Pontos |
|----------|-------|--------|--------|
| **Handler funciona** | ✅ Básico | ✅ Robusto | 10 |
| **JSON parsing** | Não checa erros | ✅ Valida erros de decode | 5 |
| **HTTP Status** | Só 200 ou hardcoded | ✅ 200/400/405 corretos | 15 |
| **Método HTTP** | Não valida | ✅ Valida `POST` apenas | 5 |
| **Headers** | Esquece `Content-Type` | ✅ Define corretamente | 5 |
| **Helper functions** | Código duplicado | ✅ `respondWithJSON()`, etc | 10 |

**Bônus** (+15 pontos extras):
- Validação de email com regex (+5)
- `defer r.Body.Close()` (+3)
- Logging de operações (+2)
- Testes unitários (+5)

---

## 🔍 O Que Observar Durante a Codificação

### 🚩 Red Flags (Descartam o candidato)
- ❌ **Código não funciona** após 10 minutos
- ❌ **Usa frameworks** (Gin, Echo) quando pediu Go nativo
- ❌ **Não cria endpoint HTTP** (só faz refatoração)
- ❌ **Pânico em JSON inválido** (não trata erros)

### ⚠️ Sinais de Pleno
- Refatoração superficial (só renomeia)
- Handler funciona mas sem validação de método
- Não usa helper functions (código duplicado)
- Constantes ainda hardcoded no código
- Não implementa `String()` method

### ✅ Sinais de Sênior
- **Método `Validate()`** separado (testável)
- **Constantes** para valores mágicos
- **Helper functions** (`respondWithJSON`, `respondWithError`)
- **Status codes** apropriados (405 para método errado)
- **`defer r.Body.Close()`** para cleanup
- **Comentários godoc** nas funções públicas
- Pensa em **extensibilidade** (fácil adicionar validações)

---

## 💬 Perguntas de Entrevista

### Sobre Refatoração

#### 1️⃣ "Por que você criou o método `Validate()`?"

**Resposta Pleno:**
- "Para organizar melhor o código"
- "Porque é boa prática"
- Explicação vaga ou genérica

**Resposta Sênior:**
- "**Separação de responsabilidades** - validação é uma preocupação diferente de processamento"
- "**Testabilidade** - posso testar validação isoladamente"
- "**Reutilização** - posso usar em outros lugares (API, CLI, etc)"
- Menciona **Single Responsibility Principle**

---

#### 2️⃣ "Por que extrair `MinAge` e `MaxAge` em constantes?"

**Resposta Pleno:**
- "É boa prática não ter números mágicos"
- Não consegue explicar benefícios concretos

**Resposta Sênior:**
- "**Single Source of Truth** - se mudar o limite, mudo em um lugar só"
- "**Auto-documentação** - `MaxAge` é mais claro que `150`"
- "**Configurabilidade** - fácil transformar em variável de ambiente depois"
- Menciona facilidade de **testes** (diferentes limites em diferentes contextos)

---

#### 3️⃣ "O que é a interface `Stringer` e por que você implementou?"

**Resposta Pleno:**
- "Não sei" ou "É uma interface do Go"
- Explica de forma vaga

**Resposta Sênior:**
- "Interface com método `String() string`"
- "Usada automaticamente por `fmt.Print`, `fmt.Println`, etc"
- "**Centraliza formatação** - um lugar para definir representação em string"
- "Facilita **logging** e **debugging**"
- Pode mencionar outras interfaces comuns: `error`, `io.Reader`, `io.Writer`

---

### Sobre o Endpoint HTTP

#### 4️⃣ "Por que você valida o método HTTP no handler?"

**Resposta Pleno:**
- "O teste pediu"
- Não vê problema em aceitar GET/PUT

**Resposta Sênior:**
- "**Princípio RESTful** - POST é para criação"
- "**Segurança** - evitar operações não intencionadas"
- "**Clareza da API** - comunicar exatamente o que é permitido"
- Retorna **405 Method Not Allowed** (não 400)
- Pode mencionar middleware em produção

---

#### 5️⃣ "Por que criar `respondWithJSON()` e `respondWithError()`?"

**Resposta Pleno:**
- "Para não repetir código"
- Explicação superficial

**Resposta Sênior:**
- "**DRY** (Don't Repeat Yourself) - evita duplicação"
- "**Consistência** - todas as respostas JSON têm mesmos headers"
- "**Manutenibilidade** - mudanças em uma função afetam todos os endpoints"
- "**Separação de concerns** - handlers focam em lógica, helpers em formato"
- Pode mencionar que em produção isso seria um **middleware** ou **package separado**

---

#### 6️⃣ "O que acontece se você não usar `defer r.Body.Close()`?"

**Resposta Pleno:**
- "Não sei" ou "Pode dar leak de memória"
- Resposta vaga

**Resposta Sênior:**
- "**Resource leak** - conexões HTTP ficam abertas"
- "**Goroutine leak** - em alta carga, esgota recursos"
- "**defer garante cleanup** mesmo se houver erro/panic"
- Pode mencionar que `http.Server` reutiliza conexões (connection pooling)
- Em produção, pode causar **file descriptor exhaustion**

---

#### 7️⃣ "Como você testaria este endpoint?"

**Resposta Pleno:**
- "Rodaria com curl manualmente"
- Não pensa em testes automatizados

**Resposta Sênior:**
- "**Testes unitários** do método `Validate()` isoladamente"
- "**Table-driven tests** com casos válidos/inválidos"
- "**httptest.ResponseRecorder** para testar handler sem servidor real"
- "Testes de integração com servidor de teste"
- Pode mencionar `go test -race` para race conditions
- **Mocking** se tivesse dependências externas (banco, etc)

---

### Perguntas de Aprofundamento

#### 8️⃣ "Como você melhoraria este código para produção?"

**Resposta Pleno:**
- "Adicionaria mais validações"
- Sugestões vagas

**Resposta Sênior (vários pontos possíveis):**
- **Logging estruturado** (zap, logrus)
- **Middleware** para logging, CORS, autenticação
- **Context** para timeouts e cancelamento
- **Validação mais robusta** (biblioteca como validator)
- **Persistência** (banco de dados)
- **Metrics** (prometheus)
- **Graceful shutdown**
- **Rate limiting**
- **Testes** (unitários + integração)
- **CI/CD pipeline**

---

#### 9️⃣ "Qual a diferença entre `errors.New()` e `fmt.Errorf()`?"

**Resposta Pleno:**
- "São a mesma coisa" ou não sabe

**Resposta Sênior:**
- "`fmt.Errorf()` permite **formatação** com verbos (%s, %d, %v)"
- "Pode usar `%w` para **error wrapping** (Go 1.13+)"
- "`errors.New()` é mais simples para mensagens estáticas"
- Menciona `errors.Is()` e `errors.As()` para unwrap

---

#### 🔟 "Por que não usar um framework como Gin ou Echo?"

**Resposta Pleno:**
- "O teste pediu net/http nativo"
- Não entende trade-offs

**Resposta Sênior:**
- **Teste pediu nativo** para avaliar conhecimento do Go stdlib
- **Trade-offs**:
  - Frameworks: mais produtivo, middleware pronto, validações
  - Nativo: mais controle, menos dependências, deploy mais leve
- "Para **microserviços simples**, net/http é suficiente"
- "Para **APIs complexas**, frameworks ajudam na produtividade"
- Conhece pelo menos 2-3 frameworks (Gin, Echo, Fiber, Chi)

---

## 📊 Pontuação Final

| Pontuação | Nível | Decisão |
|-----------|-------|---------|
| **80-100+** | ✅ **Sênior** | Aprovar |
| **60-79** | ⭐ **Pleno** | Aprovar para pleno |
| **40-59** | ⚠️ **Júnior+** | Não se encaixa |
| **<40** | ❌ **Júnior** | Rejeitar |

---

## ✅ Checklist de Avaliação Rápida

### Durante o código (10 min)
- [ ] Criou método `Validate()`?
- [ ] Implementou `String()`?
- [ ] Extraiu constantes `MinAge/MaxAge`?
- [ ] Handler valida método HTTP?
- [ ] Criou helper functions?
- [ ] Usa status codes corretos?
- [ ] Adiciona `defer r.Body.Close()`?

### Perguntas essenciais
- [ ] Explica bem por que criou `Validate()`?
- [ ] Sabe o que é interface `Stringer`?
- [ ] Entende `defer r.Body.Close()`?
- [ ] Pensa em testes automatizados?

### Decisão final
- [ ] **Sênior**: 4+ items do código + 3+ perguntas bem respondidas
- [ ] **Pleno**: 2-3 items do código + 2+ perguntas razoáveis
- [ ] **Não se encaixa**: <2 items ou não consegue explicar decisões

---

## 🎓 Conclusão

Este teste diferencia claramente **Pleno** de **Sênior** através de:

✅ **Código**: Validate() + String() + constantes + helpers = Sênior  
✅ **Explicação**: Consegue articular "por quê" das decisões = Sênior  
✅ **Visão**: Pensa em testabilidade, manutenibilidade, produção = Sênior  

**Use este guia durante a entrevista para avaliação consistente e objetiva.**
