# Write your MySQL query statement below
SELECT a.Name AS Employee
from Employee a left join Employee b
ON a.ManagerId=b.Id
WHERE a.Salary > b.Salary AND a.ManagerId IS NOT NULL