 #!/bin/bash  
commit=$(git rev-parse --verify HEAD --pretty=%H)
  echo "{\n\
    \"commit\": \"$commit\",\n\
    \"author\": \"$(git log -1 $commit --pretty=%an)\",\n\
    \"date\": \"$(git log -1 $commit --pretty=%cd)\",\n\
    \"message\": \"$(git log -1 $commit --pretty=%f)\",\n\
    \"approved-by\": \"$(git log -1 $commit --pretty=%b | grep Approved-by | awk -F ': ' '{print $NF","}' | xargs echo | sed -e 's/,$//')\"\n\
}"
