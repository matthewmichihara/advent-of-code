require 'set'

all_shared = []
priorities = (('a'..'z').zip(1..) + ('A'..'Z').zip(27..)).to_h

File.readlines('input.txt').each_slice(3).each do |group|
  shared = group.first.strip.split(//).to_set
  group.each do |rucksack|
    shared = shared & rucksack.strip.split(//).to_set
  end
  all_shared.push(*shared)
end

puts all_shared.map { |item| priorities[item] }.sum
