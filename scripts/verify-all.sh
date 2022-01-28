for file in ../test_data/*.txt; do
    echo "******  Verifying file $file ..."
    ../hxu file verify $file ;
    echo "Done."
done