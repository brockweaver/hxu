for file in ../testdata/*.txt; do
    echo "******  Verifying file $file ..."
    ../hxu file verify $file ;
    echo "Done."
done