(ns ui.core
  (:require [om.core :refer [IRender root]]
            [om.dom :refer [circle svg text]]))

(enable-console-print!)

(println "Edits to this text should show up in your developer console.")

;; define your app data so that it doesn't get over-written on reload

(defonce app-state (atom {:text "1m"}))

(defn timer-view [data owner]
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
                 (:text data))))))

(root timer-view app-state
      {:target (. js/document (getElementById "timer"))})

(defn on-js-reload []
  ;; optionally touch your app-state to force rerendering depending on
  ;; your application
  ;; (swap! app-state update-in [:__figwheel_counter] inc)
)

