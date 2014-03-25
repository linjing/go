package qsort


func quickSort(values [] int, left int, right int) {
  temp := values[left]
  p := left
  i, j := left, right

  for i <= j {
    for j >= p && values[j] >= temp {
      j = j - 1
    }

    if j >= p && values[j] < temp {
      values[p] = values[j]
      p = j
    }

    for i <= p && values[i] <= temp {
      i = i + 1 
    }
    if i <= p && values[i] > temp {
      values[p] = values[i]
      p = i
    }
  }

  values[p] = temp
  if p - left > 1 {
    quickSort(values, left, p - 1)
  }
  if right - p > 1 {
    quickSort(values, p + 1, right)
  }
}

func QSort(values [] int) {
  quickSort(values, 0, len(values) - 1)
}
