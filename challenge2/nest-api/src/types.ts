export type Position = {
    lat: number;
    lng: number;
  }
  
  export type Route = {
    id: number,
    title: string;
    startPosition: Position;
    endPosition: Position;
  };
  