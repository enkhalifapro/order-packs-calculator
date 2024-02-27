FROM node AS builder

# Set the Current Working Directory inside the container
WORKDIR /tmp/order-pack-calculator-frontend

COPY . .

RUN npm install

# Build
RUN npm run build


# Start fresh from a smaller image
FROM nginx:alpine

COPY --from=builder /tmp/order-pack-calculator-frontend/dist /app/static

COPY ./nginx.conf /etc/nginx/nginx.conf

WORKDIR /app