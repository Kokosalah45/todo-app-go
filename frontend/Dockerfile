FROM node:21-alpine

RUN npm i -g pnpm

WORKDIR /frontend

COPY ./package.json .

RUN pnpm i

COPY . .

RUN pnpm build

EXPOSE 3344

CMD ["pnpm" , "preview"]