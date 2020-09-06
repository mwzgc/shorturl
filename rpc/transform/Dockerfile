FROM scratch
WORKDIR /opt/
VOLUME /tmp

COPY transform /opt
COPY etc/transform-docker.yaml /opt/transform.yaml
ENTRYPOINT ["/opt/transform", "-f", "./transform.yaml"]