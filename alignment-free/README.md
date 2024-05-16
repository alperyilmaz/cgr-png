# Alignment Free benchmark

[AFproject](https://afproject.org/app/) provides benchmark for testing various algoritms. We applied our approach to "Benchmark -> Genome Based Phylogeny -> Full Genome Sequences -> Plants" data.

## Download data

```bash
wget https://afproject.org/media/genome/std/assembled/plants/dataset/assembled-plants.zip
unzip assembled-plants.zip
mv assembled-plants/*fasta .
```

## Count and prepare images for all genomes

If you are able to compile the Golang version of *counts2image* script, you can run the following script

```bash
ls *fasta | cut -f1 -d. | while read GENOME; do bash count_genome_fast.sh $GENOME 12; done
```
Otherwise please run the slower Python version

```bash
ls *fasta | cut -f1 -d. | while read GENOME; do bash count_genome.sh $GENOME 12; done
```

## Use NCC metric to calculate similarity for each possible pair of images

```bash
bash compare_images.sh NCC > compare_all_NCC_results
```

## Prepare distance/similarity matrix from results

This step requires *datamash* which can be installed with `sudo apt install datamash`. *tsv-pretty* is from [tsv-utils](https://github.com/eBay/tsv-utils) tool, please refer to its Github link for installation. 

```bash
function _convert_to_matrix() { awk  '{if($1 > $2){print $1,$2,$3}else{print $2,$1,$3}}' | datamash -t" " crosstab 1,2 sum 3 |  tsv-pretty -d" " -s1 | sed -e 's/N\/A/   /g'; }

cat compare_all_NCC_results | awk '{printf"%s %s %s\n",$3,$5,$6}' |  sed -e 's/_12mer.png:*//g' | _convert_to_matrix > NCC_matrix
```

# Compare results with benchmark results

The best aligment-free phylogeny result was by co-phylog algoritm, the raw data was loaded from https://afproject.org/app/benchmark/genome/std/assembled/plants/results/ and then we compared our results with *co-phylog-results*

```bash
cat co-phylog-results | tr "\t" " " | _convert_to_matrix
```

# The NCC distance matrix

NCC results are very close but NCC is similarity but co-phylog data is distance matrix, so we should extract the values from 1 to have distance

```bash
cat compare_all_NCC_results | awk '{$NF=1-$NF; print $0}' | awk '{printf"%s %s %s\n",$3,$5,$6}' |  sed -e 's/_12mer.png:*//g' | _convert_to_matrix > NCC_distance_matrix
```
