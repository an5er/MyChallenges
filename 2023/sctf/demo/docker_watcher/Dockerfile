FROM ubuntu:latest

RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release

RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor > /usr/share/keyrings/docker-archive-keyring.gpg

RUN echo "deb [signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null

RUN apt-get update

RUN apt-get install -y docker-ce docker-ce-cli containerd.io

RUN cp /usr/bin/docker /usr/local/bin/

#server
RUN apt-get install -y nodejs npm

RUN mkdir /app
COPY src/server /app/server
WORKDIR /app/server

RUN npm install

#file server node
WORKDIR /app/node
COPY src/node /app/node

#nginx proxy
RUN apt-get install -y nginx
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

#main
WORKDIR /app
COPY src/build.sh /app
RUN chmod 777 build.sh

#front
RUN rm -rf /usr/share/nginx/html
COPY src/front/dist/ /usr/share/nginx/html

ENTRYPOINT ["sh","build.sh"]

