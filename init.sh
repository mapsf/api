#!/usr/bin/env bash

if [ ! -f .env ]
then
    echo "Creating .env file..."
    cp .env.dist .env
fi

if [ ! -f api/.env ]
then
    echo "Creating api/.env file..."
    cp api/.env.dist api/.env
fi
