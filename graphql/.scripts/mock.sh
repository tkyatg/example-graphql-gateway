skip_files=('main.go')

# それっぽいファイルを無視。
files=`find . -type f -name "*.go" ! -name "*_mock.go" ! -path "./tools/*" ! -path "./.*/*" ! -name "*_test.go" $(printf "! -name %s " ${skip_files[@]})`

for file in ${files};
do
  no_ext=`echo ${file} | sed 's/\.[^\.]*$//'`
  mock=${no_ext}_mock.go
  p_name=`head -n 1 ${file} | cut -c 9-`
  mockgen -source ${file} -destination ${mock} -package $p_name && if [ $(grep -c '' ${mock}) = 5 ]; then rm -rf ${mock};fi && echo "generated ${mock}" &
done
wait