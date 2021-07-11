 #!/bin/bash
 DIRECTORY=`dirname "$ABSOLUTE_FILENAME"`
CONTAINERNAME=`cat "$DIRECTORY/container_name.txt"`

docker stop $CONTAINERNAME
