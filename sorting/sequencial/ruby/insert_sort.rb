require 'benchmark'

def insert_sort(ary, n)
  n.times do |j|
    insert = ary[j]
    i = 0
    while i < j && ary[i] < insert
      i += 1
    end

    # insert
    while i <= j
      tmp = ary[i]
      ary[i] = insert
      insert = tmp
      i += 1
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
  insert_sort(ary1, size)
  insert_sort(ary2, size)
  insert_sort(ary3, size)
end
p result / 3
