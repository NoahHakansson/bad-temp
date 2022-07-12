# syntax=docker/dockerfile:1

FROM node:16 as build

WORKDIR /app

# Download modules
COPY package.json ./
COPY package-lock.json ./
RUN npm install
# copy everything to /app
COPY . ./
# build
RUN npm run build

