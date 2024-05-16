# count and prepare images for all genomes
ls *fasta | cut -f1 -d. | while read GENOME; do bash count_genome.sh $GENOME 12; done

# compare all pairs of images for different metrics
compare -list metric | while read METRIC; do bash compare_images.sh $METRIC > compare_all_"$METRIC"_results; done

# prepare distance/similarity matrix from results
function _convert_to_matrix() { awk  '{if($1 > $2){print $1,$2,$3}else{print $2,$1,$3}}' | datamash -t" " crosstab 1,2 sum 3 |  tsv-pretty -d" " -s1 | sed -e 's/N\/A/   /g'; }

compare -list metric | while read METRIC; do cat compare_all_"$METRIC"_results | awk '{printf"%s %s %s\n",$3,$5,$6}' |  sed -e 's/_12mer.png:*//g' | _convert_to_matrix > "$METRIC"_matrix; done

# compare with benchmark results
# the best result was by co-phylog algoritm, the raw data was loaded from https://afproject.org/app/benchmark/genome/std/assembled/plants/results/

cat co-phylog-results | tr "\t" " " | _convert_to_matrix

# looks like NCC results are very close but NCC is similarity but co-phylog data is distance matrix, so we should extract the values from 1 to have distance


