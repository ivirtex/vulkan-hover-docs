# VkDeviceFaultInfoEXT(3) Manual Page

## Name

VkDeviceFaultInfoEXT - Structure specifying device fault information



## [](#_c_specification)C Specification

The `VkDeviceFaultInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultInfoEXT {
    VkStructureType                 sType;
    void*                           pNext;
    char                            description[VK_MAX_DESCRIPTION_SIZE];
    VkDeviceFaultAddressInfoEXT*    pAddressInfos;
    VkDeviceFaultVendorInfoEXT*     pVendorInfos;
    void*                           pVendorBinaryData;
} VkDeviceFaultInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a null-terminated UTF-8 string which is a human readable description of the fault.
- `pAddressInfos` is `NULL` or a pointer to an array of [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) structures describing either memory accesses which **may** have caused a page fault, or describing active instruction pointers at the time of the fault. If not `NULL`, each element of `pAddressInfos` describes the a bounded region of GPU virtual address space containing either the GPU virtual address accessed, or the value of an active instruction pointer.
- `pVendorInfos` is `NULL` or a pointer to an array of [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorInfoEXT.html) structures describing vendor-specific fault information.
- `pVendorBinaryData` is `NULL` or a pointer to `vendorBinarySize` number of bytes of data, which will be populated with a vendor-specific binary crash dump, as described in [Vendor Binary Crash Dumps](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vendor-binary-crash-dumps).

## [](#_description)Description

An implementation **should** populate as many members of [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html) as possible, given the information available at the time of the fault and the constraints of the implementation itself.

Due to hardware limitations, `pAddressInfos` describes ranges of GPU virtual address space, rather than precise addresses. The precise memory address accessed or the precise value of the instruction pointer **must** lie within the region described.

Note

Each element of `pAddressInfos` describes either:

- A memory access which may have triggered a page fault and may have contributed to device loss
- The value of an active instruction pointer at the time a fault occurred. This value may be indicative of the active pipeline or shader at the time of device loss

Comparison of the GPU virtual addresses described by `pAddressInfos` to GPU virtual address ranges reported by the `VK_EXT_device_address_binding_report` extension may allow applications to correlate between these addresses and Vulkan objects. Applications should be aware that these addresses may also correspond to resources internal to an implementation, which will not be reported via the `VK_EXT_device_address_binding_report` extension.

Valid Usage (Implicit)

- [](#VUID-VkDeviceFaultInfoEXT-sType-sType)VUID-VkDeviceFaultInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_FAULT_INFO_EXT`
- [](#VUID-VkDeviceFaultInfoEXT-pNext-pNext)VUID-VkDeviceFaultInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html), [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorInfoEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0