.PHONY: dev
dev:
	ctlptl apply -f "https://ctlptl.softonic.io/kind.yaml"
	skaffold dev
