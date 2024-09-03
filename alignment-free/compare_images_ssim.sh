#!/bin/bash
images=(*mer.png)
# Function to calculate similarity using ImageMagick's compare command
calculate_similarity() {
    local img1="$1"
    local img2="$2"
    local result=$(python compare_ssim.py "$img1" "$img2"  2>&1)
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
