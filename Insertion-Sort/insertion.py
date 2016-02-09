def insertionSort(arraylist):
	i = 0
	j = 0
	temp = 0
	length = len(arraylist)
	for  i in range(1, length):
		j = i
		while (j > 0 and arraylist[j] < arraylist[j - 1]):
			temp = arraylist[j]
			arraylist[j] = arraylist[j-1]
			arraylist[j-1] = temp
			j -= 1
		i+=1
	
arraylist = [5, 3, 2, 9, 7, 8, 1, 54,26,93,17,77,31,44,55,20, 15, 631,1424,151252,6246,24,624,624,624,426,247,245758,469,35]
insertionSort(arraylist)
print arraylist
