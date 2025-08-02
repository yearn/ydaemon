export default function compareObjects(actual: any, expected: any, tolerance: Record<string, number> = {}, path = ''): { pass: boolean; message: string } {
  // Handle null/undefined
  if (actual === null || actual === undefined || expected === null || expected === undefined) {
    const pass = actual === expected
    return {
      pass,
      message: `expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
    }
  }

  // Handle arrays
  if (Array.isArray(actual) && Array.isArray(expected)) {
    if (actual.length !== expected.length) {
      return {
        pass: false,
        message: `expected array length ${actual.length} to equal ${expected.length} at path: ${path}`
      }
    }
    
    for (let i = 0; i < actual.length; i++) {
      const result = compareObjects(actual[i], expected[i], tolerance, `${path}[${i}]`)
      if (!result.pass) return result
    }
    return { pass: true, message: '' }
  }

  // Handle objects
  if (typeof actual === 'object' && typeof expected === 'object') {
    const actualKeys = Object.keys(actual)
    const expectedKeys = Object.keys(expected)
    
    if (actualKeys.length !== expectedKeys.length) {
      return {
        pass: false,
        message: `expected object with ${actualKeys.length} keys to have ${expectedKeys.length} keys at path: ${path}`
      }
    }

    for (const key of actualKeys) {
      if (!expectedKeys.includes(key)) {
        return {
          pass: false,
          message: `expected key '${key}' to exist in expected object at path: ${path}`
        }
      }
      
      const result = compareObjects(actual[key], expected[key], tolerance, path ? `${path}/${key}` : key)
      if (!result.pass) return result
    }
    return { pass: true, message: '' }
  }

  // Handle numbers with tolerance
  if (typeof actual === 'number' && typeof expected === 'number') {
    const toleranceValue = tolerance[path]
    if (toleranceValue !== undefined) {
      const delta = Math.abs(actual - expected)
      const pass = delta <= toleranceValue
      return {
        pass,
        message: `expected ${actual} to ${pass ? 'not ' : ''}be almost ${expected} ±${toleranceValue} (delta: ${delta}) at path: ${path}`
      }
    } else {
      const pass = actual === expected
      return {
        pass,
        message: `expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
      }
    }
  }

  // Handle other types
  const pass = actual === expected
  return {
    pass,
    message: `expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
  }
}
