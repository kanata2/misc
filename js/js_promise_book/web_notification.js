function notifyMessage(msg, opts, callback) {
  if (typeof Notification === 'undefined') {
    callback(new Error('doesn\'t support Notification API'))
    return
  }
  if (Notification.permission === 'granted') {
    const notification = new Notification(msg, opts)
    callback(null, notification)
  }
  else {
    Notification.requestPermission(status => {
      if (Notification.permission !== status) {
        Notification.permission = status
      }
      if (status === 'granted') {
        const notification = new Notification(msg, opts)
        callback(null, notification)
      }
      else {
        callback(new Error('user denied'))
      }
    })
  }
}

function notifyMessageAsPromise(msg, opts) {
  return new Promise((resolve, reject) => {
    notifyMessage(msg, opts, (error, notification) => {
      if (error) {
        reject(error)
      } else {
        resolve(notification)
      }
    })
  })
}

notifyMessageAsPromise("Hi!").then(noti => {
  console.log(noti)
}).catch(error => {
  console.error(error)
})

