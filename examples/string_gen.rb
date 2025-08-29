# simple program that generates strings in bf
# only ascii letters, cant do escape chars?
# any "\" will get replaced with "\n"

def s_to_uint8(s)
    return s.split("").map{|c| if c == "\\" then 10 else c.ord end }
end

def gen(bytes)
    bytes.each do |b|
        print "#{b.chr}"
        0.upto(b-1) do
            print "+"
        end
        print ". >\n"
    end
end

if ARGV.empty?
    print "USAGE:\n ruby gen <string> > output"
end

gen(s_to_uint8(ARGV[0]))
