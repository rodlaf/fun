'use client';

import React, { createContext, useState, useEffect } from 'react';

const StateContext = createContext();

const StateProvider = ({ children, initialValue = new Array(10).fill(null) }) => {
  const [state, setState] = useState(initialValue);
  const [ws, setWs] = useState(null);

  useEffect(() => {
    const socket = new WebSocket('wss://colorfun.fly.dev:7001/ws');
    setWs(socket);

    socket.onmessage = (event) => {
      const stateUpdate = JSON.parse(event.data);
      console.log('got state update: ', stateUpdate)

      setState(prevState => {
        const newState = [...prevState];
        newState[stateUpdate.index] = stateUpdate.value;
        return newState;
      });
    };

    return () => socket.close();
  }, []);

  const sendUpdate = (index, value) => {
    if (ws) {
      const stateUpdate = { index, value };
      ws.send(JSON.stringify(stateUpdate));
    }
  };

  return (
    <StateContext.Provider value={{ state, setState, sendUpdate }}>
      {children}
    </StateContext.Provider>
  );
};

export { StateProvider, StateContext };