import React from 'react';

export default ({
    temperature,
    humidity,
    getData
}) => {
    return (
        <div className='app-wrapper'>
            <h1>Weather in second room</h1>
            <p>Temperature: {temperature}</p>
            <p>Humidity: {humidity}</p>
            <button className="app-button app-button_green" onClick={getData}>Update temperature</button>
        </div>
    );
}
