require 'ffi'

module Sum
  extend FFI::Library
  ffi_lib 'lib.dylib'
  attach_function :poll, [], :void
end

Sum.poll()

sleep 100