**VRPL:{{.Quotation.Region}}:{{.Quotation.MachineType}}:REN:AMC:{{.Quotation.BankName}}:{{.MonthYear}}:{{.Quotation.RefNo}}**                                         **Date:{{.Date}}**
<br/>                           
{{.Quotation.Address}}

Dear Sir,

**Sub: Renewal offer For Comphrensive Annual Maintenance Contract For Your Currency Counting/ Detecting Machines ({{machineNames .Quotation.Machines}})**

This is in connection with the above said contract, we would like to inform you that your AMC period has completed on **{{.Quotation.ExpiryDate}}**

In this connection, we have enclosed herewith **A.M.C.** contract offer letter effective from **Date: {{.Quotation.Period}}** for your kind consideration and confirmation.

**Hope Our above offer is in line with your requirement.**

If any clarification is required please feel free to call us on **9925204929/079-26766642**

Thanking you in the meanwhile and assuring you the best of our service and kind
attention always.

Yours faithfully,


For, **Veb Robomak (P) Ltd.**

<br/>

**Authorized Signatory**

---------------------------

**Comprehensive Annual Maintenance Contract**

VRPL shall maintain the machine specified in this agreement on the terms and condition mention as below.

**Terms & Conditions:**

1. This agreement shall remain enforced for the period off one year from the date
commencement.
2. Before taking the system under agreement as an acceptance, a test will be
carried out by our engineer in presences of the representative of the Institution.
3. Any Government levies imposed by the government; it is under stood that such
levies will payable extra.

**Services Rendered:**

We agree to provide Comprehensive Maintenance services under the agreement to
keep the machine in good working condition.

1. Maintenance will cover both preventive and breakdown calls. This will include the cost of the parts except those considered as consumable like "OK stamp" & "Ink Bottle".
2. Preventive maintenance will be provided Quarterly.
3. The service includes as under: cleaning of spindles, encoding diaphragm change,
belt checking, checking of RPM of the motor, cleaning of air filter & checking of
spindle unit.
4. Maintenance and services will be provided during normal working hours on all
the working days.
5. Break down calls will be responded within 48 hrs.

**Excluding:**

Major up gradation of the system.

Any work external to the machine such as maintenance of the attachment accessories
etc. not originally included in the contract. Repairs of malfunctioning or damaged due to rat bites, damages due to accident, transportation negligence, natural disaster or use of non-standard electrical power and power fluctuation or short circuit.

--------------


**Contract Details:**

Contract Period: **{{.Quotation.Period}}**

Model Name|Rate|GST@18%|Total With Tax|Qty|Sub Total|
----------|----|-------|--------------|---|---------|
55,L|25,L|30,R|25,R|15,L|25,R|
{{- range .Quotation.Machines}}
{{.Model}}|{{.Rate}}|{{.Gst}}|{{.TotalWithTax}}|{{.Qty}}|{{.Total}}|
{{- end}}
GrandTotal|||||{{.Quotation.Total}}|


<br/>

**Payment Details**

Payment Bank Details|A/c Name: Veb Robomak Pvt. Ltd.\nBank Name: SBI\nBranch: Old Padra Road\nA/C No: 33778781620\nIFSC Code: SBIN0010687|
-----------------------|----------------
87,L|88,L|
Payment Terms|{{.Quotation.PaymentTerms}}|

<br/>
Signed on Behalf of Institution                         Signed on Behalf of Veb Robomak Pvt Ltd

<br/>
**Authorized Signature & Stamp.**                   **Authorized Signature & Stamp.**

**Date:**                                                               **Date: {{.Date}}**

