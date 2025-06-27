# VkSamplerReductionModeCreateInfo(3) Manual Page

## Name

VkSamplerReductionModeCreateInfo - Structure specifying sampler
reduction mode



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkSamplerReductionModeCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkSamplerReductionModeCreateInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkSamplerReductionMode    reductionMode;
} VkSamplerReductionModeCreateInfo;
```

or the equivalent

``` c
// Provided by VK_EXT_sampler_filter_minmax
typedef VkSamplerReductionModeCreateInfo VkSamplerReductionModeCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `reductionMode` is a
  [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html) value
  controlling how texture filtering combines texel values.

## <a href="#_description" class="anchor"></a>Description

If the `pNext` chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerCreateInfo.html)
includes a `VkSamplerReductionModeCreateInfo` structure, then that
structure includes a mode controlling how texture filtering combines
texel values.

If this structure is not present, `reductionMode` is considered to be
`VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`.

Valid Usage (Implicit)

- <a href="#VUID-VkSamplerReductionModeCreateInfo-sType-sType"
  id="VUID-VkSamplerReductionModeCreateInfo-sType-sType"></a>
  VUID-VkSamplerReductionModeCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO`

- <a href="#VUID-VkSamplerReductionModeCreateInfo-reductionMode-parameter"
  id="VUID-VkSamplerReductionModeCreateInfo-reductionMode-parameter"></a>
  VUID-VkSamplerReductionModeCreateInfo-reductionMode-parameter  
  `reductionMode` **must** be a valid
  [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_sampler_filter_minmax](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sampler_filter_minmax.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSamplerReductionMode.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSamplerReductionModeCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
