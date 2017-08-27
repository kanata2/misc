#!/bin/bash

if ! type aws > /dev/null 2>&1; then
    echo "Please install aws command."
    exit 127
fi

if ! type jq > /dev/null 2>&1; then
    echo "Please install jq command."
    exit 127
fi

if [ $# -ne 1 ]; then
    echo "Invalid numbers of argument. Please input 1 argument."
    exit 1
fi

readonly REPOSITORY=$1
digests=$(aws ecr describe-images --repository-name "$REPOSITORY" | jq -r '.imageDetails[] | select(.imageTags == null) | "imageDigest=" + .imageDigest')

if [ -z "$digests" ]; then
    echo "There is no images to delete."
    exit 0
fi

if [ -n "$DRYRUN" ]; then
    echo "dry-run:"
    echo "  The following ids will be deleted."
    echo "$digests"
else
    digests=$( echo "$digests" | paste -sd ' ' -)
    aws ecr batch-delete-image --repository-name "$REPOSITORY" --image-ids "$digests"
fi
