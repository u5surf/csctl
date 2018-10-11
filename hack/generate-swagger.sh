#!/bin/sh

set -e

# Cloud API
swagger generate model -f cloud/api/swagger.yaml \
    -t cloud/api \
    -m types

# Provision API
swagger generate model -f cloud/provision/swagger.yaml \
    -t cloud/provision \
    -m types
