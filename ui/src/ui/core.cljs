(ns ui.core
  (:require [om.core :refer [IRender root]]
            [om.dom :refer [circle svg text]]
            [cljs-time.coerce :refer [to-date to-long to-string]]
            [cljs-time.core :refer [now interval in-seconds]]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom{:Owner "josephburnett79"
                         :State "anon-running"
                         :Start (to-string (now))
                         :Stop "0001-01-01T00:00:00Z"
                         :TimerID ""}))

(defn timer-view [data owner]
  (let [elapsed-seconds (in-seconds (interval (to-date (:Start data)) (now)))
        elapsed-minutes (int (/ elapsed-seconds 60))
        period-seconds (mod elapsed-seconds 60)]
    (reify
      IRender
      (render [_]
        (svg #js {:width "500"
                  :height "500"}
             (circle #js {:r "200"
                          :cx "250"
                          :cy "250"
                          :fill "green"})
             (circle #js {:r "180"
                          :cx "250"
                          :cy "250"
                          :fill "white"})
             (text #js {:x "120"
                        :y "300"
                        :fill "blue"
                        :style #js {:font-size "170px"}}
                   elapsed-minutes)
             (text #js {:x "350"
                        :y "300"
                        :fill "blue"}
                   period-seconds))))))

(root timer-view app-state
      {:target (. js/document (getElementById "timer"))})

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)

