module.exports = {
  testEnvironment: 'jsdom',
  setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  preset: 'ts-jest',
  moduleNameMapper: {
    // プロジェクトの実際のディレクトリ構造に合わせて調整
    '^@/(.*)$': '<rootDir>/$1', // または実際のパスに合わせて変更
  },
  transform: {
    '^.+\\.(js|jsx|ts|tsx)$': 'babel-jest',
  },
  testMatch: ['**/__tests__/**/*.+(ts|tsx|js)', '**/?(*.)+(spec|test).+(ts|tsx|js)'],
  roots: ['<rootDir>'],
};
