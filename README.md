<img src="https://img.shields.io/badge/lifecycle-experimental-orange.svg" alt="Life cycle: experimental"> <a href="https://godoc.org/github.com/clindet/bioctl"><img src="https://godoc.org/github.com/clindet/bioctl?status.svg" alt="GoDoc"></a>

## bioctl

[bioctl](https://github.com/clindet/bioctl) is a simple command line tool to facilitate the data analysis.

## Installation

```bash
# windows
wget https://github.com/clindet/bioctl/releases/download/v0.1.2/bioctl_0.1.2_Windows_64-bit.tar.gz

# osx
wget https://github.com/clindet/bioctl/releases/download/v0.1.2/bioctl_0.1.2_Darwin_64-bit.tar.gz

# linux
wget https://github.com/clindet/bioctl/releases/download/v0.1.2/bioctl_0.1.2_Linux_64-bit.tar.gz

# get latest version
go get -u github.com/clindet/bioctl
```

## Usage

### Simple parallel tasks

```bash
# concurent 2 tasks with total 8 tasks
echo 'touch /tmp/$1 /tmp/$2; sleep ${1}' > job.sh && bioctl par --cmd "sh job.sh" -t 2 --index 1,2,5-10
# concurent 4 tasks with total 8 tasks and env parse
bioctl par --cmd 'sh job.sh {{index}} {{key2}}' -t 4 --index 1,2,5-10 --env "key2:123"

# concurent 4 tasks with total 8 tasks (direct) and env parse (more log)
bioctl par --cmd 'echo {{index}} {{key2}}; sleep {{index}}' -t 4 --index 1,2,5-10 --env "key2:123" --verbose 2 --save-log

# concurent 4 tasks with total 8 tasks, env
# and not to force add {{index}} at the end of cmd
bioctl par --cmd 'echo {{key2}}; sleep {{index}}' -t 4 --index 1,2,5-10 --env "key2:123" --force-idx false --save-log

# pipe usage
echo 'sh job.sh' | bioctl par -t 4 --index 1,2,5-10 -
```

### Convert function

- support convert PubMed and SRA abstract from XML => JSON

```bash
# convert PubMed XML to JSON
bget api ncbi -q "Galectins control MTOR and AMPK in response to lysosomal damage to induce autophagy OR MTOR-independent autophagy induced by interrupted endoplasmic reticulum-mitochondrial Ca2+ communication: a dead end in cancer cells. OR The PARK10 gene USP24 is a negative regulator of autophagy and ULK1 protein stability OR Coordinate regulation of autophagy and the ubiquitin proteasome system by MTOR." | bioctl cvrt --xml2json pubmed -

# convert SRA XML to JSON
bget api ncbi -d 'sra' -q PRJNA527715 | bioctl cvrt --xml2json sra -
```

### Format function

- support to prettify JSON stream
- support convert key-value JSON to slice JSON

```bash
# json pretty
echo '{"a": {"a": 123}, "b": {"b": 567}}' | bioctl fmt --json-pretty --indent 2 -
# key:value => slice
echo '{"a": {"a": 123}, "b": {"b": 567}}' | bioctl fmt --json-to-slice --indent 4 -
```

### Simple file stat

```bash
# file stat
# equal to wc -l
bioctl fn -l *

bioctl fn -l * --verbose 0
```

### Uncompress files

```bash
mkdir a
mkdir b

zip -q -r -o myfile.zip a
tar -czvf myfile.tar.gz b
rmdir a
rmdir b

bioctl -u 'myfile.zip myfile.tar.gz'

ls -l

rmdir a
rmdir b
rm myfile.zip myfile.tar.gz
```

### Random

```bash
# UUID
bioctl rand --uuid -n 10
# string
bioctl rand --str -l 35 -n 22
# int
bioctl rand --int -l 23 -n 10
```

### Plot related

```bash
# show all themes
bioctl plot --show-themes
# show all theme names
bioctl plot --show-themes-name
# returns color theme
bioctl plot --theme red_blue
```

### Range

```bash
# range number sequence
bioctl range 1 100 2.5
bioctl range --start 2 --end 1000 --step 15

# char mode
bioctl range --mode char --step 5
bioctl range --mode char --step 3 --start-char s --end-char z --step 2 --ref-str qrstuvwxy12442z
bioctl range --mode char --step 3 --start-char s --end-char z --step 2 --ref-str qrstuvwxy12442z --sep ''
```

### Stat

```bash
bioctl stat --min 1 2 3 4 5 100
bioctl stat --max 1 2 3 4 5
bioctl stat --mean 1 3 5 7 9 26 100
bioctl stat --median 1 10 3 4 100 143 123 12 22.2
bioctl stat --mfreq 2 10 2 2 10 10 11 12 14
bioctl stat --var 2 10 2 2 10
bioctl stat --sum 1 1 2 3 2 14234 12 12 1331 23 12 12
// vector1: 1-6 vector2: 10-51
bioctl stat --pearson 1 2 3 4 5 6 10 20 30 40 50 51
bioctl stat --percentile 30 1 2 3 4 5 6 7 8 9 10
bioctl stat --freq 1 2 3 2 1 1 1 1 1 1 | sort

bioctl statdf --min _examples/test.csv --at ",0"
bioctl statdf --max _examples/test.csv --at ",0"
bioctl statdf --mean _examples/test.csv --at ",0"
bioctl statdf --median _examples/test.csv --at ",0"
bioctl statdf --mfreq _examples/test.csv --at ",0"
bioctl statdf --var _examples/test.csv --at ",0"
bioctl statdf --sum _examples/test.csv --at ",0"
bioctl statdf --pearson _examples/test.csv --at ",0:1" --print
bioctl statdf --percentile 30 _examples/test.csv --at ",0"
bioctl statdf --freq _examples/test.csv --at ",0" | sort
```

## Maintainer

- [@Jianfeng](https://github.com/Miachol)

## License

Academic Free License version 3.0

