from sort.quick.quick import exch


# 个人认为这是最简单优雅的逻辑
def partition(a, s, e):
    v = a[e]  # 取最后一个数
    i = s - 1
    for j in range(s, e):  # 不能循环到最后一个数 e，range 正好不包含 e
        #  从前往后找小于 v 的数字
        if a[j] <= v:
            # 如果小于 v 依次放到第一个元素、第二个元素...
            # 先 exch 再 + 1 会有问题，可能会走到数组外面，想想如果一个元素都不满足条件，
            # i 依然指向某个位置，
            # 此时 i 应该在数组外的 s - 1 位置，
            # 应该在确认了有小于 v 的元素再 + 1，确定数组防止元素的位置，再 exch
            # exch(a, i, j)
            # i += 1 # 如果都小于 v ，此时 i 会跑到数组外面
            # ---
            i += 1
            exch(a, i, j)
    # 遍历完成后，i 指向最后一个小于 v 的元素，
    # 把 v 放在 i + 1 的位置，返回这个位置（i + 1）
    exch(a, i + 1, e)
    return i + 1


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
