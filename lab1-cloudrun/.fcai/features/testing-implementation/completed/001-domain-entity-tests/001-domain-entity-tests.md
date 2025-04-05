# Implementação de Testes para Entidades de Domínio

## Objetivo

Implementar testes unitários abrangentes para as entidades `Weather` e `Address`, garantindo que todos os comportamentos e validações estejam funcionando corretamente.

## Escopo

### Entidade Weather

1. Teste do construtor `NewWeather`
2. Teste de conversão para Fahrenheit
3. Teste de conversão para Kelvin
4. Teste de serialização JSON
5. Teste de casos de borda (temperaturas extremas)

### Entidade Address

1. Teste do construtor `NewAddress`
2. Teste de validação de CEP
3. Teste de formatação de CEP
4. Teste de limpeza de CEP
5. Teste de casos inválidos

## Critérios de Aceitação

- [ ] Cobertura de 100% para ambas as entidades
- [ ] Testes documentados e organizados
- [ ] Uso de table-driven tests
- [ ] Validação de todos os casos de erro
- [ ] Testes executando sem falhas

## Dependências

- Pacote `testing` do Go
- Pacote `testify/assert` para asserções

## Status

- Iniciado em: [DATA]
- Status: Em Andamento
- Responsável: Sistema
