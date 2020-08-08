<template>
  <div class="dp-range-date override" :class="{'is-range': isRangePicker, 'is-open': isPickerActive}"><!-- .dp-range-date.o-h override hook for more explicit css -->
    <div v-show="isPickerActive" class="dp-click-overlay" @click="closePicker" />
    <div class="dp-field field has-addons">
      <div class="control">
        <input ref="inp" class="input dp-input " @blur="parseDate">
      </div>
      <slot name="buttons">
        <div class="control">
          <button class="dp-open-button button is-primary" @click="togglePicker">
            <i class="fa fa-calendar" />
          </button>
        </div>
      </slot>
    </div>
    <transition>
      <div v-show="isPickerActive" class="dp-picker box">
        <div class="dp-month-year field has-addons">
          <div class="control">
            <button class="button dp-next-prev">
              <i class="fa fa-chevron-left" />
            </button>
          </div>
          <div class="control">
            <div class="select">
              <select v-model="selectedMonth">
                <option :value="0">
                  Jan
                </option>
                <option :value="1">
                  Feb
                </option>
                <option :value="2">
                  Mar
                </option>
                <option :value="3">
                  Apr
                </option>
                <option :value="4">
                  May
                </option>
                <option :value="5">
                  Jun
                </option>
                <option :value="6">
                  Jul
                </option>
                <option :value="7">
                  Aug
                </option>
                <option :value="8">
                  Sept
                </option>
                <option :value="9">
                  Oct
                </option>
                <option :value="10">
                  Nov
                </option>
                <option :value="11">
                  Dec
                </option>
              </select>
            </div>
          </div>
          <div class="control">
            <div class="select">
              <select v-model="selectedYear">
                <option v-for="(year, index) in yearRange" :key="index" :value="year">{{year}}</option>
              </select>
            </div>
          </div>
          <div class="control">
            <button class="button dp-next-prev"><i class="fa fa-chevron-right"></i></button>
          </div>
        </div>
        <div class="dp-days">
          <div class="dp-day-row dp-day-header">
            <div class="dp-day dp-day-text" v-for="(dayText, index) in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" :key="index">
              {{dayText}}
            </div>
          </div>
          <div class="dp-day-row" v-for="(dayrow, index) in dayRows" :key="index">
            <div 
              class="dp-day"
              @click="dateClicked(day)"
              @mouseover="updateHoverDate(day)"
              :class="{
                'is-diff-month': day.date.getMonth() !== selectedMonth,
                'is-hovered': hoverDate != null && innerFrom != null && day.date.toLocaleDateString() > innerFrom && day.date.toLocaleDateString() < hoverDate, 
                'is-active-one': innerFrom != null && day.date.toLocaleDateString() === innerFrom.toLocaleDateString(), 
                'is-active-two': innerTo != null && day.date.toLocaleDateString() === innerTo.toLocaleDateString(),
                'is-active-between': innerFrom != null && innerTo != null && day.date.toLocaleDateString() > innerFrom && day.date.toLocaleDateString() < innerTo
                }"
              v-for="(day, index) in dayrow" :key="index">
              {{day.date.getDate()}}
            </div>
          </div>
        </div>
        <div v-if="type === 'datetime'" class="dp-month-year dp-field field has-addons">
          <div class="control">
            <div class="select">
              <select v-model="selectedHour" @change="setTimeOnFrom">
                <option :value="1">1</option>
                <option :value="2">2</option>
                <option :value="3">3</option>
                <option :value="4">4</option>
                <option :value="5">5</option>
                <option :value="6">6</option>
                <option :value="7">7</option>
                <option :value="8">8</option>
                <option :value="9">9</option>
                <option :value="10">10</option>
                <option :value="11">11</option>
                <option :value="12">12</option>
              </select>
            </div>
          </div>
          <div class="control">
            <div class="select">
              <select v-model="selectedMinute" @change="setTimeOnFrom">
                <option :value="0">00</option>
                <option :value="1">01</option>
                <option :value="2">02</option>
                <option :value="3">03</option>
                <option :value="4">04</option>
                <option :value="5">05</option>
                <option :value="6">06</option>
                <option :value="7">07</option>
                <option :value="8">08</option>
                <option :value="9">09</option>
                <option :value="10">10</option>
                <option :value="11">11</option>
                <option :value="12">12</option>
                <option :value="13">13</option>
                <option :value="14">14</option>
                <option :value="15">15</option>
                <option :value="16">16</option>
                <option :value="17">17</option>
                <option :value="18">18</option>
                <option :value="19">19</option>
                <option :value="20">20</option>
                <option :value="21">21</option>
                <option :value="22">22</option>
                <option :value="23">23</option>
                <option :value="24">24</option>
                <option :value="25">25</option>
                <option :value="26">26</option>
                <option :value="27">27</option>
                <option :value="28">28</option>
                <option :value="29">29</option>
                <option :value="30">30</option>
                <option :value="31">31</option>
                <option :value="32">32</option>
                <option :value="33">33</option>
                <option :value="34">34</option>
                <option :value="35">35</option>
                <option :value="36">36</option>
                <option :value="37">37</option>
                <option :value="38">38</option>
                <option :value="39">39</option>
                <option :value="40">40</option>
                <option :value="41">41</option>
                <option :value="42">42</option>
                <option :value="43">43</option>
                <option :value="44">44</option>
                <option :value="45">45</option>
                <option :value="46">46</option>
                <option :value="47">47</option>
                <option :value="48">48</option>
                <option :value="49">49</option>
                <option :value="50">50</option>
                <option :value="51">51</option>
                <option :value="52">52</option>
                <option :value="53">53</option>
                <option :value="54">54</option>
                <option :value="55">55</option>
                <option :value="56">56</option>
                <option :value="57">57</option>
                <option :value="58">58</option>
                <option :value="59">59</option>
              </select>
            </div>
          </div>
          <div class="control">
            <div class="select">
              <select v-model="selectedAMPM" @change="setTimeOnFrom">
                <option :value="0">AM</option>
                <option :value="12">PM</option>
              </select>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script>
