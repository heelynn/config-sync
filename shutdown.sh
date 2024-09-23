#!/bin/bash

# Get a list of processes containing "config-sync" in their command line
pids=$(pgrep -f "config-sync")

# Check if any matching processes were found
if [ -z "$pids" ]; then
    echo "No matching processes found for 'config-sync'."
else
    # Iterate over each PID
    for pid in $pids; do
        # Send the QUIT signal to each process
        kill -QUIT "$pid"
        echo "Sent QUIT signal to PID $pid."
    done
fi