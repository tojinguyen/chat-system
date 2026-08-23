FROM node:22-alpine AS deps
WORKDIR /usr/src/app

COPY package*.json .
RUN npm ci

FROM node:22-alpine AS builder
WORKDIR /usr/src/app

COPY --from=deps /usr/src/app/node_modules ./node_modules
COPY . .

RUN npm run build

ENV NODE_ENV production
RUN npm prune --production

FROM node:22-alpine AS runner
WORKDIR /usr/src/app

ENV NODE_ENV production

USER node

COPY --from=builder --chown=node:node /usr/src/app/node_modules ./node_modules
COPY --from=builder --chown=node:node /usr/src/app/dist ./dist
COPY --from=builder --chown=node:node /usr/src/app/package*.json ./

EXPOSE 3000
CMD ["node", "dist/main.js"]