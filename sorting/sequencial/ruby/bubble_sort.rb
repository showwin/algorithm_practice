require 'benchmark'

def bubble_sort(ary)
  n = ary.length
  (n - 1).times do |k|
    (n - k - 1).times do |i|
      if ary[i] > ary[i + 1]
        tmp = ary[i]
        ary[i] = ary[i + 1]
        ary[i + 1] = tmp
      end
    end
  end
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

result = Benchmark.realtime do
  bubble_sort(ary1)
  bubble_sort(ary2)
  bubble_sort(ary3)
end
p result / 3
