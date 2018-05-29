#!/bin/bash

curl --header "Content-Type: application/json" \
    --request POST \
    --data '["https://www.amazon.co.uk/gp/product/1509836071", "https://www.amazon.co.uk/gp/product/1509800662"]' \
    http://localhost:8080/upload/urls
