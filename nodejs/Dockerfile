# BASE IMAGE
FROM node:8-alpine AS base

RUN yarn config set no-progress true

WORKDIR /app
COPY package.json .
COPY yarn.lock .

## Build Image (with all deps)
FROM base AS build

# install node packages
RUN yarn install --production
# copy production node_modules aside (for later use)
RUN cp -R node_modules prod_node_modules
# install ALL node_modules, including 'devDependencies'
RUN yarn install

COPY . .
RUN yarn build

## Release Image (minimal deps)
FROM base AS release

COPY --from=build /app/prod_node_modules ./node_modules
COPY --from=build /app/lib ./lib

ENV NODE_ENV=production
ENV PORT=80
EXPOSE 80

CMD yarn start
