# Contexto Geral do Projeto: Lab1 do curso de pós graduação GoExpert da FullCycle. API de Consulta de Temperatura por CEP

## **Descrição do Projeto**

O projeto tem como objetivo criar uma **API REST eficiente** para consulta de temperaturas a partir de um CEP. A aplicação será desenvolvida em Go, utilizando serviços externos para consulta de localização e dados climáticos, com posterior publicação no **Google Cloud Run**.

A solução permitirá que aplicações externas enviem um CEP via **API HTTP**, iniciando automaticamente a consulta de localização através do ViaCEP (ou similar), seguida pela consulta de temperatura atual através do WeatherAPI (ou similar), retornando os dados formatados em diferentes unidades de temperatura.

## **Objetivo do Projeto**

- Fornecer um serviço de consulta de temperatura **simples e eficiente**.
- Implementar uma API **intuitiva** para consulta de temperatura a partir de CEP.
- Converter automaticamente as temperaturas para diferentes unidades (Celsius, Fahrenheit e Kelvin).
- Realizar o deploy da aplicação no **Google Cloud Run**, garantindo disponibilidade e escalabilidade.

## **Principais Tecnologias Utilizadas**

- **Linguagem:** Go 1.24
- **APIs:** RESTful HTTP para consulta de temperatura por CEP.
- **APIs externas:**
  - **ViaCEP** (ou similar) para conversão de CEP em localização
  - **WeatherAPI** (ou similar) para consulta de dados climáticos
- **Containers:** Docker para desenvolvimento e implantação
- **Cloud:** Google Cloud Run para hospedagem da aplicação
- **Testes:** Testes automatizados para validação do funcionamento

## **Fluxo Geral da Aplicação**

1. **Recebimento do CEP**: O usuário envia um CEP via API HTTP.
2. **Validação do CEP**: A API valida se o CEP possui formato válido (8 dígitos).
3. **Consulta de Localização**: O sistema consulta o serviço ViaCEP para obter a cidade correspondente.
4. **Consulta de Temperatura**: Com base na cidade, o sistema consulta o WeatherAPI para obter a temperatura atual.
5. **Conversão de Unidades**: A temperatura é convertida para Celsius, Fahrenheit e Kelvin.
6. **Resposta ao Usuário**: A API retorna as temperaturas nas três unidades com status adequado.

## **Tratamento de Erros**

- **CEP Inválido**: Retorno de status 422 com mensagem "invalid zipcode"
- **CEP Não Encontrado**: Retorno de status 404 com mensagem "can not find zipcode"
- **Sucesso**: Retorno de status 200 com as temperaturas nas três unidades

## **Desafios Técnicos e Soluções**

1. **Integração com APIs Externas**

- Implementação de clients HTTP resilientes para comunicação com ViaCEP e WeatherAPI.
- Tratamento adequado de erros e timeouts nas chamadas externas.

2. **Tratamento de Formatos e Conversões**

- Validação rigorosa do formato do CEP.
- Implementação precisa das fórmulas de conversão de temperatura.

3. **Deployment no Google Cloud Run**

- Configuração adequada de Dockerfile para otimização do container.
- Configuração do serviço no Google Cloud Run para escalabilidade automática.

## **Nome do Módulo**

`github.com/andrevfarias/goexpert/lab1-cloudrun`
