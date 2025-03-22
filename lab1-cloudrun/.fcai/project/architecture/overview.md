# Fluxo Completo da Aplicação de Consulta de Temperatura por CEP

## \*\*Fase 1 - Implementação da API

### **1. Recebimento do CEP pela API**

- O cliente (frontend ou outra aplicação) faz uma requisição **HTTP GET** enviando o CEP para a API.
- A API:
  1. Valida se o CEP possui formato correto (8 dígitos).
  2. Registra logs da requisição com informações de:
     - CEP requisitado.
     - Timestamp da solicitação.
     - IP do solicitante.
  3. Inicia o fluxo de processamento da consulta.

### **2. Processamento da Consulta de Localização**

- Um **serviço de localização** é acionado para processar o CEP:
  1. Faz uma requisição HTTP para a **API ViaCEP**.
  2. Extrai os dados da cidade/localidade da resposta.
  3. Trata possíveis erros:
     - CEP não encontrado (404)
     - Falhas de comunicação com a API externa

### **3. Consulta da Temperatura Atual**

- Após obter a localização, um **serviço de clima** realiza:
  1. Requisição HTTP para a **API WeatherAPI** com o nome da cidade.
  2. Extração da temperatura atual em graus Celsius.
  3. Conversão da temperatura para múltiplas unidades:
     - Fahrenheit usando a fórmula `F = C * 1,8 + 32`
     - Kelvin usando a fórmula `K = C + 273`

### **4. Retorno da Resposta ao Cliente**

- **API HTTP** responde com:
  - Status 200 e JSON com temperaturas em caso de sucesso: `{ "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }`
  - Status 422 e mensagem "invalid zipcode" para CEP em formato inválido
  - Status 404 e mensagem "can not find zipcode" para CEP não encontrado

---

## **Contexto Geral do Projeto**

- **Nome do Módulo:** `github.com/andrevfarias/goexpert/lab1-cloudrun`
- **Versão do Go:** 1.24
- **Objetivo do Projeto:**

  - Criar uma API REST eficiente para consulta de temperaturas a partir de um CEP fornecido.
  - A solução permite que aplicações enviem um CEP via API HTTP e recebam as temperaturas atuais na localidade correspondente em diferentes unidades (Celsius, Fahrenheit e Kelvin).
  - Implementar tratamento adequado de erros para CEPs inválidos ou não encontrados.
  - Realizar deploy da aplicação no Google Cloud Run para garantir disponibilidade e escalabilidade.

- **Principais Tecnologias:**

  - **Linguagem:** Go 1.24
  - **APIs Externas:**
    - ViaCEP para conversão de CEP em localização
    - WeatherAPI para consulta de dados climáticos
  - **Container:** Docker para desenvolvimento e implantação via docker-compose.yml
  - **Cloud:** Google Cloud Run para hospedagem da aplicação
  - **Testes:** Testes automatizados para validação do funcionamento
