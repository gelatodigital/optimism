#!/bin/bash

set -e

# Path to your .env file
envFile=".env"

# Check if the .env file exists
if [ ! -f "$envFile" ]; then
    echo "Error: .env file at $envFile not found"
    exit 1
fi

# Load the .env file
source "$envFile"

export CONTRACT_ADDRESSES_PATH="deployments/${DEPLOYMENT_CONTEXT}/.deploy"
export DEPLOY_CONFIG_PATH="deploy-config/${DEPLOYMENT_CONTEXT}.json"
export DEPLOYMENT_OUTFILE="$CONTRACT_ADDRESSES_PATH.updated"

forge script scripts/UpdateRoles.s.sol:UpdateRoles --sig 'updateGuardian()' --private-key $ADMIN_PK --rpc-url $RPC_URL --broadcast --legacy

