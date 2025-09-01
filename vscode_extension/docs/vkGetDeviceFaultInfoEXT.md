# vkGetDeviceFaultInfoEXT(3) Manual Page

## Name

vkGetDeviceFaultInfoEXT - Reports diagnostic fault information on the specified logical device



## [](#_c_specification)C Specification

To retrieve diagnostic information about faults that **may** have caused device loss, call:

```c++
// Provided by VK_EXT_device_fault
VkResult vkGetDeviceFaultInfoEXT(
    VkDevice                                    device,
    VkDeviceFaultCountsEXT*                     pFaultCounts,
    VkDeviceFaultInfoEXT*                       pFaultInfo);
```

## [](#_parameters)Parameters

- `device` is the logical device from which to query the diagnostic fault information.
- `pFaultCounts` is a pointer to a [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultCountsEXT.html) structure in which counts for structures describing additional fault information are returned.
- `pFaultInfo` is `NULL` or a pointer to a [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html) structure in which fault information is returned.

## [](#_description)Description

If `pFaultInfo` is `NULL`, then the counts of corresponding additional fault information structures available are returned in the `addressInfoCount` and `vendorInfoCount` members of `pFaultCounts`. Additionally, the size of any vendor-specific binary crash dump is returned in the `vendorBinarySize` member of `pFaultCounts`.

If `pFaultInfo` is not `NULL`, `pFaultCounts` **must** point to a [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultCountsEXT.html) structure with each structure count or size member (`addressInfoCount`, `vendorInfoCount`, `vendorBinarySize`) set by the application to the number of elements in the corresponding output array member of `pFaultInfo` (`pAddressInfos` and `pVendorInfos`), or to the size of the output buffer in bytes (`pVendorBinaryData`). On return, each structure count member is overwritten with the number of structures actually written to the corresponding output array member of `pFaultInfo`. Similarly, `vendorBinarySize` is overwritten with the number of bytes actually written to the `pVendorBinaryData` member of `pFaultInfo`.

If the [vendor-specific crash dumps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-deviceFaultVendorBinary) feature is not enabled, then implementations **must** set `pFaultCounts`-&gt;vendorBinarySize to zero and **must** not modify `pFaultInfo`-&gt;pVendorBinaryData.

If any `pFaultCounts` structure count member is less than the number of corresponding fault properties available, at most structure count (`addressInfoCount`, `vendorInfoCount`) elements will be written to the associated `pFaultInfo` output array. Similarly, if `vendorBinarySize` is less than the size in bytes of the available crash dump data, at most `vendorBinarySize` elements will be written to `pVendorBinaryData`.

If `pFaultInfo` is `NULL`, then subsequent calls to [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html) for the same `device` **must** return identical values in the `addressInfoCount`, `vendorInfoCount` and `vendorBinarySize` members of `pFaultCounts`.

If `pFaultInfo` is not `NULL`, then subsequent calls to [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html) for the same `device` **must** return identical values in the output members of `pFaultInfo` (`pAddressInfos`, `pVendorInfos`, `pVendorBinaryData`), up to the limits described by the structure count and buffer size members of `pFaultCounts` (`addressInfoCount`, `vendorInfoCount`, `vendorBinarySize`). If the sizes of the output members of `pFaultInfo` increase for a subsequent call to [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html), then supplementary information **may** be returned in the additional available space.

If any `pFaultCounts` structure count member is smaller than the number of corresponding fault properties available, or if `pFaultCounts`-&gt;vendorBinarySize is smaller than the size in bytes of the generated binary crash dump data, `VK_INCOMPLETE` will be returned instead of `VK_SUCCESS`, to indicate that not all the available properties were returned.

If `pFaultCounts`-&gt;vendorBinarySize is less than what is necessary to store the [binary crash dump header](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vendor-binary-crash-dumps), nothing will be written to `pFaultInfo`-&gt;pVendorBinaryData and zero will be written to `pFaultCounts`-&gt;vendorBinarySize.

Valid Usage

- [](#VUID-vkGetDeviceFaultInfoEXT-device-07336)VUID-vkGetDeviceFaultInfoEXT-device-07336  
  `device` **must** be in the *lost* state
- [](#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07337)VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07337  
  If the value referenced by `pFaultCounts->addressInfoCount` is not `0`, and `pFaultInfo->pAddressInfos` is not `NULL`, `pFaultInfo->pAddressInfos` **must** be a valid pointer to an array of `pFaultCounts->addressInfoCount` [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) structures
- [](#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07338)VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07338  
  If the value referenced by `pFaultCounts->vendorInfoCount` is not `0`, and `pFaultInfo->pVendorInfos` is not `NULL`, `pFaultInfo->pVendorInfos` **must** be a valid pointer to an array of `pFaultCounts->vendorInfoCount` [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorInfoEXT.html) structures
- [](#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07339)VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-07339  
  If the value referenced by `pFaultCounts->vendorBinarySize` is not `0`, and `pFaultInfo->pVendorBinaryData` is not `NULL`, `pFaultInfo->pVendorBinaryData` **must** be a valid pointer to an array of `pFaultCounts->vendorBinarySize` bytes

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceFaultInfoEXT-device-parameter)VUID-vkGetDeviceFaultInfoEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-parameter)VUID-vkGetDeviceFaultInfoEXT-pFaultCounts-parameter  
  `pFaultCounts` **must** be a valid pointer to a [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultCountsEXT.html) structure
- [](#VUID-vkGetDeviceFaultInfoEXT-pFaultInfo-parameter)VUID-vkGetDeviceFaultInfoEXT-pFaultInfo-parameter  
  If `pFaultInfo` is not `NULL`, `pFaultInfo` **must** be a valid pointer to a [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html) structure

Return Codes

On success, this command returns

- `VK_INCOMPLETE`
- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceFaultCountsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultCountsEXT.html), [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceFaultInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0