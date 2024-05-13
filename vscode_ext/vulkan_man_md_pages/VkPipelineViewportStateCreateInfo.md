# VkPipelineViewportStateCreateInfo(3) Manual Page

## Name

VkPipelineViewportStateCreateInfo - Structure specifying parameters of a
newly created pipeline viewport state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineViewportStateCreateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_0
typedef struct VkPipelineViewportStateCreateInfo {
    VkStructureType                       sType;
    const void*                           pNext;
    VkPipelineViewportStateCreateFlags    flags;
    uint32_t                              viewportCount;
    const VkViewport*                     pViewports;
    uint32_t                              scissorCount;
    const VkRect2D*                       pScissors;
} VkPipelineViewportStateCreateInfo;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `flags` is reserved for future use.

- `viewportCount` is the number of viewports used by the pipeline.

- `pViewports` is a pointer to an array of [VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)
  structures, defining the viewport transforms. If the viewport state is
  dynamic, this member is ignored.

- `scissorCount` is the number of <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#fragops-scissor"
  target="_blank" rel="noopener">scissors</a> and **must** match the
  number of viewports.

- `pScissors` is a pointer to an array of [VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html)
  structures defining the rectangular bounds of the scissor for the
  corresponding viewport. If the scissor state is dynamic, this member
  is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkPipelineViewportStateCreateInfo-viewportCount-01216"
  id="VUID-VkPipelineViewportStateCreateInfo-viewportCount-01216"></a>
  VUID-VkPipelineViewportStateCreateInfo-viewportCount-01216  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `viewportCount` **must** not be greater than `1`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-scissorCount-01217"
  id="VUID-VkPipelineViewportStateCreateInfo-scissorCount-01217"></a>
  VUID-VkPipelineViewportStateCreateInfo-scissorCount-01217  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `scissorCount` **must** not be greater than `1`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-viewportCount-01218"
  id="VUID-VkPipelineViewportStateCreateInfo-viewportCount-01218"></a>
  VUID-VkPipelineViewportStateCreateInfo-viewportCount-01218  
  `viewportCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewports`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-scissorCount-01219"
  id="VUID-VkPipelineViewportStateCreateInfo-scissorCount-01219"></a>
  VUID-VkPipelineViewportStateCreateInfo-scissorCount-01219  
  `scissorCount` **must** be less than or equal to
  `VkPhysicalDeviceLimits`::`maxViewports`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-x-02821"
  id="VUID-VkPipelineViewportStateCreateInfo-x-02821"></a>
  VUID-VkPipelineViewportStateCreateInfo-x-02821  
  The `x` and `y` members of `offset` member of any element of
  `pScissors` **must** be greater than or equal to `0`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-offset-02822"
  id="VUID-VkPipelineViewportStateCreateInfo-offset-02822"></a>
  VUID-VkPipelineViewportStateCreateInfo-offset-02822  
  Evaluation of (`offset.x` + `extent.width`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-offset-02823"
  id="VUID-VkPipelineViewportStateCreateInfo-offset-02823"></a>
  VUID-VkPipelineViewportStateCreateInfo-offset-02823  
  Evaluation of (`offset.y` + `extent.height`) **must** not cause a
  signed integer addition overflow for any element of `pScissors`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-scissorCount-04134"
  id="VUID-VkPipelineViewportStateCreateInfo-scissorCount-04134"></a>
  VUID-VkPipelineViewportStateCreateInfo-scissorCount-04134  
  If `scissorCount` and `viewportCount` are both not dynamic, then
  `scissorCount` and `viewportCount` **must** be identical

- <a href="#VUID-VkPipelineViewportStateCreateInfo-viewportCount-04135"
  id="VUID-VkPipelineViewportStateCreateInfo-viewportCount-04135"></a>
  VUID-VkPipelineViewportStateCreateInfo-viewportCount-04135  
  If the graphics pipeline is being created with
  `VK_DYNAMIC_STATE_VIEWPORT_WITH_COUNT` set then `viewportCount`
  **must** be `0`, otherwise `viewportCount` **must** be greater than
  `0`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-scissorCount-04136"
  id="VUID-VkPipelineViewportStateCreateInfo-scissorCount-04136"></a>
  VUID-VkPipelineViewportStateCreateInfo-scissorCount-04136  
  If the graphics pipeline is being created with
  `VK_DYNAMIC_STATE_SCISSOR_WITH_COUNT` set then `scissorCount` **must**
  be `0`, otherwise `scissorCount` **must** be greater than `0`

- <a
  href="#VUID-VkPipelineViewportStateCreateInfo-viewportWScalingEnable-01726"
  id="VUID-VkPipelineViewportStateCreateInfo-viewportWScalingEnable-01726"></a>
  VUID-VkPipelineViewportStateCreateInfo-viewportWScalingEnable-01726  
  If the `viewportWScalingEnable` member of a
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  structure included in the `pNext` chain is `VK_TRUE`, the
  `viewportCount` member of the
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)
  structure **must** be greater than or equal to
  [VkPipelineViewportStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateInfo.html)::`viewportCount`

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineViewportStateCreateInfo-sType-sType"
  id="VUID-VkPipelineViewportStateCreateInfo-sType-sType"></a>
  VUID-VkPipelineViewportStateCreateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_STATE_CREATE_INFO`

- <a href="#VUID-VkPipelineViewportStateCreateInfo-pNext-pNext"
  id="VUID-VkPipelineViewportStateCreateInfo-pNext-pNext"></a>
  VUID-VkPipelineViewportStateCreateInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html),
  [VkPipelineViewportDepthClipControlCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportDepthClipControlCreateInfoEXT.html),
  [VkPipelineViewportExclusiveScissorStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportExclusiveScissorStateCreateInfoNV.html),
  [VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html),
  [VkPipelineViewportSwizzleStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportSwizzleStateCreateInfoNV.html),
  or
  [VkPipelineViewportWScalingStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportWScalingStateCreateInfoNV.html)

- <a href="#VUID-VkPipelineViewportStateCreateInfo-sType-unique"
  id="VUID-VkPipelineViewportStateCreateInfo-sType-unique"></a>
  VUID-VkPipelineViewportStateCreateInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkPipelineViewportStateCreateInfo-flags-zerobitmask"
  id="VUID-VkPipelineViewportStateCreateInfo-flags-zerobitmask"></a>
  VUID-VkPipelineViewportStateCreateInfo-flags-zerobitmask  
  `flags` **must** be `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkGraphicsPipelineCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkGraphicsPipelineCreateInfo.html),
[VkPipelineViewportStateCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportStateCreateFlags.html),
[VkRect2D](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRect2D.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkViewport](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewport.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportStateCreateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
