#!/bin/sh

set -e

# Cloud API
swagger generate model -f pkg/cloud/api/cloud-api-swagger.yaml \
    -t pkg/cloud/api \
    -m types

# Provision API
# TODO
#swagger generate model -f pkg/cloud/api/cloud-provision-swagger.yaml \
    #-t pkg/cloud/api/provision \
    #-m types
