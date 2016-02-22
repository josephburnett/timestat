(ns ui.core
  (:require-macros [cljs.core.async.macros :refer [go-loop]])
  (:require [om.core :as om :refer [IDisplayName IRender IRenderState IInitState IWillMount]]
            [om.dom :as dom]
            [cljs-time.coerce :as time-coerce]
            [cljs-time.core :as time]
            [cljs.core.async :as async]
            [clojure.string :as string]
            [goog.math :as math]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom {:timer {:Owner "josephburnett79"
                                  :State "anon-running"
                                  :Start (time-coerce/to-string (time/now))
                                  :Stop "0001-01-01T00:00:00Z"
                                  :TimerID ""}
                          :timer-ids [{:name "Clean the kitchen"
                                       :id "clean-the-kitchen"}
                                      {:name "Take a bath"
                                       :id "take-a-bath"}]}))

(defn pie [x y r fill tenth-degrees]
  (let [points (map #(str (+ x (math/angleDx (/ % 10) r)) "," (+ y (math/angleDy (/ % 10) r)))
                    (range 0 tenth-degrees))
        points-string (string/join " " (cons (str x "," y) points))]
    (dom/polygon #js {:points points-string
                      :transform (str "rotate(-90," x "," y ")")
                      :fill fill})))

(defn circle [x y r fill]
  (dom/circle #js {:r r
                   :cx x
                   :cy y
                   :fill fill}))
  
(defn timer-minutes [data owner]
  (reify
    IDisplayName
    (display-name [_]
      "Minutes on the timer")
    IWillMount
    (will-mount [_]
      (go-loop []
        (let [start (om/get-state owner :start)
              interval-seconds (/ (time/in-millis (time/interval start (time/now))) 1000)]
          (om/update-state! owner :elapsed-minutes #(mod (/ interval-seconds 60) 60)))
        (<! (async/timeout 1000))
        (recur)))
    IInitState
    (init-state [_]
      {:elapsed-minutes 0})
    IRenderState
    (render-state [_ {x :x y :y r :r min :elapsed-minutes init :initialized}]
      (pie x y r "#d9d9d9" (* 6 10 min)))))

(defn timer-hours [data owner]
  (reify
    IWillMount
    (will-mount [_]
      (go-loop []
        (let [start (om/get-state owner :start)
              interval-seconds (/ (time/in-millis (time/interval start (time/now))) 1000)]
          (om/update-state! owner :elapsed-hours #(mod (/ interval-seconds 60 60) 24)))
        (<! (async/timeout 10000))
        (recur)))
    IInitState
    (init-state [_]
      {:elapsed-hours 0})
    IRenderState
    (render-state [_ {x :x y :y r :r hr :elapsed-hours}]
      (pie x y (/ r 1.4) "#cccccc" (* 30 10 hr)))))

(defn timer-seconds [data owner]
  (reify
    IWillMount
    (will-mount [_]
      (go-loop []
        (let [start (om/get-state owner :start)
              interval-seconds (/ (time/in-millis (time/interval start (time/now))) 1000)]
          (om/update-state! owner :elapsed-seconds #(mod interval-seconds 60)))
        (<! (async/timeout 50))
        (recur)))
    IInitState
    (init-state [_]
      {:elapsed-seconds 0})
    IRenderState
    (render-state [_ {x :x y :y r :r sec :elapsed-seconds}]
      (dom/line #js {:x1 (+ x (math/angleDx (* 6 sec) (/ r 1.4)))
                     :y1 (+ y (math/angleDy (* 6 sec) (/ r 1.4)))
                     :x2 (+ x (math/angleDx (* 6 sec) r))
                     :y2 (+ y (math/angleDy (* 6 sec) r))
                     :style #js {:stroke "#bfbfbf"
                                 :strokeWidth "4px"}
                     :transform (str "rotate(-90," x "," y")")}))))

(defn timer-text [data owner]
  (reify
    IDisplayName
    (display-name [_]
      "The time text")
    IInitState
    (init-state [_]
      {:elapsed-seconds 0
       :elapsed-minutes 0
       :elapsed-hours 0})
    IWillMount
    (will-mount [_]
      (go-loop []
        (let [start (om/get-state owner :start)
              interval-seconds (/ (time/in-millis (time/interval start (time/now))) 1000)]
          (om/update-state! owner :elapsed-minutes #(mod (/ interval-seconds 60) 60))
          (om/update-state! owner :elapsed-hours #(mod (/ interval-seconds 60 60) 24)))
        (<! (async/timeout 1000))
        (recur)))
    IRenderState
    (render-state [_ {x :x y :y sec :elapsed-seconds min :elapsed-minutes hr :elapsed-hours}]
      (let [min-only (= 0 (int hr))]
        (dom/text #js {:x (if min-only
                            (- x 50)
                            (- x 130))
                       :y (+ y 25)
                       :fill "blue"
                       :style #js {:fontSize "90px"}}
                  (if min-only 
                    (str (int min) "m")
                    (str (int hr) "h " (int min) "m")))))))

(defn timer-view [data owner]
  (reify
    IDisplayName
    (display-name [_]
      "The timer")
    IInitState
    (init-state [_]
      {:width 500
       :height 500
       :radius 200
       :x 250
       :y 250})
    IRenderState
    (render-state [_ {w :width h :height x :x y :y r :radius}]
      (let [start-date (time-coerce/to-date (get-in data [:timer :Start]))
            timer-dim {:x x :y y :r r :start start-date}]
        (dom/svg #js {:width w
                      :height h}
                 (circle x y (+ r 20) "green")
                 (circle x y r "white")
                 (om/build timer-minutes nil {:init-state timer-dim})
                 (om/build timer-hours nil {:init-state timer-dim})
                 (om/build timer-seconds nil {:init-state timer-dim})
                 (om/build timer-text nil {:init-state timer-dim}))))))

(defn menu-view [data owner]
  (reify
    IRender
    (render [_]
      (dom/ul nil
          (map #(dom/li nil (:name %)) (:timer-ids data))))))
              
(om/root timer-view app-state
         {:target (. js/document (getElementById "timer"))})

(om/root menu-view app-state
         {:target (. js/document (getElementById "menu"))})

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)
