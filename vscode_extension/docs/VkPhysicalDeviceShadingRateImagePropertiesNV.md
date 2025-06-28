# VkPhysicalDeviceShadingRateImagePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceShadingRateImagePropertiesNV - Structure describing shading rate image limits that can be supported by an implementation



## [](#_c_specification)C Specification

The `VkPhysicalDeviceShadingRateImagePropertiesNV` structure is defined as:

```c++
// Provided by VK_NV_shading_rate_image
typedef struct VkPhysicalDeviceShadingRateImagePropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         shadingRateTexelSize;
    uint32_t           shadingRatePaletteSize;
    uint32_t           shadingRateMaxCoarseSamples;
} VkPhysicalDeviceShadingRateImagePropertiesNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- []()`shadingRateTexelSize` indicates the width and height of the portion of the framebuffer corresponding to each texel in the shading rate image.
- []()`shadingRatePaletteSize` indicates the maximum number of palette entries supported for the shading rate image.
- []()`shadingRateMaxCoarseSamples` specifies the maximum number of coverage samples supported in a single fragment. If the product of the fragment size derived from the base shading rate and the number of coverage samples per pixel exceeds this limit, the final shading rate will be adjusted so that its product does not exceed the limit.

## [](#_description)Description

If the `VkPhysicalDeviceShadingRateImagePropertiesNV` structure is included in the `pNext` chain of the [VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceProperties2.html) structure passed to [vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceProperties2.html), it is filled in with each corresponding implementation-dependent property.

These properties are related to the [shading rate image](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#primsrast-shading-rate-image) feature.

Valid Usage (Implicit)

- [](#VUID-VkPhysicalDeviceShadingRateImagePropertiesNV-sType-sType)VUID-VkPhysicalDeviceShadingRateImagePropertiesNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV`

## [](#_see_also)See Also

[VK\_NV\_shading\_rate\_image](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_shading_rate_image.html), [VkExtent2D](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceShadingRateImagePropertiesNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0