package main

import (
	"fmt"
	"os"
	
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/storage"
)

func main() {
	// Test the indexer storage functionality
	chainID := uint64(1)
	vault1 := common.HexToAddress("0x1234567890123456789012345678901234567890")
	vault2 := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	
	// Initialize storage
	storage.InitializeIndexerStorage(chainID)
	
	// Test getting empty value
	block1 := storage.GetLastIndexedBlock(chainID, vault1)
	fmt.Printf("Initial block for vault1: %d (should be 0)\n", block1)
	
	// Set some values
	storage.SetLastIndexedBlock(chainID, vault1, 18000000)
	storage.SetLastIndexedBlock(chainID, vault2, 17500000)
	
	// Read them back
	block1 = storage.GetLastIndexedBlock(chainID, vault1)
	block2 := storage.GetLastIndexedBlock(chainID, vault2)
	fmt.Printf("After setting - vault1: %d, vault2: %d\n", block1, block2)
	
	// Check if file was created
	if _, err := os.Stat("data/meta/indexer/1.json"); err == nil {
		fmt.Println("✓ JSON file created successfully")
		
		// Read file contents
		content, _ := os.ReadFile("data/meta/indexer/1.json")
		fmt.Println("File contents:")
		fmt.Println(string(content))
	} else {
		fmt.Println("✗ JSON file not found")
	}
	
	// Clean up test file
	os.Remove("data/meta/indexer/1.json")
	fmt.Println("\n✓ Test completed successfully")
}