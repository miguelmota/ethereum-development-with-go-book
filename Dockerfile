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
USER ubuntu
COPY . /app
WORKDIR /app
RUN wget -nv -O- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash
RUN source ~/.nvm/nvm.sh && nvm install 11 && nvm use 11
RUN source ~/.nvm/nvm.sh && make install && make ebooks

CMD ["echo", "complete"]
