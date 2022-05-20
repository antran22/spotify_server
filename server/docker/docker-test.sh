#!/usr/bin/env sh
curl --silent --fail http://spotify-server:3000/ping | grep "pong" > /dev/null