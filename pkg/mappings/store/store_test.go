package store

import (
	"context"
	"testing"

	"github.com/loft-sh/vcluster/pkg/scheme"
	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	testingutil "github.com/loft-sh/vcluster/pkg/util/testing"
	"gotest.tools/v3/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStore(t *testing.T) {
	genericStore, err := NewStore(context.TODO(), testingutil.NewFakeClient(scheme.Scheme), testingutil.NewFakeClient(scheme.Scheme), NewMemoryBackend())
	assert.NilError(t, err)

	store, ok := genericStore.(*Store)
	assert.Equal(t, true, ok)

	gvk := corev1.SchemeGroupVersion.WithKind("Secret")
	virtualName := types.NamespacedName{
		Name:      "virtual-name",
		Namespace: "virtual-namespace",
	}
	hostName := types.NamespacedName{
		Name:      "host-name",
		Namespace: "host-namespace",
	}
	labelMapping := synccontext.LabelMapping{
		Virtual: "virtual-label",
		Host:    "host-label",
	}

	baseCtx := context.TODO()
	baseMapping := synccontext.NameMapping{
		GroupVersionKind: gvk,
		VirtualName:      virtualName,
	}

	// record reference
	err = store.RecordReference(baseCtx, synccontext.NameMapping{
		GroupVersionKind: gvk,
		HostName:         hostName,
		VirtualName:      virtualName,
	}, baseMapping)
	assert.NilError(t, err)

	// virtual -> host
	translatedHostName, ok := store.VirtualToHostName(baseCtx, synccontext.Object{
		GroupVersionKind: gvk,
		NamespacedName:   virtualName,
	})
	assert.Equal(t, true, ok)
	assert.Equal(t, hostName, translatedHostName)

	// virtual -> host
	translatedVirtualName, ok := store.HostToVirtualName(baseCtx, synccontext.Object{
		GroupVersionKind: gvk,
		NamespacedName:   hostName,
	})
	assert.Equal(t, true, ok)
	assert.Equal(t, virtualName, translatedVirtualName)

	// virtual -> host
	_, ok = store.HostToVirtualName(baseCtx, synccontext.Object{
		GroupVersionKind: gvk,
	})
	assert.Equal(t, false, ok)

	// check inner structure of store
	assert.Equal(t, 1, len(store.mappings))
	assert.Equal(t, 0, len(store.hostToVirtualLabel))
	assert.Equal(t, 0, len(store.hostToVirtualLabelCluster))
	assert.Equal(t, 0, len(store.virtualToHostLabel))
	assert.Equal(t, 0, len(store.virtualToHostLabelCluster))
	assert.Equal(t, 1, len(store.hostToVirtualName))
	assert.Equal(t, 1, len(store.virtualToHostName))

	// make sure the mapping is not added
	nameMapping := synccontext.NameMapping{
		GroupVersionKind: gvk,
		HostName:         hostName,
		VirtualName:      virtualName,
	}
	err = store.RecordReference(baseCtx, nameMapping, baseMapping)
	assert.NilError(t, err)
	assert.Equal(t, 1, len(store.mappings))
	assert.Equal(t, 0, len(store.hostToVirtualLabel))
	assert.Equal(t, 0, len(store.hostToVirtualLabelCluster))
	assert.Equal(t, 0, len(store.virtualToHostLabel))
	assert.Equal(t, 0, len(store.virtualToHostLabelCluster))
	assert.Equal(t, 1, len(store.hostToVirtualName))
	assert.Equal(t, 1, len(store.virtualToHostName))

	// validate mapping itself
	mapping, ok := store.mappings[nameMapping]
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(mapping.References))
	assert.Equal(t, 0, len(mapping.Labels))
	assert.Equal(t, 0, len(mapping.LabelsCluster))

	// map label
	err = store.RecordLabel(baseCtx, labelMapping, baseMapping)
	assert.NilError(t, err)

	// check mappings
	hostLabel, ok := store.VirtualToHostLabel(baseCtx, labelMapping.Virtual)
	assert.Equal(t, true, ok)
	assert.Equal(t, hostLabel, labelMapping.Host)
	virtualLabel, ok := store.HostToVirtualLabel(baseCtx, labelMapping.Host)
	assert.Equal(t, true, ok)
	assert.Equal(t, virtualLabel, labelMapping.Virtual)

	// validate mapping itself
	mapping, ok = store.mappings[nameMapping]
	assert.Equal(t, true, ok)
	assert.Equal(t, 0, len(mapping.References))
	assert.Equal(t, 1, len(mapping.Labels))
	assert.Equal(t, 0, len(mapping.LabelsCluster))

	// garbage collect mapping
	store.garbageCollectMappings(context.TODO())
	_, ok = store.mappings[nameMapping]
	assert.Equal(t, false, ok)
}
