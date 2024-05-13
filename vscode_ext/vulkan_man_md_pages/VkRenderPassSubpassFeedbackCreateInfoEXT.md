# VkRenderPassSubpassFeedbackCreateInfoEXT(3) Manual Page

## Name

VkRenderPassSubpassFeedbackCreateInfoEXT - Request for feedback about
the creation of subpass



## <a href="#_c_specification" class="anchor"></a>C Specification

Feedback about the creation of a subpass **can** be obtained by
including a `VkRenderPassSubpassFeedbackCreateInfoEXT` structure in the
`pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html).
`VkRenderPassSubpassFeedbackCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassSubpassFeedbackCreateInfoEXT {
    VkStructureType                        sType;
    const void*                            pNext;
    VkRenderPassSubpassFeedbackInfoEXT*    pSubpassFeedback;
} VkRenderPassSubpassFeedbackCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pSubpassFeedback` is a pointer to a
  [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html)
  structure in which feedback is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-sType-sType"
  id="VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-sType-sType"></a>
  VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_SUBPASS_FEEDBACK_CREATE_INFO_EXT`

- <a
  href="#VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-pSubpassFeedback-parameter"
  id="VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-pSubpassFeedback-parameter"></a>
  VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-pSubpassFeedback-parameter  
  `pSubpassFeedback` **must** be a valid pointer to a
  [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationControlEXT.html),
[VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSubpassFeedbackInfoEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSubpassDescription2.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassSubpassFeedbackCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
