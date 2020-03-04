.PHONY: all
all: $(addsuffix .json,\
	000100 \
	000500 \
	001000 \
	005000 \
	010000 \
	050000 \
	100000 \
)

%.json: student.yaml names.txt faculties.txt specialities.txt
	jg $< \
	  -a $* \
	  -f names=names.txt,faculties=faculties.txt,specialities=specialities.txt \
	  -o $@

.PHONY: clean
clean: $(wildcard *.json)
	rm $^
