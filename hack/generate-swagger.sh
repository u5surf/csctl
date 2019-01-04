#!/bin/sh

set -e

# See https://goswagger.io/install.html for more
# information on installing the swagger binary

# Cloud API
swagger generate model -f cloud/api/swagger.yaml \
    -t cloud/api \
    -m types

# Auth API
swagger generate model -f cloud/auth/swagger.yaml \
    -t cloud/auth \
    -m types

# Provision API
swagger generate model -f cloud/provision/swagger.yaml \
    -t cloud/provision \
    -m types
