require 'benchmark'

def quick_sort(buttom, top, ary)
  return if buttom >= top
  div = ary[buttom]
  lower = buttom
  upper = top
  while lower < upper
    lower += 1 while lower <= upper && ary[lower] <= div
    upper -= 1 while lower <= upper && ary[upper] > div
    next unless lower < upper
    tmp = ary[lower]
    ary[lower] = ary[upper]
    ary[upper] = tmp
  end
  ary[buttom] = ary[upper]
  ary[upper] = div
  quick_sort(buttom, upper - 1, ary)
  quick_sort(upper + 1, top, ary)
  ary
end

size = 10000
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

ary = []
result = Benchmark.realtime do
  ary = quick_sort(0, size - 1, ary1)
  quick_sort(0, size - 1, ary2)
  quick_sort(0, size - 1, ary3)
end
p result / 3

ary.each do |a|
  puts a
end
