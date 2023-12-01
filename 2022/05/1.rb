input = File.readlines('input.txt')

num_stacks = input.first.length / 4
stacks = []
num_stacks.times do
  stacks.push []
end

puts stacks
moves_phase = false

input.each do |line|
  if line.strip.empty?
    moves_phase = true
  end

  if moves_phase
    if match = line.match(/move (\d+) from (\d+) to (\d+)/)
      num_moves, from_stack, to_stack = match.captures
      num_moves = num_moves.to_i
      from_stack = from_stack.to_i
      to_stack = to_stack.to_i

      num_moves.times do
        stacks[to_stack-1].push(stacks[from_stack-1].pop)
      end
    end
    next
  end

  (0...num_stacks).each do |i|
    chunk = line[i*4, 4]
    if match = chunk.match(/(\[(.)\])/)
      _, crate = match.captures
      stacks[i].prepend(crate)
    end
  end
end

message = ""
stacks.each do |stack|
  message += stack.last 
end

puts message


