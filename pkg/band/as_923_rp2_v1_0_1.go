// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package band

import (
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	// AS_923_RP2_v1_0_1 is the band definition for AS923 Group 1 in the RP002-1.0.1 specification.
	AS_923_RP2_v1_0_1 = as923RP2101Band(AS_923, as923Group1Offset)
	// AS_923_2_RP2_v1_0_1 is the band definition for AS923 Group 2 in the RP002-1.0.1 specification.
	AS_923_2_RP2_v1_0_1 = as923RP2101Band(AS_923_2, as923Group2Offset)
	// AS_923_3_RP2_v1_0_1 is the band definition for AS923 Group 3 in the RP002-1.0.1 specification.
	AS_923_3_RP2_v1_0_1 = as923RP2101Band(AS_923_3, as923Group3Offset)
)

var as923RP2101Band = func(id string, offset as923GroupOffset) Band {
	return Band{
		ID: id,

		SupportsDynamicADR: true,

		MaxUplinkChannels: 16,
		UplinkChannels:    as923DefaultChannels(offset),

		MaxDownlinkChannels: 16,
		DownlinkChannels:    as923DefaultChannels(offset),

		SubBands: as923SubBandParameters(offset),

		DataRates: map[ttnpb.DataRateIndex]DataRate{
			ttnpb.DataRateIndex_DATA_RATE_0: makeLoRaDataRate(
				12, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(59, 0),
			),
			ttnpb.DataRateIndex_DATA_RATE_1: makeLoRaDataRate(
				11, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(59, 0),
			),
			ttnpb.DataRateIndex_DATA_RATE_2: makeLoRaDataRate(
				10, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(123, 19),
			),
			ttnpb.DataRateIndex_DATA_RATE_3: makeLoRaDataRate(
				9, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(123, 61),
			),
			ttnpb.DataRateIndex_DATA_RATE_4: makeLoRaDataRate(
				8, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 133),
			),
			ttnpb.DataRateIndex_DATA_RATE_5: makeLoRaDataRate(
				7, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250),
			),
			ttnpb.DataRateIndex_DATA_RATE_6: makeLoRaDataRate(
				7, 250000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250),
			),
			ttnpb.DataRateIndex_DATA_RATE_7: makeFSKDataRate(50000, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		},
		MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_5,
		StrictCodingRate:    true,

		ReceiveDelay1:        defaultReceiveDelay1,
		ReceiveDelay2:        defaultReceiveDelay2,
		JoinAcceptDelay1:     defaultJoinAcceptDelay1,
		JoinAcceptDelay2:     defaultJoinAcceptDelay2,
		MaxFCntGap:           defaultMaxFCntGap,
		ADRAckLimit:          defaultADRAckLimit,
		ADRAckDelay:          defaultADRAckDelay,
		MinRetransmitTimeout: defaultRetransmitTimeout - defaultRetransmitTimeoutMargin,
		MaxRetransmitTimeout: defaultRetransmitTimeout + defaultRetransmitTimeoutMargin,

		DefaultMaxEIRP: 16,
		TxOffset: []float32{
			0,
			-2,
			-4,
			-6,
			-8,
			-10,
			-12,
			-14,
		},

		FreqMultiplier:   100,
		ImplementsCFList: true,
		CFListType:       ttnpb.CFListType_FREQUENCIES,

		Rx1Channel: channelIndexIdentity,
		Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, dwellTime bool) (ttnpb.DataRateIndex, error) {
			so := int8(offset)
			if so > 5 {
				so = 5 - so
			}
			si := int8(idx) - so

			minDR := ttnpb.DataRateIndex_DATA_RATE_0
			if dwellTime {
				minDR = ttnpb.DataRateIndex_DATA_RATE_2
			}
			switch {
			case si <= int8(minDR):
				return minDR, nil
			case si >= 7:
				return ttnpb.DataRateIndex_DATA_RATE_7, nil
			}
			return ttnpb.DataRateIndex(si), nil
		},

		GenerateChMasks: generateChMask16,
		ParseChMask:     parseChMask16,

		DefaultRx2Parameters: Rx2Parameters{
			DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_2,
			Frequency:     as923DefaultRX2Frequency(offset),
		},

		Beacon: Beacon{
			DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_3,
			CodingRate:    Cr4_5,
			Frequencies:   as923BeaconFrequencies(offset),
		},
		PingSlotFrequencies: as923BeaconFrequencies(offset),

		TxParamSetupReqSupport: true,

		// Based on LoRaMac-node 4.6.0.
		// https://github.com/Lora-net/LoRaMac-node/blob/fe8247e2b84101fe701531a5f9ef14f035743af4/src/mac/region/RegionAS923.h#L161-L164
		// https://github.com/Lora-net/LoRaMac-node/blob/fe8247e2b84101fe701531a5f9ef14f035743af4/src/mac/region/RegionCommon.h#L97-L100
		BootDwellTime: DwellTime{
			Uplinks:   boolPtr(true),
			Downlinks: boolPtr(false),
		},
	}
}
