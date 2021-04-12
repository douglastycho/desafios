export interface Position {
  lat: number;
  lng: number;
}
export interface Route {
  id: number,
  title: string;
  startPosition: Position;
  endPosition: Position;
}
