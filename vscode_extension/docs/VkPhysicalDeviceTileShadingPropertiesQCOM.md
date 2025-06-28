# VkPhysicalDeviceTileShadingPropertiesQCOM(3) Manual Page

## Name

VkPhysicalDeviceTileShadingPropertiesQCOM - Structure describing properties supported by VK\_QCOM\_tile\_shading



## [](#_c_specification)C Specification

The `VkPhysicalDeviceTileShadingPropertiesQCOM` structure is defined as:

```c++
// Provided by VK_QCOM_tile_shading
typedef struct VkPhysicalDeviceTileShadingPropertiesQCOM {
    VkStructureType    sType;
    void*              pNext;
    uint32_t           maxApronSize;
    VkBool32           preferNonCoherent;
    VkExtent2D         tileGranularity;
    VkExtent2D         maxTileShadingRate;
} VkPhysicalDeviceTileShadingPropertiesQCOM;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`maxApronSize` is the maximum value supported which can be specified for [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html)::`apronSize` or `width` and `height`.
- []()`preferNonCoherent` indicates that the implementation prefers tile attachments declared in shaders with the `NonCoherentTileAttachmentReadQCOM` decoration. Use of the decoration **may** offer performance or power advantages.
- []()`tileGranularity` provides a guarantee on the granularity of each tile. Each tile will have dimensions that are a multiple of this granularity in width and height.
- []()`maxTileShadingRate` is the maximum value of `TileShadingRateQCOM` and **must** be a power of 2.

## [](#_description)Description

If the `VkPhysicalDeviceTileShadingPropertiesQCOM` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceTileShadingPropertiesQCOM-sType-sType)VUID-VkPhysicalDeviceTileShadingPropertiesQCOM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_TILE_SHADING_PROPERTIES_QCOM`

## [](#_see_also)See Also

[VK\_QCOM\_tile\_shading](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_QCOM_tile_shading.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceTileShadingPropertiesQCOM)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0