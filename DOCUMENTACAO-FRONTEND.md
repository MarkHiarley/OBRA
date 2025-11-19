# 📚 DOCUMENTAÇÃO PARA O FRONTEND - RELATÓRIOS

**Data:** 19 de novembro de 2025  
**Versão:** 1.0  
**API Base URL:** `http://localhost:9090`

---

## 🔐 AUTENTICAÇÃO

Todos os endpoints requerem autenticação via Bearer Token.

### Login
```http
POST /login
Content-Type: application/json

{
  "email": "usuario@exemplo.com",
  "senha": "senha123"
}
```

**Resposta:**
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Use o `access_token` em todas as requisições:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## 📸 1. RELATÓRIO FOTOGRÁFICO

### Descrição
Retorna um relatório contendo apenas:
- Cabeçalho da empresa
- Resumo da obra (sem valores financeiros)
- Lista de todas as fotos da obra

### Endpoint
```http
GET /relatorios/fotografico/:obra_id
Authorization: Bearer {token}
```

### Parâmetros
| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `obra_id` | integer | Sim | ID da obra |

### Exemplo de Requisição
```javascript
const obraId = 5;
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";

fetch(`http://localhost:9090/relatorios/fotografico/${obraId}`, {
  method: 'GET',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  }
})
.then(response => response.json())
.then(data => {
  console.log('Relatório:', data.data);
})
.catch(error => console.error('Erro:', error));
```

### Resposta de Sucesso (200)
```json
{
  "data": {
    "cabecalho_empresa": {
      "nome_empresa": "EMPRESA CONSTRUTORA",
      "logotipo": null
    },
    "resumo_obra": {
      "nome_obra": "Casa Residencial - Fortaleza",
      "localizacao": ",  -  - Fortaleza - CE",
      "contrato_numero": "CONTR-2024-001",
      "lote": null,
      "descricao_breve": null,
      "foto_obra": null,
      "informacoes_gerais": "Relatório fotográfico da execução da obra"
    },
    "fotos": [
      {
        "id": 8,
        "url": "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAY...",
        "titulo_legenda": "Foto do período: tarde",
        "data": "2024-11-08T00:00:00Z",
        "observacao": "Validação completa do campo foto base64",
        "categoria": "DIARIO"
      }
    ]
  }
}
```

### Tipos TypeScript
```typescript
interface RelatorioFotografico {
  data: {
    cabecalho_empresa: {
      nome_empresa: string;
      logotipo: string | null;
    };
    resumo_obra: {
      nome_obra: string;
      localizacao: string;
      contrato_numero: string | null;
      lote: string | null;
      descricao_breve: string | null;
      foto_obra: string | null;  // Base64 ou URL da foto principal da obra
      informacoes_gerais: string;
    };
    fotos: Foto[];
  };
}

interface Foto {
  id: number;
  url: string;  // Base64 ou URL da imagem
  titulo_legenda: string | null;
  data: string | null;  // ISO 8601 date
  observacao: string | null;
  categoria: string;
}
```

### Possíveis Erros
| Código | Descrição |
|--------|-----------|
| 400 | ID da obra inválido |
| 401 | Token inválido ou ausente |
| 404 | Obra não encontrada |
| 500 | Erro interno do servidor |

---

## 📅 2. DIÁRIO DE OBRAS SEMANAL

### Descrição
Gera páginas semanais do diário de obras para um período específico.
- O campo `descricao` vem **vazio** (null) para o usuário preencher manualmente
- Cada semana tem 7 dias (última semana pode ter menos)
- Retorna os dias que tiveram trabalho registrado

### Endpoint
```http
POST /diarios/semanal
Authorization: Bearer {token}
Content-Type: application/json
```

### Body da Requisição
```json
{
  "obra_id": 5,
  "data_inicio": "2024-11-01",
  "data_fim": "2024-11-30"
}
```

### Parâmetros
| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `obra_id` | integer | Sim | ID da obra |
| `data_inicio` | string | Sim | Data inicial (formato: YYYY-MM-DD) |
| `data_fim` | string | Sim | Data final (formato: YYYY-MM-DD) |

### Exemplo de Requisição
```javascript
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...";

