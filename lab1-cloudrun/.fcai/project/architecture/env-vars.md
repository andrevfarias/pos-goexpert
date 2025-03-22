# Uso de Variáveis de Ambiente no Projeto

## **Visão Geral**

O projeto utilizará **variáveis de ambiente** para tornar a configuração mais flexível, segura e adaptável a diferentes ambientes de execução. Em vez de definir valores sensíveis diretamente no código, utilizaremos variáveis para armazenar informações como **chaves de API** e **configurações de serviços externos**.

## **Motivação para Uso de Variáveis de Ambiente**

- **Flexibilidade**: Permite configurar o sistema para diferentes ambientes (desenvolvimento, teste, produção) sem modificar o código.
- **Segurança**: Evita armazenar credenciais sensíveis no código-fonte, reduzindo riscos de exposição acidental.
- **Facilidade de Deploy**: As configurações podem ser alteradas diretamente via `.env` ou no ambiente Cloud Run.

## **Principais Configurações Usadas no Projeto**

### **Configurações de Serviços Externos e APIs**

- `WEATHER_API_KEY`: Chave de API para acesso ao serviço WeatherAPI.
- `WEATHER_API_BASE_URL`: URL base para a API WeatherAPI (opcional, pode ter um valor padrão).
- `VIACEP_API_BASE_URL`: URL base para a API ViaCEP (opcional, pode ter um valor padrão).
- `API_TIMEOUT_SECONDS`: Timeout para requisições às APIs externas (valor padrão: 10).

### **Configurações Gerais**

- `APP_ENV`: Define o ambiente de execução (`development`, `staging`, `production`).

## **Carregamento das Variáveis no Ambiente Docker**

- As variáveis serão carregadas automaticamente a partir de um arquivo `.env`, garantindo que a configuração possa ser facilmente alterada sem modificar os arquivos do código.
- No **Docker Compose**, as variáveis serão passadas para os contêineres; os valores default podem ser especificados no serviço através da seção: environment.

## **Carregamento no Google Cloud Run**

- No ambiente de produção no Google Cloud Run, as variáveis de ambiente serão configuradas no momento da implantação.
- Valores sensíveis como chaves de API podem ser armazenados no Secret Manager do Google Cloud e referenciados no serviço Cloud Run.
- O Cloud Run permite atualizar variáveis de ambiente sem a necessidade de reconstruir a imagem do container.
