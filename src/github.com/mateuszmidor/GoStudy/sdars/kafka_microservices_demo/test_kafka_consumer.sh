#!/bin/bash

kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0
