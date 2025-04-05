# Documentação da API

## Descrição

Documentação completa da API de Consulta de Temperatura por CEP, incluindo endpoints, parâmetros, respostas e exemplos.

## Status

✅ Concluído

## Objetivos

- ✅ Documentar todos os endpoints da API
  - ✅ Endpoint de consulta de temperatura
  - ✅ Endpoint de health check
- ✅ Documentar formatos de requisição e resposta
- ✅ Documentar códigos de status HTTP
- ✅ Documentar exemplos de uso
- ✅ Documentar estruturas de dados

## Tecnologias e Ferramentas

- OpenAPI/Swagger 3.0.0
- YAML para definição da documentação
- Markdown para documentação adicional

## Documentação Implementada

### Endpoints

1. GET /temperature

   - Parâmetros:
     - zipcode (query): CEP no formato de 8 dígitos
   - Respostas:
     - 200: Temperatura encontrada
     - 400: CEP não informado
     - 422: CEP inválido
     - 404: CEP não encontrado
     - 500: Erro interno

2. GET /health
   - Respostas:
     - 200: API funcionando normalmente

### Estruturas de Dados

1. Temperature
   - temp_c: Temperatura em Celsius (float)
   - temp_f: Temperatura em Fahrenheit (float)
   - temp_k: Temperatura em Kelvin (float)

## Métricas de Qualidade

- ✅ Todos os endpoints documentados
- ✅ Todos os parâmetros documentados
- ✅ Todos os códigos de status documentados
- ✅ Exemplos incluídos para todas as operações
- ✅ Documentação em formato padrão OpenAPI

## Próximos Passos

1. [ ] Adicionar autenticação e autorização
2. [ ] Documentar limites de taxa
3. [ ] Adicionar exemplos de uso com curl
4. [ ] Implementar interface Swagger UI

## Observações

- Documentação segue especificação OpenAPI 3.0.0
- Exemplos incluídos para facilitar o entendimento
- Documentação mantida junto ao código fonte
- Formato YAML escolhido para melhor legibilidade
