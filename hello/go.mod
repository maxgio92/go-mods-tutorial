module example.com/hello

go 1.17

replace example.com/please => ../please

require example.com/greetings v0.0.0-00010101000000-000000000000

replace example.com/greetings => ../greetings
