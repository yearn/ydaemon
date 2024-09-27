# Adding a Vault to the Fancy Section on yearn.fi

Contact: **@dudesahn**, **@Schlagonia**, **@mil0xeth**

----

To add a vault to the Fancy section (highlighted vaults) on yearn.fi, follow these steps:

1. Locate the `vault.json` file:
   - Find the appropriate file in the `data/meta/vaults/{chainID}` directory.
   - The file name should match the vault's address (e.g., `0x1234...abcd.json`).

2. Open the `vault.json` file in a text editor.

3. Find the `metadata` section in the JSON structure.

4. Update or add the following field:
   - Set `"isHighlighted": true`

5. Example of the relevant part of the JSON structure:
   ```json
   {
     "metadata": {
       "isHighlighted": true,
       // ... other metadata fields ...
     }
   }
   ```

6. Save the changes to the `vault.json` file.

7. Commit the changes to the repository and create a pull request for the update.

8. Notify @saltyfacu or @Majorfi to review, merge and deploy the changes.

Note: Ensure that the vault meets all necessary criteria for inclusion in the Fancy section before making this change. The Fancy section is typically reserved for vaults that are particularly noteworthy or important, so consider carefully whether the vault qualifies for this designation.
