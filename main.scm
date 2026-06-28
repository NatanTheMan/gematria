(import scheme 
        (chicken base)
        (chicken io)
        (chicken process-context))

(define units (list "" "א" "ב" "ג" "ד" "ה" "ו" "ז" "ח" "ט"))
(define dozens (list "" "י" "כ" "ל" "מ" "נ" "ס" "ע" "פ" "צ"))
(define hundreds (list "" "ק" "ר" "ש" "ת"))

(define (uni n)
  (list-ref units n))

(define (dozen n)
  (cond 
    ((= n 15) "וט")
    ((= n 16) "זט")
    (else (list-ref dozens (quotient n 10)))))

(define (hundred n)
  (list-ref hundreds (quotient n 100)))

(define (gematria n)
  (let ((d (remainder n 100)))
    (if (< n 500)
      (if (or (= d 16) (= d 15))
        (string-append
          (dozen d)
          (hundred n))
        (string-append
          (uni (remainder n 10))
          (dozen d)
          (hundred n)))

    "not supported")))

(define (main)
  (let* ((arg (car (command-line-arguments)))
        (num (string->number arg)))

    (if num
      (display (gematria num))
      (display "not num"))))

(main)
