FROM node:20.18-alpine AS build
WORKDIR /api
COPY . .
RUN npm install
RUN npm run build

RUN ls

FROM node:20.18-alpine
COPY --from=build /api/dist .
COPY --from=build /api/package.json .
COPY --from=build /api/public .
RUN npm install --omit=dev
EXPOSE 3000
CMD ["npm","start"]