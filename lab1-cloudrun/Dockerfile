FROM golang:1.24-alpine

WORKDIR /app

# Instalação de ferramentas de desenvolvimento
RUN apk add --no-cache \
    git \
    make \
    curl \
    gcc \
    musl-dev

# Copia o arquivo go.mod (e potencialmente go.sum) para o contêiner
COPY go.mod ./
COPY go.sum ./

# Baixa todas as dependências. Serão armazenadas em cache se os arquivos go.mod e go.sum não mudarem
RUN go mod download

# Manter o contêiner em execução
CMD ["tail", "-f", "/dev/null"] 