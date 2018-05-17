**Date:**{{datesFmt .Date}}
<br/>
{{- if eq .Quotation.QuotationType "NEW"}}
**REF:**VRPL:{{.Quotation.Region}}:{{.Quotation.MachineType}}:NEW:AMC:{{.Quotation.RefNo}}
{{- else if eq .Quotation.QuotationType "REN"}}
**REF:**VRPL:{{.Quotation.Region}}:{{.Quotation.MachineType}}:REN:AMC:{{.Quotation.RefNo}}
{{- else}}
**REF:**No Refernce Found For this Quotation Type
{{- end}}
<br/>                          
{{.Quotation.Address}}

Dear Sir,

{{- if eq .Quotation.QuotationType "NEW"}}
**Sub: Offer for Non Comphrensive Annual Maintenance Contract for Your Note Counting Machine ({{machineNames .Quotation.Machines}})**

We take privilege to introduce our self as a leading Manufacturer and Maintenance House. For past three decades we have been providing services in Nationalized Banks, Co-operative Banks and Corporate. We are supplying Note Counting Machines and providing maintenance services for the same and other brands as well.

In this connection we are pleased to submit our fresh annual maintenance contract for your Note Counting Machine for your kind consideration.
{{- else if eq .Quotation.QuotationType "REN"}}
**Sub: Renewal offer for Non Comphrensive Annual Maintenance Contract for Your Note Counting Machine ({{machineNames .Quotation.Machines}})**

This is in connection with the above said contract, we would like to inform you that your AMC period has completed on **{{.Quotation.ExpiryDate}}**.

In this connection, we have enclosed herewith **A.M.C.** contract offer letter effective from **Date: {{datesFmt .Quotation.Period}}** for your kind consideration and confirmation.

**Hope Our above offer is in line with your requirement.**
{{- else}}
**Sub: Subject Not Found For This Quotation Type**

Content Not Found for this Quotation Type
{{- end}}

If any clarification needed please feel free to call us on: **079-26424229/99252 04929**.

Thanking you in the meanwhile and assuring you the best of our service and kind attention always.

Yours faithfully,


For, **Veb Robomak (P) Ltd.**

<br/>

**Authorized Signatory**

---------------------------

**Non Comprehensive Annual Maintenance Contract**

VRPL shall maintain the machine specified in this agreement on the terms and condition mention as below.

**Terms & Conditions:**

1. This agreement shall remain enforced for the period off one year from the date commencement.
2. Before taking the system under agreement as an acceptance, a test will be carried out by our engineer in presences of the representative of the Institution.
3. Any Government levies imposed by the government; it is understood that such levies will payable extra.

**Services Rendered:**

We agree to provide Comprehensive Maintenance services under the agreement to keep the machine in good working condition.

1. Maintenance will cover both preventive and breakdown calls. 
2. Preventive maintenance will be provided Quarterly.
3. The service includes as under: cleaning of spindles, encoding diaphragm change, belt checking, checking of RPM of the motor, cleaning of air filter & checking of
spindle unit.
4. Maintenance and services will be provided during normal working hours on all the working days.
5. Break down calls will be responded within 48 hrs.
6. Any parts damaged or replaced cost for the same will be borne by the Institution.

--------------


**Contract Details:**

Contract Period: **{{datesFmt .Quotation.Period}}**

Model Name|Rate|GST@18%|Total With Tax|Qty|Sub Total|
----------|----|-------|--------------|---|---------|
50,L|30,R|25,R|30,R|14,L|32,R|
{{- range .Quotation.Machines}}
{{.Model}}|{{indianCurr .Rate}}|{{indianCurrF .Gst}}|{{indianCurrF .TotalWithTax}}|{{.Qty}}|{{indianCurrF .Total}}|
{{- end}}
{{- if .Quotation.RoundOff}}
RoundOff|||||{{indianCurrF .Quotation.RoundOff}}|
{{- end}}
GrandTotal|||||{{indianCurrF .Quotation.Total}}|


<br/>

**Payment Details**

{{if eq .Quotation.Region "BRD"}}
Payment Bank Details|A/c Name: Veb Robomak Pvt. Ltd.\nBank Name: SBI\nBranch: Old Padra Road\nA/C No: 33778781620\nIFSC Code: SBIN0010687|
{{- else}}
Payment Bank Details|A/c Name: Veb Robomak Pvt. Ltd.\nBank Name: SBI\nBranch: Polytechnic , Ahmedabad\nA/C No: 30956461892\nIFSC Code: SBIN0001043|
{{- end}}
-----------------------|----------------
80,L|101,L|
Payment Terms|{{.Quotation.PaymentTerms}}|

<br/>
Signed on Behalf of Institution                Signed on Behalf of Veb Robomak Pvt Ltd

<br/>

**Authorized Signature & Stamp.**           **Authorized Signature & Stamp.**

**Date:**                                                       **Date: {{datesFmt .Date}}**
