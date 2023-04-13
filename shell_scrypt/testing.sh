#!/bin/sh

# Define your function here
name=$1
secondname=$2
Hello () {
   echo "Hello $1 World $1 $2"
}

# Invoke your function
# Hello $name $secondname
#for (( c=1; c<=5; c++ ));
#do 
#   echo "Welcome $c times"
#done

i=80
while [ "$i" -le 101 ]; do
    echo "halo $i"
    i=$(( i + 1 ))
done 
