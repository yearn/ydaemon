# Retiring a Vault in yDaemon

Contact: **@dudesahn**, **@Schlagonia**, **@mil0xeth**

----

To retire a vault in yDaemon, you need to update the vault's metadata in the corresponding `vault.json` file. Follow these steps:

1. Locate the `vault.json` file:
   - Find the appropriate file in the `data/meta/vaults/{chainID}` directory.
   - The file name should match the vault's address (e.g., `0x1234...abcd.json`).

2. Open the `vault.json` file in a text editor.

3. Find the `metadata` section in the JSON structure.

4. Update the following fields:
   - Set `"isRetired": true`
   - Set `"hideAlways": true` (if you want to hide the vault from UI)

5. Save the changes to the `vault.json` file.

6. Commit the changes to the repository and create a pull request for the update.

7. Notify @saltyfacu or @Majorfi to review, merge and deploy the changes.
