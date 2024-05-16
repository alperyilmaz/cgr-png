#!/usr/bin/bash

GENOME=$1
KMER=$2
kmc -b -fm -k"$KMER" -ci0 -cs5000000 "$GENOME".fasta "$GENOME"_counts .
kmc_dump "$GENOME"_counts  "$GENOME"_"$KMER"mer_counts.tsv
cat "$GENOME"_"$KMER"mer_counts.tsv | python ../counts2image.py "$KMER" > "$GENOME"_"$KMER"mer.png
