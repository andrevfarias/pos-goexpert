# Tecnologias e Pacotes Go Utilizados no Projeto

## **Principais Recursos da Biblioteca Padrão do Go**

### **1. Logging**

- `log/slog`: Biblioteca moderna para logging estruturado.

### **2. HTTP e REST**

- `net/http`: Biblioteca padrão para trabalhar com requisições HTTP.
- `encoding/json`: Para codificação e decodificação de JSON.
- `context`: Para gerenciamento de contexto em requisições.

## **Principais Pacotes Externos Utilizados**

### **1. Web Framework e API REST**

- `github.com/go-chi/chi/v5`: Roteador minimalista e eficiente para APIs HTTP.
- `github.com/go-chi/cors`: Middleware para gerenciar CORS (Cross-Origin Resource Sharing).

### **2. Clientes HTTP**

- `github.com/go-resty/resty/v2`: Cliente HTTP elegante e simplificado para chamadas a APIs externas (ViaCEP e WeatherAPI).

### **3. Variáveis de Ambiente**

- `github.com/joho/godotenv`: Para carregar variáveis de ambiente de um arquivo `.env`.

### **4. Validação**

- `github.com/go-playground/validator/v10`: Para validação do formato do CEP e outros dados de entrada.

### **5. Mocks para Testes**

- `github.com/stretchr/testify`: Framework de testes para facilitar a escrita de testes unitários e de integração.
