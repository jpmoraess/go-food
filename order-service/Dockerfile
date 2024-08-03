# Etapa de construção
FROM golang:1.22.5-alpine AS builder

# Configura o diretório de trabalho
WORKDIR /app

# Copia os arquivos de módulo do Go e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código fonte para o contêiner
COPY . .

# Compila o aplicativo Go
RUN CGO_ENABLED=0 GOOS=linux go build -o order ./cmd/main.go

################################################

# Etapa final
FROM scratch

# Copia o binário do estágio de construção
COPY --from=builder /app/order /

# Define o ponto de entrada do contêiner
ENTRYPOINT ["/order"]