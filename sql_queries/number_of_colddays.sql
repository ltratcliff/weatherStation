SELECT
  COUNT(*)
FROM
  (
  SELECT
    date(Dateutc/1000, 'unixepoch', 'localtime') as d, 
    min(Tempf) 
  FROM
    weather 
  GROUP BY d having min(Tempf) < 32
  );
