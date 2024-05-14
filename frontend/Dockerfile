FROM node:21-alpine

RUN npm i -g pnpm

COPY ./package.json /frontend

WORKDIR /frontend

RUN pnpm i

COPY . /frontend

RUN pnpm build

EXPOSE 8000

CMD ["pnpm" , "start"]