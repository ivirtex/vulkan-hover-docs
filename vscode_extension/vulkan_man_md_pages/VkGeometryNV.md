# VkGeometryNV(3) Manual Page

## Name

VkGeometryNV - Structure specifying a geometry in a bottom-level
acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkGeometryNV` structure describes geometry in a bottom-level
acceleration structure and is defined as:

``` c
// Provided by VK_NV_ray_tracing
typedef struct VkGeometryNV {
    VkStructureType       sType;
    const void*           pNext;
    VkGeometryTypeKHR     geometryType;
    VkGeometryDataNV      geometry;
    VkGeometryFlagsKHR    flags;
} VkGeometryNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `geometryType` specifies the
  [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html) which this geometry refers
  to.

- `geometry` contains the geometry data as described in
  [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html).

- `flags` has [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsKHR.html)
  describing options for this geometry.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkGeometryNV-geometryType-03503"
  id="VUID-VkGeometryNV-geometryType-03503"></a>
  VUID-VkGeometryNV-geometryType-03503  
  `geometryType` **must** be `VK_GEOMETRY_TYPE_TRIANGLES_NV` or
  `VK_GEOMETRY_TYPE_AABBS_NV`

Valid Usage (Implicit)

- <a href="#VUID-VkGeometryNV-sType-sType"
  id="VUID-VkGeometryNV-sType-sType"></a>
  VUID-VkGeometryNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_GEOMETRY_NV`

- <a href="#VUID-VkGeometryNV-pNext-pNext"
  id="VUID-VkGeometryNV-pNext-pNext"></a>
  VUID-VkGeometryNV-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkGeometryNV-geometryType-parameter"
  id="VUID-VkGeometryNV-geometryType-parameter"></a>
  VUID-VkGeometryNV-geometryType-parameter  
  `geometryType` **must** be a valid
  [VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html) value

- <a href="#VUID-VkGeometryNV-geometry-parameter"
  id="VUID-VkGeometryNV-geometry-parameter"></a>
  VUID-VkGeometryNV-geometry-parameter  
  `geometry` **must** be a valid
  [VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html) structure

- <a href="#VUID-VkGeometryNV-flags-parameter"
  id="VUID-VkGeometryNV-flags-parameter"></a>
  VUID-VkGeometryNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkGeometryFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagBitsKHR.html) values

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_ray_tracing](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_ray_tracing.html),
[VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureInfoNV.html),
[VkGeometryDataNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryDataNV.html),
[VkGeometryFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryFlagsKHR.html),
[VkGeometryTypeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGeometryTypeKHR.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkGeometryNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
