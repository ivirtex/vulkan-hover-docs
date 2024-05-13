# VkMicromapBuildSizesInfoEXT(3) Manual Page

## Name

VkMicromapBuildSizesInfoEXT - Structure specifying build sizes for a
micromap



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkMicromapBuildSizesInfoEXT` structure describes the required build
sizes for a micromap and scratch buffers and is defined as:

``` c
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapBuildSizesInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       micromapSize;
    VkDeviceSize       buildScratchSize;
    VkBool32           discardable;
} VkMicromapBuildSizesInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `micromapSize` is the size in bytes required in a
  [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMicromapEXT.html) for a build or update operation.

- `buildScratchSize` is the size in bytes required in a scratch buffer
  for a build operation.

- `discardable` indicates whether or not the micromap object may be
  destroyed after an acceleration structure build or update. A false
  value means that acceleration structures built with this micromap
  **may** contain references to the data contained therein, and the
  application **must** not destroy the micromap until ray traversal has
  concluded. A true value means that the information in the micromap
  will be copied by value into the acceleration structure, and the
  micromap **may** be destroyed after the acceleration structure build
  concludes.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkMicromapBuildSizesInfoEXT-sType-sType"
  id="VUID-VkMicromapBuildSizesInfoEXT-sType-sType"></a>
  VUID-VkMicromapBuildSizesInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_BUILD_SIZES_INFO_EXT`

- <a href="#VUID-VkMicromapBuildSizesInfoEXT-pNext-pNext"
  id="VUID-VkMicromapBuildSizesInfoEXT-pNext-pNext"></a>
  VUID-VkMicromapBuildSizesInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetMicromapBuildSizesEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMicromapBuildSizesInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
