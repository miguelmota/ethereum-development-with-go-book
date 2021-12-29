FROM ubuntu:18.04

RUN apt-get update -y
RUN apt-get install -y \
  build-essential \
  curl \
  wget \
  python \
  pkg-config \
  libcairo2-dev \
  libjpeg-dev \
  libgif-dev \
  libglib2.0-0 \
  libgl-dev \
  libnss3-dev \
  libx11-6 \
  libx11-xcb1 \
  libxcb1 \
  libxcursor1 \
  libxdamage1 \
  libxext6 \
  libxfixes3 \
  libxi6 \
  libxrandr2 \
  libxrender1 \
  libxss1 \
  libxtst6 \
  libxcb1-dev \
  libxcomposite-dev

RUN wget -nv -O- https://download.calibre-ebook.com/linux-installer.sh | bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 ubuntu

RUN mkdir -p /usr/local/nvm
ENV NVM_DIR /usr/local/nvm
ENV NODE_VERSION 11.15.0

RUN curl https://raw.githubusercontent.com/creationix/nvm/v0.39.1/install.sh | bash \
    && source $NVM_DIR/nvm.sh \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use default

ENV NODE_PATH $NVM_DIR/v$NODE_VERSION/lib/node_modules
ENV PATH $NVM_DIR/versions/node/v$NODE_VERSION/bin:$PATH

COPY Makefile .
RUN make install-gitbook

USER ubuntu
RUN mkdir -p /home/ubuntu/app
COPY . /home/ubuntu/app
WORKDIR /home/ubuntu/app

RUN make install-modules
RUN make build
RUN make ebooks

CMD ["echo", "complete"]
