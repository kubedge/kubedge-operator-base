FROM alpine:3.15

ENV OPERATOR=/usr/local/bin/kubedge-base-operator \
    USER_UID=1001 \
    USER_NAME=kubedge-base-operator

# install operator binary
COPY build/_output/bin/kubedge-base-operator ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

COPY build/test-pods /opt/kubedge-operators/test-pods

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
