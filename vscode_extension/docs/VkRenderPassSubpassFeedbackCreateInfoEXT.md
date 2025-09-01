# VkRenderPassSubpassFeedbackCreateInfoEXT(3) Manual Page

## Name

VkRenderPassSubpassFeedbackCreateInfoEXT - Request for feedback about the creation of subpass



## [](#_c_specification)C Specification

Feedback about the creation of a subpass **can** be obtained by including a `VkRenderPassSubpassFeedbackCreateInfoEXT` structure in the `pNext` chain of [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html). `VkRenderPassSubpassFeedbackCreateInfoEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassSubpassFeedbackCreateInfoEXT {
    VkStructureType                        sType;
    const void*                            pNext;
    VkRenderPassSubpassFeedbackInfoEXT*    pSubpassFeedback;
} VkRenderPassSubpassFeedbackCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pSubpassFeedback` is a pointer to a [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackInfoEXT.html) structure in which feedback is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-sType-sType)VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_SUBPASS_FEEDBACK_CREATE_INFO_EXT`
- [](#VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-pSubpassFeedback-parameter)VUID-VkRenderPassSubpassFeedbackCreateInfoEXT-pSubpassFeedback-parameter  
  `pSubpassFeedback` **must** be a valid pointer to a [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_subpass\_merge\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subpass_merge_feedback.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationControlEXT.html), [VkRenderPassSubpassFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackInfoEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSubpassDescription2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassDescription2.html), [vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassSubpassFeedbackCreateInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0