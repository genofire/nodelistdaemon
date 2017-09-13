#!/bin/bash
host=$1
port=$2
remote="circleci@${host}"
echo "deploying..."
ssh -p $port $remote sudo systemctl stop nodelistdaemon;
RETVAL=$?
[ $RETVAL -ne 0 ] && exit 1
scp -q -P $port ~/.go_workspace/bin/nodelistdaemon $remote:~/bin/nodelistdaemon;
RETVAL=$?
ssh -p $port $remote sudo systemctl start nodelistdaemon;
[ $RETVAL -eq 0 ] && RETVAL=$?
[ $RETVAL -ne 0 ] && exit 1
echo "deployed"
