require 'benchmark'

def comb_sort(ary, n)
  gap = (n * 10) / 13
  flg = 1
  while gap > 1 || flg != 1
    gap = (gap * 10) / 13
    gap = 1 if gap == 0
    flg = 1
    (n - gap).times do |i|
      next if ary[i] <= ary[i + gap]
      flg = 0
      tmp = ary[i]
      ary[i] = ary[i + gap]
      ary[i + gap] = tmp
    end
  end
  ary
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
  comb_sort(ary1, size)
  comb_sort(ary2, size)
  comb_sort(ary3, size)
end
p result / 3
