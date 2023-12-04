changelog() {
    local name=$(awk '/^module/{print $2}' ../go.mod)
    local version=$(cat ../.version)
    local title="# $name v$version"

cat << EOF > ../CHANGELOG
$title

$(git log --pretty=format:%s)

$(cat ../CHANGELOG)
EOF
}
