# VkAabbPositionsKHR(3) Manual Page

## Name

VkAabbPositionsKHR - Structure specifying two opposing corners of an
axis-aligned bounding box



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAabbPositionsKHR` structure is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAabbPositionsKHR {
    float    minX;
    float    minY;
    float    minZ;
    float    maxX;
    float    maxY;
    float    maxZ;
} VkAabbPositionsKHR;
```

or the equivalent

``` c
// Provided by VK_NV_ray_tracing
typedef VkAabbPositionsKHR VkAabbPositionsNV;
```

## <a href="#_members" class="anchor"></a>Members

- `minX` is the x position of one opposing corner of a bounding box.

- `minY` is the y position of one opposing corner of a bounding box.

- `minZ` is the z position of one opposing corner of a bounding box.

- `maxX` is the x position of the other opposing corner of a bounding
  box.

- `maxY` is the y position of the other opposing corner of a bounding
  box.

- `maxZ` is the z position of the other opposing corner of a bounding
  box.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkAabbPositionsKHR-minX-03546"
  id="VUID-VkAabbPositionsKHR-minX-03546"></a>
  VUID-VkAabbPositionsKHR-minX-03546  
  `minX` **must** be less than or equal to `maxX`

- <a href="#VUID-VkAabbPositionsKHR-minY-03547"
  id="VUID-VkAabbPositionsKHR-minY-03547"></a>
  VUID-VkAabbPositionsKHR-minY-03547  
  `minY` **must** be less than or equal to `maxY`

- <a href="#VUID-VkAabbPositionsKHR-minZ-03548"
  id="VUID-VkAabbPositionsKHR-minZ-03548"></a>
  VUID-VkAabbPositionsKHR-minZ-03548  
  `minZ` **must** be less than or equal to `maxZ`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAabbPositionsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
