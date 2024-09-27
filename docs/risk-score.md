# Risk Score

Contact: **Spalen** (security team)

----

The risk score is a comprehensive evaluation system used to assess the security and reliability of Yearn vaults. It consists of multiple factors that contribute to the overall risk level of a vault. The risk score is crucial for users to understand the potential risks associated with investing in a particular vault.

## Risk Level

The risk level is an indicator of the overall security of the vault, calculated based on multiple factors. It ranges from 1 to 5, where 1 represents the highest security and 5 indicates the lowest. This score is derived from the individual risk scores described below.

## Risk Score Components

The risk score is composed of the following components:

1. **Review**: Evaluates the number of Sources of Trust (SoTs) involved in the strategy review process.
   - Score 1: 5 SoTs
   - Score 2: 4 SoTs
   - Score 3: 3 SoTs
   - Score 4: 2 SoTs
   - Score 5: 1 SoT
   SoTs include: internal strategist, peer reviews, expert peer reviews, ySec security reviews, and ySec recurring security reviews.

2. **Testing**: Assesses the testing coverage of the strategy.
   - Score 1: 95%+ coverage
   - Score 2: 90%+ coverage
   - Score 3: 80%+ coverage
   - Score 4: 70%+ coverage
   - Score 5: Below 70% coverage

3. **Complexity**: Measures the complexity of the strategy based on its source lines of code (sLOC).
   - Score 1: 0-150 sLOC
   - Score 2: 150-300 sLOC
   - Score 3: 300-450 sLOC
   - Score 4: 450-600 sLOC
   - Score 5: 750+ sLOC

4. **Risk Exposure**: Determines the potential for and frequency of losses in the strategy.
   - Score 1: No lossable cases, only gains
   - Score 2: Loss of funds up to 0-10% (e.g., deposit/withdrawal fees)
   - Score 3: Loss of funds up to 10-15% (e.g., protocol-specific IL exposure)
   - Score 4: Loss of funds up to 15-70% (e.g., adding liquidity to single-sided curve stable pools)
   - Score 5: Loss of funds up to 70-100% (e.g., leveraging cross assets and getting liquidated)

5. **Protocol Integration**: Counts the number of external protocols integrated into the strategy.
   - Score 1: Interacts with 1 external protocol
   - Score 2: Interacts with 2 external protocols
   - Score 3: Interacts with 3 external protocols
   - Score 4: Interacts with 4 external protocols
   - Score 5: Interacts with 5 or more external protocols

6. **Centralization Risk**: Evaluates the dependency on privileged roles and potential for centralized control.
   - Score 1: Fully permissionless, no dependency on privileged roles
   - Score 2: Has privileged roles but minimal risk of rug pulls
   - Score 3: Involves privileged roles with less frequent use and lower rug pull risk
   - Score 4: Frequent off-chain management with safeguards against rug pulls
   - Score 5: Heavy reliance on off-chain management, potential exposure to rug pulls

7. **External Protocol Audit**: Considers the number of audits conducted on the external protocols used.
   - Score 1: 4 or more audits by trusted firms or security researchers
   - Score 2: 3 audits
   - Score 3: 2 audits
   - Score 4: 1 audit
   - Score 5: No audits

8. **External Protocol Centralization**: Assesses the centralization of the external protocols used.
   - Score 1: Multisig with known trusted people and timelock, or governanceless/immutable contracts
   - Score 2: Multisig with known trusted people
   - Score 3: Multisig with known people but low threshold
   - Score 4: Multisig with unknown addresses or potential for partial harm
   - Score 5: EOA owner, unverified contracts, or high potential for complete harm

9. **External Protocol TVL**: Evaluates the Total Value Locked (TVL) in the external protocols.
   - Score 1: TVL of $480M or more
   - Score 2: TVL between $120M and $480M
   - Score 3: TVL between $40M and $120M
   - Score 4: TVL between $10M and $40M
   - Score 5: TVL of $10M or less

10. **External Protocol Longevity**: Considers how long the external protocol contracts have been deployed.
    - Score 1: 24 months or more
    - Score 2: Between 18 and 24 months
    - Score 3: Between 12 and 18 months
    - Score 4: Between 6 and 12 months
    - Score 5: Less than 6 months

11. **External Protocol Type**: Evaluates the purpose and nature of the external protocols used.
    - Score 1: Blue-chip protocols (e.g., AAVE, Compound, Uniswap, Curve, Convex, Balancer)
    - Score 2: Slightly modified forked blue-chip protocols
    - Score 3: Non-forked AMM lending/borrowing protocols, leveraged farming protocols, new concepts
    - Score 4: Cross-chain applications (bridges, yield aggregators, lending/borrowing)
    - Score 5: Protocols with main expertise in off-chain operations (e.g., RWA protocols)

## Updating Risk Scores

Risk scores are manually updated and stored in the vault metadata. To update a risk score:

1. Locate the vault's JSON file in the `data/meta/vaults/` directory.
2. Find the `metadata` object within the vault's JSON structure.
3. Update the `riskScore` object with the new values for each component.
4. If necessary, update the `riskLevel` field in the metadata object.
5. Optionally, add a `comment` field to the `riskScore` object to provide justification for the given risk scores.
6. Commit the changes to the repository and create a pull request for the update.
7. Notify @saltyfacu or @Majorfi to review, merge and deploy the changes.

