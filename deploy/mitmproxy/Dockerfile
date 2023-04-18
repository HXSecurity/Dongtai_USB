FROM quay.io/centos/centos:stream9
WORKDIR /opt
COPY . .
RUN yum -y install *epel* &&  yum -y install libpcap-devel \
  && ls -al /usr/lib64/libpcap.so.* \
  && ALIMARCH=`arch` && curl -L https://dongtai-helm-charts.oss-ap-northeast-1.aliyuncs.com/apk/mitmweb -o /opt/mitmweb \
  && ln -s /usr/lib64/libpcap.so.1.10.0 /usr/lib64/libpcap.so.0.8 && chmod +x /opt/mitmweb
ENTRYPOINT ./mitmweb --set block_global=${BLOCK_GLOBAL}  --listen-host 0.0.0.0 --listen-port 10802 --web-host 0.0.0.0 --web-port 8000 -k -s headerannotion.py --mode upstream:http://${XRAY_URL}:9811