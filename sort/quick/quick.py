import random

def wrap(partition):
    def qs(a, s, e):
        if s < e:
            p = partition(a, s, e)
            qs(a, s, p - 1)
            qs(a, p + 1, e)
    return qs

def partition(a, s, e):
    t = a[e] # 取最后一个数
    i = s - 1
    for j in range(s, e): # 不能循环到 e，range 正好不包含 e
        if a[j] <= t: # 先 exch 再 + 1 会有问题，可能会走到数组外面，应该确认了 + 1 再 exch
            i += 1
            exch(a, i, j)
    exch(a, i + 1, e)
    return i + 1

def partition1(a, s, e):
    t = a[s] # 第一个数
    while s < e:
        # 先从后往前找个小于 t 的把第一个位置占住，把自己的位置空出来
        while s < e:
            if a[e] <= t:
                a[s] = a[e] # 此时 a[s] 所指的是一个空位
                break
            else:
                e -= 1
        while s < e:
            if a[s] > t:
                a[e] = a[s] # 此时 a[e] 所指的是一个空位，因为之前已经赋值给 a[s]
                break
            else:
                s += 1
    # i == j 时指针相等，因此必然会指向一个空位
    a[s] = t
    return s

# 简写
def partition1_1(a, i, j):
    t = a[i] # 第一个数
    while i < j:
        # find element < t from right to left
        while i < j:
            if a[j] <= t:
                a[i] = a[j] # 此时 a[i] 所指的是一个空位
                break
            j -= 1
        while i < j:
            if a[i] > t:
                a[j] = a[i] # 此时 a[j] 所指的是一个空位，因为之前已经赋值给 a[i]
                break
            i += 1
    # i == j 时指针相等，因此必然会指向一个空位，把 t 放进去
    a[i] = t
    return i

# 这个逻辑很难写啊
'''
重点是 while True 里面的逻辑，当 >= 或 <= 的时候就返回，不能自增或自减
因为必须有一个指针一直指着分隔元素，然后交替从左边或右边找元素换到正确的方向
注意点就是：
    - <=、>= 操作符
    - 不能乱递增 i ++、j--
'''
def partition2(a, s, e):
    t = a[s]
    i = s
    j = e
    while True:
        while i < j: # 必须拿 j 的一个拷贝来找，才能不丢失选择元素的位置
            if a[j] <= t: 
                break
            else:
                j -= 1
        while i < j:
            if a[i] >= t: 
                break
            else:
                i += 1
        if i < j:
            exch(a, i, j)
        else:
            return j # 返回 i、j 哪个都行

def exch(a, i, j):
    if i == j: return
    t = a[i]
    a[i] = a[j]
    a[j] = t

if __name__ == "__main__":
    a = [5, 4, 3, 2, 1]
    wrap(partition)(a, 0, len(a) - 1)
    print(a)

    random.shuffle(a)
    wrap(partition1)(a, 0, len(a) - 1)
    print(a)

    random.shuffle(a)
    wrap(partition1_1)(a, 0, len(a) - 1)
    print(a)

    random.shuffle(a)
    # print(a)
    # print('---')
    # a = [3, 1, 2, 4, 5]
    # print(a)
    wrap(partition2)(a, 0, len(a) - 1)
    print(a)