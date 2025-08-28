# simple program that generates strings in bf
# only ascii letters, cant do escape chars?

def s_to_uint8(s)
    return s.split("").map{|c| c.ord }
end

def gen(bytes)
    bytes.each do |b|
        # print "#{b}"
        0.upto(b-1) do
            print "+"
        end
        print ". (#{b.chr}) >\n"
    end
end

if ARGV.empty?
    print "USAGE:\n ruby gen <string> > output"
end

gen(s_to_uint8(ARGV[0]))
