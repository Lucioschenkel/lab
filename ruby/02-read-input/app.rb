if __FILE__ == $0
  print "Enter your name: "
  # .chomp strips the newline character
  name = gets.chomp
  puts "Hello, #{name}"
end
