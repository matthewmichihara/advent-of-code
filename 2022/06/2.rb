input = File.readlines('input.txt').first

input.split(//).each_index do |index|
  if input[index, 14].split(//).uniq.size == 14
    puts index + 14
    return
  end
end
