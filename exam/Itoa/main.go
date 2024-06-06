package main

func itoa(n int) string{
	res:=""
	if n ==0 {
	  return "0"
	}
  
	for n >0{
	  res =string(n%10 + '0') +res
	  n/=10
	}
	return res
  }
  