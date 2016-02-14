(ns ui.core
  (:require [om.core :refer [IRender root]]
            [om.dom :refer [circle polygon svg text]]
            [cljs-time.coerce :refer [to-date to-long to-string]]
            [cljs-time.core :refer [now interval in-seconds]]
            [clojure.string :refer [join]]
            [goog.math :refer [angleDx angleDy]]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom{:Owner "josephburnett79"
                         :State "anon-running"
                         :Start (to-string (now))
                         :Stop "0001-01-01T00:00:00Z"
                         :TimerID ""}))

(defn timer-view [data owner]
  (let [r 200
        x 250
        y 250
        interval-seconds (in-seconds (interval (to-date (:Start data)) (now)))
        elapsed-minutes (mod (int (/ interval-seconds 60)) 60)
        elapsed-hours (int (/ interval-seconds 60 60))
        elapsed-seconds (mod interval-seconds 60)]
    (reify
      IRender
      (render [_]
        (svg #js {:width "500"
                  :height "500"}
             (circle #js {:r (+ 20 r)
                          :cx x
                          :cy y
                          :fill "green"})
             (circle #js {:r r
                          :cx x
                          :cy y
                          :fill "white"})
             (when-not (= 0 elapsed-minutes)
               (let [minute-points (map #(str (+ x (angleDx % r)) "," (+ y (angleDy % r)))
                                        (range 0 (* 6 elapsed-minutes) 6))
                     points-string (join " " (cons (str x "," y) minute-points))]
                 (polygon #js {:points points-string
                               :transform (str "rotate(-90," x "," y ")")
                               :fill "#d9d9d9"})))
             (text #js {:x "120"
                        :y "300"
                        :fill "blue"
                        :style #js {:font-size "170px"}}
                   elapsed-minutes)
             (text #js {:x "350"
                        :y "300"
                        :fill "blue"}
                   elapsed-seconds))))))

(root timer-view app-state
      {:target (. js/document (getElementById "timer"))})

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)

