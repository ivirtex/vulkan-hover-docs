# VkAccelerationStructureGeometryAabbsDataKHR(3) Manual Page

## Name

VkAccelerationStructureGeometryAabbsDataKHR - Structure specifying
axis-aligned bounding box geometry in a bottom-level acceleration
structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureGeometryAabbsDataKHR` structure is defined
as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureGeometryAabbsDataKHR {
    VkStructureType                  sType;
    const void*                      pNext;
    VkDeviceOrHostAddressConstKHR    data;
    VkDeviceSize                     stride;
} VkAccelerationStructureGeometryAabbsDataKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `data` is a device or host address to memory containing
  [VkAabbPositionsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAabbPositionsKHR.html) structures containing
  position data for each axis-aligned bounding box in the geometry.

- `stride` is the stride in bytes between each entry in `data`. The
  stride **must** be a multiple of `8`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03545"
  id="VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03545"></a>
  VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03545  
  `stride` **must** be a multiple of `8`

- <a href="#VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03820"
  id="VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03820"></a>
  VUID-VkAccelerationStructureGeometryAabbsDataKHR-stride-03820  
  `stride` **must** be less than or equal to 2<sup>32</sup>-1

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureGeometryAabbsDataKHR-sType-sType"
  id="VUID-VkAccelerationStructureGeometryAabbsDataKHR-sType-sType"></a>
  VUID-VkAccelerationStructureGeometryAabbsDataKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_GEOMETRY_AABBS_DATA_KHR`

- <a href="#VUID-VkAccelerationStructureGeometryAabbsDataKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureGeometryAabbsDataKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureGeometryAabbsDataKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkAccelerationStructureGeometryDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryDataKHR.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureGeometryAabbsDataKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
