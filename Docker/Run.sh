 #!/bin/bash
 DIRECTORY=`dirname "$ABSOLUTE_FILENAME"`
CONTAINERNAME=`cat "$DIRECTORY/container_name.txt"`

echo $DIRECTORY

echo $CONTAINERNAME

docker run --name $CONTAINERNAME -e POSTGRES_PASSWORD=test -d -p 5432:5432 postgres

sleep 10
