#!/opt/rh/rh-python38/root/bin/python3

import datetime
import sqlite3
import pytz

today = datetime.date.today()
tomorrow = today + datetime.timedelta(days=1)
est = pytz.timezone('US/Eastern')
#fmt = '%Y-%m-%d %H:%M:%S %Z%z'
fmt = '%H:%M:%S'

s = f'''
SELECT
    Date,
    Tempf,
    FeelsLike,
    Dailyrainin
FROM
    WEATHER
WHERE
    DATE BETWEEN '{today.isoformat()} 04:00:00+00:00' AND '{tomorrow.isoformat()} 04:00:00+00:00'
ORDER BY
    Date;
'''


conn = sqlite3.connect('weather.db')
c = conn.cursor()

c.execute(s)
rows = c.fetchall()
print(f"{'Time': <6}  {'Temp': <5} {'Feels': <6} Rain")
for row in rows:
    dd = datetime.datetime.fromisoformat(row[0])
    print(dd.astimezone(est).strftime(fmt), row[1], f'{row[2]: <6}', row[3])
c.close()
