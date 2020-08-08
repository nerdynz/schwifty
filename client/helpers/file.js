function getBase64 (file) {
  let promise = new Promise((resolve, reject) => {
    let reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = function () {
      resolve(reader.result)
    }
    reader.onerror = function (error) {
      resolve(error)
    }
  })
  return promise
}

function getFileFromEvent (ev) {
  let file
  if (ev.dataTransfer.items) {
    // Use DataTransferItemList interface to access the file(s)
    for (let i = 0; i < ev.dataTransfer.items.length; i++) {
      // If dropped items aren't files, reject them
      if (ev.dataTransfer.items[i].kind === 'file') {
        file = ev.dataTransfer.items[i].getAsFile()
      }
    }
  } else {
    // Use DataTransfer interface to access the file(s)
    for (let i = 0; i < ev.dataTransfer.files.length; i++) {
      if (ev.dataTransfer.items[i].kind === 'file') {
        file = ev.dataTransfer.items[i].getAsFile()
      }
    }
  }
  return file
}

export { getBase64, getFileFromEvent }
