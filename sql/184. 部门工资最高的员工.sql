select  
 Department.name AS 'Department',
 Employee.name  AS 'Employee',
 Employee.salary
from Employee 
left join  Department on Department.id = Employee.departmentId
where 
    (Employee.salary,  Employee.departmentId) in (
      select       
         max(Employee.salary), Employee.departmentId
      from  Employee
      group by Employee.departmentId
    )


    