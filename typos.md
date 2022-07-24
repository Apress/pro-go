# Typos for *Pro Go*

This file contains smaller errors that are unlikely to prevent the code examples from working and which I will correct in the next edition of the book. See [this](errata.md) file for mistakes that prevent the code examples from working as expected.

---

**Chapter 4**

On page 70, the title for Listing 4-12 is:

> Using Constants in the main.go File in the basicFeatures Folder

It should be:

> Using ***Variables*** in the main.go File in the basicFeatures Folder

(Thanks to Michael for reporting this problem)

---

On page 77, the text following the listing contains this sentence:

>   The listing defines variables named `discount` and `salesperson`, neither of which is used in the rest of 
the code. 

The name of the second variable is `salesPerson`, with an uppercase `P`.

(Thanks to Tawfik Khalifeh for reporting this problem)

---

**Chapter 5**

On page 102, Table 5-7 contains the following text in the description for the `Floor` function:

>   The largest integer that is less than 27.1, for example, is 28.

It should be:

>   The largest integer that is less than 27.1, for example, is ***27***.

(Thanks to Francisco Vicedo and Korbai Zoltán for reporting this problem)

---

On page 114, the text following Listing 5-31 reads:

> The `Itoa` function accepts an `int` value, which is explicitly converted to an `int64` and passed to the 
`ParseInt` function.

This should be the `FormatInt` function.

(Thanks to Tawfik Khalifeh for reporting this problem)

---

**Chapter 6**

On page 125, Figure 6-5 uses a semi-colon instead of a colon in the exploding view of the statement.

(Thanks to Saleh for reporting this problem)

---

On page 134, Listing 6-18 contains a slice, but should use an array, like this:

    products := [3]string { "Kayak", "Lifejacket", "Soccer Ball"}

(Thanks to Tawfik Khalifeh for reporting this problem)

---

**Chapter 7**

On page 163, Listing 7-18 contains this statement:

    moreNames := []string { "Hat Gloves"}

It should be:

    moreNames := []string { "Hat", "Gloves"}

(Thanks to Francisco Vicedo for reporting this problem)

---

On pages 170-174, the formatting of the output is inconsistent. The listings include statements like this one:

    fmt.Println("allNames", allNames)

These statements should include a colon, like this:

    fmt.Println("allNames:", allNames)


(Thanks to Tawfik Khalifeh for reporting this problem)

---

**Chapter 8**


On page 206, the paragraph following Listing 8-14 contains the following sentence:

> The output from Listing 8-14 shows that the changes made to the values in the swpValues function do not affect the variables defined in the main function:

This should be:

> The output from Listing 8-14 shows that the changes made to the values in the ***swapValues*** function do not affect the variables defined in the main function:


(Thanks to Tawfik Khalifeh for reporting this problem)

---

On page 214, the calcTotalPrice function doesn't calculate a total price. It actually just adds the prices to the minSpend amount.

(Thanks to Francisco Vicedo for reporting this problem)

---

On page 217, the description of the results order doesn't match the actual output.

(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 9**

On page 223, listing 9-5 contains this statement in two locations;

     fmt.Println("Function assigned:", calcFunc == nil)

This is confusing and should be:

     fmt.Println("Function unassigned:", calcFunc == nil)


(Thanks to Francisco Vicedo for reporting this problem)

---

On page 240, this sentence:

>   On occasion, you may find encounter the contrary issue, which is the need to avoid early evaluation to ensure that the current value is used by a function

should be:

>   On occasion, you may ~~find~~ encounter the contrary issue, which is the need to avoid early evaluation to ensure that the current value is used by a function

(Thanks to Tawfik Khalifeh for reporting this problem)


---

**Chapter 10**

On page 249, this sentence:

>   For Listing 10-5, the zero type for the price field is 0, because the field type is float64; the code produces the following output when compiled and executed

should be:

> For Listing 10-5, the zero ***value*** for the price field is 0, because the field type is float64; the code produces the following output when compiled and executed

(Thanks to Tawfik Khalifeh for reporting this problem)

----

On page 272, the summary starts with:

> In this chapter, I describe the Go struct featur

This should be:

>   In this chapter, I ***described*** the Go struct featur

(Thanks to Francisco Vicedo for reporting this problem)

---

**Chapter 11**

On page 289, this text is incorrect:

> This interface has been given the name Expense, and the interface body contains a single method signature. 

The interface defines two methods.

(Thanks to Joseph Kashi and Francisco Vicedo for reporting this problem)

---

On page 296, this sentence:

>   Compile the project; you will receive the following error, telling you that a pointer receive is required:

should be:

>   Compile the project; you will receive the following error, telling you that a pointer ***receiver*** is required:

(Thanks to Tawfik Khalifeh for reporting this problem)


---

On page 302, this sentence:

>   Go allows the user of the empty interface—which means an interface that defines no methods—to represent 
any type, which can be a useful way to group disparate types that share no common features, as shown in 
Listing 11-32.

should be:

>   Go allows the ***use*** of the empty interface—which means an interface that defines no methods—to represent 
any type, which can be a useful way to group disparate types that share no common features, as shown in 
Listing 11-32.

(Thanks to Tawfik Khalifeh for reporting this problem)

---

**Chapter 12**

On page 319, the wrong chapter is specified for the description of the `FormatFloat` function. It should be Chapter 5.

(Thanks to Francisco Vicedo for reporting this problem)


---

**Chapter 13**

On page 336, the text:

> The *Product field is embedded, which means that its name ***its*** its type

should be:

> The *Product field is embedded, which means that its name ***is*** its type

(Thanks to Joseph Kashi and Francisco Vicedo for reporting this problem)

---

**Chapter 16**

On page 417, Table 16-4 gives the impression that the `Title` function transforms a string value, rather than creating a new string.

(Thanks to Francisco Vicedo for reporting this problem)

---

On Page 421, Listing 16-8 contains two statements that use the `LastIndex` function and omits the `LastIndexAny` function.

(Thanks to Francisco Vicedo for reporting this problem)
