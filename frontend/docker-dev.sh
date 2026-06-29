#!/bin/sh

pnpm install --config.minimumReleaseAge=0 --config.confirmModulesPurge=false

exec pnpm dev --host