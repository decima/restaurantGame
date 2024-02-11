
doc.generate:
	docker run --rm \
        -v $PWD:/local openapitools/openapi-generator-cli generate \
        -i /local/documentation/api.yaml \
        -g cwiki \
        -o /local/var/cwiki