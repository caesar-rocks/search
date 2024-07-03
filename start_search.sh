#!/bin/bash

# exposes default port 7700 to local host
docker  run --rm -it -p 7700:7700 --name search -d getmeili/meilisearch:v1.9