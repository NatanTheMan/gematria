(import scheme 
        (chicken base)
        (chicken io)
        (chicken process-context))

(define units-list (list "" "א" "ב" "ג" "ד" "ה" "ו" "ז" "ח" "ט"))
(define dozens-list (list "" "י" "כ" "ל" "מ" "נ" "ס" "ע" "פ" "צ"))
(define hundreds-list (list "" "ק" "ר" "ש" "ת"))

(define (to-unit n)
  (list-ref units-list n))

(+ 1 2)
(define (to-dozen n)
  (cond 
    ((= n 15) "וט")
    ((= n 16) "זט")
    (else (list-ref dozens-list (quotient n 10)))))

(define (to-hundred n)
  (list-ref hundreds-list (quotient n 100)))

(define (gematria n)
  (let ((d (remainder n 100))
        (u (remainder n 10)))

    (if (or (= d 16) (= d 15))
      (string-append
        (to-dozen d)
        (to-hundred n))

      (string-append
        (to-unit u)
        (to-dozen d)
        (to-hundred n)))))

(define (main)
  (let* ((arg (car (command-line-arguments)))
         (num (string->number arg)))

    (if num
      (begin
        (display (gematria num))
        (newline))
      (display "not num"))))

(main)
