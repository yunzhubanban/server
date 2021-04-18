FROM alpine:latest

LABEL maintainer="云助班班科技创新委员会 <yunzhubanban@qianjunakasumi.ren>" copyright="Copyright (c) 2021-现在，云助班班科技创新委员会"

WORKDIR /

COPY bin/linux_amd64/main /main

VOLUME [ "/data", "/config", "/.yunzhubanban" ]

CMD /main
