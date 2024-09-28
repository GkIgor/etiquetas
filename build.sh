platforms=("windows/amd64" "linux/amd64")
package="br-atacadao.corp/etiquetas"
IFS='/'
package_name=""

for part in $package; do
    package_name="$part"
done

unset IFS

echo $package_name
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=$package_name'-'$GOOS'-'$GOARCH

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o build/$output_name $package
done