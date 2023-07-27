import { Box, List, ThemeIcon } from "@mantine/core";
import { CheckCircleFillIcon } from "@primer/octicons-react";
import useSWR from "swr";
import "./App.css";
import AddTodo from './components/AddTodo';

export interface Todo {
  id: number;
  title: string;
  body: string;
  done: Boolean;
}

export const ENDPOINT = "http://localhost:8080";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json())


function App() {
  const { data, mutate } = useSWR<Todo[]>("api/getalltodos", fetcher);

  async function markTodoAdDone() {
    const updated = await fetch(`${ENDPOINT}/api/createtodo`, {
      method: "POST",
    }).then((r) => r.json);
    mutate(updated);
  }
  return (
    <Box>
      <List>
        {data?.map((todo) => {
          return (
            <List.Item onClick={() => markTodoAdDone()}
              key={`todo_list__${todo.id}`}
              icon={
                todo.done ? (
                  <ThemeIcon color="teal" size={24} radius="xl">
                    <CheckCircleFillIcon size={20}/>
                  </ThemeIcon>
              ):(
                <ThemeIcon color="gray" size={24} radius="xl">
                  <CheckCircleFillIcon size={20}/>
                </ThemeIcon>
              )
            }
            >
              {todo.title}
            </List.Item>
          );
        })}
      </List>
      <AddTodo mutate={mutate} />
    </Box>
  );
}

export default App;