import parseTime from 'parsetime'
const monthNames = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']

function rdGetDOY (now) {
  let start = new Date(now.getFullYear(), 0, 0)
  let diff = (now - start) + ((start.getTimezoneOffset() - now.getTimezoneOffset()) * 60 * 1000)
  let oneDay = 1000 * 60 * 60 * 24
  let day = Math.floor(diff / oneDay)
  return day
}

function rdFmt (date) {
  let day = date.getDate()
  let monthIndex = date.getMonth()
  let year = date.getFullYear()
  return '' + day + ' ' + monthNames[monthIndex] + ' ' + year
}

function rdFmtTime (date) {
  let hours = date.getHours()
  let minutes = date.getMinutes()
  let amPm = 'AM'
  if (hours > 11) {
    amPm = 'PM'
  }
  if (hours === 0) {
    hours = 12
  }
  if (hours > 12) {
    hours = hours - 12
  }
  if (minutes < 10) {
    minutes = '0' + minutes
  }
  return hours + ':' + minutes + ' ' + amPm
}

export default {
  // COMPONENT
  // ______________________________________
  name: 'RangeDatePicker',
  props: {
    type: {
      type: String,
      required: true,
      validator: function (value) {
        // The value must match one of these strings
        return ['date', 'datetime', 'rangedate'].indexOf(value) !== -1
      }
    },
    cssPrefix: {
      type: String,
      default: ''
    },
    startDate: Date,
    endDate: Date,
    value: {
      type: String | Date
    },
    from: Date,
    to: Date,
    isRange: Boolean
  },
  components: {},
  computed: {
    isRangePicker () {
      return this.type.indexOf('range') >= 0
    },
    fmttedRangeDate () {
      let fmt = ''
      if (this.innerFrom) {
        fmt += rdFmt(this.innerFrom)
      }
      if (this.innerTo) {
        fmt += ' - ' + rdFmt(this.innerTo)
      }
      return fmt
    },
    selectedDate () {
      return new Date(this.selectedYear, this.selectedMonth, 1)
    },
    selectedDateDayOfYear () {
      return rdGetDOY(this.selectedDate)
    },
    yearRange () {
      let start = new Date().getFullYear()
      let end = new Date(new Date().getFullYear() + 10)
      if (this.startDate) {
        start = new Date(this.startDate)
      }
      if (this.endDate) {
        end = new Date(this.endDate)
      }
      let years = []
      while (start < end) {
        years.push(start)
        start++
      }
      return years
    },
    dayRows () {
      // THIS always starts on sunday.. so if your month starts on a monday then sunday will be a 0 array item
      let days = []
      let i = 0
      let selectedDay = this.selectedDate.getDay()
      let monthEnd = new Date(this.selectedDate.getFullYear(), this.selectedDate.getMonth() + 1, 0).getDate()
      // padd start of month
      while (i < selectedDay && i < 32) {
        days.push(0)
        i++
      }
      i = 1 // one based cause dates don't start at 0
      while (i <= monthEnd && i < 32) {
        days.push(i)
        i++
      }
      let dayRows = []
      let dayRow = []
      for (let i = 0; i < days.length; i++) {
        if (i % 7 === 0 && i !== 0) {
          dayRows.push(dayRow)
          dayRow = []
        }
        dayRow.push({
          dayNum: days[i],
          date: new Date(this.selectedDate.getFullYear(), this.selectedDate.getMonth(), i - selectedDay + 1) // +1 for 0 based index with offset of padded days
        })
      }
      // padd end of month
      let padEnd = 7 - dayRow.length
      i = 1
      while (i <= padEnd && i < 32) {
        dayRow.push({
          dayNum: 0,
          date: new Date(this.selectedDate.getFullYear(), this.selectedDate.getMonth() + 1, i)
        })
        i++
      }
      dayRows.push(dayRow)
      return dayRows
    },
    pfx () {
      return this.cssPrefix
    },
    hasFrom () {
      return (this.innerFrom !== null)
    },
    hasTo () {
      return (this.innerTo !== null)
    }
  },
  methods: {
    parseDate (ev) {
      let val = ev.target.value

      let parsed = parseTime(val)
      if (parsed.absolute) {
        this.setFrom(new Date(parsed.absolute), true)
      }

      if (val.indexOf('-') > 0) {
        let toVal = val.split('-')
        if (toVal.length > 0) {
          val = toVal[1]
        }

        let parsed = parseTime(val)
        if (parsed.absolute) {
          this.setTo(new Date(parsed.absolute))
        }
      }

      this.closePicker()
    },
    dateClicked (day) {
      this.hoverDate = null
      if (this.isRangePicker) {
        if (this.hasTo) {
          this.setFrom(day.date)
          this.setTo(null)
          return
        }

        if (this.hasFrom && this.innerFrom.getTime() < day.date.getTime()) {
          this.setTo(day.date)
          this.closePicker()
          this.$emit('input', {from: this.innerFrom, to: this.innerTo})
          return
        }
      }

      this.setFrom(day.date)
      if (this.isRangePicker) {
        this.$emit('input', {from: this.innerFrom, to: this.innerTo})
      } else {
        this.$emit('input', this.innerFrom)
        if (this.type !== 'datetime') {
          this.closePicker()
        }
      }
    },
    updateStartDate (day) {
      this.startDate = day.date
    },
    updateHoverDate (day) {
      if (!this.isRangePicker) {
        return
      }
      if (this.innerFrom !== null && this.innerTo === null) {
        this.hoverDate = day.date
      }
    },
    closePicker () {
      this.isPickerActive = false
      this.bindEsc()
    },
    togglePicker () {
      this.isPickerActive = !this.isPickerActive
      this.bindEsc()
    },
    bindEsc () {
      if (this.isPickerActive) {
        this.listener = document.addEventListener('keyup', (evt) => {
          if (evt.keyCode === 27) {
            this.closePicker()
          }
        })
      } else {
        document.removeEventListener('keyup', this.listener)
      }
    },
    setFrom (val, isParsed) {
      this.innerFrom = val
      if (this.type === 'datetime' && !isParsed) {
        this.setTimeOnFrom()
      } else {
        this.$emit('from-changed', this.innerFrom)
        this.formatField()
      }
    },
    setTimeOnFrom () {
      let dt = this.innerFrom ? new Date(this.innerFrom) : new Date()
      let hours = this.selectedHour
      if (this.selectedAMPM === 0 && hours === 12) {
        hours = 0
      } else if (this.selectedAMPM === 12 && hours < 12) {
        hours = this.selectedHour + this.selectedAMPM
      }
      dt.setHours(hours)
      dt.setMinutes(this.selectedMinute)
      this.innerFrom = dt
      this.$emit('from-changed', this.innerFrom)
      this.formatField()
    },
    setTo (val) {
      this.innerTo = val
      this.$emit('to-changed', this.innerTo)
      this.formatField()
    },
    formatField () {
      let result = ''
      if (this.innerFrom) {
        result = rdFmt(this.innerFrom)
        if (this.type.indexOf('time') > -1) {
          result += ' ' + rdFmtTime(this.innerFrom)
        }
      }
      if (this.isRangePicker) {
        result += ' - '
      }
      if (this.innerTo) {
        result += rdFmt(this.innerTo)
        if (this.type.indexOf('time') > -1) {
          result += ' ' + rdFmtTime(this.innerTo)
        }
      }
      this.$refs.inp.value = result
    }
  },
  watch: {},
  data () {
    return {
      selectedMonth: new Date().getMonth(),
      selectedYear: new Date().getFullYear(),
      selectedHour: 12,
      selectedMinute: 0,
      selectedAMPM: 0,
      start: null,
      hoverDate: null,
      isPickerActive: false,
      innerFrom: null,
      innerTo: null,
      listener: null
    }
  },

  // LIFECYCLE METHODS
  // ______________________________________
  beforeCreate () {
  },
  created () {
    if (Object.hasOwnProperty(this.value, 'to')) {
      this.innerFrom = this.value.from
      this.innerTo = this.value.to
    }
    if (this.from || this.value) {
      this.innerFrom = this.from || new Date(this.value)
      let hours = this.innerFrom.getHours()
      if (hours > 12) {
        hours = hours - 12
        this.amPm = 12
      }
      this.selectedHour = hours
      this.selectedMinute = this.innerFrom.getMinutes()
    }
    if (this.to) {
      this.innerTo = this.to
    }
  },
  beforeMount () {
  },
  mounted () {
    this.formatField()
  },
  beforeUpdate () {
  },
  updated () {
  },
  beforeDestroy () {
  },
  destroyed () {
  }
}
</script>

