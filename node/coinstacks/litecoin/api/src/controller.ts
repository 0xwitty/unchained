import { bech32 } from 'bech32'
import { Blockbook } from '@shapeshiftoss/blockbook'
import { Service } from '../../../common/api/src/utxo/service'
import { UTXO } from '../../../common/api/src/utxo/controller'
import { Logger } from '@shapeshiftoss/logger'

const INDEXER_URL = process.env.INDEXER_URL
const INDEXER_WS_URL = process.env.INDEXER_WS_URL
const INDEXER_API_KEY = process.env.INDEXER_API_KEY

if (!INDEXER_URL) throw new Error('INDEXER_URL env var not set')
if (!INDEXER_WS_URL) throw new Error('INDEXER_WS_URL env var not set')

export const logger = new Logger({
  namespace: ['unchained', 'coinstacks', 'litecoin', 'api'],
  level: process.env.LOG_LEVEL,
})

const blockbook = new Blockbook({ httpURL: INDEXER_URL, wsURL: INDEXER_WS_URL, apiKey: INDEXER_API_KEY, logger })

const isXpub = (pubkey: string): boolean => {
  return pubkey.startsWith('Ltub') || pubkey.startsWith('Mtub') || pubkey.startsWith('zpub')
}

export const formatAddress = (address: string): string => {
  if (bech32.decodeUnsafe(address.toLowerCase())?.prefix === 'ltc') return address.toLowerCase()

  return address
}

export const service = new Service({ blockbook, isXpub, addressFormatter: formatAddress })

// assign service to be used for all instances of UTXO
UTXO.service = service