fetch('http://localhost:9090/diarios/semanal', {
  method: 'POST',
  headers: {
    'Authorization': `Bearer ${token}`,
    'Content-Type': 'application/json'
  },
  body: JSON.stringify({
    obra_id: 5,
    data_inicio: '2024-11-01',
    data_fim: '2024-11-30'
  })
})
.then(response => response.json())
.then(data => {
  console.log('Diário:', data.data);
  console.log('Total de semanas:', data.data.semanas.length);
})
.catch(error => console.error('Erro:', error));
```

### Resposta de Sucesso (200)
```json
{
  "data": {
    "dados_obra": {
      "nome_obra": "Casa Residencial - Fortaleza",
      "localizacao": ",  -  - Fortaleza - CE",
      "contrato_numero": "CONTR-2024-001",
      "contratante": "Não informado",
      "contratada": "Não informado"
    },
    "semanas": [
      {
        "numero": 1,
        "data_inicio": "2024-11-01",
        "data_fim": "2024-11-07",
        "descricao": null,
        "dias_trabalho": []
      },
      {
        "numero": 2,
        "data_inicio": "2024-11-08",
        "data_fim": "2024-11-14",
        "descricao": null,
        "dias_trabalho": []
      },
      {
        "numero": 3,
        "data_inicio": "2024-11-15",
        "data_fim": "2024-11-21",
        "descricao": null,
        "dias_trabalho": []
      }
    ]
  }
}
```

### Tipos TypeScript
```typescript
interface DiarioSemanal {
  data: {
    dados_obra: {
      nome_obra: string;
      localizacao: string;
      contrato_numero: string | null;
      contratante: string;
      contratada: string;
    };
    semanas: Semana[];
  };
}

