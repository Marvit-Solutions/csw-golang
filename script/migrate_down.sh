#!/bin/zsh

# Parse command-line arguments
while getopts ":u:p:d:" opt; do
        case $opt in
                u) user="$OPTARG"
                ;;
                p) password="$OPTARG"
                ;;
                d) database="$OPTARG"
                ;;
                \?) echo "Invalid option: -$OPTARG" >&2
        exit 1
        ;;
esac
done

# Check if user, password, and database values are provided
if [ -z "$user" ] || [ -z "$password" ] || [ -z "$database" ]; then
        echo "Error: Please provide values for user (-u), password (-p), and database (-d)." >&2
        exit 1
fi

# Perform migration with the provided user, password, and database
migrate -source file://library/migration -database "postgres://$user:$password@localhost:5432/$database?sslmode=disable" down
