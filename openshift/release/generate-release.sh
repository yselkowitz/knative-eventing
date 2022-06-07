#!/usr/bin/env bash

source $(dirname $0)/resolve.sh

release=$1

artifacts_dir="openshift/release/artifacts"
rm -rf $artifacts_dir
mkdir -p $artifacts_dir

if [ "$release" == "ci" ]; then
    image_prefix="registry.ci.openshift.org/openshift/knative-nightly:knative-eventing-"
    tag=""
else
    image_prefix="registry.ci.openshift.org/openshift/knative-${release}:knative-eventing-"
    tag=""
fi

eventing_core="${artifacts_dir}/eventing-core.yaml"
eventing_crds="${artifacts_dir}/eventing-crds.yaml"
in_memory_channel="${artifacts_dir}/in-memory-channel.yaml"
mt_channel_broker="${artifacts_dir}/mt-channel-broker.yaml"
eventing_sugar_controller="${artifacts_dir}/eventing-sugar-controller.yaml"
eventing_post_install="${artifacts_dir}/eventing-post-install.yaml"

# Eventing CRDs
resolve_resources config/core/resources "${eventing_crds}" "$image_prefix" "$tag"
# Eventing core
resolve_resources config "${eventing_core}" "$image_prefix" "$tag"
# Eventing post-install
resolve_resources config/post-install "${eventing_post_install}" "$image_prefix" "$tag"
# Sugar controller
resolve_resources config/sugar "${eventing_sugar_controller}" "$image_prefix" "$tag"
# In memory channel
resolve_resources config/channels/in-memory-channel "${in_memory_channel}" "$image_prefix" "$tag"
resolve_resources config/channels/in-memory-channel/configmaps "${in_memory_channel}" "$image_prefix" "$tag"
resolve_resources config/channels/in-memory-channel/deployments "${in_memory_channel}" "$image_prefix" "$tag"
resolve_resources config/channels/in-memory-channel/resources "${in_memory_channel}" "$image_prefix" "$tag"
resolve_resources config/channels/in-memory-channel/roles "${in_memory_channel}" "$image_prefix" "$tag"
resolve_resources config/channels/in-memory-channel/webhooks "${in_memory_channel}" "$image_prefix" "$tag"
# MT Broker
resolve_resources config/brokers/mt-channel-broker "${mt_channel_broker}" "$image_prefix" "$tag"
