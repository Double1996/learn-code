select 
FirstName, LastName, City, State
from 
Person
left join (select PersonId as Pid, City,State  from Address) as t_p on t_p.Pid = Person.PersonId