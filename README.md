# golang-rag

Chat em linha de comando escrito em Go que conversa com qualquer servidor
compatível com a API da OpenAI (Ollama, LM Studio, OpenAI, Groq, etc.).
Para trocar de fornecedor, basta alterar o `.env` — nenhuma mudança de código.

## Configuração do `.env`

O projeto carrega as configurações de um arquivo `.env` na raiz. Se ele não
existir, crie um:

```bash
touch .env
```

### Variáveis disponíveis

| Variável             | Obrigatória | Descrição                                                                 | Padrão se omitida          |
| -------------------- | ----------- | ------------------------------------------------------------------------- | -------------------------- |
| `OPENAI_BASE_URL`    | Sim*        | Endereço do servidor compatível com a API da OpenAI                       | `https://api.openai.com/v1`|
| `OPENAI_API_KEY`     | Depende     | Chave de acesso. Servidores locais (Ollama, LM Studio) não exigem autenticação | —                    |
| `OPENAI_MODEL`       | Sim*        | Identificador do modelo de chat                                           | `gpt-4o-mini`              |
| `SYSTEM_PROMPT_FILE` | Não         | Caminho de um arquivo `.md`/texto cujo conteúdo vira o *system prompt*    | *(sem system prompt)*      |

\* Se quiser usar a OpenAI oficial, os padrões já funcionam — basta definir
`OPENAI_API_KEY`.

> **Observação:** variáveis já exportadas no ambiente têm precedência sobre o
> `.env` (convenção *twelve-factor*).

### Exemplo 1 — Ollama local (padrão deste projeto)

Com o [Ollama](https://ollama.com) instalado e o serviço rodando (ele escuta
na porta `11434` por padrão):

```env
OPENAI_BASE_URL=http://localhost:11434/v1
OPENAI_API_KEY=dummy-key
OPENAI_MODEL=gemma3:latest
```

- O `OPENAI_MODEL` precisa corresponder a um nome listado por `ollama list`.
  Baixe o modelo com `ollama pull gemma3` (ou outro de sua preferência).
- A chave pode ser qualquer valor (`dummy-key`): o servidor local não valida
  autenticação, mas enviar o cabeçalho vazio pode ser rejeitado por alguns
  servidores — por isso um valor qualquer é o mais seguro.

### Exemplo 2 — Outros fornecedores

Qualquer endpoint que fale o protocolo da OpenAI funciona sem alterar código.

```env
# OpenAI (oficial)
OPENAI_BASE_URL=https://api.openai.com/v1
OPENAI_API_KEY=sk-sua-chave-aqui
OPENAI_MODEL=gpt-4o-mini
```

```env
# Ollama Cloud
OPENAI_BASE_URL=https://ollama.com/v1
OPENAI_API_KEY=sua-chave-ollama
OPENAI_MODEL=nome-do-modelo
```

```env
# LM Studio (local)
OPENAI_BASE_URL=http://localhost:1234/v1
OPENAI_API_KEY=dummy-key
OPENAI_MODEL=nome-do-modelo-carregado
```

Consulte a documentação de cada fornecedor para saber o endereço correto da
API e os identificadores de modelo disponíveis.

## Como rodar

```bash
go run ./cmd/rag
```

Digite suas mensagens no terminal. Use `Q` (ou `exit`/`quit`) para sair.
