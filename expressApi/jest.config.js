/** @type {import('ts-jest').JestConfigWithTsJest} **/
export default {
  testEnvironment: "node",
  verbose: false,
  collectCoverage: true,
  coverageReporters: ["html"],
  transform: {
    "^.+.tsx?$": ["ts-jest", {}],
  },
};