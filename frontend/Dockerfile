FROM node:25

WORKDIR /app

RUN npm install -g pnpm@10.17.0

COPY . .

RUN chmod +x docker-dev.sh

EXPOSE 5173

CMD ["./docker-dev.sh"]