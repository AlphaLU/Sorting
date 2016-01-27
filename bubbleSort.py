def bubbleSort(l):
	for num in range(len(l)-1,0,-1):
		for i in range(num):
			if l[i] > l[i+1]:
				tempo = l[i]
				l[i] = l[i+1]
				l[i+1] = tempo
		
l = [2, 1 , 5, 9 , 6]
bubbleSort(l)
print(l)
