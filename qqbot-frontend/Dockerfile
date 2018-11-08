# build stage
FROM node:10.5.0-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:1.13.12-alpine
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
