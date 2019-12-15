package codegen

//go:generate docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g go -o /local/client -c /local/config.json
