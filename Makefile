.PHONY: clean
clean:
	rm -f tests/types_borsh.go

.PHONY: gen
gen:
	go run ./tests/gen/

.PHONY: regen
regen: clean gen