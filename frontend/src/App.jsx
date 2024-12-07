import React, { useState } from "react";
import "./App.css";
import { Schedule } from "../wailsjs/go/main/App";

const App = () => {
    const [timer, setTimer] = useState("");
    const [action, setAction] = useState("");
    const [showCustomTimer, setShowCustomTimer] = useState(false);
    const [showCustomAction, setShowCustomAction] = useState(false);
    const [customTimer, setCustomTimer] = useState("");
    const [customAction, setCustomAction] = useState("");

    const handlePresetTimer = (value) => {
        setTimer(value);
        setShowCustomTimer(false);
    };

    const handleCustomTimer = (value) => {
        setCustomTimer(value);
        setTimer("");
    };

    const handlePresetAction = (value) => {
        setAction(value);
        setShowCustomAction(false);
    };

    const handleCustomAction = (value) => {
        setCustomAction(value);
        setAction("");
    };

    const scheduleTask = () => {
        const selectedTimer = customTimer || timer;
        const selectedAction = customAction || action;
        Schedule(selectedTimer, selectedAction)
    };

    return (
        <div className="container">
            <h1>App</h1>

            <div className="flowchart">
                <div className="timer-section">
                    <button
                        className={timer === "30" ? "selected" : ""}
                        onClick={() => handlePresetTimer("30")}
                    >
                        30 seconds
                    </button>
                    <button
                        className={timer === "60" ? "selected" : ""}
                        onClick={() => handlePresetTimer("60")}
                    >
                        1 minute
                    </button>
                    <button
                        className={timer === "300" ? "selected" : ""}
                        onClick={() => handlePresetTimer("300")}
                    >
                        5 minutes
                    </button>
                    <button
                        className={showCustomTimer ? "selected" : ""}
                        onClick={() => setShowCustomTimer(true)}
                    >
                        Custom
                    </button>
                    {showCustomTimer && (
                        <input
                            type="number"
                            placeholder="Custom time (sec)"
                            value={customTimer}
                            onChange={(e) => handleCustomTimer(e.target.value)}
                        />
                    )}
                </div>

                <div className="arrow-down"></div>

                <div className="action-section">
                    <button
                        className={action === "Shutdown" ? "selected" : ""}
                        onClick={() => handlePresetAction("Shutdown")}
                    >
                        Shutdown
                    </button>
                    <button
                        className={action === "Restart" ? "selected" : ""}
                        onClick={() => handlePresetAction("Restart")}
                    >
                        Restart
                    </button>
                    <button
                        className={action === "Sleep" ? "selected" : ""}
                        onClick={() => handlePresetAction("Sleep")}
                    >
                        Sleep
                    </button>
                    <button
                        className={showCustomAction ? "selected" : ""}
                        onClick={() => setShowCustomAction(true)}
                    >
                        Custom
                    </button>
                    {showCustomAction && (
                        <input
                            type="text"
                            placeholder="Custom shell command"
                            value={customAction}
                            onChange={(e) => handleCustomAction(e.target.value)}
                        />
                    )}
                </div>
            </div>

            <button id="scheduleBtn" onClick={scheduleTask}>
                Schedule
            </button>
        </div>
    );
};

export default App;
