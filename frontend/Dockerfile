FROM node:18-alpine as dev

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
CMD ["npm", "run", "dev", "--", "--host"]
EXPOSE 3000


FROM node:18-alpine as builder

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build
