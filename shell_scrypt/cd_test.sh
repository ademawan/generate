
# !/bin/bash
 
# A simple bash script to move up to desired directory level directly
 

test=$1
echo "$test"
cd "$test"
cat "./file1"


weather=$(curl -s http://103.13.207.248/v1/user/1)
echo "test $weather"
