# VkRenderPassCreationFeedbackCreateInfoEXT(3) Manual Page

## Name

VkRenderPassCreationFeedbackCreateInfoEXT - Request feedback about the
creation of render pass



## <a href="#_c_specification" class="anchor"></a>C Specification

To obtain feedback about the creation of a render pass, include a
`VkRenderPassCreationFeedbackCreateInfoEXT` structure in the `pNext`
chain of [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html). The
`VkRenderPassCreationFeedbackCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationFeedbackCreateInfoEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkRenderPassCreationFeedbackInfoEXT*    pRenderPassFeedback;
} VkRenderPassCreationFeedbackCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `pRenderPassFeedback` is a pointer to a
  [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackInfoEXT.html)
  structure in which feedback is returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkRenderPassCreationFeedbackCreateInfoEXT-sType-sType"
  id="VUID-VkRenderPassCreationFeedbackCreateInfoEXT-sType-sType"></a>
  VUID-VkRenderPassCreationFeedbackCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_FEEDBACK_CREATE_INFO_EXT`

- <a
  href="#VUID-VkRenderPassCreationFeedbackCreateInfoEXT-pRenderPassFeedback-parameter"
  id="VUID-VkRenderPassCreationFeedbackCreateInfoEXT-pRenderPassFeedback-parameter"></a>
  VUID-VkRenderPassCreationFeedbackCreateInfoEXT-pRenderPassFeedback-parameter  
  `pRenderPassFeedback` **must** be a valid pointer to a
  [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackInfoEXT.html)
  structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_subpass_merge_feedback](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_subpass_merge_feedback.html),
[VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreateInfo2.html),
[VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationControlEXT.html),
[VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassCreationFeedbackInfoEXT.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateRenderPass2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkRenderPassCreationFeedbackCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
