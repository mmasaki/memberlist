require 'mkmf'
find_executable('go')

$objs = []
def $objs.empty?; false ;end

create_makefile("memberlist/memberlist")

case `#{CONFIG['CC']} --version`
when /Free Software Foundation/
  ldflags = '-Wl,--unresolved-symbols=ignore-all'
when /clang/
  ldflags = '-undefined dynamic_lookup'
end

File.open('Makefile', 'a') do |f|
  f.write <<-EOS.gsub(/^ {8}/, "\t")
$(DLLIB): Makefile $(srcdir)/memberlist.go $(srcdir)/wrapper.go
        CGO_CFLAGS='$(INCFLAGS)' CGO_LDFLAGS='#{ldflags}' \
          go build -p 4 -buildmode=c-shared -o $(DLLIB)
  EOS
end
