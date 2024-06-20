#!/bin/sh
services="user video interaction social api"
for service in $services
do
  make SERVICE=${service}
done
