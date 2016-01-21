#
# This dose not work
#
require 'benchmark'
require 'parallel'

MAX_THREAD = 4
SIZE = 10

def bubble_sort(ary)
  tasks = []
  (SIZE - 1).times do |i|
    tasks[i] = SIZE - i - 1
  end
  done = Array.new(SIZE, Array.new(SIZE, 0))

  Parallel.map(tasks, in_threads: 2) do |k|
    k.times do |j|
      true while done[k][j+1] != 1 && k != (SIZE - 1) # wait
      if ary[j] > ary[j + 1]
        tmp = ary[j]
        ary[j] = ary[j + 1]
        ary[j + 1] = tmp
      end
      done[k - 1][j] = 1
    end
  end

  ary
end

f1 = File.open("./../../data/#{SIZE}/seed0.txt", 'r')
ary1 = Array.new(SIZE)
index = 0
f1.each_line do |row|
  ary1[index] = row.chomp.to_i
  index += 1
end
f2 = File.open("./../../data/#{SIZE}/seed1.txt", 'r')
ary2 = Array.new(SIZE)
index = 0
f2.each_line do |row|
  ary2[index] = row.chomp.to_i
  index += 1
end
f3 = File.open("./../../data/#{SIZE}/seed2.txt", 'r')
ary3 = Array.new(SIZE)
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
