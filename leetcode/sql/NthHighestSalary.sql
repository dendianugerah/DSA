CREATE FUNCTION getNthHighestSalary(N INT) returns INT
BEGIN
    SET N = N - 1;
    RETURN(
        SELECT DISTINCT (Employee.salary) FROM Employee ORDER BY Employee.salary DESC 
        LIMIT 1 OFFSET N
    );
END