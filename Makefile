.PHONY: test
test:
	go test -v ./...

.PHONY: gen
gen:
	$(MAKE) -C testdata

.PHONY: clean
clean:
	$(MAKE) -C testdata clean

.PHONY: bench
bench: $(wildcard testdata/*.json)
	go test -timeout 0 -bench . -args $(sort $^)

.PHONY: report
report:
	docker-compose up