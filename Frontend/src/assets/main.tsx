import React from 'react'
import ReactDOM from 'react-dom/client'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'

import App from './App'
import Home from './components/pages/home'
import Reglas from './components/pages/reglas'


// Rutas
const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/reglas",
    element: <Reglas/>,
  },
  {
    path: "/persona",
    element: <App/>,
  },
]);


const rootElement = document.getElementById('root') as HTMLElement;

ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)

