CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  set n = n -1;
  RETURN (
      # Write your MySQL query statement below.
     SELECT  -- 第N 高的薪水
            salary
      FROM 
            employee
      GROUP BY 
            salary
      ORDER BY 
            salary DESC
      LIMIT N, 1 -- 
  );
END