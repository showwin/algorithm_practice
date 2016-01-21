require 'benchmark'

SEED_SIZE = 1000000

def linear_search(ary, target)
  index = 0
  ary.each_with_index do |e, i|
    next unless e == target
    index = i
    break
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
    linear_search(seed[i], rand(SEED_SIZE))
  end
end
p "Time: #{result / 10}"
