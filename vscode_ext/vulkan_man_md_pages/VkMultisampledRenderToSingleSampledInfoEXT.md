# VkMultisampledRenderToSingleSampledInfoEXT(3) Manual Page

## Name

VkMultisampledRenderToSingleSampledInfoEXT - Structure containing info
for multisampled rendering to single-sampled attachments in a subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html) or
[VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html) includes a
`VkMultisampledRenderToSingleSampledInfoEXT` structure, then that
structure describes how multisampled rendering is performed on single
sampled attachments in that subpass.

The `VkMultisampledRenderToSingleSampledInfoEXT` structure is defined
as:

``` c
// Provided by VK_EXT_multisampled_render_to_single_sampled
typedef struct VkMultisampledRenderToSingleSampledInfoEXT {
    VkStructureType          sType;
    const void*              pNext;
    VkBool32                 multisampledRenderToSingleSampledEnable;
    VkSampleCountFlagBits    rasterizationSamples;
} VkMultisampledRenderToSingleSampledInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `multisampledRenderToSingleSampledEnable` controls whether
  multisampled rendering to single-sampled attachments is performed as
  described <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#multisampled-render-to-single-sampled"
  target="_blank" rel="noopener">below</a>.

- `rasterizationSamples` is a
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) specifying the
  number of samples used in rasterization.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-06878"
  id="VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-06878"></a>
  VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-06878  
  The value of `rasterizationSamples` **must** not be
  `VK_SAMPLE_COUNT_1_BIT`

- <a href="#VUID-VkMultisampledRenderToSingleSampledInfoEXT-pNext-06880"
  id="VUID-VkMultisampledRenderToSingleSampledInfoEXT-pNext-06880"></a>
  VUID-VkMultisampledRenderToSingleSampledInfoEXT-pNext-06880  
  If added to the `pNext` chain of
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html), each `imageView` member of
  any element of
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pColorAttachments`,
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pDepthAttachment`, or
  [VkRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInfo.html)::`pStencilAttachment` that is
  not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) **must** have a format that
  supports the sample count specified in `rasterizationSamples`

Valid Usage (Implicit)

- <a href="#VUID-VkMultisampledRenderToSingleSampledInfoEXT-sType-sType"
  id="VUID-VkMultisampledRenderToSingleSampledInfoEXT-sType-sType"></a>
  VUID-VkMultisampledRenderToSingleSampledInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MULTISAMPLED_RENDER_TO_SINGLE_SAMPLED_INFO_EXT`

- <a
  href="#VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-parameter"
  id="VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-parameter"></a>
  VUID-VkMultisampledRenderToSingleSampledInfoEXT-rasterizationSamples-parameter  
  `rasterizationSamples` **must** be a valid
  [VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html) value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_multisampled_render_to_single_sampled](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_multisampled_render_to_single_sampled.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkSampleCountFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleCountFlagBits.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMultisampledRenderToSingleSampledInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
