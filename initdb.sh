#!/bin/bash
sqlite3 weather.db
sqlite3 weather.db < db.schema
