# VkMicromapBuildInfoEXT(3) Manual Page

## Name

VkMicromapBuildInfoEXT - Structure specifying the data used to build a
micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapBuildInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapBuildInfoEXT {
    VkStructureType                     sType;
    const void*                         pNext;
    VkMicromapTypeEXT                   type;
    VkBuildMicromapFlagsEXT             flags;
    VkBuildMicromapModeEXT              mode;
    VkMicromapEXT                       dstMicromap;
    uint32_t                            usageCountsCount;
    const VkMicromapUsageEXT*           pUsageCounts;
    const VkMicromapUsageEXT* const*    ppUsageCounts;
    VkDeviceOrHostAddressConstKHR       data;
    VkDeviceOrHostAddressKHR            scratchData;
    VkDeviceOrHostAddressConstKHR       triangleArray;
    VkDeviceSize                        triangleArrayStride;
} VkMicromapBuildInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `type` is a [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html) value
  specifying the type of micromap being built.

- `flags` is a bitmask of
  [VkBuildMicromapFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagBitsEXT.html)
  specifying additional parameters of the micromap.

- `mode` is a [VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapModeEXT.html)
  value specifying the type of operation to perform.

- `dstMicromap` is a pointer to the target micromap for the build.

- `usageCountsCount` specifies the number of usage counts structures
  that will be used to determine the size of this micromap.

- `pUsageCounts` is a pointer to an array of
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures.

- `ppUsageCounts` is a pointer to an array of pointers to
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures.

- `data` is the device or host address to memory which contains the data
  for the micromap.

- `scratchData` is the device or host address to memory that will be
  used as scratch memory for the build.

- `triangleArray` is the device or host address to memory containing the
  [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html) data

- `triangleArrayStride` is the stride in bytes between each element of
  `triangleArray`

## <a href="#_description" class="anchor"></a>Description

Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid
pointer, the other **must** be `NULL`. The elements of the non-`NULL`
array describe the total counts used to build each micromap. Each
element contains a `count` which is the number of micromap triangles of
that `format` and `subdivisionLevel` contained in the micromap. Multiple
elements with the same `format` and `subdivisionLevel` are allowed and
the total count for that `format` and `subdivisionLevel` is the sum of
the `count` for each element.

Each micromap triangle refers to one element in `triangleArray` which
contains the `format` and `subdivisionLevel` for that particular
triangle as well as a `dataOffset` in bytes which is the location
relative to `data` where that triangle’s micromap data begins. The data
at `triangleArray` is laid out as a 4 byte unsigned integer for the
`dataOffset` followed by a 2 byte unsigned integer for the subdivision
level then a 2 byte unsigned integer for the format. In practice,
compilers compile [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html) to
match this pattern.

For opacity micromaps, the data at `data` is packed as either one bit
per element for `VK_OPACITY_MICROMAP_FORMAT_2_STATE_EXT` or two bits per
element for `VK_OPACITY_MICROMAP_FORMAT_4_STATE_EXT` and is packed from
LSB to MSB in each byte. The data at each index in those bytes is
interpreted as discussed in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#ray-opacity-micromap"
target="_blank" rel="noopener">Ray Opacity Micromap</a>.

For displacement micromaps, the data at `data` is interpreted as
discussed in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#displacement-micromap-encoding"
target="_blank" rel="noopener">Displacement Micromap Encoding</a>.

Valid Usage

- <a href="#VUID-VkMicromapBuildInfoEXT-pUsageCounts-07516"
  id="VUID-VkMicromapBuildInfoEXT-pUsageCounts-07516"></a>
  VUID-VkMicromapBuildInfoEXT-pUsageCounts-07516  
  Only one of `pUsageCounts` or `ppUsageCounts` **can** be a valid
  pointer, the other **must** be `NULL`

