#! /bin/bash


#man 7 singal

#NO TRAP SIGKILL AND SIGSTOP
# trap "echo Recive KILL signal" SIGKILL
trap "echo Recive INTERRUPT signal" SIGINT

:w
echo "pid is $$"

while (( COUNT < 10 )); do
    sleep 10
    (( COUNT ++ ))
    echo $COUNT
done
exit 0
