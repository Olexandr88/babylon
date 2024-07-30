package btcstaking_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/babylonlabs-io/babylon/btcstaking"

	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/require"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
)

func generateTxFromOutputs(r *rand.Rand, info *btcstaking.IdentifiableStakingInfo) (*wire.MsgTx, int, int) {
	numOutputs := r.Int31n(18) + 2

	stakingOutputIdx := int(r.Int31n(numOutputs))
	opReturnOutputIdx := int(datagen.RandomIntOtherThan(r, int(stakingOutputIdx), int(numOutputs)))

	tx := wire.NewMsgTx(2)
	for i := 0; i < int(numOutputs); i++ {
		if i == stakingOutputIdx {
			tx.AddTxOut(info.StakingOutput)
		} else if i == opReturnOutputIdx {
			tx.AddTxOut(info.OpReturnOutput)
		} else {
			tx.AddTxOut(wire.NewTxOut((r.Int63n(1000000000) + 10000), datagen.GenRandomByteArray(r, 32)))
		}
	}

	return tx, stakingOutputIdx, opReturnOutputIdx
}

// Property: Every staking tx generated by our generator should be properly parsed by
// our parser
func FuzzGenerateAndParseValidV0StakingTransaction(f *testing.F) {
	// lot of seeds as test is pretty fast and we want to test a lot of different values
	datagen.AddRandomSeedsToFuzzer(f, 1000)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		// 3 - 10 covenants
		numCovenantKeys := uint32(r.Int31n(7) + 3)
		quroum := uint32(r.Intn(int(numCovenantKeys)) + 1)
		stakingAmount := btcutil.Amount(r.Int63n(1000000000) + 10000)
		stakingTime := uint16(r.Int31n(math.MaxUint16-1) + 1)
		tag := datagen.GenRandomByteArray(r, btcstaking.TagLen)
		net := &chaincfg.MainNetParams

		sc := GenerateTestScenario(r, t, 1, numCovenantKeys, quroum, stakingAmount, stakingTime)

		outputs, err := btcstaking.BuildV0IdentifiableStakingOutputs(
			tag,
			sc.StakerKey.PubKey(),
			sc.FinalityProviderKeys[0].PubKey(),
			sc.CovenantPublicKeys(),
			quroum,
			stakingTime,
			stakingAmount,
			net,
		)

		require.NoError(t, err)
		require.NotNil(t, outputs)

		tx, stakingOutputIdx, opReturnOutputIdx := generateTxFromOutputs(r, outputs)

		// ParseV0StakingTx and IsPossibleV0StakingTx should be consistent and recognize
		// the same tx as a valid staking tx
		require.True(t, btcstaking.IsPossibleV0StakingTx(tx, tag))

		parsedTx, err := btcstaking.ParseV0StakingTx(
			tx,
			tag,
			sc.CovenantPublicKeys(),
			quroum,
			net,
		)
		require.NoError(t, err)
		require.NotNil(t, parsedTx)

		require.Equal(t, outputs.StakingOutput.PkScript, parsedTx.StakingOutput.PkScript)
		require.Equal(t, outputs.StakingOutput.Value, parsedTx.StakingOutput.Value)
		require.Equal(t, stakingOutputIdx, parsedTx.StakingOutputIdx)

		require.Equal(t, outputs.OpReturnOutput.PkScript, parsedTx.OpReturnOutput.PkScript)
		require.Equal(t, outputs.OpReturnOutput.Value, parsedTx.OpReturnOutput.Value)
		require.Equal(t, opReturnOutputIdx, parsedTx.OpReturnOutputIdx)

		require.Equal(t, tag, parsedTx.OpReturnData.Tag)
		require.Equal(t, uint8(0), parsedTx.OpReturnData.Version)
		require.Equal(t, stakingTime, parsedTx.OpReturnData.StakingTime)

		require.Equal(t, schnorr.SerializePubKey(sc.StakerKey.PubKey()), parsedTx.OpReturnData.StakerPublicKey.Marshall())
		require.Equal(t, schnorr.SerializePubKey(sc.FinalityProviderKeys[0].PubKey()), parsedTx.OpReturnData.FinalityProviderPublicKey.Marshall())
	})
}

// TODO Negative test cases