- <a href="#VUID-VkMicromapBuildInfoEXT-type-07517"
  id="VUID-VkMicromapBuildInfoEXT-type-07517"></a>
  VUID-VkMicromapBuildInfoEXT-type-07517  
  If `type` is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` the `format`
  member of [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) **must** be a
  valid value from `VkOpacityMicromapFormatEXT`

- <a href="#VUID-VkMicromapBuildInfoEXT-type-07518"
  id="VUID-VkMicromapBuildInfoEXT-type-07518"></a>
  VUID-VkMicromapBuildInfoEXT-type-07518  
  If `type` is `VK_MICROMAP_TYPE_OPACITY_MICROMAP_EXT` the `format`
  member of [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html) **must**
  be a valid value from `VkOpacityMicromapFormatEXT`

- <a href="#VUID-VkMicromapBuildInfoEXT-type-08704"
  id="VUID-VkMicromapBuildInfoEXT-type-08704"></a>
  VUID-VkMicromapBuildInfoEXT-type-08704  
  If `type` is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` the `format`
  member of [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) **must** be a
  valid value from `VkDisplacementMicromapFormatNV`

- <a href="#VUID-VkMicromapBuildInfoEXT-type-08705"
  id="VUID-VkMicromapBuildInfoEXT-type-08705"></a>
  VUID-VkMicromapBuildInfoEXT-type-08705  
  If `type` is `VK_MICROMAP_TYPE_DISPLACEMENT_MICROMAP_NV` the `format`
  member of [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTriangleEXT.html) **must**
  be a valid value from `VkDisplacementMicromapFormatNV`

Valid Usage (Implicit)

- <a href="#VUID-VkMicromapBuildInfoEXT-sType-sType"
  id="VUID-VkMicromapBuildInfoEXT-sType-sType"></a>
  VUID-VkMicromapBuildInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_BUILD_INFO_EXT`

- <a href="#VUID-VkMicromapBuildInfoEXT-pNext-pNext"
  id="VUID-VkMicromapBuildInfoEXT-pNext-pNext"></a>
  VUID-VkMicromapBuildInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

- <a href="#VUID-VkMicromapBuildInfoEXT-type-parameter"
  id="VUID-VkMicromapBuildInfoEXT-type-parameter"></a>
  VUID-VkMicromapBuildInfoEXT-type-parameter  
  `type` **must** be a valid [VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html)
  value

- <a href="#VUID-VkMicromapBuildInfoEXT-flags-parameter"
  id="VUID-VkMicromapBuildInfoEXT-flags-parameter"></a>
  VUID-VkMicromapBuildInfoEXT-flags-parameter  
  `flags` **must** be a valid combination of
  [VkBuildMicromapFlagBitsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagBitsEXT.html) values

- <a href="#VUID-VkMicromapBuildInfoEXT-pUsageCounts-parameter"
  id="VUID-VkMicromapBuildInfoEXT-pUsageCounts-parameter"></a>
  VUID-VkMicromapBuildInfoEXT-pUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `pUsageCounts` is not `NULL`,
  `pUsageCounts` **must** be a valid pointer to an array of
  `usageCountsCount` [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html)
  structures

- <a href="#VUID-VkMicromapBuildInfoEXT-ppUsageCounts-parameter"
  id="VUID-VkMicromapBuildInfoEXT-ppUsageCounts-parameter"></a>
  VUID-VkMicromapBuildInfoEXT-ppUsageCounts-parameter  
  If `usageCountsCount` is not `0`, and `ppUsageCounts` is not `NULL`,
  `ppUsageCounts` **must** be a valid pointer to an array of
  `usageCountsCount` valid pointers to
  [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html) structures

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkBuildMicromapFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapFlagsEXT.html),
[VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuildMicromapModeEXT.html),
[VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressConstKHR.html),
[VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceOrHostAddressKHR.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html),
[VkMicromapTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapTypeEXT.html),
[VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapUsageEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkBuildMicromapsEXT.html),
[vkCmdBuildMicromapsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBuildMicromapsEXT.html),
[vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapBuildInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
