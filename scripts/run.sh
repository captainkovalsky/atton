 #!/bin/bash

#  Put atton with script
#  run with `./run.sh`
#  See logs

  while true; do
    r=$(( $RANDOM % 60 + 60))
    echo -e "fetching data ... \\r\\n"
    ./atton parse
    sleep $r
done