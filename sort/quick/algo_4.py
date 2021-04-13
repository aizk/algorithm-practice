from sort.quick.quick import exch


def partition(a, s, e):
    v = a[s]
    i = s
    j = e + 1
    while True:
        while True:
            i = i + 1
            if a[i] >= v:
                break
            # 注意边界
            if i == e:
                break
        while True:
            j = j - 1
            if a[j] < v:
                break
            # 注意边界
            if j == s:
                break
        if i >= j:
            break
        exch(a, i, j)
    exch(a, s, j)
    return j


def sort(a, s, e):
    if s >= e:
        return
    j = partition(a, s, e)
    sort(a, s, j - 1)
    sort(a, j + 1, e)


if __name__ == '__main__':
    a = [2, 4, 3, 1, 5, 3, 8]
    sort(a, 0, len(a) - 1)
    print(a)
