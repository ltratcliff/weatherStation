# Ambient Weather Station API Consumer

Ambient Weather API query and store in local sqlite3 db

## Requirements

.env file with MAC, APPKEY and APIKEY vars

```sh
cat > .env <<EOF                                                                                                                            14:15:44
MAC=00:00:00:00:00
APIKEY=fillmein
APPKEY=fillmein
EOF
```

Sqlite3 database with schema applied

```sh
./initdb.sh
```

:tada:
