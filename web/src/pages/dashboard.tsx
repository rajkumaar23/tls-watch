import { useAuth } from "@/context/auth"
import { useEffect } from "react"

export function Dashboard () {
    const {user} = useAuth();

    return (
        <>
        <p>henlo, {user?.name.toLowerCase()}</p>
        </>
    )
}