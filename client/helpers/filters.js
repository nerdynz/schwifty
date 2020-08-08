function byUUID (array, uuid) {
  let findUUID = function (obj) {
    return obj.UUID === uuid
  }
  return array.find(findUUID)
}

function indexByUUID (array, uuid) {
  let findUUID = function (obj) {
    return obj.UUID === uuid
  }
  return array.findIndex(findUUID)
}

function byULID (array, uuid) {
  let findULID = function (obj) {
    return obj.ULID === uuid
  }
  return array.find(findULID)
}

function indexByULID (array, ulid) {
  let findULID = function (obj) {
    return obj.ULID === ulid
  }
  return array.findIndex(findULID)
}

function indexByField (array, fieldName, value) {
  if (!array) {
    return null
  }
  let find = function (obj) {
    return obj[fieldName] === value
  }
  return array.findIndex(find)
}

function byField (array, fieldName, value) {
  if (!array) {
    return null
  }
  let find = function (obj) {
    return obj[fieldName] === value
  }
  return array.find(find)
}

function changeSortByULID (records, fromULID, toULID, property) {
  let from = indexByULID(records, fromULID)
  let to = indexByULID(records, toULID)
  return changeSort(records, from, to, property)
}

function changeSortByField (records, fieldName, fromVal, toVal, property, isReverse) {
  let from = indexByField(records, fieldName, fromVal)
  let to = indexByField(records, fieldName, toVal)
  return changeSort(records, from, to, property, isReverse)
}

function changeSort (records, from, to, property, isReverse) {
  let newRecords = arrayMove(records, from, to)
  newRecords.forEach((record, index) => {
    if (property) {
      if (record && Object.prototype.hasOwnProperty.call(record, property)) {
        record[property] = (isReverse ? newRecords.length - index : (index + 1)) * 50
        if (Object.prototype.hasOwnProperty.call(record, 'IsDirty')) {
          record.IsDirty = true
        }
      }
    } else if (record && Object.prototype.hasOwnProperty.call(record, 'SortPosition')) {
      record.SortPosition = (isReverse ? newRecords.length - index : (index + 1)) * 50
      if (Object.prototype.hasOwnProperty.call(record, 'IsDirty')) {
        record.IsDirty = true
      }
    }
  })
  return newRecords
}

function changeSortReverse (records, from, to, property) {
  return changeSort(records, to, from, property, true)
}

function getMaxSort (records, property) {
  let sort = 0
  records.forEach((record, index) => {
    let curSort = 0
    if (property) {
      if (Object.prototype.hasOwnProperty.call(record, property)) {
        curSort = record[property]
      }
    } else {
      curSort = record.SortPosition
    }

    if (curSort > sort) {
      sort = curSort
    }
  })
  return sort
}

export { byUUID, indexByUUID, indexByField, indexByULID, byULID, byField, changeSort, changeSortReverse, changeSortByField, changeSortByULID, getMaxSort }

function arrayMove (arr, from, to) {
  if (to >= arr.length) {
    let k = to - arr.length + 1
    while (k--) {
      arr.push(undefined)
    }
  }
  arr.splice(to, 0, arr.splice(from, 1)[0])
  return arr // for testing
}
