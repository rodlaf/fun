import { StateProvider } from "@/components/stateProvider";
import Rotator from "@/components/rotator";

export default async function Home() {
    const res = await fetch("http://0.0.0.0:8080/getState", { cache: 'no-store' })
    const data = await res.json()
    const state = data.state

    return (
        <StateProvider initialValue={state}>
            <div className="flex flex-col items-center justify-center min-h-screen p-2 bg-gray-100">
                <h1 className="text-2xl font-bold text-center mb-6 text-blue-600">Welcome to ColorFun!</h1>
                <div className="text-lg mb-2">These buttons are a shared state of all users!</div>
                <div className="text-md mb-4">This is a demo of <a href="https://github.com/rodlaf/fun" className="text-blue-500 underline">React Server Components and WebSockets</a>.</div>
                <div className="grid grid-cols-2 gap-4">
                    {
                        data.state.map((item, index) => {
                            return <div key={index} className="p-2 bg-white shadow rounded"><Rotator index={index} initialValue={item} /></div>
                        })
                    }
                </div>
            </div>
        </StateProvider>
    );
}
