// We only need Bun to test, here's an easy way to extend bun with our own matchers.
// This file provides custom Jest-like matchers (like toAlmostBe) that work with Bun's test runner.

import { expect as bunExpect, test as bunTest } from 'bun:test'
import compareObjects from './compareObjects'

declare module 'bun:test' {
  interface Matchers<T> {
    toAlmostBe(expected: number, tolerance?: number): void
    toAlmostEqual(expected: any, options?: { tolerance?: Record<string, number> }): void
  }
}

bunExpect.extend({
  toAlmostBe(received: unknown, expected: number, tolerance = 1e-6) {
    if (typeof received !== 'number' || typeof expected !== 'number') {
      return {
        pass: false,
        message: () => `expected both values to be numbers, got ${typeof received} and ${typeof expected}`
      }
    }

    const pass = Math.abs(received - expected) <= tolerance
    return {
      pass,
      message: () =>
        `expected ${received} to ${pass ? 'not ' : ''}be almost ${expected} ±${tolerance}`
    }
  },

  toAlmostEqual(received: unknown, expected: any, options: { tolerance?: Record<string, number> } = {}) {
    const { tolerance = {} } = options
    
    if (typeof received !== 'object' || typeof expected !== 'object') {
      return {
        pass: false,
        message: () => `expected both values to be objects, got ${typeof received} and ${typeof expected}`
      }
    }

    return compareObjects(received, expected, tolerance)
  }
})

export const expect = bunExpect
export const test = bunTest