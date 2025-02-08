FROM node:16 AS build

WORKDIR /app

# COPY package.json yarn.lock ./
COPY package.json ./

RUN yarn install

COPY . .

# Передаём REACT_APP_BACKEND_URL в процесс сборки
ARG REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND_URL=$REACT_APP_BACKEND_URL

RUN yarn build

# Используем Nginx для отдачи статики
FROM nginx:alpine

# Копируем собранные файлы из первого этапа в директорию для статики Nginx
COPY --from=build /app/build /usr/share/nginx/html

EXPOSE 80

# Запускаем Nginx
CMD ["nginx", "-g", "daemon off;"]
