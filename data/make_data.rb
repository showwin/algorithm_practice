size = 100_000
max = 1_000_000
3.times do |i|
  f = File.open("./#{size}/seed#{i}.txt", 'w')
  size.times do |k|
    f.write("#{rand(max)}\n")
  end
end
