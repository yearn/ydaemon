export default function deepCompareObjects(actual: any, expected: any, options: { tolerance?: Record<string, number>, message?: string } = {}, path = ''): { pass: boolean; message: string } {
  const { tolerance = {}, message } = options

  // Handle null/undefined
  if (actual === null || actual === undefined || expected === null || expected === undefined) {
    const pass = actual === expected
    return {
      pass,
      message: `${message ? message + '\n' : ''}expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
    }
  }

  // Handle arrays
  if (Array.isArray(actual) && Array.isArray(expected)) {
    if (actual.length !== expected.length) {
      return {
        pass: false,
        message: `${message ? message + '\n' : ''}expected array length ${actual.length} to equal ${expected.length} at path: ${path}`
      }
    }
    
    for (let i = 0; i < actual.length; i++) {
      const result = deepCompareObjects(actual[i], expected[i], options, `${path}[${i}]`)
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
        message: `${message ? message + '\n' : ''}expected object with ${actualKeys.length} keys to have ${expectedKeys.length} keys at path: ${path}`
      }
    }

    for (const key of actualKeys) {
      if (!expectedKeys.includes(key)) {
        return {
          pass: false,
          message: `${message ? message + '\n' : ''}expected key '${key}' to exist in expected object at path: ${path}`
        }
      }
      
      const result = deepCompareObjects(actual[key], expected[key], options, path ? `${path}/${key}` : key)
      if (!result.pass) return result
    }
    return { pass: true, message: '' }
  }

  // Handle numbers with percentage tolerance
  if ((typeof actual === 'number' || typeof actual === 'bigint') && 
      (typeof expected === 'number' || typeof expected === 'bigint')) {
    const toleranceDecimal = tolerance[path]
    if (toleranceDecimal !== undefined) {
      // Convert to numbers for tolerance comparison
      const actualNum = Number(actual)
      const expectedNum = Number(expected)
      const delta = Math.abs(actualNum - expectedNum)
      const toleranceValue = Math.abs(expectedNum) * toleranceDecimal
      const pass = delta <= toleranceValue
      const percentageDelta = expectedNum !== 0 ? (delta / Math.abs(expectedNum)) * 100 : 0
      return {
        pass,
        message: `${message ? message + '\n' : ''}expected ${actual} to ${pass ? 'not ' : ''}be almost ${expected} ±${(toleranceDecimal * 100).toFixed(1)}% (delta: ${percentageDelta.toFixed(2)}%) at path: ${path}`
      }
    } else {
      const pass = actual === expected
      return {
        pass,
        message: `${message ? message + '\n' : ''}expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
      }
    }
  }

  // Handle other types
  const pass = actual === expected
  return {
    pass,
    message: `${message ? message + '\n' : ''}expected ${actual} to ${pass ? 'not ' : ''}equal ${expected} at path: ${path}`
  }
}
