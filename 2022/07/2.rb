input = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k"

class Node
  attr_accessor :name, :type, :children, :size, :parent

  def initialize(name, type, size, parent)
    @name = name
    @type = type
    @children = []
    @size = size
    @parent = parent
  end
end

def print_tree(node, prefix='')
  case node.type
  when :dir
    puts "#{prefix}- #{node.name} (dir) size=#{node.size}"
    node.children.each do |child|
      print_tree(child, prefix + '  ')
    end
  when :file
    puts "#{prefix}- #{node.name} (file, size=#{node.size})"
  else
    raise 'Invalid state'
  end
end

def calculate_sum(node)
  case node.type
  when :dir
    sum = 0
    node.children.each do |child|
      sum += calculate_sum(child)
    end
    node.size = sum
    return sum
  when :file
    return node.size
  else
    raise 'Invalid state'
  end
  raise 'Invalid state'
end

def answer_sum(node)
  sum = 0
  case node.type
  when :dir
    if node.size <= 100_000
      sum += node.size
    end

    node.children.each do |child|
      sum += answer_sum(child)
    end
  when :file
    return 0
  else
    raise 'Invalid state'
  end
  return sum
end

root = Node.new('', :dir, nil, nil)
root.children << Node.new('/', :dir, nil, root)
curr = root

line_index = 0
while line_index < input.lines.size do
  line = input.lines[line_index]
  if line.strip == '$ cd ..'
    curr = curr.parent
    line_index += 1
  elsif m = line.match(/\$ cd (.+)/)
    dir_name = m.captures.first
    curr.children.each do |child|
      if child.name == dir_name
        curr = child
      end
    end
    line_index += 1
  elsif m = line.match(/\$ ls/)
    line_index += 1
    line = input.lines[line_index]
    until line == nil || line.start_with?('$')
      if m = line.match(/dir (.+)/)
        dir_name = m.captures.first
        curr.children << Node.new(dir_name, :dir, nil, curr)
      elsif m = line.match(/(\d+) (.+)/)
        size, file_name = m.captures
        curr.children << Node.new(file_name, :file, size.to_i, curr)
      else
        raise "Invalid state"
      end
      line_index += 1
      line = input.lines[line_index]
    end
  else
    raise "Invalid state"
  end
end

calculate_sum(root)
puts answer_sum(root)
total_used_space = root.size
puts total_used_space
required_space = 70000000 - total_used_space
puts required_space
