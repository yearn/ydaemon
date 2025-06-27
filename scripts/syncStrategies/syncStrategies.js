    /**
     * This script synchronizes strategy states between fresh and stale directories. 
     * in PR #496, the strategies .json files were deleted and recreated by running yDaemon. This did not catch older strategies,
     * so I needed to create this script to update the stale strategies file and replace all that matched with the fresh strategies.
     * It is easiest to do this with 2 worktrees to create the new .json files and still keep the old ones. 
     * Run this script with the old json files in a new folder called "stale", and the new ones in a folder called "fresh". 
     * New files will be created in the refreshed directory with the same name as the stale files, but with the updated strategy states.
     *
     * @type {any}
     * @constant
     * @name syncStrateges
     * @description The parsed JSON object read from the file specified by `freshPath`.
     */

const fs = require('fs')
const path = require('path')

const freshDir = path.join(__dirname, 'fresh')
const staleDir = path.join(__dirname, 'stale')
const refreshedDir = path.join(__dirname, 'refreshed')

if (!fs.existsSync(refreshedDir)) {
  fs.mkdirSync(refreshedDir)
}

fs.readdirSync(freshDir).forEach((file) => {
  const freshPath = path.join(freshDir, file)
  const stalePath = path.join(staleDir, file)
  if (fs.existsSync(stalePath)) {
    const fresh = JSON.parse(fs.readFileSync(freshPath, 'utf8'))
    const stale = JSON.parse(fs.readFileSync(stalePath, 'utf8'))
    if (fresh.strategies && stale.strategies) {
      for (const key of Object.keys(stale.strategies)) {
        if (fresh.strategies[key]) {
          const src = fresh.strategies[key]
          stale.strategies[key].isActive = src.isActive
          stale.strategies[key].isInQueue = src.isInQueue
          stale.strategies[key].isRetired = src.isRetired
          stale.strategies[key].status = src.status
        }
      }
    }
    const outPath = path.join(refreshedDir, file)
    fs.writeFileSync(outPath, JSON.stringify(stale, null, 2))
    console.log(`${outPath} created!`)
  }
})
