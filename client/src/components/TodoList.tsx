// Starting Code
import { Flex, Spinner, Stack, Text } from "@chakra-ui/react";

import TodoItem from "./TodoItem";
import { useQuery } from "@tanstack/react-query";
import { BASE_URL } from "../App";

export type Todo = {
	_id: number;
	body: string;
	completed: boolean;
};

const TodoList = () => {
	const { data: todos, isLoading } = useQuery<Todo[]>({
		queryKey: ["todos"],
		queryFn: async () => {
            const res = await fetch(BASE_URL + "/todos");
            if (!res.ok) {
                const errorData = await res.json();
                throw new Error( errorData.error||"Failed to fetch todos");
            }
            return res.json();
        },
	});
	const allCompleted = todos?.every(todo => todo.completed);
	return (
		<>
			<Text fontSize={"4xl"} textTransform={"uppercase"} fontWeight={"bold"} textAlign={"center"} my={2}
            bgGradient='linear(to-l, #7928CA, #FF0080)'
            bgClip='text'>
				Today's Tasks
			</Text>
			{isLoading && (
				<Flex justifyContent={"center"} my={4}>
					<Spinner size={"xl"} />
				</Flex>
			)}
			{!isLoading &&todos==null&& (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
					<img src='/go.png' alt='Go logo' width={70} height={70} />
				</Stack>
			)}
			{/* Show this message when all tasks are completed */}
			{!isLoading && todos && todos.length > 0 && allCompleted && (
				<Stack alignItems={"center"} gap='3'>
					<Text fontSize={"xl"} textAlign={"center"} color={"gray.500"}>
						All tasks completed! 🤞
					</Text>
				</Stack>
			)}
			<Stack gap={3}>
				{todos?.map((todo) => (
					<TodoItem key={todo._id} todo={todo} />
				))}
			</Stack>
		</>
	);
};
export default TodoList;