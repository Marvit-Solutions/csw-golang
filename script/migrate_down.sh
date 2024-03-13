#!/bin/zsh

# Parse command-line arguments
while getopts ":p:d:" opt; do
case $opt in
p) password="$OPTARG"
;;
d) database="$OPTARG"
;;
\?) echo "Invalid option: -$OPTARG" >&2
exit 1
;;
esac
done

# Check if both password and database values are provided
if [ -z "$password" ] || [ -z "$database" ]; then
echo "Error: Please provide values for password (-p) and database (-d)." >&2
exit 1
fi

# Perform migration with the provided password and database
migrate -source file://library/migration -database "postgres://postgres:$password@localhost:5432/$database?sslmode=disable" down
