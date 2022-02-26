FROM golang:1.17.7-alpine3.15

ENV PATH="$PATH:/bin/bash" \
  BENTO4_BIN="/opt/bento4/bin" \
  PATH="$PATH:/opt/bento4/bin"

# FFMPEG
RUN apk add --update ffmpeg bash curl make

# Install Bento
WORKDIR /tmp/bento4
ENV BENTO4_BASE_URL="http://zebulon.bok.net/Bento4/source/" \
  BENTO4_VERSION="1-6-0-639" \
  BENTO4_CHECKSUM="d3351ffb425aadc698227ba84f5ec3401cad795a" \
  BENTO4_TARGET="" \
  BENTO4_PATH="/opt/bento4" \
  BENTO4_TYPE="SRC"
# download and unzip bento4
RUN apk add --update --upgrade curl python3 unzip bash gcc g++ scons && \
  curl -O -s ${BENTO4_BASE_URL}/Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
  sha1sum -b Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip | grep -o "^$BENTO4_CHECKSUM " && \
  mkdir -p ${BENTO4_PATH} && \
  unzip Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip -d ${BENTO4_PATH} && \
  rm -rf Bento4-${BENTO4_TYPE}-${BENTO4_VERSION}${BENTO4_TARGET}.zip && \
  apk del unzip

WORKDIR /go/src

#vamos mudar para o endpoint correto. Usando top apenas para segurar o processo rodando
ENTRYPOINT [ "top" ]