<style lang="css">
.dp-range-date .dp-click-overlay {
  position: fixed;
  background: rgba(0, 0, 0, 0.1);
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  z-index: 10001;
}
.dp-range-date {
  position: relative;
  width: 100%;
}

.dp-range-date .dp-picker .dp-field{
  margin-top: 0.5rem;
}

.dp-range-date.is-open .dp-input {
  position: relative;
  width: 100%;
  z-index: 10003;
}
.dp-range-date .dp-picker{
  background: white;
  position: absolute;
  margin-top: 0.5rem;
  padding: 0.5rem;
  z-index: 10002;
}
.dp-range-date .dp-month-year {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center;
}
.dp-range-date .dp-month-year .dp-month-year-spacer{
  display: block;
  width: 0.5rem;
}
.dp-range-date .dp-days .dp-day-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.dp-range-date .dp-days .dp-day-row .dp-day {
  cursor: pointer;
  width: 14.25%;
  text-align: center;
  padding: 0.5rem 0.75rem;
  margin-left: -1px;
}
.dp-range-date .dp-days .dp-day-row .dp-day.is-diff-month{
  color: grey;
}
.dp-range-date .dp-days .dp-day-row .dp-day:hover,
.dp-range-date .dp-days .dp-day-row .dp-day.is-hovered {
  background: red;
}
.dp-range-date .dp-days .dp-day-row .dp-day.is-active-one {
  background: purple;
}
.dp-range-date .dp-days .dp-day-row .dp-day.is-active-two {
  background: orange;
}
.dp-range-date .dp-days .dp-day-row .dp-day.is-active-between {
  background: yellow;
}
.dp-range-date .dp-days .dp-day-row .dp-day.dp-day-text {
  font-size: 0.7rem;
  color: grey;
  text-transform: uppercase;
  background: white !important;
}
.dp-range-date .dp-next-prev {
  color: #3273dc;
}

</style>
