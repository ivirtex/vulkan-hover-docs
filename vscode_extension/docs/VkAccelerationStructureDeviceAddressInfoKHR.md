# VkAccelerationStructureDeviceAddressInfoKHR(3) Manual Page

## Name

VkAccelerationStructureDeviceAddressInfoKHR - Structure specifying the acceleration structure to query an address for



## [](#_c_specification)C Specification

The `VkAccelerationStructureDeviceAddressInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureDeviceAddressInfoKHR {
    VkStructureType               sType;
    const void*                   pNext;
    VkAccelerationStructureKHR    accelerationStructure;
} VkAccelerationStructureDeviceAddressInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructure` specifies the acceleration structure whose address is being queried.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureDeviceAddressInfoKHR-sType-sType)VUID-VkAccelerationStructureDeviceAddressInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_DEVICE_ADDRESS_INFO_KHR`
- [](#VUID-VkAccelerationStructureDeviceAddressInfoKHR-pNext-pNext)VUID-VkAccelerationStructureDeviceAddressInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAccelerationStructureDeviceAddressInfoKHR-accelerationStructure-parameter)VUID-VkAccelerationStructureDeviceAddressInfoKHR-accelerationStructure-parameter  
  `accelerationStructure` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetAccelerationStructureDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureDeviceAddressKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureDeviceAddressInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0