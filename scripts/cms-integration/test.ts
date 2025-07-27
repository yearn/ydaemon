import { expect as bunExpect, test as bunTest } from 'bun:test'

declare module 'bun:test' {
  interface Matchers<T> {
    toAlmostBe(expected: number, tolerance?: number): void
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
  }
})

export const expect = bunExpect
export const test = bunTest
