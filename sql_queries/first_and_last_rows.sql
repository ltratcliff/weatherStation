SELECT
  MIN(date(Dateutc/1000, 'unixepoch', 'localtime')) as mind,
  MAX(date(Dateutc/1000, 'unixepoch', 'localtime')) as maxd 
FROM weather;
