package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/babylonchain/babylon/x/btcstaking/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

func (k Keeper) AddBTCStakingConsumerEvent(ctx context.Context, consumerID string, event *types.BTCStakingConsumerEvent) error {
	store := k.btcStakingConsumerEventStore(ctx, consumerID)
	storeKey := []byte(consumerID)

	// If the consumer already has events, append the new event to the existing list
	// TODO: repeated marshalling/unmarshalling is inefficient; consider refactoring
	var events types.BTCStakingIBCPacket
	if store.Has(storeKey) {
		eventsBytes := store.Get(storeKey)
		k.cdc.MustUnmarshal(eventsBytes, &events)
	}

	switch {
	case event.GetNewFp() != nil:
		events.NewFp = append(events.NewFp, event.GetNewFp())
	case event.GetActiveDel() != nil:
		events.ActiveDel = append(events.ActiveDel, event.GetActiveDel())
	case event.GetSlashedDel() != nil:
		events.SlashedDel = append(events.SlashedDel, event.GetSlashedDel())
	case event.GetUnbondedDel() != nil:
		events.UnbondedDel = append(events.UnbondedDel, event.GetUnbondedDel())
	default:
		return fmt.Errorf("unrecognized event type for event %+v", event)
	}

	eventsBytes := k.cdc.MustMarshal(&events)
	store.Set(storeKey, eventsBytes)

	return nil
}

// GetBTCStakingConsumerIBCPacket gets BTC staking consumer IBC packet for a given consumer ID.
func (k Keeper) GetBTCStakingConsumerIBCPacket(ctx context.Context, consumerID string) *types.BTCStakingIBCPacket {
	store := k.btcStakingConsumerEventStore(ctx, consumerID)
	storeKey := []byte(consumerID)
	if !store.Has(storeKey) {
		return nil
	}

	var events types.BTCStakingIBCPacket
	eventsBytes := store.Get(storeKey)
	k.cdc.MustUnmarshal(eventsBytes, &events)
	return &events
}

// GetAllBTCStakingConsumerIBCPackets retrieves all BTC staking consumer IBC packets from the store,
// returning a map where the keys are consumer IDs and the values are the corresponding BTCStakingIBCPacket.
func (k Keeper) GetAllBTCStakingConsumerIBCPackets(ctx context.Context) map[string]*types.BTCStakingIBCPacket {
	storeAdaptor := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	evStore := prefix.NewStore(storeAdaptor, types.BTCStakingEventKey)
	iter := evStore.Iterator(nil, nil)
	defer iter.Close()

	// Initialize the map to hold consumer ID keys and IBC packet values.
	allEvents := make(map[string]*types.BTCStakingIBCPacket)
	for ; iter.Valid(); iter.Next() {
		var events types.BTCStakingIBCPacket
		k.cdc.MustUnmarshal(iter.Value(), &events)
		consumerID := string(iter.Key())
		allEvents[consumerID] = &events
	}

	return allEvents
}

func (k Keeper) DeleteBTCStakingConsumerIBCPacket(ctx context.Context, consumerID string) {
	store := k.btcStakingConsumerEventStore(ctx, consumerID)
	storeKey := []byte(consumerID)
	store.Delete(storeKey)
}

// btcStakingConsumerEventStore returns the KVStore of the btc staking event
// prefix: BTCStakingEventKey || consumer ID
// key: BTCStakingConsumerEventType
// value: BTCStakingConsumerEvents (a list of BTCStakingConsumerEvent)
func (k Keeper) btcStakingConsumerEventStore(ctx context.Context, consumerID string) prefix.Store {
	storeAdaptor := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	evStore := prefix.NewStore(storeAdaptor, types.BTCStakingEventKey)
	return prefix.NewStore(evStore, []byte(consumerID))
}
