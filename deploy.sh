#!/bash/sh

namespace=leo
workload=jwt
image_name=burger43/jwt-generate-server
short_hash=$(git rev-parse --short HEAD)
image_name_version=$image_name:$short_hash
echo $short_hash


function set_image {
    local namespace=$1
    local workload=$2
    local hash=$3
    # echo "kubectl -n ${namespace} set image deploy ${workload} ${workload}=$image_name_version --record"
    kubectl -n ${namespace} set image deploy ${workload} ${workload}=$image_name_version --record
}

function check_status {
    local namespace=$1
    local workload=$2
    kubectl -n ${namespace} rollout status deploy ${workload}
}

set_image $namespace $workload $image_name_version

check_status $namespace $workload

