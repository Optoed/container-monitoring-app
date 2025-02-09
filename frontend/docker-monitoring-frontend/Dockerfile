FROM node:16 AS build

WORKDIR /app

COPY package.json ./

RUN yarn install

COPY . .

ARG REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND_URL=$REACT_APP_BACKEND_URL

RUN yarn build

FROM nginx:alpine

COPY --from=build /app/build /usr/share/nginx/html

# Копируем кастомную конфигурацию nginx
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]
