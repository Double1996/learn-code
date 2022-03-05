SQL89 获得积分最多的人(一)https://www.nowcoder.com/practice/1bfe3870034e4efeb4b4aa6711316c3b?tpId=82&tqId=38359&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3DSQL%25E7%25AF%2587%26topicId%3D82&difficulty=undefined&judgeStatus=undefined&tags=&title=



SELECT 
    u.name,
    SUM(grade_num) AS grade_sum
FROM user u
JOIN grade_info g
ON u.id =g.user_id
GROUP BY u.id
order by grade_sum desc
limit 0,1;