interface Semana {
  numero: number;
  data_inicio: string;  // YYYY-MM-DD
  data_fim: string;     // YYYY-MM-DD
  descricao: string | null;  // VAZIO para o usuário preencher
  dias_trabalho: string[];   // Array de datas (YYYY-MM-DD)
}
```

### Possíveis Erros
| Código | Descrição |
|--------|-----------|
| 400 | Dados inválidos (obra_id, datas incorretas) |
| 401 | Token inválido ou ausente |
| 404 | Obra não encontrada |
| 500 | Erro interno do servidor |

---

## 🎨 COMO IMPLEMENTAR NO FRONTEND

### Relatório Fotográfico

#### Página de Visualização
```javascript
// React Example
function RelatorioFotografico({ obraId }) {
  const [relatorio, setRelatorio] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchRelatorio = async () => {
      try {
        const token = localStorage.getItem('token');
        const response = await fetch(
          `http://localhost:9090/relatorios/fotografico/${obraId}`,
          {
            headers: {
              'Authorization': `Bearer ${token}`,
              'Content-Type': 'application/json'
            }
          }
        );
        const data = await response.json();
        setRelatorio(data.data);
      } catch (error) {
        console.error('Erro ao buscar relatório:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchRelatorio();
  }, [obraId]);

  if (loading) return <div>Carregando...</div>;

  return (
    <div>
      {/* Cabeçalho */}
      <header>
        <h1>{relatorio.cabecalho_empresa.nome_empresa}</h1>
        {relatorio.cabecalho_empresa.logotipo && (
          <img src={relatorio.cabecalho_empresa.logotipo} alt="Logo" />
        )}
      </header>

      {/* Resumo da Obra */}
      <section>
        <h2>{relatorio.resumo_obra.nome_obra}</h2>
        
        {/* Foto principal da obra (se existir) */}
        {relatorio.resumo_obra.foto_obra && (
          <div className="foto-obra-principal">
            <img 
              src={relatorio.resumo_obra.foto_obra} 
              alt="Foto da obra"
              style={{ maxWidth: '100%', marginBottom: '1rem' }}
            />
          </div>
        )}
        
        <p>Localização: {relatorio.resumo_obra.localizacao}</p>
        <p>Contrato: {relatorio.resumo_obra.contrato_numero}</p>
        {relatorio.resumo_obra.lote && (
          <p>Lote: {relatorio.resumo_obra.lote}</p>
        )}
        {relatorio.resumo_obra.descricao_breve && (
          <p>{relatorio.resumo_obra.descricao_breve}</p>
        )}
      </section>

      {/* Galeria de Fotos */}
      <section>
        <h3>Fotos da Obra</h3>
        <div className="gallery">
          {relatorio.fotos.map(foto => (
            <div key={foto.id} className="foto-item">
              <img src={foto.url} alt={foto.titulo_legenda || 'Foto'} />
              <p>{foto.titulo_legenda}</p>
              <small>{new Date(foto.data).toLocaleDateString()}</small>
              {foto.observacao && <p>{foto.observacao}</p>}
            </div>
          ))}
        </div>
      </section>
    </div>
  );
}
```

### Diário de Obras

#### Página de Seleção de Período
```javascript
// React Example
function DiarioObras({ obraId }) {
  const [dataInicio, setDataInicio] = useState('');
  const [dataFim, setDataFim] = useState('');
  const [diario, setDiario] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleGerarDiario = async () => {
    setLoading(true);
    try {
      const token = localStorage.getItem('token');
      const response = await fetch(
        'http://localhost:9090/diarios/semanal',
        {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            obra_id: obraId,
            data_inicio: dataInicio,
            data_fim: dataFim
          })
        }
      );
      const data = await response.json();
      setDiario(data.data);
    } catch (error) {
      console.error('Erro ao gerar diário:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      {/* Seleção de Período */}
      <section>
        <h2>Gerar Diário de Obras</h2>
        <label>
          Data Início:
          <input 
            type="date" 
            value={dataInicio}
            onChange={(e) => setDataInicio(e.target.value)}
          />
        </label>
        <label>
          Data Fim:
          <input 
            type="date" 
            value={dataFim}
            onChange={(e) => setDataFim(e.target.value)}
          />
        </label>
        <button onClick={handleGerarDiario} disabled={loading}>
          {loading ? 'Gerando...' : 'Gerar Diário'}
        </button>
      </section>

      {/* Páginas Semanais */}
      {diario && (
        <section>
          <h3>{diario.dados_obra.nome_obra}</h3>
          <p>{diario.dados_obra.localizacao}</p>
          
          {diario.semanas.map(semana => (
            <div key={semana.numero} className="semana-card">
              <h4>Semana {semana.numero}</h4>
              <p>Período: {semana.data_inicio} até {semana.data_fim}</p>
              
              {/* Campo editável para descrição */}
              <label>
                Descrição (O que foi executado):
                <textarea 
                  placeholder="Descreva os serviços executados nesta semana..."
                  defaultValue={semana.descricao || ''}
                  rows={10}
                />
              </label>
              
              {semana.dias_trabalho && semana.dias_trabalho.length > 0 && (
                <p>
                  Dias de trabalho: {semana.dias_trabalho.join(', ')}
                </p>
              )}
              
              <button>Salvar Descrição</button>
            </div>
          ))}
        </section>
      )}
    </div>
  );
}
```

---

## 📋 FLUXO COMPLETO

### Relatório Fotográfico
```
1. Usuário seleciona uma obra
   ↓
2. Frontend chama GET /relatorios/fotografico/:obra_id
   ↓
3. Backend retorna: cabeçalho + resumo + fotos
   ↓
4. Frontend exibe em formato de relatório
   ↓
5. Usuário pode imprimir ou exportar PDF
```

### Diário de Obras
```
1. Usuário seleciona uma obra
   ↓
2. Usuário escolhe período (data início → data fim)
   ↓
3. Frontend chama POST /diarios/semanal
   ↓
4. Backend retorna: páginas semanais com descrição VAZIA
   ↓
5. Frontend exibe semanas como cards editáveis
   ↓
6. Usuário preenche descrição de cada semana manualmente
   ↓
7. Usuário salva cada descrição
   ↓
8. Frontend pode imprimir ou exportar PDF
```

---

## 🎯 BOAS PRÁTICAS

### 1. Validação de Datas
```javascript
// Validar datas antes de enviar
if (new Date(dataFim) < new Date(dataInicio)) {
  alert('Data final deve ser maior que data inicial');
  return;
}
```

### 2. Tratamento de Erros
```javascript
try {
  const response = await fetch(url, options);
  
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Erro ao buscar dados');
  }
  
  const data = await response.json();
  // processar dados...
} catch (error) {
  console.error('Erro:', error);
  // Exibir mensagem ao usuário
}
```

### 3. Loading States
```javascript
// Sempre mostrar feedback visual
const [loading, setLoading] = useState(false);

// Durante a requisição
setLoading(true);
// ... fetch ...
setLoading(false);
```

### 4. Formatação de Datas
```javascript
// Formatar data ISO para formato local
const formatarData = (dataISO) => {
  if (!dataISO) return 'N/A';
  return new Date(dataISO).toLocaleDateString('pt-BR');
};
```

### 5. Imagens Base64
```javascript
// As fotos vêm em Base64, podem ser usadas diretamente
<img src={foto.url} alt="Foto da obra" />
// URL já vem no formato: data:image/jpeg;base64,/9j/4AAQ...
```

---

## 🔗 ENDPOINTS COMPLETOS

| Endpoint | Método | Descrição |
|----------|--------|-----------|
| `/login` | POST | Autenticação |
| `/relatorios/fotografico/:obra_id` | GET | Relatório Fotográfico |
| `/diarios/semanal` | POST | Diário de Obras Semanal |

---

## 📞 SUPORTE

Em caso de dúvidas:
1. Verifique se o token está válido
2. Confirme o formato das datas (YYYY-MM-DD)
3. Valide o `obra_id`
4. Verifique os logs do navegador para erros

---

**Versão:** 1.0  
**Data:** 19 de novembro de 2025  
**Status:** Pronto para integração
