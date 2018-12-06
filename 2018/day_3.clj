(require '[clojure.string :as str])
(require '[clojure.math.combinatorics :as combo])

(defn parse_rectangle
  [definition]
  (let [values (->
                definition
                (str/replace " @ " "#")
                (str/replace "," "#")
                (str/replace "x" "#")
                (str/replace ": " "#")
                (str/split #"#"))
        start_x (read-string (get values 2))
        start_y (read-string (get values 3))
        width (read-string (get values 4))
        height (read-string (get values 5))]
    {:start_x (+ 1 start_x)
     :start_y (+ 1 start_y)
     :end_x (+ start_x width)
     :end_y (+ start_y height)
     }))

(assert (=
         (parse_rectangle "#1 @ 871,327: 16x20")
         {
          :start_x 872
          :start_y 328
          :end_x 887
          :end_y 347}))
(assert (=
         (parse_rectangle "#123 @ 3,2: 5x4")
         {
          :start_x 4
          :start_y 3
          :end_x 8
          :end_y 6}))

(defn parse_rectangles []
  (->>
   (slurp "day_3_input")
   (str/split-lines)
   (map parse_rectangle)))


(defn point_within_rectangle?
  [rectangle point]
  (let [start_x (get rectangle :start_x)
        end_x (get rectangle :end_x)
        start_y (get rectangle :start_y)
        end_y (get rectangle :end_y)
        point_x (get point :x)
        point_y (get point :y)]
    (and (<= start_x point_x)
         (<= start_y point_y)
         (>= end_x point_x)
         (>= end_y point_y)
         )))

(assert (= (point_within_rectangle?
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20}
            {:x 18
             :y 17})
           true))

(assert (= (point_within_rectangle?
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20}
            {:x 23
             :y 17})
           false))

(assert (= (point_within_rectangle?
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20}
            {:x 10
             :y 10})
           true))

(assert (= (point_within_rectangle?
            {:start_x 5
             :start_y 5
             :end_x 15
             :end_y 15}
            {:x 10
             :y 10})
           true))

(defn rectangles_overlap?
  [rectangle_left rectangle_right]
  (let [point_up_left {:x (get rectangle_left :start_x)
                       :y (get rectangle_left :start_y)}
        point_up_right {:x (get rectangle_left :end_x)
                        :y (get rectangle_left :start_y)}
        point_down_right {:x (get rectangle_left :end_x)
                          :y (get rectangle_left :end_y)}
        point_down_left {:x (get rectangle_left :start_x)
                         :y (get rectangle_left :end_y)}
        ]
    (->>
     [point_up_left point_up_right point_down_right point_down_left]
     (filter (fn [point] (point_within_rectangle? rectangle_right point)))
     (not= [])
     )))

(assert (= (rectangles_overlap?
            {:start_x 15
             :start_y 15
             :end_x 20
             :end_y 20}
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20}))
        true)

(assert (= (rectangles_overlap?
            {:start_x 150
             :start_y 150
             :end_x 200
             :end_y 200}
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20})
           false))

(assert (= (rectangles_overlap?
            {:start_x 5
             :start_y 5
             :end_x 15
             :end_y 15}
            {:start_x 0
             :start_y 10
             :end_x 10
             :end_y 20})
           true))

(defn compare_point
  [compare_func point_left point_right]
  (let [left_x (get point_left :x)
        right_x (get point_right :x)
        left_y (get point_left :y)
        right_y (get point_right :y)]
    {:x (compare_func left_x right_x)
     :y (compare_func left_y right_y)}))

(assert (= (compare_point
            max
            {:x 10
             :y 20}
            {:x 20
             :y 40})
           {:x 20
            :y 40}))

(assert (= (compare_point
            min
            {:x 10
             :y 20}
            {:x 20
             :y 40})
           {:x 10
            :y 20}))

(defn overlaping_rectangle
  [rectangle_left rectangle_right]
  (let [start_left {:x (get rectangle_left :start_x)
                     :y (get rectangle_left :start_y)}
         start_right {:x (get rectangle_right :start_x)
                      :y (get rectangle_right :start_y)}
         end_left {:x (get rectangle_left :end_x)
                   :y (get rectangle_left :end_y)}
         end_right {:x (get rectangle_right :end_x)
                    :y (get rectangle_right :end_y)}
         start (compare_point max start_left start_right)
         end (compare_point min end_left end_right)
         ]
     {:start_x (get start :x)
      :start_y (get start :y)
      :end_x (get end :x)
      :end_y (get end :y)}))

(assert (= (overlaping_rectangle
            {:start_x 15
             :start_y 15
             :end_x 30
             :end_y 30}
            {:start_x 10
             :start_y 10
             :end_x 20
             :end_y 20}
            )
           {:start_x 15
            :start_y 15
            :end_x 20
            :end_y 20}
           ))

(assert (= (overlaping_rectangle
            {:start_x 2
             :start_y 4
             :end_x 5
             :end_y 7}
            {:start_x 4
             :start_y 2
             :end_x 7
             :end_y 5}
            )
           {:start_x 4
            :start_y 4
            :end_x 5
            :end_y 5}
           ))

(defn all_point_in_rectangle
  [rectangle]
  (let [start_x (get rectangle :start_x)
        end_x (get rectangle :end_x)
        start_y (get rectangle :start_y)
        end_y (get rectangle :end_y)]
    (set (for [x (range start_x end_x)
               y (range start_y end_y)]
          {:x x
           :y y}))))

;; (assert (= (all_point_in_rectangle
;;             {:start_x 15
;;              :start_y 18
;;              :end_x 17
;;              :end_y 19}
;;             )
;;            #{{:x 15
;;               :y 18}
;;              {:x 16
;;               :y 18}
;;              {:x 17
;;               :y 18}
;;              {:x 15
;;               :y 19}
;;              {:x 16
;;               :y 19}
;;              {:x 17
;;               :y 19}}
;;            ))

(defn overlaping_rectangles
  [rectangles]
  (->>
   (combo/combinations rectangles 2)
   (filter #(apply rectangles_overlap? %))
   (map #(apply overlaping_rectangle %))
   (set)
   )
  )

(assert (= (overlaping_rectangles
            [{:start_x 15
              :start_y 15
              :end_x 20
              :end_y 20}
             {:start_x 10
              :start_y 10
              :end_x 20
              :end_y 20}
             {:start_x 100
              :start_y 100
              :end_x 200
              :end_y 200}
             ])
           #{{:start_x 15
             :start_y 15
             :end_x 20
             :end_y 20}}))

(defn find_all_overlaping_inches
  [rectangles]
  (->>
   rectangles
   (overlaping_rectangles)
   (map all_point_in_rectangle)
   (reduce into)
   (count)
   )
  )

;; (-> (parse_rectangles) find_all_overlaping_inches)
