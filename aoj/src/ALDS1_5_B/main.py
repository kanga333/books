from copy import copy

MAX = 1000000000
count = 0

def printarray(x):
    print(" ".join([str(v) for v in x]))

def merge(a, left, mid, right):
    l = copy(a[left:mid])
    r = copy(a[mid:right])
    l.append(MAX)
    r.append(MAX)

    li, ri = 0, 0
    for i in range(left, right):
        global count
        count += 1
        if l[li] <= r[ri]:
            a[i] = l[li]
            li += 1
        else:
            a[i] = r[ri]
            ri += 1

def merge_sort(a, left, right):
    if left+1 >= right:
        return
    mid = int((left + right) / 2)
    merge_sort(a, left, mid)
    merge_sort(a, mid, right)
    merge(a, left, mid, right)
    return

# ALDS1_5_B: マージソート
def main():
    n = int(input())
    a = [int(v) for v in input().split(" ")]
    merge_sort(a, 0, n)
    printarray(a)
    global count
    print(count)

if __name__ == '__main__':
    main()
