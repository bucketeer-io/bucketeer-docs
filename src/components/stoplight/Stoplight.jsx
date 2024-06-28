import React from "react";
import { API } from "@stoplight/elements";
import "./Stoplight.css";

function Stoplight({ apiDescriptionUrl }) {
  return (
    <div className="stoplight">
      <API 
        apiDescriptionUrl={apiDescriptionUrl} 
        router="hash" 
        hideSchemas
      />
    </div>
  );
}

export default Stoplight;
