#!/usr/bin/env bash

# ******************************************************
# DESC    : run echo server
# AUTHOR  : wongoo
# VERSION : 1.0
# LICENCE : LGPL V3
# EMAIL   : gelnyang@163.com
# MOD     : 2019-04-29
# ******************************************************

export APP_CONF_FILE=`pwd`/config.toml
export APP_LOG_CONF_FILE=`pwd`/log.xml

go run *.go
