#!/bin/bash

diff -U3 --speed-large-files --ignore-space-change --ignore-blank-lines old.sql new.sql
