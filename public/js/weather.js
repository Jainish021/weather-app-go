const weatherForm = document.querySelector("form")
const search = document.querySelector("input")

weatherForm.addEventListener("submit", (e) => {
    e.preventDefault()
    const location = search.value
    const content = document.getElementById("main_content")
    const message1 = document.getElementById("message1")
    const weather_icon = document.getElementById("weather_icon")
    const temprature = document.getElementById("temprature")
    const feelslike = document.getElementById("feelslike")
    const humidity = document.getElementById("humidity")
    const precip = document.getElementById("precip")
    const weather_description = document.getElementById("weather_description")
    const wind_speed = document.getElementById("wind_speed")
    const wind_dir = document.getElementById("wind_dir")
    const pressure = document.getElementById("pressure")
    const uv_index = document.getElementById("uv_index")
    const visibility = document.getElementById("visibility")
    message1.textContent = "Loading..."

    fetch("/weather?address=" + location).then((response) => {
        response.json().then((data) => {
            if (data.error) {
                message1.textContent = data.error
            } else {
                message1.textContent = data.Location
                if (content.className === "d-none") {
                    content.className = content.className.replace("d-none", "d-block")
                }
                weather_icon.className = "img-thumbnail w-5"
                weather_icon.src = data.WeatherIcon
                temprature.textContent = data.Temperature + "F"
                feelslike.textContent = data.Feelslike + "F"
                humidity.textContent = data.Humidity + "%"
                precip.textContent = data.Precip + "%"
                weather_description.textContent = data.WeatherDescription
                wind_speed.textContent = data.WindSpeed + " MPH"
                wind_dir.textContent = data.WindDir
                pressure.textContent = data.Pressure
                uv_index.textContent = data.UVIndex
                visibility.textContent = data.Visibility + " Miles"
            }
        })
    })
})