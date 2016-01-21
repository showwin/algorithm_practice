require 'benchmark'

def merge_sort(bottom, top, ary)
  return ary[bottom..top] if bottom >= top
  mid = (bottom + top) / 2
  ary1 = merge_sort(bottom, mid, ary)
  ary2 = merge_sort(mid + 1, top, ary)

  i1 = 0
  l1 = ary1.length
  i2 = 0
  l2 = ary2.length
  ary_new = Array.new(l1 + l2)
  (l1 + l2).times do
    if i1 < l1 && (i2 >= l2 || ary1[i1] <= ary2[i2])
      ary_new[i1 + i2] = ary1[i1]
      i1 += 1
    else
      ary_new[i1 + i2] = ary2[i2]
      i2 += 1
    end
  end

  ary_new
end

size = 1_000_000
f1 = File.open("./../../data/#{size}/seed0.txt", 'r')
ary1 = Array.new(size)
index = 0
f1.each_line do |row|
  ary1[index] = row.chomp.to_i
  index += 1
end
f2 = File.open("./../../data/#{size}/seed1.txt", 'r')
ary2 = Array.new(size)
index = 0
f2.each_line do |row|
  ary2[index] = row.chomp.to_i
  index += 1
end
f3 = File.open("./../../data/#{size}/seed2.txt", 'r')
ary3 = Array.new(size)
index = 0
f3.each_line do |row|
  ary3[index] = row.chomp.to_i
  index += 1
end

result = Benchmark.realtime do
  merge_sort(0, size - 1, ary1)
  merge_sort(0, size - 1, ary2)
  merge_sort(0, size - 1, ary3)
end
p result / 3
