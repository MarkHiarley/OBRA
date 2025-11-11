# Bug: Validação contraditória de `aprovado_por_id` com `status_aprovacao`

Status: Draft

Data: 2025-11-10

Resumo
------
Há validação contraditória relacionada ao campo `aprovado_por_id` quando `status_aprovacao = "PENDENTE"`.

Comportamento observado (bug):
- Enviar uma requisição com `status_aprovacao = "PENDENTE"` e `aprovado_por_id` com um ID válido (ex.: 13) retorna erro 400 "deve ser nulo".
- Em alguns fluxos, `aprovado_por_id = 0` ou omitido é tratado de forma inconsistente.

Comportamento esperado (regras de negócio)
-----------------------------------------
Tabela resumida:

| status_aprovacao | aprovado_por_id enviado | Resultado esperado |
|------------------|-------------------------:|-------------------:|
| PENDENTE         | omitido (não enviado)   | 200 OK (persistir NULL) |
| PENDENTE         | null                    | 200 OK (persistir NULL) |
| PENDENTE         | 0                       | 200 OK (opcional: tratar como NULL) |
| PENDENTE         | 13 (id válido)          | 400 Bad Request (deve ser nulo) |
| APROVADO         | omitido                 | 400 Bad Request (obrigatório) |
| APROVADO         | 0                       | 400 Bad Request (obrigatório) |
| APROVADO         | 13 (id válido)          | 200 OK (válido) |
| REJEITADO        | omitido                 | 400 Bad Request (obrigatório) |
| REJEITADO        | 0                       | 400 Bad Request (obrigatório) |
| REJEITADO        | 13 (id válido)          | 200 OK (válido) |

Notas:
- Do ponto de vista do banco de dados, o valor `0` normalmente não existe como FK. A aplicação deve converter `0` → `NULL` quando isso for sinal de "ausência".
- A constraint de BD (CHECK) pode exigir `aprovado_por_id IS NULL` para `PENDENTE`; portanto a aplicação deve normalizar/informar NULL quando necessário.

Arquivos/locais prováveis afetados
----------------------------------
- `internal/models/diario.go` (struct `DiarioObra` — campo `AprovadoPorID`)
- `internal/controllers/diario.go` (parsing do JSON / normalização do campo)
- `internal/usecases/diario.go` (validação de regras de negócio, onde `status_aprovacao` é avaliado)
- `internal/services/diario.go` (persistência em DB — conversão para `nil` nos parâmetros SQL)
- `migrations/000017_fix_diario_aprovador.up.sql` (já existe algo semelhante — revisar)

Causa provável
--------------
- O campo é tratado como `null.Int` (ou similar) e a validação aplicada exige `NULL` para `PENDENTE`, mas a aplicação não está normalizando `0` → `NULL` nem tratando apropriadamente quando o campo é omitido.
- Em outro fluxo, a validação está sendo aplicada antes da normalização, causando rejeição inconsistentes.

Correção proposta (passos + snippets)
-------------------------------------
1) Model: tornar o campo JSON-friendly (opções):
   - Alternativa A (recomendada): usar ponteiro `*int` com tag `omitempty` no JSON.

```go
// internal/models/diario.go (trecho)
AprovadoPorID *int `json:"aprovado_por_id,omitempty"`
```

   - Alternativa B: manter `null.Int` mas garantir `json:"aprovado_por_id,omitempty"` e normalizar `0` para `Valid=false`.

2) Controller: Normalizar o valor recebido antes de validação/usecase

```go
// Exemplo no controller, após bind:
if diario.AprovadoPorID != nil && *diario.AprovadoPorID == 0 {
    diario.AprovadoPorID = nil
}
```

Se usar `null.Int`, então:
```go
if diario.AprovadoPorID.Valid && diario.AprovadoPorID.Int64 == 0 {
    diario.AprovadoPorID = null.Int{}
}
```

3) Usecase: Validação condicional baseada em `status_aprovacao`

```go
func validateAprovador(status string, aprovadoPorID *int) error {
    s := strings.ToUpper(strings.TrimSpace(status))
    switch s {
    case "PENDENTE":
        // aceita nil
        if aprovadoPorID != nil && *aprovadoPorID > 0 {
            return fmt.Errorf("aprovado_por_id deve ser nulo quando status_aprovacao = PENDENTE")
        }
    case "APROVADO", "REJEITADO":
        if aprovadoPorID == nil || *aprovadoPorID <= 0 {
            return fmt.Errorf("aprovado_por_id é obrigatório quando status_aprovacao = %s", s)
        }
    default:
        // comportamento padrão
    }
    return nil
}
```

Se estiver usando `null.Int`, converta antes para `*int` ou adapte a validação para `null.Int`.

4) Service (persistência): montar parâmetros SQL com `nil` quando `AprovadoPorID == nil`

```go
var aprovadoPor interface{}
if diario.AprovadoPorID != nil {
    aprovadoPor = int64(*diario.AprovadoPorID)
} else {
    aprovadoPor = nil
}
// use 'aprovadoPor' como arg para Exec/QueryRow
```

5) DB Constraint: revisar `migrations/000017_fix_diario_aprovador.up.sql` e garantir que a constraint CHECK aceita os casos esperados. Não recomendar alterar a constraint para permitir `0` — prefira normalizar `0` → NULL na aplicação.

Testes sugeridos
----------------
Criar testes unitários para a função de validação e testes de integração para o endpoint (`POST /diarios` / `PUT /diarios/:id`) com os seguintes payloads:

- PENDENTE + omitido -> 200
- PENDENTE + null -> 200
- PENDENTE + 0 -> 200 (após normalizar para NULL)
- PENDENTE + 13 -> 400
- APROVADO + omitido -> 400
- APROVADO + 0 -> 400
- APROVADO + 13 -> 200
- REJEITADO idem APROVADO

Exemplo de teste unitário (pseudocódigo):

```go
func TestValidateAprovador(t *testing.T) {
    // PENDENTE + nil => ok
    if err := validateAprovador("PENDENTE", nil); err != nil {
        t.Fatal(err)
    }
    // PENDENTE + 13 => error
    id := 13
    if err := validateAprovador("PENDENTE", &id); err == nil {
        t.Fatal("expected error when aprovado_por_id present for PENDENTE")
    }
}
```

Checklist de deploy
-------------------
- Aplicar mudanças no código
- rodar `go build` e `go test ./...`
- subir nova imagem / container
- executar migrations se necessário (revisões de constraints)

Mudanças no repositório
-----------------------
Eu posso gerar um patch com as mudanças no model/controller/usecase e adicionar testes. Indique se devo aplicar as alterações agora (eu edito os arquivos) ou apenas gerar o arquivo Markdown para enviar ao desenvolvedor.

---

"Autor": Automatizado (relatório gerado pelo assistente de desenvolvimento)