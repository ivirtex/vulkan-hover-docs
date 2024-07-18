# VkAccelerationStructureTrianglesOpacityMicromapEXT(3) Manual Page

## Name

VkAccelerationStructureTrianglesOpacityMicromapEXT - Structure
specifying an opacity micromap in a bottom-level acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureTrianglesOpacityMicromapEXT` structure is
defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkAccelerationStructureTrianglesOpacityMicromapEXT {
    VkStructureType                     sType;
    void*                               pNext;
    VkIndexType                         indexType;
    VkDeviceOrHostAddressConstKHR       indexBuffer;
    VkDeviceSize                        indexStride;
    uint32_t                            baseTriangle;
    uint32_t                            usageCountsCount;
    const VkMicromapUsageEXT*           pUsageCounts;
    const VkMicromapUsageEXT* const*    ppUsageCounts;
    VkMicromapEXT                       micromap;
} VkAccelerationStructureTrianglesOpacityMicromapEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `indexType` is the type of triangle indices used when indexing this
  micromap

- `indexBuffer` is the address containing the triangle indices

- `indexStride` is the byte stride between triangle indices

- `baseTriangle` is the base value added to the non-negative triangle
  indices

- `usageCountsCount` specifies the number of usage counts structures
  that will be used to determine the size of this micromap.

- `pUsageCounts` is a pointer to an array of
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures.

- `ppUsageCounts` is a pointer to an array of pointers to
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures.

- `micromap` is the handle to the micromap object to include in this
  geometry

## <a href="#_description" class="anchor"></a>Description

If `VkAccelerationStructureTrianglesOpacityMicromapEXT` is included in
the `pNext` chain of a
[VkAccelerationStructureGeometryTrianglesDataKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureGeometryTrianglesDataKHR.html)
structure, that geometry will reference that micromap.

For each triangle in the geometry, the acceleration structure build
fetches an index from `indexBuffer` using `indexType` and `indexStride`.
If that value is the unsigned cast of one of the values from
[VkOpacityMicromapSpecialIndexEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkOpacityMicromapSpecialIndexEXT.html)
then that triangle behaves as described for that special value in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-opacity-micromap"
target="_blank" rel="noopener">Ray Opacity Micromap</a>. Otherwise that
triangle uses the opacity micromap information from `micromap` at that
index plus `baseTriangle`.

Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid
pointer, the other **must** be `NULL`. The elements of the non-`NULL`
array describe the total count used to build this geometry. For a given
`format` and `subdivisionLevel` the number of triangles in this geometry
matching those values after indirection and special index handling
**must** be equal to the sum of matching `count` provided.

If `micromap` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), then every value
read from `indexBuffer` **must** be one of the values in
`VkOpacityMicromapSpecialIndexEXT`.

Valid Usage

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-07335"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-07335"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-07335  
  Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid
  pointer, the other **must** be `NULL`

Valid Usage (Implicit)

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-sType-sType"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-sType-sType"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_TRIANGLES_OPACITY_MICROMAP_EXT`

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-parameter"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-parameter"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-parameter"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-parameter"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-pUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `pUsageCounts` is not `NULL`,
  `pUsageCounts` **must** be a valid pointer to an array of
  `usageCountsCount` [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html)
  structures

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-ppUsageCounts-parameter"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-ppUsageCounts-parameter"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-ppUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `ppUsageCounts` is not `NULL`,
  `ppUsageCounts` **must** be a valid pointer to an array of
  `usageCountsCount` valid pointers to
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures

- <a
  href="#VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-micromap-parameter"
  id="VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-micromap-parameter"></a>
  VUID-VkAccelerationStructureTrianglesOpacityMicromapEXT-micromap-parameter  
  If `micromap` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `micromap`
  **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html),
[VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureTrianglesOpacityMicromapEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
