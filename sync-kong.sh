#!/bin/bash

docker run -i --rm --network=cu-get-ref-backend_default -v ${PWD}:/deck kong/deck --kong-addr=http://kong:8001/ -s /deck/kong.yaml sync
