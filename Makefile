define gen_data
	tools/protoc-3.9.0-osx-x86_64/bin/protoc --go_out=paths=source_relative:./gen -I=. $(1)
endef

gen_protos:
	$(call gen_data,protos/perfetto/common/descriptor.proto)
	$(call gen_data,protos/perfetto/trace/perfetto_trace.proto)
	$(call gen_data,protos/perfetto/trace_processor/trace_processor.proto)
