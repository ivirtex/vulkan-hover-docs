# VkMicromapBuildSizesInfoEXT(3) Manual Page

## Name

VkMicromapBuildSizesInfoEXT - Structure specifying build sizes for a micromap



## [](#_c_specification)C Specification

The `VkMicromapBuildSizesInfoEXT` structure describes the required build sizes for a micromap and scratch buffers and is defined as:

```c++
// Provided by VK_EXT_opacity_micromap
typedef struct VkMicromapBuildSizesInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       micromapSize;
    VkDeviceSize       buildScratchSize;
    VkBool32           discardable;
} VkMicromapBuildSizesInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `micromapSize` is the size in bytes required in a [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) for a build or update operation.
- `buildScratchSize` is the size in bytes required in a scratch buffer for a build operation.
- `discardable` indicates whether or not the micromap object may be destroyed after an acceleration structure build or update. A false value means that acceleration structures built with this micromap **may** contain references to the data contained therein, and the application **must** not destroy the micromap until ray traversal has concluded. A true value means that the information in the micromap will be copied by value into the acceleration structure, and the micromap **may** be destroyed after the acceleration structure build concludes.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkMicromapBuildSizesInfoEXT-sType-sType)VUID-VkMicromapBuildSizesInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MICROMAP_BUILD_SIZES_INFO_EXT`
- [](#VUID-VkMicromapBuildSizesInfoEXT-pNext-pNext)VUID-VkMicromapBuildSizesInfoEXT-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMicromapBuildSizesInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0