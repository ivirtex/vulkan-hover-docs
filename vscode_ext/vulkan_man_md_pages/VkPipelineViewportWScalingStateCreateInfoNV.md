# VkPipelineViewportWScalingStateCreateInfoNV(3) Manual Page

## Name

VkPipelineViewportWScalingStateCreateInfoNV - Structure specifying
parameters of a newly created pipeline viewport W scaling state



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkPipelineViewportWScalingStateCreateInfoNV` structure is defined
as:

``` c
// Provided by VK_NV_clip_space_w_scaling
typedef struct VkPipelineViewportWScalingStateCreateInfoNV {
    VkStructureType                sType;
    const void*                    pNext;
    VkBool32                       viewportWScalingEnable;
    uint32_t                       viewportCount;
    const VkViewportWScalingNV*    pViewportWScalings;
} VkPipelineViewportWScalingStateCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `viewportWScalingEnable` controls whether viewport **W** scaling is
  enabled.

- `viewportCount` is the number of viewports used by **W** scaling, and
  **must** match the number of viewports in the pipeline if viewport
  **W** scaling is enabled.

- `pViewportWScalings` is a pointer to an array of
  [VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html) structures defining
  the **W** scaling parameters for the corresponding viewports. If the
  viewport **W** scaling state is dynamic, this member is ignored.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkPipelineViewportWScalingStateCreateInfoNV-sType-sType"
  id="VUID-VkPipelineViewportWScalingStateCreateInfoNV-sType-sType"></a>
  VUID-VkPipelineViewportWScalingStateCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_PIPELINE_VIEWPORT_W_SCALING_STATE_CREATE_INFO_NV`

- <a
  href="#VUID-VkPipelineViewportWScalingStateCreateInfoNV-viewportCount-arraylength"
  id="VUID-VkPipelineViewportWScalingStateCreateInfoNV-viewportCount-arraylength"></a>
  VUID-VkPipelineViewportWScalingStateCreateInfoNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_clip_space_w_scaling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_clip_space_w_scaling.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkViewportWScalingNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkViewportWScalingNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPipelineViewportWScalingStateCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
