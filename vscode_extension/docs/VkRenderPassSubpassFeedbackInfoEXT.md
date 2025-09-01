# VkRenderPassSubpassFeedbackInfoEXT(3) Manual Page

## Name

VkRenderPassSubpassFeedbackInfoEXT - Feedback about the creation of subpass



## [](#_c_specification)C Specification

The `VkRenderPassSubpassFeedbackInfoEXT` structure is defined as:

Warning

This functionality is deprecated by [Vulkan Version 1.4](#versions-1.4). See [Deprecated Functionality](#deprecation-dynamicrendering) for more information.

```c++
// Provided by VK_EXT_subpass_merge_feedback
typedef struct VkRenderPassSubpassFeedbackInfoEXT {
    VkSubpassMergeStatusEXT    subpassMergeStatus;
    char                       description[VK_MAX_DESCRIPTION_SIZE];
    uint32_t                   postMergeIndex;
} VkRenderPassSubpassFeedbackInfoEXT;
```

## [](#_members)Members

- `subpassMergeStatus` is a `VkSubpassMergeStatusEXT` value specifying information about whether the subpass is merged with the previous subpass and the reason why it is not merged.
- `description` is an array of `VK_MAX_DESCRIPTION_SIZE` `char` containing a null-terminated UTF-8 string which provides additional details.
- `postMergeIndex` is the subpass index after the subpass merging.

## [](#_description)Description

## [](#_see_also)See Also

[VK\_EXT\_subpass\_merge\_feedback](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_subpass_merge_feedback.html), [VkRenderPassSubpassFeedbackCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassSubpassFeedbackCreateInfoEXT.html), [VkSubpassMergeStatusEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubpassMergeStatusEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkRenderPassSubpassFeedbackInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0