# VkMultisampledRenderToSingleSampledInfoEXT(3) Manual Page

## Name

VkMultisampledRenderToSingleSampledInfoEXT - Structure containing info for multisampled rendering to single-sampled attachments in a subpass



## [](#_c_specification)C Specification

If the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html) or [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html) includes a `VkMultisampledRenderToSingleSampledInfoEXT` structure, then that structure describes how multisampled rendering is performed on single sampled attachments in that subpass.

The `VkMultisampledRenderToSingleSampledInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkMultisampledRenderToSingleSampledInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkBool32                 multisampledRenderToSingleSampledEnable;
    VkSampleCountFlagBits    rasterizationSamples;
} VkMultisampledRenderToSingleSampledInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `multisampledRenderToSingleSampledEnable` controls whether multisampled rendering to single-sampled attachments is performed as described [below](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#multisampled-render-to-single-sampled).
- `rasterizationSamples` is a [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) specifying the number of samples used in rasterization.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-06878)VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-06878  
  The value of `rasterizationSamples` **must** not be `VK_SAMPLE_COUNT_1_BIT`
- [](#VUID-VkMultisampledRenderToSingleSampledInfoEXT-pNext-06880)VUID-VkMultisampledRenderToSingleSampledInfoEXT-pNext-06880  
  If added to the `pNext` chain of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html), each `imageView` member of any element of [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pColorAttachments`, [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pDepthAttachment`, or [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInfo.html)::`pStencilAttachment` that is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) **must** have a format that supports the sample count specified in `rasterizationSamples`

Valid Usage (Implicit)

- [](#VUID-VkMultisampledRenderToSingleSampledInfoEXT-sType-sType)VUID-VkMultisampledRenderToSingleSampledInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_INFO_EXT`
- [](#VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-parameter)VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-parameter  
  `rasterizationSamples` **must** be a valid [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html) value

## [](#_see_also)See Also

[VK\_EXT\_multisampled\_render\_to\_single\_sampled](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_multisampled_render_to_single_sampled.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSampleCountFlagBits.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMultisampledRenderToSingleSampledInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0