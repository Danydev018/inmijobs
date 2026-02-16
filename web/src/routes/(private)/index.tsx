import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(private)/')({
  component: App,
})

function App() {
  return (
    <div className="p-6 md:p-10 max-w-7xl mx-auto">
      <header className="mb-8">
        <h1 className="text-3xl font-bold text-[#1F2937]">
          ¡Bienvenido de nuevo a <span className="text-[#F97316]">Inmijobs</span>!
        </h1>
        <p className="text-[#6B7280] mt-2">
          Aquí tienes un resumen de las últimas oportunidades y comunidades para ti.
        </p>
      </header>

      {/* Cards de ejemplo para rellenar el Dashboard */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-white p-6 rounded-3xl border border-[#E5E7EB] shadow-sm hover:shadow-md transition-shadow">
          <div className="w-12 h-12 bg-[#FFF3E6] rounded-2xl flex items-center justify-center mb-4">
            <span className="text-[#F97316] font-bold">03</span>
          </div>
          <h3 className="font-bold text-lg">Postulaciones</h3>
          <p className="text-[#6B7280] text-sm">Tienes 3 procesos activos esta semana.</p>
        </div>

        <div className="bg-white p-6 rounded-3xl border border-[#E5E7EB] shadow-sm hover:shadow-md transition-shadow">
          <div className="w-12 h-12 bg-[#F3E8FF] rounded-2xl flex items-center justify-center mb-4">
            <span className="text-[#8B5CF6] font-bold">12</span>
          </div>
          <h3 className="font-bold text-lg">Nuevas Vacantes</h3>
          <p className="text-[#6B7280] text-sm">Basadas en tus preferencias de perfil.</p>
        </div>
      </div>
    </div>
  )
}