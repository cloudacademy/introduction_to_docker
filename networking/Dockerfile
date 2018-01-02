FROM ubuntu:16.04

RUN apt update && apt install -y \
    arp-scan \
    iputils-ping \
    iproute2 \
    curl
    
COPY webapp /

CMD ["/bin/bash"]

# This container is used for the networking demo
# arp-scan --interface=eth0 --localnet