# Errata for *Pro Go*

This file contains errors that are likely to prevent the code examples from working as described in ths book. See [this](typos.md) file for small mistakes that I intend to correct in the next edition.


---

** Chapter 18**

From page 469, the 1.20.1 release of Go changes the generation of random numbers so that a seed value does not have to be explicitly specified, meaning the `Seed` function is not required. 

(Thanks to Ken Rubotzky for reporting this problem)

---

**Chapter 33**

On Page 909, the getAnonymousFieldMethods function is shown with the wrong parameters and result. This statement:

    func getAnonymousFieldMethods(target []reflect.Type) reflect.Method {

should be:

    func getAnonymousFieldMethods(target reflect.Type) []reflect.Method {


(Thanks to Yaroslav Lazarev for reporting this problem)

---