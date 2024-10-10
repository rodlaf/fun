FROM golang:1.20-alpine AS backend

WORKDIR /ws

COPY ws/ .

RUN go build -o ws ws.go

FROM node:18-alpine AS frontend

WORKDIR /fun

COPY fun/ .

RUN npm install
RUN npm run build

# copy executable from backend to frontend container
COPY --from=backend /ws/ws .

CMD ["sh", "-c", "./ws & npm run start"]