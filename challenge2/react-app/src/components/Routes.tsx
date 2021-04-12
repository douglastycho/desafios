import { useEffect, useState } from "react";


import { Route } from "../util/models";

const API_URL = process.env.REACT_APP_API_URL;
type Props = {

};
export const Routes = (props: Props) => {
  const [routes, setRoutes] = useState<Route[]>([]);

  useEffect(() => {
    fetch(`${API_URL}/routes`)
      .then((data) => data.json())
      .then((data) => setRoutes(data));
  }, []);

  return (
    <div className="container">
      {routes.map((route) => (
        <div className="item">
          <div className="title">
            <h3>{route.title}:</h3> 
          </div>
          <div className="start">
            <h4>Rota Inicial</h4>
              <h5>Latitude: {route.startPosition.lat} |  Longitude : {route.startPosition.lng}</h5>
          </div>
          <div className="end"> 
            <h4>Rota Final</h4>
              <h5>Latitude: {route.endPosition.lat} | Longitude: {route.endPosition.lng}</h5> 
          </div>
          </div>
      ))}
    </div>
  );
};