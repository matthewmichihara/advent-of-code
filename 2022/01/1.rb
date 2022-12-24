puts "hello ruby"

input = File.open('input.txt').read
elves = input.split(/\n{2,}/)
puts elves.map { |elf|
  elf.split().map {|s| s.to_i}.sum
}.max
