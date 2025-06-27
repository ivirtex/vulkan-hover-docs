# VkDeviceFaultCountsEXT(3) Manual Page

## Name

VkDeviceFaultCountsEXT - Structure specifying device fault information



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceFaultCountsEXT` structure is defined as:

``` c
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultCountsEXT {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           addressInfoCount;
    uint32_t           vendorInfoCount;
    VkDeviceSize       vendorBinarySize;
} VkDeviceFaultCountsEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `addressInfoCount` is the number of
  [VkDeviceFaultAddressInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultAddressInfoEXT.html)
  structures describing either memory accesses which **may** have caused
  a page fault, or the addresses of active instructions at the time of
  the fault.

- `vendorInfoCount` is the number of
  [VkDeviceFaultVendorInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceFaultVendorInfoEXT.html)
  structures describing vendor-specific fault information.

- `vendorBinarySize` is the size in bytes of a vendor-specific binary
  crash dump, which may provide additional information when imported
  into external tools.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDeviceFaultCountsEXT-sType-sType"
  id="VUID-VkDeviceFaultCountsEXT-sType-sType"></a>
  VUID-VkDeviceFaultCountsEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEVICE_FAULT_COUNTS_EXT`

- <a href="#VUID-VkDeviceFaultCountsEXT-pNext-pNext"
  id="VUID-VkDeviceFaultCountsEXT-pNext-pNext"></a>
  VUID-VkDeviceFaultCountsEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_fault](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_fault.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceFaultInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceFaultCountsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
