FROM node:14-alpine as build

COPY . /project

WORKDIR /project

RUN npm i

ARG userurl

ARG contenturl

ENV REACT_APP_USER_URL=$userurl

ENV REACT_APP_CONTENT_URL=${contenturl}

RUN npm run build

# production env
FROM nginx:stable-alpine

COPY --from=build /project/build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]