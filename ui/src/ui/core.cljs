(ns ui.core
  (:require-macros [cljs.core.async.macros :refer [go-loop]])
  (:require [om.core :as om :refer [IDisplayName IRender IRenderState IInitState IWillMount]]
            [om.dom :as dom]
            [cljs-time.coerce :as time-coerce]
            [cljs-time.core :as time]
            [cljs.core.async :as async]
            [clojure.string :as string]
            [goog.math :as math]
            [ajax.core :refer [GET]]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom {:timer {}
                          :timer-ids []
                          :dimensions {
                                       :width 500
                                       :height 500
                                       :x 250
                                       :y 250
                                       :r 200}}))

(defn load-timer-async []
  (GET "/timer" {:handler #(let [cursor (om/root-cursor app-state)]
                             (om/update! cursor :timer %))
                 :keywords? true
                 :response-format :json}))

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

(defn self-refreshing-component [& {:keys [name interval-millis render-fn]}]
  (fn [data owner]
    (reify
      IDisplayName
      (display-name [_]
        name)
      IWillMount
      (will-mount [_]
        (go-loop []
          (let [data (om/root-cursor app-state)]
            (when (contains? (:timer data) :Start)
              (let [start (time-coerce/to-date (get-in data [:timer :Start]))
                    interval-seconds (/ (time/in-millis (time/interval start (time/now))) 1000)]
                (om/update-state! owner :elapsed-seconds #(mod interval-seconds 60))
                (om/update-state! owner :elapsed-minutes #(mod (/ interval-seconds 60) 60))
                (om/update-state! owner :elapsed-hours #(mod (/ interval-seconds 60 60) 24))))
            (<! (async/timeout interval-millis))
            (recur))))
      IInitState
      (init-state [_]
        {:elapsed-seconds 0
         :elapsed-minutes 0
         :elapsed-hours 0})
      IRenderState
      (render-state [owner state]
        (render-fn data state)))))

(def timer-minutes
  (self-refreshing-component
   :name "Minutes on the timer"
   :interval-millis 1000
   :render-fn (fn [data {min :elapsed-minutes}]
                (let [{x :x y :y r :r} (:dimensions data)]
                  (pie x y r "#d9d9d9" (* 6 10 min))))))

(def timer-hours
  (self-refreshing-component
   :name "Hours on the timer"
   :interval-millis 10000
   :render-fn (fn [data {hr :elapsed-hours}]
                (let [{x :x y :y r :r} (:dimensions data)]
                  (pie x y (/ r 1.4) "#cccccc" (* 30 10 hr))))))

(def timer-seconds
  (self-refreshing-component
   :name "Seconds on the timer"
   :interval-millis 50
   :render-fn (fn [data {sec :elapsed-seconds}]
                (let [{x :x y :y r :r} (:dimensions data)]
                  (dom/line #js {:x1 (+ x (math/angleDx (* 6 sec) (/ r 1.4)))
                                 :y1 (+ y (math/angleDy (* 6 sec) (/ r 1.4)))
                                 :x2 (+ x (math/angleDx (* 6 sec) r))
                                 :y2 (+ y (math/angleDy (* 6 sec) r))
                                 :style #js {:stroke "#bfbfbf"
                                             :strokeWidth "4px"}
                                 :transform (str "rotate(-90," x "," y")")})))))

(def timer-text
  (self-refreshing-component
   :name "Timer text"
   :interval-millis 1000
   :render-fn (fn [data {min :elapsed-minutes hr :elapsed-hours}]
                (let [{x :x y :y} (:dimensions data)
                      min-only (= 0 (int hr))]
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
    IRender
    (render [_]
      (let [cursors (select-keys data [:dimensions :timer])
            {w :width h :height x :x y :y r :r} (:dimensions cursors)]
        (dom/svg #js {:width w
                      :height h}
                 (circle x y (+ r 20) "green")
                 (circle x y r "white")
                 (om/build timer-minutes cursors)
                 (om/build timer-hours cursors)
                 (om/build timer-seconds cursors)
                 (om/build timer-text cursors))))))

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

(load-timer-async)

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)
