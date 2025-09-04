#!/usr/bin/env bun

/**
 * Simple script to compare strategy lists between Kong-powered yDaemon and production
 * Usage: bun scripts/kong-integration/compare-strategies.ts [chainId] [localPort]
 */

const chainId = process.argv[2] || "all";
const localPort = process.argv[3] || "3001";

// Supported chains from daemon logs
const supportedChains = ["1", "10", "42161", "100", "137", "146", "250", "8453", "747474"];

if (chainId === "all") {
  console.log(`🔍 Comparing strategies for ALL chains`);
  console.log(`📍 Local port: ${localPort}`);
  console.log(`📍 Testing chains: ${supportedChains.join(", ")}`);
  console.log();
} else {
  const localUrl = `http://localhost:${localPort}/${chainId}/strategies/all?limit=1000`;
  const prodUrl = `https://ydaemon.yearn.fi/${chainId}/strategies/all?limit=1000`;
  
  console.log(`🔍 Comparing strategies for chain ${chainId}`);
  console.log(`📍 Local (Kong):  ${localUrl}`);
  console.log(`📍 Production:     ${prodUrl}`);
  console.log();
}

async function fetchStrategies(url: string): Promise<string[]> {
  const response = await fetch(url);
  const strategies = await response.json();
  return strategies.map((s: any) => s.address.toLowerCase()).sort();
}

async function checkKongDirectly(strategyAddress: string): Promise<boolean> {
  const query = `{
    strategies(yearn: true) {
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
  return data.data.strategies.some((s: any) => s.address.toLowerCase() === strategyAddress.toLowerCase());
}

async function compareChain(chain: string) {
  const localUrl = `http://localhost:${localPort}/${chain}/strategies/all?limit=1000`;
  const prodUrl = `https://ydaemon.yearn.fi/${chain}/strategies/all?limit=1000`;
  
  console.log(`\n=== Chain ${chain} ===`);
  console.log(`📍 Local:  ${localUrl}`);
  console.log(`📍 Prod:   ${prodUrl}`);
  
  try {
    const [kongStrategies, prodStrategies] = await Promise.all([
      fetchStrategies(localUrl),
      fetchStrategies(prodUrl),
    ]);

    const kongSet = new Set(kongStrategies);
    const prodSet = new Set(prodStrategies);
    
    const missingFromKong = prodStrategies.filter(s => !kongSet.has(s));
    const missingFromProd = kongStrategies.filter(s => !prodSet.has(s));
    const common = kongStrategies.filter(s => prodSet.has(s));

    console.log(`📊 Kong: ${kongStrategies.length}, Prod: ${prodStrategies.length}, Common: ${common.length}`);
    
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
      kongCount: kongStrategies.length,
      prodCount: prodStrategies.length,
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
    
    let reportLines = [`Strategy Comparison Report - ${new Date().toISOString()}`, "", "Chain | Kong | Prod | Common | Missing Kong | Missing Prod | Status", "------|------|------|--------|--------------|--------------|-------"];
    
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
            const localUrl = `http://localhost:${localPort}/${result.chain}/strategies/all?limit=1000`;
            const prodUrl = `https://ydaemon.yearn.fi/${result.chain}/strategies/all?limit=1000`;
            const [kongStrategies, prodStrategies] = await Promise.all([
              fetchStrategies(localUrl),
              fetchStrategies(prodUrl),
            ]);
            const kongSet = new Set(kongStrategies);
            const missingFromKong = prodStrategies.filter(s => !kongSet.has(s));
            missingFromKong.forEach(addr => reportLines.push(`    ${addr}`));
          } catch (e) {
            reportLines.push(`    Error fetching details: ${e}`);
          }
        }
        
        if (result.missingFromProd > 0) {
          reportLines.push(`  Missing from Production (${result.missingFromProd}):`);
          try {
            const localUrl = `http://localhost:${localPort}/${result.chain}/strategies/all?limit=1000`;
            const prodUrl = `https://ydaemon.yearn.fi/${result.chain}/strategies/all?limit=1000`;
            const [kongStrategies, prodStrategies] = await Promise.all([
              fetchStrategies(localUrl),
              fetchStrategies(prodUrl),
            ]);
            const prodSet = new Set(prodStrategies);
            const missingFromProd = kongStrategies.filter(s => !prodSet.has(s));
            missingFromProd.forEach(addr => reportLines.push(`    ${addr}`));
          } catch (e) {
            reportLines.push(`    Error fetching details: ${e}`);
          }
        }
      }
    }
    
    // Save report to file
    const reportPath = "/home/murdertxxth/git/ydaemon/scripts/kong-integration/compare-strategies.txt";
    await Bun.write(reportPath, reportLines.join("\n"));
    console.log(`\n📄 Report saved: ${reportPath}`);
    
  } else {
    // Single chain mode
    const localUrl = `http://localhost:${localPort}/${chainId}/strategies/all?limit=1000`;
    const prodUrl = `https://ydaemon.yearn.fi/${chainId}/strategies/all?limit=1000`;
    
    console.log("📥 Fetching strategy lists...");
    
    const [kongStrategies, prodStrategies] = await Promise.all([
      fetchStrategies(localUrl),
      fetchStrategies(prodUrl),
    ]);

    console.log("📊 Strategy counts:");
    console.log(`   Kong (local):  ${kongStrategies.length} strategies`);
    console.log(`   Production:    ${prodStrategies.length} strategies`);
    console.log();

    // Find differences
    const kongSet = new Set(kongStrategies);
    const prodSet = new Set(prodStrategies);
    
    const missingFromKong = prodStrategies.filter(s => !kongSet.has(s));
    const missingFromProd = kongStrategies.filter(s => !prodSet.has(s));
    const common = kongStrategies.filter(s => prodSet.has(s));

    console.log("📈 Differences:");
    console.log(`   Common strategies:        ${common.length}`);
    console.log(`   Missing from Kong:        ${missingFromKong.length}`);
    console.log(`   Missing from Prod:        ${missingFromProd.length}`);
    console.log();

    if (missingFromKong.length > 0) {
      console.log("🔻 Strategies in production but missing from Kong:");
      missingFromKong.slice(0, 10).forEach(s => console.log(`   ${s}`));
      if (missingFromKong.length > 10) console.log(`   ... and ${missingFromKong.length - 10} more`);
      console.log();
    }

    if (missingFromProd.length > 0) {
      console.log("🔺 Strategies in Kong but missing from production:");
      missingFromProd.slice(0, 10).forEach(s => console.log(`   ${s}`));
      if (missingFromProd.length > 10) console.log(`   ... and ${missingFromProd.length - 10} more`);
      console.log();
    }

    console.log();
    console.log("✅ Comparison complete!");
    
    // Save to files for further analysis
    await Bun.write(`/tmp/kong_strategies_${chainId}.txt`, kongStrategies.join("\n"));
    await Bun.write(`/tmp/prod_strategies_${chainId}.txt`, prodStrategies.join("\n"));
    console.log(`📁 Files saved:`);
    console.log(`   Kong strategies:     /tmp/kong_strategies_${chainId}.txt`);
    console.log(`   Production strategies: /tmp/prod_strategies_${chainId}.txt`);
  }
}

main().catch(console.error);