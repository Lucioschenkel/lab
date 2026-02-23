if __FILE__ == $0
  # Using a block makes sure the file is automatically closed
  open('test.txt') { |f| 
    puts f.readlines
  }
end
