const assert = require('assert');
it('should use `done` for test', done => {
  setTimeout(() => {
    assert(true);
    done();
  }, 0);
});

it('should use `done` for test?', done => {
  const promise = Promise.resolve(42);
  promise.then(value => {
    assert(true)
  }).then(done, done)  
})

describe('Promise', () => {
  it('should return a promise object', () => {
    const promise = Promise.resolve(42)
    return promise.then(value => { assert(value === 42) })
  })
})

function shouldRejected(promise) {
  return {
    'catch': (fn) => promise.then(() => {
      throw new Error('Expected promise to be rejected but it was fullfiled')
    }, reason => {
      fn.call(promise, reason)
    })
  }
}
it('should be rejected', () => {
  const promise = Promise.reject(new Error('human error'))
  return shouldRejected(promise).catch((error) => {
    assert(error.message === 'human error')
  })
})

