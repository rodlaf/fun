import { StateProvider } from "@/components/stateProvider";
import Rotator from "@/components/rotator";

export default async function Home() {
    const res = await fetch("http://0.0.0.0:8080/getState", { cache: 'no-store' })
    const data = await res.json()
    const state = data.state

    return (
        <StateProvider initialValue={state}>
            <div className="flex flex-col items-center justify-center min-h-screen py-2">
                <h1 className="text-xl font-bold text-center mb-4">Welcome to Fun!</h1>
                <div className="grid grid-cols-2 gap-4">
                    {
                        data.state.map((item, index) => {
                            return <div key={index}><Rotator index={index} initialValue={item} /></div>
                        })
                    }
                </div>
            </div>
        </StateProvider>
    );
}
