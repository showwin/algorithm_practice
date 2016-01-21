require 'benchmark'

SEED_SIZE = 1000000

def binary_search(ary, target)
  left = 0
  right = ary.length - 1
  mid = ary.length / 2
  while left < right
    mid = (left + right) / 2
    if ary[mid] == target
      break
    elsif ary[mid] > target
      right = mid
    else
      left = mid + 1
    end
  end
end

def make_seed(size)
  hash = {}
  10.times do |i|
    ary = Array.new(size)
    size.times do |j|
      ary[j] = rand(size)
    end
    hash[i] = ary
  end
  hash
end

seed = make_seed(SEED_SIZE)

result = Benchmark.realtime do
  10.times do |i|
    binary_search(seed[i].sort, rand(SEED_SIZE))
  end
end
p "Time: #{result / 10}"
