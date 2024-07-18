# VkPhysicalDeviceShadingRateImagePropertiesNV(3) Manual Page

## Name

VkPhysicalDeviceShadingRateImagePropertiesNV - Structure describing
shading rate image limits that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceShadingRateImagePropertiesNV` structure is defined
as:

``` c
// Provided by VK_NV_shading_rate_image
typedef struct VkPhysicalDeviceShadingRateImagePropertiesNV {
    VkStructureType    sType;
    void*              pNext;
    VkExtent2D         shadingRateTexelSize;
    uint32_t           shadingRatePaletteSize;
    uint32_t           shadingRateMaxCoarseSamples;
} VkPhysicalDeviceShadingRateImagePropertiesNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- <span id="limits-shadingRateTexelSize"></span> `shadingRateTexelSize`
  indicates the width and height of the portion of the framebuffer
  corresponding to each texel in the shading rate image.

- <span id="limits-shadingRatePaletteSize"></span>
  `shadingRatePaletteSize` indicates the maximum number of palette
  entries supported for the shading rate image.

- <span id="limits-shadingRateMaxCoarseSamples"></span>
  `shadingRateMaxCoarseSamples` specifies the maximum number of coverage
  samples supported in a single fragment. If the product of the fragment
  size derived from the base shading rate and the number of coverage
  samples per pixel exceeds this limit, the final shading rate will be
  adjusted so that its product does not exceed the limit.

## <a href="#_description" class="anchor"></a>Description

If the `VkPhysicalDeviceShadingRateImagePropertiesNV` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

These properties are related to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#primsrast-shading-rate-image"
target="_blank" rel="noopener">shading rate image</a> feature.

Valid Usage (Implicit)

- <a href="#VUID-VkPhysicalDeviceShadingRateImagePropertiesNV-sType-sType"
  id="VUID-VkPhysicalDeviceShadingRateImagePropertiesNV-sType-sType"></a>
  VUID-VkPhysicalDeviceShadingRateImagePropertiesNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SHADING_RATE_IMAGE_PROPERTIES_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkExtent2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExtent2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceShadingRateImagePropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
