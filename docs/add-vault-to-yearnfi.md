# Adding a Vault to yearn.fi

Contact: **@dudesahn**, **@Schlagonia**, **@mil0xeth**

----

To add a new vault to yearn.fi, follow these steps:

1. Locate the `vault.json` file:
   - Find the appropriate file in the `data/meta/vaults/{chainID}` directory.
   - The file name should match the vault's address (e.g., `0x1234...abcd.json`).

2. Open the `vault.json` file in a text editor.

3. Find the `metadata` section in the JSON structure.

4. Update or add the following fields:
   - Ensure `"isRetired": false`
   - Ensure `"isHidden": false`
   - Under the `inclusion` object, set `"isYearn": true`

5. Example of the relevant part of the JSON structure:
   ```json
   {
     "metadata": {
       "isRetired": false,
       "isHidden": false,
       // ... other metadata fields ...
       "inclusion": {
         "isSet": true,
         "isYearn": true,
         // ... other inclusion fields ...
       }
       // ... other metadata fields ...
     }
   }
   ```

6. Save the changes to the `vault.json` file.

7. Commit the changes to the repository and create a pull request for the update.

8. Notify @saltyfacu or @Majorfi to review, merge and deploy the changes.

Note: Ensure that the vault meets all necessary criteria for inclusion in yearn.fi before making these changes. Also, review other metadata fields to ensure they are correctly set for the specific vault you're adding.

