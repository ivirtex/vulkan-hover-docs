# VkAccelerationStructureDeviceAddressInfoKHR(3) Manual Page

## Name

VkAccelerationStructureDeviceAddressInfoKHR - Structure specifying the
acceleration structure to query an address for



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureDeviceAddressInfoKHR` structure is defined
as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureDeviceAddressInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkAccelerationStructureKHR    accelerationStructure;
} VkAccelerationStructureDeviceAddressInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructure` specifies the acceleration structure whose
  address is being queried.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureDeviceAddressInfoKHR-sType-sType"
  id="VUID-VkAccelerationStructureDeviceAddressInfoKHR-sType-sType"></a>
  VUID-VkAccelerationStructureDeviceAddressInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DEVICE_ADDRESS_INFO_KHR`

- <a href="#VUID-VkAccelerationStructureDeviceAddressInfoKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureDeviceAddressInfoKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureDeviceAddressInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

- <a
  href="#VUID-VkAccelerationStructureDeviceAddressInfoKHR-accelerationStructure-parameter"
  id="VUID-VkAccelerationStructureDeviceAddressInfoKHR-accelerationStructure-parameter"></a>
  VUID-VkAccelerationStructureDeviceAddressInfoKHR-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureDeviceAddressInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
