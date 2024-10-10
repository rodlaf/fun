'use client';

import { useContext, useEffect, useState } from 'react';
import { StateContext } from './stateProvider';

export default function Rotator({ index, initialValue }) {
    const { state, sendUpdate } = useContext(StateContext);
    const [value, setValue] = useState(initialValue);

    useEffect(() => {
        setValue(state[index]);
    }, [state]);

    const handleClick = () => {
        const newValue = (value + 1) % 6;
        setValue(newValue);
        sendUpdate(index, newValue);
    };

    const getColor = (value) => {
        const colors = [
            'bg-red-500',
            'bg-orange-500',
            'bg-yellow-500',
            'bg-green-500',
            'bg-blue-500',
            'bg-purple-500'
        ];
        return colors[value];
    };

    return (
        <button
                key={index}
                onClick={handleClick}
                className={`w-32 h-12 px-4 py-2 ${getColor(value)} text-white rounded hover:bg-opacity-75 focus:outline-none focus:ring-2 focus:ring-opacity-50`}
        >
                {/* {value} */}
        </button>
    );
}