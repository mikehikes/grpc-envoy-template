FROM envoyproxy/envoy:v1.22.0

COPY envoy.yaml /etc/envoy/envoy.yaml
EXPOSE 8099
ENTRYPOINT [ "/usr/local/bin/envoy" ]
CMD [ "-c /etc/envoy/envoy.yaml", "-l trace", "--log-path /tmp/envoy_info.log" ]