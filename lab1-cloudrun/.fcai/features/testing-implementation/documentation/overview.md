# Implementação de Testes Automatizados

## Visão Geral

Esta feature tem como objetivo implementar uma cobertura abrangente de testes automatizados para garantir a qualidade e confiabilidade do código.

## Objetivos

1. Implementar testes unitários para as entidades de domínio
2. Implementar testes unitários para os casos de uso
3. Implementar testes de integração para os adaptadores
4. Configurar ferramentas de cobertura de código
5. Estabelecer pipeline de testes automatizados

## Componentes a Serem Testados

### Domínio

- Entidade `Weather`
  - Conversões de temperatura
  - Serialização JSON
- Entidade `Address`
  - Validação de CEP
  - Formatação de CEP

### Casos de Uso

- `GetTemperatureByZipCode`
  - Fluxo de sucesso
  - Cenários de erro
  - Validações de entrada

### Adaptadores

- Cliente ViaCEP
- Cliente WeatherAPI
- Handlers HTTP

## Ferramentas e Práticas

- Uso do pacote `testing` do Go
- Mocks com `testify/mock`
- Asserções com `testify/assert`
- Cobertura de código com `go test -cover`
- Testes de tabela (table-driven tests)
- Subtestes organizados

## Critérios de Aceitação

1. Cobertura mínima de 80% para o código de domínio
2. Testes de integração funcionando em ambiente CI
3. Documentação clara dos testes
4. Mocks apropriados para dependências externas
