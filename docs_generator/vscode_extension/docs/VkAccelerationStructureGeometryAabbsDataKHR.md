# VkAccelerationStructureGeometryAabbsDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryAabbsDataKHR - Structure specifying axis-aligned bounding box geometry in a bottom-level acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureGeometryAabbsDataKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryAabbsDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    data;
    VkDeviceSize                     stride;
} VkAccelerationStructureGeometryAabbsDataKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `data` is a device or host address to memory containing [VkAabbPositionsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAabbPositionsKHR.html) structures containing position data for each axis-aligned bounding box in the geometry.
- `stride` is the stride in bytes between each entry in `data`. The stride **must** be a multiple of `8`.

## [](#_description)Description

Valid Usage

- [](#VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03545)VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03545  
  `stride` **must** be a multiple of `8`
- [](#VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03820)VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03820  
  `stride` **must** be less than or equal to 232-1

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureGeometryAabbsDataKHR-sType-sType)VUID-VkAccelerationStructureGeometryAabbsDataKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR`
- [](#VUID-VkAccelerationStructureGeometryAabbsDataKHR-pNext-pNext)VUID-VkAccelerationStructureGeometryAabbsDataKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureGeometryDataKHR.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureGeometryAabbsDataKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0