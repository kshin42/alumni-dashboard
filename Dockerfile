FROM ubuntu:20.04

WORKDIR /src

# Install Golang
RUN apt-get update && apt-get install -y curl jq
RUN curl -OL https://golang.org/dl/go1.18.linux-amd64.tar.gz
RUN mv go1.18.linux-amd64.tar.gz /usr/local/
RUN tar -xf /usr/local/go1.18.linux-amd64.tar.gz -C /usr/local
RUN rm -rf /usr/local/go1.18.linux-amd64.tar.gz
ENV GOPATH="/src"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

# Install nodejs
RUN curl -sL https://deb.nodesource.com/setup_16.x -o /src/nodesource_setup.sh
RUN bash /src/nodesource_setup.sh
RUN apt install nodejs
RUN rm -rf /src/nodesource_setup.sh

# Clone Repo
RUN apt-get install -y git
RUN git clone https://github.com/kshin42/alumni-dashboard.git

# Install Mysql
RUN DEBIAN_FRONTEND=noninteractive apt install -q -y mysql-server
RUN service mysql start && mysqladmin -u root password root



