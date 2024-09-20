#!/bin/sh

chmod +x conf-sync

echo "conf-sync starting..."

# Get the directory of the current script
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

# Remove the last directory name to get the parent directory
BASE_DIR="$(dirname "$CURRENT_DIR")"

# App file path
APP="$BASE_DIR/bin/conf-sync"

# App arguments
APP_ARGS="-config $BASE_DIR/conf/application.yaml"

# Command to execute the app
COMMAND="$APP $APP_ARGS"

# Execute the command
eval $COMMAND