#!/usr/bin/env bun

/**
 * Simple script to compare vault lists between Kong-powered yDaemon and production
 * Usage: bun scripts/kong-integration/compare-vaults.ts [chainId] [localPort]
 */

const chainId = process.argv[2] || "all";
const localPort = process.argv[3] || "3001";

// Supported chains from daemon logs
const supportedChains = ["1", "10", "42161", "100", "137", "146", "250", "8453", "747474"];

if (chainId === "all") {
  console.log(`🔍 Comparing vaults for ALL chains`);
  console.log(`📍 Local port: ${localPort}`);
  console.log(`📍 Testing chains: ${supportedChains.join(", ")}`);
  console.log();
} else {
  const localUrl = `http://localhost:${localPort}/${chainId}/vaults/all?limit=1000`;
  const prodUrl = `https://ydaemon.yearn.fi/${chainId}/vaults/all?limit=1000`;
  
  console.log(`🔍 Comparing vaults for chain ${chainId}`);
  console.log(`📍 Local (Kong):  ${localUrl}`);
  console.log(`📍 Production:     ${prodUrl}`);
  console.log();
}

async function fetchVaults(url: string): Promise<string[]> {
  const response = await fetch(url);
  const vaults = await response.json();
  return vaults.map((v: any) => v.address.toLowerCase()).sort();
}

async function checkKongDirectly(vaultAddress: string): Promise<boolean> {
  const query = `{
    vaults(yearn: true) {
      address
      chainId
    }
  }`;

  const response = await fetch("https://kong.yearn.farm/api/gql", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query }),
  });

  const data = await response.json();
  return data.data.vaults.some((v: any) => v.address.toLowerCase() === vaultAddress.toLowerCase());
}

async function compareChain(chain: string) {
  const localUrl = `http://localhost:${localPort}/${chain}/vaults/all?limit=1000`;
  const prodUrl = `https://ydaemon.yearn.fi/${chain}/vaults/all?limit=1000`;
  
  console.log(`\n=== Chain ${chain} ===`);
  console.log(`📍 Local:  ${localUrl}`);
  console.log(`📍 Prod:   ${prodUrl}`);
  
  try {
    const [kongVaults, prodVaults] = await Promise.all([
      fetchVaults(localUrl),
      fetchVaults(prodUrl),
    ]);

    const kongSet = new Set(kongVaults);
    const prodSet = new Set(prodVaults);
    
    const missingFromKong = prodVaults.filter(v => !kongSet.has(v));
    const missingFromProd = kongVaults.filter(v => !prodSet.has(v));
    const common = kongVaults.filter(v => prodSet.has(v));

    console.log(`📊 Kong: ${kongVaults.length}, Prod: ${prodVaults.length}, Common: ${common.length}`);
    
    if (missingFromKong.length > 0) {
      console.log(`❌ ${missingFromKong.length} missing from Kong`);
    }
    
    if (missingFromProd.length > 0) {
      console.log(`❌ ${missingFromProd.length} missing from Prod`);
    }
    
    if (missingFromKong.length === 0 && missingFromProd.length === 0) {
      console.log(`✅ Perfect match!`);
    }
    
    return {
      chain,
      kongCount: kongVaults.length,
      prodCount: prodVaults.length,
      commonCount: common.length,
      missingFromKong: missingFromKong.length,
      missingFromProd: missingFromProd.length,
      success: true
    };
  } catch (error) {
    console.log(`❌ Error: ${error}`);
    return {
      chain,
      kongCount: 0,
      prodCount: 0,
      commonCount: 0,
      missingFromKong: 0,
      missingFromProd: 0,
      success: false,
      error: error.toString()
    };
  }
}

