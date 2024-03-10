# /bin/bash


curl -s $url_json | jq .accounts[] | jq 'select($name_filter == "$value_filter")'
