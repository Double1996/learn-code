select 
DIsTINCT L1.Num as Num 
from Logs as  L1
from Logs as  L2
from Logs as  L3
where L1.id = L2.id -1 
and  L2.id = L3.id -1 
and  L1.Num = L2.Num
and  L2.Num = L3.Num 