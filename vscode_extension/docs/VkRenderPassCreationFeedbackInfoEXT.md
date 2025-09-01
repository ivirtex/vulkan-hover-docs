# VkRenderPassCreationFeedbackInfoEXT(3) Manual Page

## Name

VkRenderPassCreationFeedbackInfoEXT - Feedback about the creation of a render pass



## [](#_c_specification)C Specification

The `VkRenderPassCreationFeedbackInfoEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassCreationFeedbackInfoEXT {
    uint32_t    postMergeSubpassCount;
} VkRenderPassCreationFeedbackInfoEXT;
```

## [](#_members)Members

- `postMergeSubpassCount` is the subpass count after merge.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_EXT\_subpass\_merge\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subpass_merge_feedback.html), [VkRenderPassCreationFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassCreationFeedbackCreateInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassCreationFeedbackInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0