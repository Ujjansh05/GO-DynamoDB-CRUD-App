package routes

import ()


type Router struct{
	config *Config
	router *chi.Mux
}

func NewRouter() *Router{
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}



func (r *Router) SetRouter() *chi.Mux{

}

func (r *Router) setConfigsRouter(){

}

func RouterHealth(){


}

func RouterProduct(){

}

func EnableTimeout(){

}

func EnableCORS(){

}

func EnableRecover(){

}

func EnableRequestID(){

}

func EnableRealIP(){
	
}