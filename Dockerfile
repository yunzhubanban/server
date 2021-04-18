FROM scratch

LABEL maintainer="云助班班科技创新委员会 <yunzhubanban@qianjunakasumi.ren>" copyright="Copyright (c) 2021-现在，云助班班科技创新委员会"

WORKDIR /

ADD bin/linux_amd64/main /main
RUN chmod +x /main

CMD ./main
