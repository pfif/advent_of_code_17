(require '[clojure.string :as str])

(defn parse_drifts [advent_input]
  (map read-string (str/split-lines advent_input)))

(defn resulting_frequency [drifts]
  (reduce + drifts))

(println (resulting_frequency (parse_drifts (slurp "day_1_part_1_input")))) ;; First problem

(defn find_duplicate_frequency [frequencies last_frequency drifts]
  (let [last_frequency (+ last_frequency (first drifts))]
    (if (frequencies last_frequency)
      last_frequency
      (recur (conj frequencies last_frequency) last_frequency (rest drifts)))))

(println (find_duplicate_frequency #{0} 0 (cycle (parse_drifts (slurp "day_1_part_1_input"))))) ;; Second problem
