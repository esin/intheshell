FROM golang:1.13-buster as builder

WORKDIR /src

COPY intheshell.go /src/

RUN go build intheshell.go

FROM ubuntu:18.04

RUN apt-get update && apt-get install -y openssh-server && \
    mkdir /var/run/sshd && chmod -x /etc/update-motd.d/* && \
    useradd -m -s /usr/local/bin/intheshell ghost && \
    /etc/init.d/ssh stop && \
    sed -ri 's/ghost:(!)?:/ghost:U6aMy0wojraho:/' /etc/shadow && \
    sed -ri 's/Port 22/Port 22222/' /etc/ssh/sshd_config &&\
    sed -ri 's/#PasswordAuthentication yes/PasswordAuthentication yes/' /etc/ssh/sshd_config &&\
    sed -ri 's/PermitEmptyPasswords no/PermitEmptyPasswords yes/' /etc/ssh/sshd_config &&\
    sed -ri 's@Subsystem sftp /usr/lib/openssh/sftp-server@@' /etc/ssh/sshd_config && \
    sed -ri 's/X11Forwarding no/X11Forwarding yes/' /etc/ssh/sshd_config && \
    echo "AllowUsers ghost" >> /etc/ssh/sshd_config && \
    echo "AllowTcpForwarding no" >> /etc/ssh/sshd_config && \
    sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

EXPOSE 22222

COPY --from=builder /src/intheshell /usr/local/bin/

CMD ["/usr/sbin/sshd", "-D"]
