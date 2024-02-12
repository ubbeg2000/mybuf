dirs=$(find gen/go -type d)
dirs=($dirs)

for dir in "${dirs[@]}" 
do
    num_files=$(find "$dir" -maxdepth 1 -type f | wc -l)
    if [[ $num_files -lt 1 ]]; then
        continue 1
    fi
    if [[ $dir == *"mock"* ]]; then
        continue 1
    fi

    files=$(find "$dir" -maxdepth 1 -type f -exec basename {} \;)
    files=($files)
    for file in "${files[@]}" 
    do
        # echo $file
        mockgen -source="$dir/$file" -destination="$dir/mocks/$file"
    done
done