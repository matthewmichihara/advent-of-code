require 'set'

def covers(a1, a2, b1, b2)
  return (a1 <= b1 && a2 >= b2) || (b1 <= a1 && b2 >= a2)
end

puts File.readlines('input.txt').filter { |line|
  m = line.match /(\d+)-(\d+),(\d+)-(\d+)/
  covers(m[1].to_i, m[2].to_i, m[3].to_i, m[4].to_i)
}.count
