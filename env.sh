DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export GOPATH=${DIR}/src/go
export PATH=$PATH:$(go env GOPATH)/bin

alias go="./pants go :: -- "
