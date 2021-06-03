#!/bin/bash

# exporting the status of chain
# 1. make sure daemon is turned off as you can't export status while daemon is actively writing on the database
# 2. run the command to export the status

tsukid export --home="$HOME/.tsukid"

# Note: should not forget to set home flag.