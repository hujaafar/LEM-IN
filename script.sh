for file in testFiles/*.txt; do
    echo "Running main.go with $file"
    go run main.go "$file"
done