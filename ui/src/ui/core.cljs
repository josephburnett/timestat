(ns ui.core
  (:require-macros [cljs.core.async.macros :refer [go-loop]])
  (:require [om.core :refer [IRenderState IInitState IWillMount root transact!]]
            [om.dom :refer [circle line polygon svg text]]
            [cljs-time.coerce :refer [to-date to-long to-string]]
            [cljs-time.core :refer [now interval in-millis]]
            [cljs.core.async :refer [timeout]]
            [clojure.string :refer [join]]
            [goog.math :refer [angleDx angleDy]]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom {:timer {:Owner "josephburnett79"
                                  :State "anon-running"
                                  :Start (to-string (now))
                                  :Stop "0001-01-01T00:00:00Z"
                                  :TimerID ""}
                          :elapsed-seconds 0
                          :elapsed-minutes 0
                          :elapsed-hours 0}))

(defn pie [x y r fill tenth-degrees]
  (let [points (map #(str (+ x (angleDx (/ % 10) r)) "," (+ y (angleDy (/ % 10) r)))
                    (range 0 tenth-degrees))
        points-string (join " " (cons (str x "," y) points))]
    (polygon #js {:points points-string
                  :transform (str "rotate(-90," x "," y ")")
                  :fill fill})))

(defn timer-view [data owner]
  (reify
    IWillMount
    (will-mount [_]
      (go-loop []
        (let [interval-seconds (/ (in-millis (interval (to-date (get-in data [:timer :Start])) (now))) 1000)]
          (transact! data :elapsed-seconds #(mod interval-seconds 60))
          (transact! data :elapsed-minutes #(mod (/ interval-seconds 60) 60))
          (transact! data :elapsed-hours #(mod (/ interval-seconds 60 60) 24))
          (<! (timeout 10))
          (recur))))
    IRenderState
    (render-state [_ _]
      (let [r 200
            x 250
            y 250
                                        ;interval-seconds (in-seconds (interval (to-date (:Start data)) (now)))
            elapsed-minutes (:elapsed-minutes data) ;(mod (int (/ interval-seconds 60)) 60)
            elapsed-hours (:elapsed-hours data) ; (int (/ interval-seconds 60 60))
            elapsed-seconds (:elapsed-seconds data)] ; (mod interval-seconds 60)]
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
             (pie x y r "#d9d9d9" (* 6 10 elapsed-minutes))
             (pie x y (/ r 1.4) "#cccccc" (* 30 10 elapsed-hours))
             (line #js {:x1 (+ x (angleDx (* 6 elapsed-seconds) (/ r 1.4)))
                        :y1 (+ y (angleDy (* 6 elapsed-seconds) (/ r 1.4)))
                        :x2 (+ x (angleDx (* 6 elapsed-seconds) r))
                        :y2 (+ y (angleDy (* 6 elapsed-seconds) r))
                        :style #js {:stroke "#bfbfbf"
                                    :strokeWidth "4px"}
                        :transform (str "rotate(-90," x "," y")")})
             (text #js {:x "190"
                        :y "280"
                        :fill "blue"
                        :style #js {:fontSize "90px"}}
                   (if (= 0 (int elapsed-hours))
                     (str (int elapsed-minutes) "m")
                     (str (int elapsed-hours) "h " (int elapsed-minutes) "m")))
             (text #js {:x "350"
                        :y "310"
                        :fill "blue"}
                   (int elapsed-seconds)))))))

(root timer-view app-state
      {:target (. js/document (getElementById "timer"))})

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)

