input = File.readlines('input.txt').first

input.split(//).each_index do |index|
  if input[index, 4].split(//).uniq.size == 4
    puts index + 4
    return
  end
end
