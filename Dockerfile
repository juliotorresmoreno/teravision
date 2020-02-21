FROM ubuntu:18.04

RUN mkdir /opt/teravision
ADD bin/teravision /opt/teravision
ADD env.yml /opt/teravision
RUN chmod +x /opt/teravision/teravision

WORKDIR /opt/teravision

EXPOSE 1323

ENTRYPOINT [ "/opt/teravision/teravision" ]