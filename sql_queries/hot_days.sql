SELECT
  date(Dateutc/1000, 'unixepoch', 'localtime') as d, 
  max(Tempf) 
FROM
   weather 
GROUP BY d having max(Tempf) > 99;
