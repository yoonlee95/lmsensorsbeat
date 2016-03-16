BEATNAME=lmsensorsbeat
BEAT_DIR=github.com/eskibars
SYSTEM_TESTS=false
TEST_ENVIRONMENT=false
ES_BEATS=./vendor/github.com/elastic/beats
GOPACKAGES=$(shell glide novendor)
PREFIX?=.

# Path to the libbeat Makefile
-include $(ES_BEATS)/libbeat/scripts/Makefile

.PHONY: init
init:
	glide update  --no-recursive
	make update
	git init
	git add .

.PHONY: install-cfg
install-cfg:
	mkdir -p $(PREFIX)
	cp etc/lmsensorsbeat.template.json     $(PREFIX)/lmsensorsbeat.template.json
	cp etc/lmsensorsbeat.yml               $(PREFIX)/lmsensorsbeat.yml
	cp etc/lmsensorsbeat.yml               $(PREFIX)/lmsensorsbeat-linux.yml
	cp etc/lmsensorsbeat.yml               $(PREFIX)/lmsensorsbeat-binary.yml
	cp etc/lmsensorsbeat.yml               $(PREFIX)/lmsensorsbeat-darwin.yml
	cp etc/lmsensorsbeat.yml               $(PREFIX)/lmsensorsbeat-win.yml

.PHONY: update-deps
update-deps:
	glide update  --no-recursive
