FROM scratch
WORKDIR /opt/
VOLUME /tmp

COPY shorturl /opt
COPY etc/shorturl-api-docker.yaml /opt/shorturl-api.yaml
ENTRYPOINT ["/opt/shorturl", "-f", "shorturl-api.yaml"]