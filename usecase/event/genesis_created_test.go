package event_test

import (
	event_entity "github.com/WilliamXieCrypto/chain-indexing/entity/event"
	"github.com/WilliamXieCrypto/chain-indexing/usecase/model/genesis"
	jsoniter "github.com/json-iterator/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/WilliamXieCrypto/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeGenesisCreated", func() {
		It("should able to encode and decode to the same event", func() {
			var anyGenesis genesis.Genesis
			err := jsoniter.UnmarshalFromString(GENESIS_JSON, &anyGenesis)
			Expect(err).To(BeNil())
			event := event_usecase.NewGenesisCreated(anyGenesis)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.GENESIS_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.GenesisCreated)
			Expect(typedEvent.Name()).To(Equal(event_usecase.GENESIS_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.Height()).To(Equal(int64(0)))

			Expect(typedEvent.Genesis).To(Equal(anyGenesis))
		})
	})
})

const GENESIS_JSON = `{
  "genesis_time": "2020-10-13T08:55:58.046949Z",
  "chain_id": "testnet-croeseid-1",
  "initial_height": "1",
  "consensus_params": {
	"block": {
	  "max_bytes": "22020096",
	  "max_gas": "-1",
	  "time_iota_ms": "1000"
	},
	"evidence": {
	  "max_age_num_blocks": "100000",
	  "max_age_duration": "172800000000000"
	},
	"validator": {
	  "pub_key_types": [
		"ed25519"
	  ]
	},
	"version": {}
  },
  "app_hash": "",
  "app_state": {
	"auth": {
	  "params": {
		"max_memo_characters": "256",
		"tx_sig_limit": "7",
		"tx_size_cost_per_byte": "10",
		"sig_verify_cost_ed25519": "590",
		"sig_verify_cost_secp256k1": "1000"
	  },
	  "accounts": [
		{
		  "@type": "/cosmos.auth.v1beta1.BaseAccount",
		  "address": "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
		  "pub_key": null,
		  "account_number": "0",
		  "sequence": "0"
		},
		{
		  "@type": "/cosmos.auth.v1beta1.BaseAccount",
		  "address": "tcro1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dsdjun5p",
		  "pub_key": null,
		  "account_number": "0",
		  "sequence": "0"
		},
		{
		  "@type": "/cosmos.auth.v1beta1.BaseAccount",
		  "address": "tcro16kqr009ptgken6qsxnzfnyjfsq6q97g3fxwppq",
		  "pub_key": null,
		  "account_number": "0",
		  "sequence": "0"
		},
		{
		  "@type": "/cosmos.vesting.v1beta1.DelayedVestingAccount",
		  "base_vesting_account": {
			"base_account": {
			  "address": "tcro1e3mpg4kkz9j5h4r28fl74mzmggmjw5e9rece0k",
			  "pub_key": null,
			  "account_number": "0",
			  "sequence": "0"
			},
			"original_vesting": [
			  {
				"denom": "basetcro",
				"amount": "2000000000000000000"
			  }
			],
			"delegated_free": [],
			"delegated_vesting": [],
			"end_time": "1603788958"
		  }
		},
		{
		  "@type": "/cosmos.auth.v1beta1.BaseAccount",
		  "address": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q",
		  "pub_key": null,
		  "account_number": "0",
		  "sequence": "0"
		},
		{
		  "@type": "/cosmos.auth.v1beta1.BaseAccount",
		  "address": "tcro1fe4c63sez0hcuawryzxhc52ksr8rd5t7shvyrv",
		  "pub_key": null,
		  "account_number": "0",
		  "sequence": "0"
		}
	  ]
	},
	"bank": {
	  "params": {
		"send_enabled": [],
		"default_send_enabled": true
	  },
	  "balances": [
		{
		  "address": "tcro1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dsdjun5p",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "20000000000000"
			}
		  ]
		},
		{
		  "address": "tcro1fe4c63sez0hcuawryzxhc52ksr8rd5t7shvyrv",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "4000000000000000000"
			}
		  ]
		},
		{
		  "address": "tcro1j7pej8kplem4wt50p4hfvndhuw5jprxxn5625q",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "2000000000000000000"
			}
		  ]
		},
		{
		  "address": "tcro1e3mpg4kkz9j5h4r28fl74mzmggmjw5e9rece0k",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "2000000000000000000"
			}
		  ]
		},
		{
		  "address": "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "20000000000000"
			}
		  ]
		},
		{
		  "address": "tcro16kqr009ptgken6qsxnzfnyjfsq6q97g3fxwppq",
		  "coins": [
			{
			  "denom": "basetcro",
			  "amount": "20000000000000"
			}
		  ]
		}
	  ],
	  "supply": [],
	  "denom_metadata": [
		{
		  "description": "The native token of Crypto.com app.",
		  "denom_units": [
			{
			  "denom": "basetcro",
			  "exponent": 0,
			  "aliases": [
				"carson"
			  ]
			},
			{
			  "denom": "tcro",
			  "exponent": 8,
			  "aliases": []
			}
		  ],
		  "base": "basetcro",
		  "display": "tcro"
		}
	  ]
	},
	"chainmain": {},
	"distribution": {
	  "delegator_starting_infos": [],
	  "delegator_withdraw_infos": [],
	  "fee_pool": {
		"community_pool": []
	  },
	  "outstanding_rewards": [],
	  "params": {
		"base_proposer_reward": "0.010000000000000000",
		"bonus_proposer_reward": "0.040000000000000000",
		"community_tax": "0",
		"withdraw_addr_enabled": true
	  },
	  "previous_proposer": "",
	  "validator_accumulated_commissions": [],
	  "validator_current_rewards": [],
	  "validator_historical_rewards": [],
	  "validator_slash_events": []
	},
	"evidence": {
	  "evidence": []
	},
	"genutil": {
	  "gen_txs": [
		{
		  "body": {
			"messages": [
			  {
				"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
				"description": {
				  "moniker": "jotaro",
				  "identity": "",
				  "website": "",
				  "security_contact": "",
				  "details": ""
				},
				"commission": {
				  "rate": "0.100000000000000000",
				  "max_rate": "0.200000000000000000",
				  "max_change_rate": "0.010000000000000000"
				},
				"min_self_delegation": "1",
				"delegator_address": "tcro16kqr009ptgken6qsxnzfnyjfsq6q97g3fxwppq",
				"validator_address": "tcrocncl16kqr009ptgken6qsxnzfnyjfsq6q97g3uedcer",
				"pubkey": "tcrocnclconspub1zcjduepq5xp88wqmrhkg3xuyl6vcey3d93kw6cdglkmq4ley3ysvjfx90jnqlvaxpc",
				"value": {
				  "denom": "basetcro",
				  "amount": "10000000000000"
				}
			  }
			],
			"memo": "6989460472eb5f65faf4fdfff910f1bea800b10e@10.10.5.7:26656",
			"timeout_height": "0",
			"extension_options": [],
			"non_critical_extension_options": []
		  },
		  "auth_info": {
			"signer_infos": [
			  {
				"public_key": {
				  "@type": "/cosmos.crypto.secp256k1.PubKey",
				  "key": "AqdoWLPCuN3Jf1FSrsvadkqaUcfK+/0Q2Eo0wCZA7aVC"
				},
				"mode_info": {
				  "single": {
					"mode": "SIGN_MODE_DIRECT"
				  }
				},
				"sequence": "0"
			  }
			],
			"fee": {
			  "amount": [],
			  "gas_limit": "200000",
			  "payer": "",
			  "granter": ""
			}
		  },
		  "signatures": [
			"Pk0xhx2h8bFzGu1RYqWbahjoyHkoQWXkuQNefIc3RBwOXY2lNDz3Y0NDBOAxtV6QhXwNtSpbe2riqFrzX+RZ+g=="
		  ]
		},
		{
		  "body": {
			"messages": [
			  {
				"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
				"description": {
				  "moniker": "Stater",
				  "identity": "",
				  "website": "",
				  "security_contact": "",
				  "details": ""
				},
				"commission": {
				  "rate": "0.100000000000000000",
				  "max_rate": "0.200000000000000000",
				  "max_change_rate": "0.010000000000000000"
				},
				"min_self_delegation": "1",
				"delegator_address": "tcro1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dsdjun5p",
				"validator_address": "tcrocncl1fja5nsxz7gsqw4zccuuy8r7pjnjmc7dscdl2vz",
				"pubkey": "tcrocnclconspub1zcjduepqp6el28dgz0fs6pm2265v4xmx0uys65zf5s2av6r5gh0hmcv5j64qg6zj4w",
				"value": {
				  "denom": "basetcro",
				  "amount": "10000000000000"
				}
			  }
			],
			"memo": "7d06e2ca8363902b4cd8adb214d2d00b0464f9cd@10.10.5.7:26656",
			"timeout_height": "0",
			"extension_options": [],
			"non_critical_extension_options": []
		  },
		  "auth_info": {
			"signer_infos": [
			  {
				"public_key": {
				  "@type": "/cosmos.crypto.secp256k1.PubKey",
				  "key": "AwjuBMBhxVRLvgSl+LVAVc3Yb8nHCKoCwellBB/VYmKq"
				},
				"mode_info": {
				  "single": {
					"mode": "SIGN_MODE_DIRECT"
				  }
				},
				"sequence": "0"
			  }
			],
			"fee": {
			  "amount": [],
			  "gas_limit": "200000",
			  "payer": "",
			  "granter": ""
			}
		  },
		  "signatures": [
			"dvTk7U8iIFvrnUrZtE1CuvD/0vydaWcAwgYFWFUNHdQgei6SoRe+4q1XLr04rK3rcS+v1uIbBnpeTgSOfc5eKQ=="
		  ]
		},
		{
		  "body": {
			"messages": [
			  {
				"@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
				"description": {
				  "moniker": "Argenteus",
				  "identity": "",
				  "website": "",
				  "security_contact": "",
				  "details": ""
				},
				"commission": {
				  "rate": "0.100000000000000000",
				  "max_rate": "0.200000000000000000",
				  "max_change_rate": "0.010000000000000000"
				},
				"min_self_delegation": "1",
				"delegator_address": "tcro16yzcz3ty94awr7nr2txek9dp2klp2av9egkgxn",
				"validator_address": "tcrocncl16yzcz3ty94awr7nr2txek9dp2klp2av9vh437s",
				"pubkey": "tcrocnclconspub1zcjduepq9y742s8duh0eqenjdpc24k785c4en90gypn2vs2x3gxjd45xpp5svldq9u",
				"value": {
				  "denom": "basetcro",
				  "amount": "10000000000000"
				}
			  }
			],
			"memo": "da67a7ce620cd8488761f9d2479fe58e729bda6b@10.10.5.7:26656",
			"timeout_height": "0",
			"extension_options": [],
			"non_critical_extension_options": []
		  },
		  "auth_info": {
			"signer_infos": [
			  {
				"public_key": {
				  "@type": "/cosmos.crypto.secp256k1.PubKey",
				  "key": "A8P2XDF3GuIF0lI3z4y3Mwy1kjOZ8u8YRNAnfpRVrzCN"
				},
				"mode_info": {
				  "single": {
					"mode": "SIGN_MODE_DIRECT"
				  }
				},
				"sequence": "0"
			  }
			],
			"fee": {
			  "amount": [],
			  "gas_limit": "200000",
			  "payer": "",
			  "granter": ""
			}
		  },
		  "signatures": [
			"JiHQ6dPidlXDrLL5CW5dY5WaTbBYsuac1h4VlgzPBUh36T245L1fY1TyUjJ42LcmQA1YtuDIzt/VVBTwWZyHdw=="
		  ]
		}
	  ]
	},
	"gov": {
	  "deposit_params": {
		"max_deposit_period": "43200s",
		"min_deposit": [
		  {
			"denom": "basetcro",
			"amount": "10000000"
		  }
		]
	  },
	  "deposits": [],
	  "proposals": [],
	  "starting_proposal_id": "1",
	  "tally_params": {
		"quorum": "0.334000000000000000",
		"threshold": "0.500000000000000000",
		"veto_threshold": "0.334000000000000000"
	  },
	  "votes": [],
	  "voting_params": {
		"voting_period": "43200s"
	  }
	},
	"mint": {
	  "minter": {
		"annual_provisions": "0.000000000000000000",
		"inflation": "0.013000000000000000"
	  },
	  "params": {
		"blocks_per_year": "6311520",
		"goal_bonded": "0.670000000000000000",
		"inflation_max": "0.020000000000000000",
		"inflation_min": "0.007000000000000000",
		"inflation_rate_change": "0.013000000000000000",
		"mint_denom": "basetcro"
	  }
	},
	"params": null,
	"slashing": {
	  "missed_blocks": [],
	  "params": {
		"downtime_jail_duration": "3600s",
		"min_signed_per_window": "0.500000000000000000",
		"signed_blocks_window": "2000",
		"slash_fraction_double_sign": "0.050000000000000000",
		"slash_fraction_downtime": "0.001"
	  },
	  "signing_infos": []
	},
	"staking": {
	  "delegations": [],
	  "exported": false,
	  "last_total_power": "0",
	  "last_validator_powers": [],
	  "params": {
		"bond_denom": "basetcro",
		"historical_entries": 100,
		"max_entries": 7,
		"max_validators": 150,
		"unbonding_time": "600s"
	  },
	  "redelegations": [],
	  "unbonding_delegations": [],
	  "validators": []
	},
	"upgrade": {}
  }
}`
