(ns etalage.env
  (:require
    [selmer.parser :as parser]
    [clojure.tools.logging :as log]
    [etalage.dev-middleware :refer [wrap-dev]]))

(def defaults
  {:init
   (fn []
     (parser/cache-off!)
     (log/info "\n-=[etalage started successfully using the development profile]=-"))
   :stop
   (fn []
     (log/info "\n-=[etalage has shut down successfully]=-"))
   :middleware wrap-dev})
