#!/usr/bin/env bash

# https://platform.openai.com/docs/guides/images/generations
curl https://api.openai.com/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $GPT_APIKEY" \
  -d '{
    "prompt": "orange tabby cat on windowsill",
    "model": "dall-e-2",
    "n": 1,
    "size": "1024x1024"
  }'