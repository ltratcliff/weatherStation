SELECT
  COUNT(*)
FROM
  (
  SELECT
    date(Dateutc/1000, 'unixepoch', 'localtime') as d
  FROM
    weather 
  GROUP BY d 
  );
