# VkSamplerReductionModeCreateInfo(3) Manual Page

## Name

VkSamplerReductionModeCreateInfo - Structure specifying sampler reduction mode



## [](#_c_specification)C Specification

The `VkSamplerReductionModeCreateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkSamplerReductionModeCreateInfo {
    VkStructureType           sType;
    const void*               pNext;
    VkSamplerReductionMode    reductionMode;
} VkSamplerReductionModeCreateInfo;
```

or the equivalent

```c++
// Provided by VK_EXT_sampler_filter_minmax
typedef VkSamplerReductionModeCreateInfo VkSamplerReductionModeCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `reductionMode` is a [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html) value controlling how texture filtering combines texel values.

## [](#_description)Description

If the `pNext` chain of [VkSamplerCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerCreateInfo.html) includes a `VkSamplerReductionModeCreateInfo` structure, then that structure includes a mode controlling how texture filtering combines texel values.

If this structure is not present, `reductionMode` is considered to be `VK_SAMPLER_REDUCTION_MODE_WEIGHTED_AVERAGE`.

Valid Usage (Implicit)

- [](#VUID-VkSamplerReductionModeCreateInfo-sType-sType)VUID-VkSamplerReductionModeCreateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_SAMPLER_REDUCTION_MODE_CREATE_INFO`
- [](#VUID-VkSamplerReductionModeCreateInfo-reductionMode-parameter)VUID-VkSamplerReductionModeCreateInfo-reductionMode-parameter  
  `reductionMode` **must** be a valid [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html) value

## [](#_see_also)See Also

[VK\_EXT\_sampler\_filter\_minmax](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_sampler_filter_minmax.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkSamplerReductionMode](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSamplerReductionMode.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkSamplerReductionModeCreateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0