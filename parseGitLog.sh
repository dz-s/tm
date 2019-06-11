for commit in $(git rev-parse --short HEAD --pretty=%H);do
  echo -e "{\n\
    \"commit\": \"$commit\",\n\
    \"author\": \"$(git log -1 $commit --pretty=%an)\",\n\
    \"date\": \"$(git log -1 $commit --pretty=%cd)\",\n\
    \"message\": \"$(git log -1 $commit --pretty=%f)\",\n\
    \"approved-by\": \"$(git log -1 $commit --pretty=%b | grep Approved-by | awk -F ': ' '{print $NF","}' | xargs echo | sed -e 's/,$//')\"\n\
},"
done | \
perl -pe 'BEGIN{print "["}' | \
sed -e '$s/},/}]/'