#!/bin/zsh

# Parse command-line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        -name) migration_name="$2"; shift ;;
        *) echo "Invalid argument: $1"; exit 1 ;;
    esac
    shift
done

# Check if migration name is provided
if [ -z "$migration_name" ]; then
  echo "Error: Please provide a migration name using -name." >&2
  exit 1
fi

# Create migration with the provided name
migrate create -ext sql -dir library/migration "$migration_name"
