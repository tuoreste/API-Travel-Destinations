# Build Stage
FROM node:18-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --include=dev

COPY . .
RUN npm run build

# Nginx Stage
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

# Expose the port Nginx runs on
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
