# Pipeline CI/CD

## Descrição

Pipeline de Integração Contínua e Entrega Contínua (CI/CD) para automatizar testes, build e deploy da aplicação no Google Cloud Run.

## Status

✅ Configurado

## Objetivos

- ✅ Automatizar execução de testes
  - ✅ Testes unitários
  - ✅ Verificação de formatação de código
  - ✅ Análise de cobertura de código
- ✅ Automatizar build da aplicação
  - ✅ Build da imagem Docker
  - ✅ Push para Google Container Registry
- ✅ Automatizar deploy
  - ✅ Deploy no Google Cloud Run
  - ✅ Configuração de variáveis de ambiente
  - ✅ Configuração de acesso público

## Tecnologias e Ferramentas

- GitHub Actions
- Google Cloud Platform
  - Cloud Run
  - Container Registry
- Docker
- Codecov para análise de cobertura

## Pipeline Implementado

### Jobs

1. test

   - Checkout do código
   - Setup do Go
   - Verificação de formatação
   - Execução de testes
   - Upload de cobertura de código

2. deploy (apenas na main)
   - Autenticação no GCP
   - Build da imagem Docker
   - Push para Container Registry
   - Deploy no Cloud Run

## Métricas de Qualidade

- ✅ Testes automatizados em cada push/PR
- ✅ Verificação de formatação de código
- ✅ Análise de cobertura de código
- ✅ Deploy automatizado na main

## Próximos Passos

1. [ ] Adicionar análise estática de código
2. [ ] Implementar testes de integração no pipeline
3. [ ] Configurar ambientes de staging/produção
4. [ ] Adicionar notificações de status

## Observações

- Pipeline configurado para executar em pushes na main e pull requests
- Deploy automático apenas em pushes na main
- Necessário configurar secrets no GitHub:
  - GCP_PROJECT_ID
  - GCP_SA_KEY
  - WEATHER_API_KEY
