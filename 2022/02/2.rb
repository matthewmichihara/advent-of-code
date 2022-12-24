def score(opponent, result)
  me = case opponent
       when 'A'
        case result
        when 'X' then 'Z'
        when 'Y' then 'X'
        when 'Z' then 'Y'
        else raise 'invalid state'
        end
       when 'B'
         case result
         when 'X' then 'X'
         when 'Y' then 'Y'
         when 'Z' then 'Z'
         else raise 'invalid state'
         end
       when 'C'
         case result
         when 'X' then 'Y'
         when 'Y' then 'Z'
         when 'Z' then 'X'
         else raise 'invalid state'
         end
       else raise 'invalid state'
       end

  selection_score = case me
                    when 'X' then 1
                    when 'Y' then 2
                    when 'Z' then 3
                    else raise "invalid state"
                    end

  game_score = case opponent
               when 'A'
                 case me
                 when 'X' then 3
                 when 'Y' then 6
                 when 'Z' then 0
                 else raise 'invalid state'
                 end
               when 'B'
                 case me
                 when 'X' then 0
                 when 'Y' then 3
                 when 'Z' then 6
                 else raise 'invalid state'
                 end
               when 'C'
                 case me
                 when 'X' then 6
                 when 'Y' then 0
                 when 'Z' then 3
                 else raise 'invalid state'
                 end
               else raise 'invalid state'
               end

  game_score + selection_score
end

input = File.open('input.txt').read
puts input.lines.map { |line|
  opponent, me = line.split
  score(opponent, me)
}.sum

