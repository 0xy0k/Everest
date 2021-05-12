#!/bin/bash

tsukid query customgov proposals --log_level=debug
tsukid query customgov proposals --reverse=false --log_level=debug
tsukid query customgov proposals --limit=1 --log_level=debug