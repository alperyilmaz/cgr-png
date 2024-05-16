#!/bin/bash
METRIC=$1
# Define the list of image files
images=("halophilu_12mer.png" "parvulum_12mer.png" "sinensis_12mer.png" "rubella_12mer.png" "vinifera_12mer.png" "camaldule_12mer.png" "grandis_12mer.png" "thalian_12mer.png" "lyrata_12mer.png" "raimondii_12mer.png" "rapa_12mer.png" "cacao_12mer.png" "papaya_12mer.png" "clementin_12mer.png")

# Function to calculate similarity using ImageMagick's compare command
calculate_similarity() {
    local img1="$1"
    local img2="$2"
    local result=$(compare -metric $METRIC "$img1" "$img2" null: 2>&1)
    echo "$result"
}

# Iterate over pairs of images
for ((i = 0; i < ${#images[@]}; i++)); do
    for ((j = i + 1; j < ${#images[@]}; j++)); do
        img1="${images[$i]}"
        img2="${images[$j]}"
        similarity=$(calculate_similarity "$img1" "$img2")
        echo "Similarity between $img1 and $img2: $similarity"
    done
done
