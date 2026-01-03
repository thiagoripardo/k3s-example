# Visão geral do projeto

## Estrutura geral da base de código

- **Raiz (`/workspace/k3s-example`)**
  - `README.md`: descrição breve do projeto.
  - `go.work`: workspace Go com os módulos `api/bacon` e `api/validator`.
  - `api/`: diretório principal com os serviços e manifests.
- **`api/bacon/`**
  - `main.go`: inicializa o servidor do serviço bacon.
  - `server/server.go`: endpoints HTTP (`/`, `/generator`, `/health`), geração de bacon aleatório e chamada ao validator.
  - `domain/model/bacon.go`: modelos JSON, lista de bacons e helper `AsBuffer`.
  - `k8s/local-clusters/...`: manifests de Deployment, Service, Ingress e HPA para dev/prod.
- **`api/validator/`**
  - `main.go`: inicializa o servidor do validator.
  - `server/server.go`: endpoint `/validate` para validação do payload e `/health`.
  - `domain/model/bacon.go`: modelo de request (`BaconRequest`).
  - `k8s/local-clusters/...`: manifests de Deployment, Service e HPA para dev/prod.
- **`api/cloudflared/`**
  - `k8s/dev` e `k8s/prod`: manifests para o tunnel do Cloudflare com `ConfigMap` e `Secret`.

## Pontos importantes para saber

### 1) Fluxo principal entre serviços

- O serviço **bacon** expõe:
  - `GET /` (mensagem de boas-vindas)
  - `GET /generator` (gera um bacon aleatório, valida com o serviço validator, e retorna JSON)
  - `GET /health` (healthcheck)
- O serviço **validator** expõe:
  - `POST /validate` (verifica se `InstanceID`, `BaconName` e `Description` estão preenchidos)
  - `GET /health`
- O **bacon** chama o **validator** via Service interno do cluster, na porta exposta pelo Service.

### 2) Modelos e payloads

- O **bacon** define `BaconResponse` com `InstanceID` e `Bacon`, e serializa para JSON usando `AsBuffer()`.
- O **validator** recebe `BaconRequest` com a mesma estrutura de `InstanceID` e `Bacon`.

### 3) Kubernetes / k3s

- **Bacon**
  - Service expõe a porta 3000 para a aplicação em 8080.
  - Ingress com Traefik expõe o host `dev.lovelybacon.com`.
  - HPA escala entre 2 e 4 réplicas por CPU/Mem.
- **Validator**
  - Service expõe a porta 3001 para a aplicação em 8080.
  - HPA escala entre 2 e 4 réplicas por CPU/Mem.
- **Cloudflared**
  - Deployments para `dev` e `prod` configuram o túnel e as rotas de ingresso.

## Dicas para aprender mais rápido

1. **Entenda primeiro os dois serviços e o contrato entre eles**
   - Leia `api/bacon/server/server.go` e `api/validator/server/server.go` para entender endpoints, payloads e validações.
2. **Mapeie o fluxo HTTP e DNS no cluster**
   - Veja como o `Service` do validator expõe a porta e como o bacon consome esse endpoint.
3. **Conecte runtime e infraestrutura**
   - Compare portas, probes e Services/Ingresses nos manifests de Kubernetes.
4. **Aprenda o domínio do projeto**
   - Os modelos em `domain/model` são pequenos e ajudam a visualizar o payload JSON.
