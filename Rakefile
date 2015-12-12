require 'bundler'
Bundler::GemHelper.install_tasks
require 'rake/extensiontask'

task :default => [:compile]

spec = eval File.read('memberlist.gemspec')
Rake::ExtensionTask.new('memberlist', spec) do |ext|
  ext.ext_dir = 'ext/memberlist'
  ext.lib_dir = File.join(*['lib', 'memberlist', ENV['FAT_DIR']].compact)
  ext.source_pattern = "*.{c,cpp,go}"
end
