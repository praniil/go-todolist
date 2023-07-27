import { useState } from "react"
import { Modal, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
function AddTodo(){
    const [open, setOpen] = useState(false)

    const form = useForm({
        initialValues: {
            title: "",
            body: "",
        },
    });

    

    return(
        <>
        <Modal opened={open} onClose={() => setOpen(false)} title="Create Todo">
            <form onSubmit={form.onSubmit(createTodo)}/>
            <TextInput
            required
            mb={12}
            label = "Todo"
            placeholder = "What do you want to do?"
            {...form.getInputProps("title")}
            />
        </Modal>
        </>
    )

}

export default AddTodo;