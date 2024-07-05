FROM docker.io/golang:alpine AS builder
ARG GARM_REF

LABEL stage=builder

RUN apk add musl-dev gcc libtool m4 autoconf g++ make libblkid util-linux-dev git linux-headers upx
RUN git config --global --add safe.directory /build

ADD . /build/garm
RUN cd /build/garm && git checkout ${GARM_REF}
RUN git clone https://github.com/cloudbase/garm-provider-azure /build/garm-provider-azure
RUN git clone https://github.com/cloudbase/garm-provider-openstack /build/garm-provider-openstack
RUN git clone https://github.com/cloudbase/garm-provider-lxd /build/garm-provider-lxd
RUN git clone https://github.com/cloudbase/garm-provider-incus /build/garm-provider-incus
RUN git clone https://github.com/mercedes-benz/garm-provider-k8s /build/garm-provider-k8s
RUN git clone https://github.com/cloudbase/garm-provider-aws /build/garm-provider-aws
RUN git clone https://github.com/cloudbase/garm-provider-gcp /build/garm-provider-gcp
RUN git clone https://github.com/cloudbase/garm-provider-equinix /build/garm-provider-equinix

RUN cd /build/garm && go build -o /bin/garm \
    -tags osusergo,netgo,sqlite_omit_load_extension \
    -ldflags "-linkmode external -extldflags '-static' -s -w -X github.com/cloudbase/garm/util/appdefaults.Version=$(git describe --tags --match='v[0-9]*' --dirty --always)" \
    /build/garm/cmd/garm && upx /bin/garm
RUN mkdir -p /opt/garm/providers.d
RUN cd /build/garm-provider-azure && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-azure . && upx /opt/garm/providers.d/garm-provider-azure
RUN cd /build/garm-provider-openstack && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-openstack . && upx /opt/garm/providers.d/garm-provider-openstack
RUN cd /build/garm-provider-lxd && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-lxd . && upx /opt/garm/providers.d/garm-provider-lxd
RUN cd /build/garm-provider-incus && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-incus . && upx /opt/garm/providers.d/garm-provider-incus
RUN cd /build/garm-provider-k8s/cmd/garm-provider-k8s && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-k8s . && upx /opt/garm/providers.d/garm-provider-k8s
RUN cd /build/garm-provider-aws && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-aws . && upx /opt/garm/providers.d/garm-provider-aws
RUN cd /build/garm-provider-gcp && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-gcp . && upx /opt/garm/providers.d/garm-provider-gcp
RUN cd /build/garm-provider-equinix && go build -ldflags="-linkmode external -extldflags '-static' -s -w" -o /opt/garm/providers.d/garm-provider-equinix . && upx /opt/garm/providers.d/garm-provider-equinix

FROM scratch

COPY --from=builder /bin/garm /bin/garm
COPY --from=builder /opt/garm/providers.d/garm-provider-openstack /opt/garm/providers.d/garm-provider-openstack
COPY --from=builder /opt/garm/providers.d/garm-provider-lxd /opt/garm/providers.d/garm-provider-lxd
COPY --from=builder /opt/garm/providers.d/garm-provider-incus /opt/garm/providers.d/garm-provider-incus
COPY --from=builder /opt/garm/providers.d/garm-provider-k8s /opt/garm/providers.d/garm-provider-k8s
COPY --from=builder /opt/garm/providers.d/garm-provider-azure /opt/garm/providers.d/garm-provider-azure
COPY --from=builder /opt/garm/providers.d/garm-provider-aws /opt/garm/providers.d/garm-provider-aws
COPY --from=builder /opt/garm/providers.d/garm-provider-gcp /opt/garm/providers.d/garm-provider-gcp
COPY --from=builder /opt/garm/providers.d/garm-provider-equinix /opt/garm/providers.d/garm-provider-equinix
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/bin/garm", "-config", "/etc/garm/config.toml"]
