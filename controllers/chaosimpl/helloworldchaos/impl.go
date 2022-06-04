package helloworldchaos

import (
    "context"

    "github.com/chaos-mesh/chaos-mesh/api/v1alpha1"
    "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/utils"
    "github.com/chaos-mesh/chaos-mesh/pkg/chaosdaemon/pb"
    "github.com/go-logr/logr"
    "go.uber.org/fx"
    "sigs.k8s.io/controller-runtime/pkg/client"
    impltypes "github.com/chaos-mesh/chaos-mesh/controllers/chaosimpl/types"
)

type Impl struct {
    client.Client
    Log     logr.Logger
    decoder *utils.ContainerRecordDecoder
}

// Apply applies KernelChaos
func (impl *Impl) Apply(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
    impl.Log.Info("Apply helloworld chaos")
    decodedContainer, err := impl.decoder.DecodeContainerRecord(ctx, records[index], obj)
    if err != nil {
        return v1alpha1.NotInjected, err
    }
    pbClient := decodedContainer.PbClient
    containerId := decodedContainer.ContainerId

    _, err = pbClient.ExecHelloWorldChaos(ctx, &pb.ExecHelloWorldRequest{
        ContainerId: containerId,
    })
    if err != nil {
        return v1alpha1.NotInjected, err
    }

    return v1alpha1.Injected, nil
}

// Recover means the reconciler recovers the chaos action
func (impl *Impl) Recover(ctx context.Context, index int, records []*v1alpha1.Record, obj v1alpha1.InnerObject) (v1alpha1.Phase, error) {
    impl.Log.Info("Recover helloworld chaos")
    return v1alpha1.NotInjected, nil
}

func NewImpl(c client.Client, log logr.Logger, decoder *utils.ContainerRecordDecoder) *impltypes.ChaosImplPair {
    return &impltypes.ChaosImplPair{
        Name:   "helloworldchaos",
        Object: &v1alpha1.HelloWorldChaos{},
        Impl: &Impl{
            Client:  c,
            Log:     log.WithName("helloworldchaos"),
            decoder: decoder,
        },
        ObjectList: &v1alpha1.HelloWorldChaosList{},
    }
}

var Module = fx.Provide(
    fx.Annotated{
        Group:  "impl",
        Target: NewImpl,
    },
)