def mergeSort(l):
	print("[+] Splitting up", l)
	if len(l) > 1:
		middle = len(l)/2
		righthalf = l[middle:]
		lefthalf = l[:middle]
	
		mergeSort(right)
		mergeSort(left)
	
		i = 0
		j = 0
		k = 0
		
		while i < len(left) and j < len(right):
			if left[i] < right[j]:
				l[k] = left[i]
				i = i + 1
			else:
				l[k] = right[j]
				j= j + 1
			k = k + 1
		
		while i < len(left):
			l[k] = left[i]
			i = i + 1
			k = k + 1
		
		while j < len(right):
			l[k] = right[j]
			j = j + 1
			k = k + 1
			
		print ("[*] Merging : ", l)
		
l = [85, 1, 69, 98, 4555, 32, 984 , 25252 ,65151 ,3, 494, 4, 9, 151 ,31 ,65415 ,4894 ,89489 ,1231 ,454 ,4984 ,4984 ,3231 ,959, 18484 ,269 ,61591, -1]
mergeSort(l)
print(l)
