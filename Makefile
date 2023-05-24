.PHONY: start-envoy
start-envoy:
	cat envoy/config/envoy-override.yaml
	docker run -p 9902:9902 -p 10000:10000 -v `pwd`/envoy/config:/config  --rm envoyproxy/envoy:v1.26.1 -c /config/xds.yaml