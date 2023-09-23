10 =a
20 =b
$a =c
1 wait

lp:
	1000 =SAFE_POINT
	$a 1 + =a
	1 4 / wait
	$a 100 > end!
lp!

end:
10 wait