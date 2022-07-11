#!/bin/bash

cd /root/ironfish/monitor/
nohup ./collect > collect.log 2>&1 &
