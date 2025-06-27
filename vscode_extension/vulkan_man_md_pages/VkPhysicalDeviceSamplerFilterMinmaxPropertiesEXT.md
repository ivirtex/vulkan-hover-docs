# VkPhysicalDeviceSamplerFilterMinmaxProperties(3) Manual Page

## Name

VkPhysicalDeviceSamplerFilterMinmaxProperties - Structure describing
sampler filter minmax limits that can be supported by an implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPhysicalDeviceSamplerFilterMinmaxProperties` structure is defined
as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkPhysicalDeviceSamplerFilterMinmaxProperties {
    VkStructureType    sType;
    void*              pNext;
    VkBool32           filterMinmaxSingleComponentFormats;
    VkBool32           filterMinmaxImageComponentMapping;
} VkPhysicalDeviceSamplerFilterMinmaxProperties;
```

or the equivalent

``` c
// Provided by VK_EXT_sampler_filter_minmax
typedef VkPhysicalDeviceSamplerFilterMinmaxProperties VkPhysicalDeviceSamplerFilterMinmaxPropertiesEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

## <a href="#_description" class="anchor"></a>Description

- <span id="extension-limits-filterMinmaxSingleComponentFormats"></span>
  `filterMinmaxSingleComponentFormats` is a boolean value indicating
  whether a minimum set of required formats support min/max filtering.

- <span id="extension-limits-filterMinmaxImageComponentMapping"></span>
  `filterMinmaxImageComponentMapping` is a boolean value indicating
  whether the implementation supports non-identity component mapping of
  the image when doing min/max filtering.

If the `VkPhysicalDeviceSamplerFilterMinmaxProperties` structure is
included in the `pNext` chain of the
[VkPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceProperties2.html)
structure passed to
[vkGetPhysicalDeviceProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceProperties2.html),
it is filled in with each corresponding implementation-dependent
property.

If `filterMinmaxSingleComponentFormats` is `VK_TRUE`, the following
formats **must** support the
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_FILTER_MINMAX_BIT` feature with
`VK_IMAGE_TILING_OPTIMAL`, if they support
`VK_FORMAT_FEATURE_SAMPLED_IMAGE_BIT`:

- `VK_FORMAT_R8_UNORM`

- `VK_FORMAT_R8_SNORM`

- `VK_FORMAT_R16_UNORM`

- `VK_FORMAT_R16_SNORM`

- `VK_FORMAT_R16_SFLOAT`

- `VK_FORMAT_R32_SFLOAT`

- `VK_FORMAT_D16_UNORM`

- `VK_FORMAT_X8_D24_UNORM_PACK32`

- `VK_FORMAT_D32_SFLOAT`

- `VK_FORMAT_D16_UNORM_S8_UINT`

- `VK_FORMAT_D24_UNORM_S8_UINT`

- `VK_FORMAT_D32_SFLOAT_S8_UINT`

If the format is a depth/stencil format, this bit only specifies that
the depth aspect (not the stencil aspect) of an image of this format
supports min/max filtering, and that min/max filtering of the depth
aspect is supported when depth compare is disabled in the sampler.

If `filterMinmaxImageComponentMapping` is `VK_FALSE` the component
mapping of the image view used with min/max filtering **must** have been
created with the `r` component set to the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-image-views-identity-mappings"
target="_blank" rel="noopener">identity swizzle</a>. Only the `r`
component of the sampled image value is defined and the other component
values are undefined. If `filterMinmaxImageComponentMapping` is
`VK_TRUE` this restriction does not apply and image component mapping
works as normal.

Valid Usage (Implicit)

- <a
  href="#VUID-VkPhysicalDeviceSamplerFilterMinmaxProperties-sType-sType"
  id="VUID-VkPhysicalDeviceSamplerFilterMinmaxProperties-sType-sType"></a>
  VUID-VkPhysicalDeviceSamplerFilterMinmaxProperties-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PHYSICAL_DEVICE_SAMPLER_FILTER_MINMAX_PROPERTIES`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sampler_filter_minmax](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sampler_filter_minmax.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceSamplerFilterMinmaxProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
