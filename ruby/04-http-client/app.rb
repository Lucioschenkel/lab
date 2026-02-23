require 'uri'
require 'net/http'

uri = URI('https://pokeapi.co/api/v2/pokemon/pikachu')
res = Net::HTTP.get_response(uri)
puts res.body if res.is_a?(Net::HTTPSuccess)
