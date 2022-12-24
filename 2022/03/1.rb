require 'set'

all_shared = []
priorities = (('a'..'z').zip(1..) + ('A'..'Z').zip(27..)).to_h

File.foreach('input.txt') do |line|
  first = line[0...line.length/2].split(//).to_set
  second = line[line.length/2..].split(//).to_set

  shared = first & second
  all_shared.push(*shared)
end

puts all_shared.map { |item| priorities[item] }.sum
