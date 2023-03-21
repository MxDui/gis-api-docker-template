FROM node:lts-alpine

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

RUN apk add --update postgresql-client jpeg-dev
RUN apk add --update --virtual .tmp-build-deps \
    gcc libc-dev linux-headers postgresql-dev musl-dev zlib zlib-dev

# Install postgis and gdal
RUN apk add --no-cache postgis
RUN apk add --no-cache gdal

# install proj
RUN apk add --no-cache proj-util


EXPOSE 3000

CMD ["npm", "start"]