#!/bin/sh


rsync -avr . commander.devguard.io: --exclude .env
