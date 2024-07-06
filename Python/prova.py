from collections import Counter

el = [1, 1, 1, 2, 3, 4]
hashmap = Counter(el)
hashmap = sorted(hashmap.items(), key=lambda x: x[1], reverse=True)

a = [item[0] for item in hashmap[:3]]
