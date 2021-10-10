#!/bin/bash

tsukid query customgov proposals
tsukid query customgov proposals --reverse
tsukid query customgov proposals --limit=1
tsukid query customgov proposals --limit=1 --output=json --reverse
