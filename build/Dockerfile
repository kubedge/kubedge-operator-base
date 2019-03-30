FROM alpine:3.8

ENV OPERATOR=/usr/local/bin/baseoperator-operator \
    USER_UID=1001 \
    USER_NAME=baseoperator-operator

# install operator binary
COPY build/_output/bin/baseoperator-operator ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

COPY build/test-pods /opt/kubedge-operators/test-pods

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
