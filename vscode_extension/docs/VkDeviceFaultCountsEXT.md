# VkDeviceFaultCountsEXT(3) Manual Page

## Name

VkDeviceFaultCountsEXT - Structure specifying device fault information



## [](#_c_specification)C Specification

The `VkDeviceFaultCountsEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultCountsEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           addressInfoCount;
    uint32_t           vendorInfoCount;
    VkDeviceSize       vendorBinarySize;
} VkDeviceFaultCountsEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `addressInfoCount` is the number of [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressInfoEXT.html) structures describing either memory accesses which **may** have caused a page fault, or the addresses of active instructions at the time of the fault.
- `vendorInfoCount` is the number of [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultVendorInfoEXT.html) structures describing vendor-specific fault information.
- `vendorBinarySize` is the size in bytes of a vendor-specific binary crash dump, which may provide additional information when imported into external tools.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDeviceFaultCountsEXT-sType-sType)VUID-VkDeviceFaultCountsEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_FAULT_COUNTS_EXT`
- [](#VUID-VkDeviceFaultCountsEXT-pNext-pNext)VUID-VkDeviceFaultCountsEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceFaultInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultCountsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0