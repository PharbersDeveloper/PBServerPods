#!/bin/bash

# Remove dupulicate package
rm -rf node_modules/typegoose/node_modules/\@types

# first build
tsc
