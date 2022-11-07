SELECT 
  date('now'), 
  MAX(Tempf), 
  MIN(Tempf), 
  AVG(Tempf) 
FROM 
  weather 
WHERE
  date(Dateutc/1000, 'unixepoch', 'localtime') = date('now');

SELECT 
  date('now', '-1 year'), 
  MAX(Tempf), 
  MIN(Tempf), 
  AVG(Tempf) 
FROM 
  weather 
WHERE
  date(Dateutc/1000, 'unixepoch', 'localtime') = date('now', '-1 year');
