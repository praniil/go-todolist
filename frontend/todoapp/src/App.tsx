import { Box, List, ThemeIcon } from "@mantine/core";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import useSWR from "swr";
import "./App.css";
import AddTodo from './components/AddTodo';

function App() {

  return (
    <>
    <AddTodo/>
    </>
  );
}

export default App;
