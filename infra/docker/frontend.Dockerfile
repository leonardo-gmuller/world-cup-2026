FROM --platform=linux/amd64 node:20-bookworm-slim AS builder

WORKDIR /app

COPY package.json package-lock.json ./

RUN npm install --include=optional

COPY . .

RUN npm run build

FROM --platform=linux/amd64 nginx:1.27-alpine

COPY interface/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/interface/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]