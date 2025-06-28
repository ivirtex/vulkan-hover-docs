# VkRenderPassCreationFeedbackCreateInfoEXT(3) Manual Page

## Name

VkRenderPassCreationFeedbackCreateInfoEXT - Request feedback about the creation of render pass



## [](#_c_specification)C Specification

To obtain feedback about the creation of a render pass, include a `VkRenderPassCreationFeedbackCreateInfoEXT` structure in the `pNext` chain of [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html). The `VkRenderPassCreationFeedbackCreateInfoEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationFeedbackCreateInfoEXT {
    VkStructureType                         sType;
    const void*                             pNext;
    VkRenderPassCreationFeedbackInfoEXT*    pRenderPassFeedback;
} VkRenderPassCreationFeedbackCreateInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pRenderPassFeedback` is a pointer to a [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackInfoEXT.html) structure in which feedback is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkRenderPassCreationFeedbackCreateInfoEXT-sType-sType)VUID-VkRenderPassCreationFeedbackCreateInfoEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_RENDER_PASS_CREATION_FEEDBACK_CREATE_INFO_EXT`
- [](#VUID-VkRenderPassCreationFeedbackCreateInfoEXT-pRenderPassFeedback-parameter)VUID-VkRenderPassCreationFeedbackCreateInfoEXT-pRenderPassFeedback-parameter  
  `pRenderPassFeedback` **must** be a valid pointer to a [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackInfoEXT.html) structure

## [](#_see_also)See Also

[VK\_EXT\_subpass\_merge\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subpass_merge_feedback.html), [VkRenderPassCreateInfo2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreateInfo2.html), [VkRenderPassCreationControlEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationControlEXT.html), [VkRenderPassCreationFeedbackInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackInfoEXT.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCreateRenderPass2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCreateRenderPass2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassCreationFeedbackCreateInfoEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0