def binarySearch(a, target):
    s = 0
    e = len(a) - 1
    while s <= e:
        mid = int(s + (e - s) / 2)
        if a[mid] == target:
            return mid
        elif a[mid] < target:
            s = mid + 1
        else:
            e = mid - 1
    return -1


if __name__ == '__main__':
    a = [1, 3, 5, 7, 9, 11, 13, 15]
    print(binarySearch(a, 5))
    print(binarySearch(a, 15))

