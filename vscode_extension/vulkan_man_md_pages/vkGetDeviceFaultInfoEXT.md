# vkGetDeviceFaultInfoEXT(3) Manual Page

## Name

vkGetDeviceFaultInfoEXT - Reports diagnostic fault information on the
specified logical device



## <a href="#_c_specification" class="anchor"></a>C Specification

To retrieve diagnostic information about faults that **may** have caused
device loss, call:

``` c
// Provided by VK_EXT_device_fault
VkResult vkGetDeviceFaultInfoEXT(
    VkDevice                                    device,
    VkDeviceFaultCountsEXT*                     pFaultCounts,
    VkDeviceFaultInfoEXT*                       pFaultInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device from which to query the diagnostic
  fault information.

- `pFaultCounts` is a pointer to a
  [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultCountsEXT.html) structure in
  which counts for structures describing additional fault information
  are returned.

- `pFaultInfo` is `NULL` or a pointer to a
  [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html) structure in which
  fault information is returned.

## <a href="#_description" class="anchor"></a>Description

If `pFaultInfo` is `NULL`, then the counts of corresponding additional
fault information structures available are returned in the
`addressInfoCount` and `vendorInfoCount` members of `pFaultCounts`.
Additionally, the size of any vendor-specific binary crash dump is
returned in the `vendorBinarySize` member of `pFaultCounts`.

If `pFaultInfo` is not `NULL`, `pFaultCounts` **must** point to a
[VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultCountsEXT.html) structure with
each structure count or size member (`addressInfoCount`,
`vendorInfoCount`, `vendorBinarySize`) set by the application to the
number of elements in the corresponding output array member of
`pFaultInfo` (`pAddressInfos` and `pVendorInfos`), or to the size of the
output buffer in bytes (`pVendorBinaryData`). On return, each structure
count member is overwritten with the number of structures actually
written to the corresponding output array member of `pFaultInfo`.
Similarly, `vendorBinarySize` is overwritten with the number of bytes
actually written to the `pVendorBinaryData` member of `pFaultInfo`.

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-deviceFaultVendorBinary"
target="_blank" rel="noopener">vendor-specific crash dumps</a> feature
is not enabled, then implementations **must** set
`pFaultCounts`-\>vendorBinarySize to zero and **must** not modify
`pFaultInfo`-\>pVendorBinaryData.

If any `pFaultCounts` structure count member is less than the number of
corresponding fault properties available, at most structure count
(`addressInfoCount`, `vendorInfoCount`) elements will be written to the
associated `pFaultInfo` output array. Similarly, if `vendorBinarySize`
is less than the size in bytes of the available crash dump data, at most
`vendorBinarySize` elements will be written to `pVendorBinaryData`.

If `pFaultInfo` is `NULL`, then subsequent calls to
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html) for the same
`device` **must** return identical values in the `addressInfoCount`,
`vendorInfoCount` and `vendorBinarySize` members of `pFaultCounts`.

If `pFaultInfo` is not `NULL`, then subsequent calls to
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html) for the same
`device` **must** return identical values in the output members of
`pFaultInfo` (`pAddressInfos`, `pVendorInfos`, `pVendorBinaryData`), up
to the limits described by the structure count and buffer size members
of `pFaultCounts` (`addressInfoCount`, `vendorInfoCount`,
`vendorBinarySize`). If the sizes of the output members of `pFaultInfo`
increase for a subsequent call to
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html), then
supplementary information **may** be returned in the additional
available space.

If any `pFaultCounts` structure count member is smaller than the number
of corresponding fault properties available, or if
`pFaultCounts`-\>vendorBinarySize is smaller than the size in bytes of
the generated binary crash dump data, `VK_INCOMPLETE` will be returned
instead of `VK_SUCCESS`, to indicate that not all the available
properties were returned.

If `pFaultCounts`-\>vendorBinarySize is less than what is necessary to
store the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vendor-binary-crash-dumps"
target="_blank" rel="noopener">binary crash dump header</a>, nothing
will be written to `pFaultInfo`-\>pVendorBinaryData and zero will be
written to `pFaultCounts`-\>vendorBinarySize.

Valid Usage

- <a href="#VUID-vkGetDeviceFaultInfoEXT-device-07336"
  id="VUID-vkGetDeviceFaultInfoEXT-device-07336"></a>
  VUID-vkGetDeviceFaultInfoEXT-device-07336  
  `device` **must** be in the *lost* state

- <a href="#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07337"
  id="VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07337"></a>
  VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07337  
  If the value referenced by `pFaultCounts->addressInfoCount` is not
  `0`, and `pFaultInfo->pAddressInfos` is not `NULL`,
  `pFaultInfo->pAddressInfos` **must** be a valid pointer to an array of
  `pFaultCounts->addressInfoCount`
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)
  structures

- <a href="#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07338"
  id="VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07338"></a>
  VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07338  
  If the value referenced by `pFaultCounts->vendorInfoCount` is not `0`,
  and `pFaultInfo->pVendorInfos` is not `NULL`,
  `pFaultInfo->pVendorInfos` **must** be a valid pointer to an array of
  `pFaultCounts->vendorInfoCount`
  [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorInfoEXT.html)
  structures

- <a href="#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07339"
  id="VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07339"></a>
  VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07339  
  If the value referenced by `pFaultCounts->vendorBinarySize` is not
  `0`, and `pFaultInfo->pVendorBinaryData` is not `NULL`,
  `pFaultInfo->pVendorBinaryData` **must** be a valid pointer to an
  array of `pFaultCounts->vendorBinarySize` bytes

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceFaultInfoEXT-device-parameter"
  id="VUID-vkGetDeviceFaultInfoEXT-device-parameter"></a>
  VUID-vkGetDeviceFaultInfoEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-parameter"
  id="VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-parameter"></a>
  VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-parameter  
  `pFaultCounts` **must** be a valid pointer to a
  [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultCountsEXT.html) structure

- <a href="#VUID-vkGetDeviceFaultInfoEXT-pFaultInfo-parameter"
  id="VUID-vkGetDeviceFaultInfoEXT-pFaultInfo-parameter"></a>
  VUID-vkGetDeviceFaultInfoEXT-pFaultInfo-parameter  
  If `pFaultInfo` is not `NULL`, `pFaultInfo` **must** be a valid
  pointer to a [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_INCOMPLETE`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultCountsEXT.html),
[VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceFaultInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
