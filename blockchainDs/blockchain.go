package blockchainDs

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) GetWholeChain() []*Block {
    return bc.blocks
}

func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.blocks[len(bc.blocks) - 1]
    newBlock := NewBlock(data, prevBlock.Hash)
    bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}
