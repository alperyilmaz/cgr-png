# cgr-png
Lossless compression of kmer counts with chaos game representation. Here are the steps for generating 12mer image for human genome.

## Download data

```bash
wget https://ftp.ebi.ac.uk/pub/databases/gencode/Gencode_human/release_45/GRCh38.primary_assembly.genome.fa.gz
gunzip GRCh38.primary_assembly.genome.fa.gz
```

## Prepare kmer count data

We used KMC for kmer counting. You can install KMC in Debian/Ubuntu distros by `sudo apt install kmc`

```bash
kmc -b -fm -k12 -ci0 -cs5000000 GRCh38.primary_assembly.genome.fa  hg38_kmc_12mer_counts .
kmc_dump hg38_kmc_12mer_counts hg38_kmc_12mer_counts.tsv
```

## Convert kmer counts to PNG image

We provided Python script (might be slow) to convert kmer counts to image

```bash
cat hg38_kmc_12mer_counts.tsv | python counts2image.py 12 > hg38_12mer.png
```

For faster processing we also provided equivalent Golang version. Please compile the script and make it executable before using it.

```bash
# run following two lines only once
go build counts2image_fast.go
chmod 755 counts2image_fast.go

cat hg38_kmc_12mer_counts.tsv | ./counts2image_fast 12 > go-hg38_12mer.png
```

## Extract kmer counts from PNG images

```bash
python image2counts.py hg38_12mer.png > 12mer-hg38.tsv

# go version, build and make executable for once
go build image2counts_fast.go
chmod 755 image2counts_fast

./image2counts_fast hg38_12mer.png > 12mer-hg38-go.tsv
```

## Evidence for losslessness 

For a small set of kmer counts let's prepare an image and then extract kmer counts from that image and see if we get the same kmer count values.

```
$ cat small-set-12mers.tsv
TACGTTTCGCAT	3
AAGAAAAGGCCG	121
GAAAACTGAAAC	2818
TAATATTCATAT	1048
GTGGATAGCTGC	64
AGACGAAGGTAT	32
CTTAGTGAAGGT	154
CTCTTGTTTTAC	369
TGCGCAGCCCGC	28
TGCTTAGAGCAC	156

# convert kmers into image (the compression)
$ cat small-set-12mers.tsv | python counts2image.py 12 >| small-set-12mer.png

# extract kmer counts from image (skip kmers which have zero count)
$ python image2counts.py small-set-12mer.png | awk '$2!=0'
AAGAAAAGGCCG	121
GAAAACTGAAAC	2818
AGACGAAGGTAT	32
GTGGATAGCTGC	64
TAATATTCATAT	1048
TGCGCAGCCCGC	28
TACGTTTCGCAT	3
TGCTTAGAGCAC	156
CTTAGTGAAGGT	154
CTCTTGTTTTAC	369
```

