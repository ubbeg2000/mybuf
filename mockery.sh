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

    mockery --dir "$dir" --all --output "$dir/mocks"
done