async function main() {
  if (chainId === "all") {
    console.log("📥 Testing all chains...");
    
    const results = [];
    for (const chain of supportedChains) {
      const result = await compareChain(chain);
      results.push(result);
    }
    
    console.log("\n🏁 Summary:");
    console.log("Chain | Kong | Prod | Common | Missing Kong | Missing Prod | Status");
    console.log("------|------|------|--------|--------------|--------------|-------");
    
    let reportLines = [`Vault Comparison Report - ${new Date().toISOString()}`, "", "Chain | Kong | Prod | Common | Missing Kong | Missing Prod | Status", "------|------|------|--------|--------------|--------------|-------"];
    
    for (const result of results) {
      const status = result.success ? 
        (result.missingFromKong === 0 && result.missingFromProd === 0 ? "✅ Match" : "⚠️  Diff") : 
        "❌ Error";
      const line = `${result.chain.padEnd(5)} | ${result.kongCount.toString().padEnd(4)} | ${result.prodCount.toString().padEnd(4)} | ${result.commonCount.toString().padEnd(6)} | ${result.missingFromKong.toString().padEnd(12)} | ${result.missingFromProd.toString().padEnd(12)} | ${status}`;
      console.log(line);
      reportLines.push(line);
    }
    
    // Add detailed differences section
    reportLines.push("", "=".repeat(80), "DETAILED DIFFERENCES BY CHAIN", "=".repeat(80));
    
    for (const result of results) {
      if (result.missingFromKong > 0 || result.missingFromProd > 0) {
        reportLines.push(``, `Chain ${result.chain}:`);
        
        if (result.missingFromKong > 0) {
          reportLines.push(`  Missing from Kong (${result.missingFromKong}):`);
          // We need to re-fetch to get the actual addresses
          try {
            const localUrl = `http://localhost:${localPort}/${result.chain}/vaults/all?limit=1000`;
            const prodUrl = `https://ydaemon.yearn.fi/${result.chain}/vaults/all?limit=1000`;
            const [kongVaults, prodVaults] = await Promise.all([
              fetchVaults(localUrl),
              fetchVaults(prodUrl),
            ]);
            const kongSet = new Set(kongVaults);
            const missingFromKong = prodVaults.filter(v => !kongSet.has(v));
            missingFromKong.forEach(addr => reportLines.push(`    ${addr}`));
          } catch (e) {
            reportLines.push(`    Error fetching details: ${e}`);
          }
        }
        
        if (result.missingFromProd > 0) {
          reportLines.push(`  Missing from Production (${result.missingFromProd}):`);
          try {
            const localUrl = `http://localhost:${localPort}/${result.chain}/vaults/all?limit=1000`;
            const prodUrl = `https://ydaemon.yearn.fi/${result.chain}/vaults/all?limit=1000`;
            const [kongVaults, prodVaults] = await Promise.all([
              fetchVaults(localUrl),
              fetchVaults(prodUrl),
            ]);
            const prodSet = new Set(prodVaults);
            const missingFromProd = kongVaults.filter(v => !prodSet.has(v));
            missingFromProd.forEach(addr => reportLines.push(`    ${addr}`));
          } catch (e) {
            reportLines.push(`    Error fetching details: ${e}`);
          }
        }
      }
    }
    
    // Save report to file
    const reportPath = "/home/murdertxxth/git/ydaemon/scripts/kong-integration/compare-vaults.txt";
    await Bun.write(reportPath, reportLines.join("\n"));
    console.log(`\n📄 Report saved: ${reportPath}`);
    
  } else {
    // Single chain mode
    const localUrl = `http://localhost:${localPort}/${chainId}/vaults/all?limit=1000`;
    const prodUrl = `https://ydaemon.yearn.fi/${chainId}/vaults/all?limit=1000`;
    
    console.log("📥 Fetching vault lists...");
    
    const [kongVaults, prodVaults] = await Promise.all([
      fetchVaults(localUrl),
      fetchVaults(prodUrl),
    ]);

    console.log("📊 Vault counts:");
    console.log(`   Kong (local):  ${kongVaults.length} vaults`);
    console.log(`   Production:    ${prodVaults.length} vaults`);
    console.log();

    // Find differences
    const kongSet = new Set(kongVaults);
    const prodSet = new Set(prodVaults);
    
    const missingFromKong = prodVaults.filter(v => !kongSet.has(v));
    const missingFromProd = kongVaults.filter(v => !prodSet.has(v));
    const common = kongVaults.filter(v => prodSet.has(v));

    console.log("📈 Differences:");
    console.log(`   Common vaults:        ${common.length}`);
    console.log(`   Missing from Kong:    ${missingFromKong.length}`);
    console.log(`   Missing from Prod:    ${missingFromProd.length}`);
    console.log();

    if (missingFromKong.length > 0) {
      console.log("🔻 Vaults in production but missing from Kong:");
      missingFromKong.slice(0, 10).forEach(v => console.log(`   ${v}`));
      if (missingFromKong.length > 10) console.log(`   ... and ${missingFromKong.length - 10} more`);
      console.log();
    }

    if (missingFromProd.length > 0) {
      console.log("🔺 Vaults in Kong but missing from production:");
      missingFromProd.slice(0, 10).forEach(v => console.log(`   ${v}`));
      if (missingFromProd.length > 10) console.log(`   ... and ${missingFromProd.length - 10} more`);
      console.log();
    }

    // Test specific vault mentioned in issue (only for Ethereum)
    if (chainId === "1") {
      const testVault = "0x00CB87656196dD835b9E4d67018aE0477a1De8C1";
      console.log(`🧪 Testing specific vault: ${testVault}`);
      
      const inKong = kongSet.has(testVault.toLowerCase());
      const inProd = prodSet.has(testVault.toLowerCase());
      
      console.log(`   ${inKong ? "✅" : "❌"} ${inKong ? "Found" : "NOT found"} in Kong`);
      console.log(`   ${inProd ? "✅" : "❌"} ${inProd ? "Found" : "NOT found"} in production`);

      // Check directly in Kong GraphQL
      console.log();
      console.log("🔍 Checking Kong GraphQL directly for test vault...");
      const inKongDirect = await checkKongDirectly(testVault);
      console.log(`   ${inKongDirect ? "✅" : "❌"} ${inKongDirect ? "Found" : "NOT found"} in Kong GraphQL directly`);
    }

    console.log();
    console.log("✅ Comparison complete!");
    
    // Save to files for further analysis
    await Bun.write(`/tmp/kong_vaults_${chainId}.txt`, kongVaults.join("\n"));
    await Bun.write(`/tmp/prod_vaults_${chainId}.txt`, prodVaults.join("\n"));
    console.log(`📁 Files saved:`);
    console.log(`   Kong vaults:     /tmp/kong_vaults_${chainId}.txt`);
    console.log(`   Production vaults: /tmp/prod_vaults_${chainId}.txt`);
  }
}

main().catch(console.error);